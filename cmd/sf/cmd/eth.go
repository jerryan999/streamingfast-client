package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/streamingfast/eth-go"
	sf "github.com/streamingfast/streamingfast-client"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pbtransform "github.com/streamingfast/streamingfast-client/pb/sf/ethereum/transform/v1"
	pbcodec "github.com/streamingfast/streamingfast-client/pb/sf/ethereum/type/v2"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var ethSfCmd = &cobra.Command{
	Use:   "eth [flags] [<start_block>] [<end_block>]",
	Short: `StreamingFast Ethereum client`,
	Long:  usage,
	Args:  cobra.MaximumNArgs(2),
	RunE:  ethSfRunE,
}

func init() {
	RootCmd.AddCommand(ethSfCmd)

	defaultEthEndpoint := "api.streamingfast.io:443"
	if e := os.Getenv("STREAMINGFAST_ENDPOINT"); e != "" {
		defaultEthEndpoint = e
	}
	ethSfCmd.Flags().StringP("endpoint", "e", defaultEthEndpoint, "The endpoint to connect the stream of blocks (default value set by STREAMINGFAST_ENDPOINT env var, can be overriden by network-specific flags like --bsc)")

	// Endpoint settings
	ethSfCmd.Flags().Bool("bsc", false, "When set, will change the default endpoint to Binance Smart Chain")
	ethSfCmd.Flags().Bool("polygon", false, "When set, will change the default endpoint to Polygon (previously Matic)")
	ethSfCmd.Flags().Bool("heco", false, "When set, will change the default endpoint to Huobi Eco Chain")
	ethSfCmd.Flags().Bool("fantom", false, "When set, will change the default endpoint to Fantom Opera Mainnet")
	ethSfCmd.Flags().Bool("xdai", false, "When set, will change the default endpoint to xDai Chain")

	// Transforms
	ethSfCmd.Flags().Bool("light-block", false, "When set, returned blocks will be stripped of some information")
	ethSfCmd.Flags().StringSlice("log-filter-multi", nil, "Advanced filter. List of 'address[+address[+...]]:eventsig[+eventsig[+...]]' pairs, ex: 'dead+beef:1234+5678,:0x44,0x12:' results in 3 filters. Mutually exclusive with --log-filter-addresses and --log-filter-event-sigs.")
	ethSfCmd.Flags().StringSlice("log-filter-addresses", nil, "Basic filter. List of addresses with which to filter blocks. Mutually exclusive with --log-filter-multi.")
	ethSfCmd.Flags().StringSlice("log-filter-event-sigs", nil, "Basic filter. List of event signatures with which to filter blocks. Mutually exclusive with --log-filter-multi.")

	ethSfCmd.Flags().StringSlice("call-filter-multi", nil, "Advanced filter. List of 'address[+address[+...]]:eventsig[+eventsig[+...]]' pairs, ex: 'dead+beef:1234+5678,:0x44,0x12:' results in 3 filters. Mutually exclusive with --call-filter-addresses and --call-filter-sigs.")
	ethSfCmd.Flags().StringSlice("call-filter-addresses", nil, "Basic filter. List of addresses with which to filter blocks. Mutually exclusive with --call-filter-multi.")
	ethSfCmd.Flags().StringSlice("call-filter-sigs", nil, "Basic filter. List of method signatures with which to filter blocks. Mutually exclusive with --call-filter-multi.")
}

func ethSfRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	startCursor := viper.GetString("global-start-cursor")

	endpoint := viper.GetString("eth-cmd-endpoint")
	bscNetwork := viper.GetBool("eth-cmd-bsc")
	polygonNetwork := viper.GetBool("eth-cmd-polygon")
	hecoNetwork := viper.GetBool("eth-cmd-heco")
	fantomNetwork := viper.GetBool("eth-cmd-fantom")
	xdaiNetwork := viper.GetBool("eth-cmd-xdai")
	outputFlag := viper.GetString("global-output-key")
	redisHost := viper.GetString("global-redis-host")
	redisPassword := viper.GetString("global-redis-password")
	redisDB := viper.GetInt("global-redis-db")
	yamlEnable := viper.GetBool("global-yaml-enable")

	network := "eth"
	switch {
	case bscNetwork:
		endpoint = "bsc.streamingfast.io:443"
		network = "bsc"
	case polygonNetwork:
		endpoint = "polygon.streamingfast.io:443"
		network = "polygon"
	case hecoNetwork:
		endpoint = "heco.streamingfast.io:443"
		network = "heco"
	case fantomNetwork:
		endpoint = "fantom.streamingfast.io:443"
		network = "fantom"
	case xdaiNetwork:
		endpoint = "xdai.streamingfast.io:443"
		network = "xdai"
	}
	if endpoint == "" {
		return fmt.Errorf("unable to resolve endpoint")
	}

	inputs, err := checkArgs(startCursor, args)
	if err != nil {
		return err
	}
	zlog.Debug("arguments", zap.Reflect("args", inputs))

	if !noMoreThanOneTrue(bscNetwork, polygonNetwork, hecoNetwork, fantomNetwork, xdaiNetwork) {
		return fmt.Errorf("cannot set more than one network flag (ex: --polygon, --bsc)")
	}

	transforms := []*anypb.Any{}

	if viper.GetBool("eth-cmd-light-block") {
		t, err := lightBlockTransform()
		if err != nil {
			return fmt.Errorf("unable to create light block transform: %w", err)
		}
		transforms = append(transforms, t)
	}

	multiFilter := viper.GetStringSlice("eth-cmd-log-filter-multi")
	addrFilters := viper.GetStringSlice("eth-cmd-log-filter-addresses")
	sigFilters := viper.GetStringSlice("eth-cmd-log-filter-event-sigs")

	if yamlEnable {
		zlog.Info("yaml enabled, it will override all other in CMD")

		// read the yaml file
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			zlog.Error("yamlFile.Get err   #%v ", zap.Error(err))
			zlog.Debug("use command line options instead")
		}

		// unmarshal the yaml file
		var yamlConfig YamlConfig
		err = yaml.Unmarshal(yamlFile, &yamlConfig)
		if err != nil {
			zlog.Error("Unmarshal: %v", zap.Error(err))
			zlog.Debug("use command line options instead")
		}

		// get the target network
		found := false
		var outputKey string
		zlog.Info("network: ", zap.String("network", network))
		zlog.Info("Compare log filter: ", zap.String("kind", string(KindBlockFilter)))
		for _, dataSource := range yamlConfig.DataSources {
			if dataSource.Kind == KindBlockFilter && dataSource.Network == Network(network) {
				sigFilters = dataSource.Topics
				addrFilters = dataSource.Address
				outputKey = dataSource.OutputKey
				found = true
				break
			}
		}
		if found {
			zlog.Warn("using yaml config for log filter, command line options will be ignored")
			// update output key
			outputFlag = outputKey

		} else {
			zlog.Warn("no yaml config found for log filter, use command line options instead")
		}

	}
	writer, closer, err := blockWriter(inputs.Range, outputFlag, redisHost, redisPassword, redisDB)
	if err != nil {
		return fmt.Errorf("unable to setup writer: %w", err)
	}
	defer closer()

	hasMultiFilter := len(multiFilter) != 0
	hasSingleFilter := len(addrFilters) != 0 || len(sigFilters) != 0

	switch {
	case hasMultiFilter && hasSingleFilter:
		return fmt.Errorf("options --log-filter-{addresses|event-sigs} are incompatible with --log-filter-multi, don't use both")
	case hasMultiFilter:
		mft, err := parseMultiLogFilter(multiFilter)
		if err != nil {
			return err
		}
		if mft != nil {
			transforms = append(transforms, mft)
		}
		break
	case hasSingleFilter:

		t, err := parseSingleLogFilter(addrFilters, sigFilters)
		if err != nil {
			return fmt.Errorf("unable to create log filter transform: %w", err)
		}

		if t != nil {
			transforms = append(transforms, t)
		}
		break
	}

	multiCallFilter := viper.GetStringSlice("eth-cmd-call-filter-multi")
	addrCallFilters := viper.GetStringSlice("eth-cmd-call-filter-addresses")
	sigCallFilters := viper.GetStringSlice("eth-cmd-call-filter-sigs")

	hasMultiCallFilter := len(multiCallFilter) != 0
	hasSingleCallFilter := len(addrCallFilters) != 0 || len(sigCallFilters) != 0

	switch {
	case hasMultiCallFilter && hasSingleCallFilter:
		return fmt.Errorf("options --call-filter-{addresses|sigs} are incompatible with --call-filter-multi, don't use both")
	case hasMultiCallFilter:
		mft, err := parseMultiCallToFilter(multiCallFilter)
		if err != nil {
			return err
		}
		if mft != nil {
			transforms = append(transforms, mft)
		}
		break
	case hasSingleCallFilter:
		t, err := parseSingleCallToFilter(addrCallFilters, sigCallFilters)
		if err != nil {
			return fmt.Errorf("unable to create log filter transform: %w", err)
		}

		if t != nil {
			transforms = append(transforms, t)
		}
		break
	}

	return launchStream(ctx, streamConfig{
		writer:      writer,
		stats:       newStats(),
		brange:      inputs.Range,
		cursor:      startCursor,
		endpoint:    endpoint,
		handleForks: viper.GetBool("global-handle-forks"),
		transforms:  transforms,
	},
		func() proto.Message {
			return &pbcodec.Block{}
		},
		func(message proto.Message) sf.BlockRef {
			return message.(*pbcodec.Block).AsRef()
		})
}

func lightBlockTransform() (*anypb.Any, error) {
	transform := &pbtransform.LightBlock{}
	return anypb.New(transform)
}

func basicLogFilter(addrs []eth.Address, sigs []eth.Hash) *pbtransform.CombinedFilter {
	var addrBytes [][]byte
	var sigsBytes [][]byte

	for _, addr := range addrs {
		b := addr.Bytes()
		addrBytes = append(addrBytes, b)
	}

	for _, sig := range sigs {
		b := sig.Bytes()
		sigsBytes = append(sigsBytes, b)
	}

	f := &pbtransform.LogFilter{
		Addresses:       addrBytes,
		EventSignatures: sigsBytes,
	}
	return &pbtransform.CombinedFilter{
		LogFilters: []*pbtransform.LogFilter{f},
	}
}

func parseSingleLogFilter(addrFilters []string, sigFilters []string) (*anypb.Any, error) {
	var addrs []eth.Address
	for _, addrString := range addrFilters {
		addr := eth.MustNewAddress(addrString)
		addrs = append(addrs, addr)
	}

	var sigs []eth.Hash
	for _, sigString := range sigFilters {
		sig := eth.MustNewHash(sigString)
		sigs = append(sigs, sig)
	}

	transform := basicLogFilter(addrs, sigs)
	return anypb.New(transform)
}

func parseMultiLogFilter(in []string) (*anypb.Any, error) {

	mf := &pbtransform.CombinedFilter{}

	for _, filter := range in {
		parts := strings.Split(filter, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("option --log-filter-multi must be of type address_hash+address_hash+address_hash:event_sig_hash+event_sig_hash (repeated, separated by comma)")
		}
		var addrs []eth.Address
		for _, a := range strings.Split(parts[0], "+") {
			if a != "" {
				addr := eth.MustNewAddress(a)
				addrs = append(addrs, addr)
			}
		}
		var sigs []eth.Hash
		for _, s := range strings.Split(parts[1], "+") {
			if s != "" {
				sig := eth.MustNewHash(s)
				sigs = append(sigs, sig)
			}
		}

		mf.LogFilters = append(mf.LogFilters, basicLogFilter(addrs, sigs).LogFilters...)
	}

	t, err := anypb.New(mf)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func parseSingleCallToFilter(addrFilters []string, sigFilters []string) (*anypb.Any, error) {
	var addrs []eth.Address
	for _, addrString := range addrFilters {
		addr := eth.MustNewAddress(addrString)
		addrs = append(addrs, addr)
	}

	var sigs []eth.Hash
	for _, sigString := range sigFilters {
		sig := eth.MustNewHash(sigString)
		sigs = append(sigs, sig)
	}

	transform := basicCallToFilter(addrs, sigs)
	return anypb.New(transform)
}

func parseMultiCallToFilter(in []string) (*anypb.Any, error) {

	mf := &pbtransform.MultiCallToFilter{}

	for _, filter := range in {
		parts := strings.Split(filter, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("option --log-filter-multi must be of type address_hash+address_hash+address_hash:event_sig_hash+event_sig_hash (repeated, separated by comma)")
		}
		var addrs []eth.Address
		for _, a := range strings.Split(parts[0], "+") {
			if a != "" {
				addr := eth.MustNewAddress(a)
				addrs = append(addrs, addr)
			}
		}
		var sigs []eth.Hash
		for _, s := range strings.Split(parts[1], "+") {
			if s != "" {
				sig := eth.MustNewHash(s)
				sigs = append(sigs, sig)
			}
		}

		mf.CallFilters = append(mf.CallFilters, basicCallToFilter(addrs, sigs).CallFilters...)
	}

	t, err := anypb.New(mf)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func basicCallToFilter(addrs []eth.Address, sigs []eth.Hash) *pbtransform.CombinedFilter {
	var addrBytes [][]byte
	var sigsBytes [][]byte

	for _, addr := range addrs {
		b := addr.Bytes()
		addrBytes = append(addrBytes, b)
	}

	for _, sig := range sigs {
		b := sig.Bytes()
		sigsBytes = append(sigsBytes, b)
	}

	x := &pbtransform.CallToFilter{
		Addresses:  addrBytes,
		Signatures: sigsBytes,
	}
	return &pbtransform.CombinedFilter{
		CallFilters: []*pbtransform.CallToFilter{x},
	}
}
