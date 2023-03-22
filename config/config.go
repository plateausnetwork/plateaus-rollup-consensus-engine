package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfg Config

type Config struct {
	PlateausLotteryContractAddress          string    `mapstructure:"plateaus_lottery_contract_address"`
	PolygonLotteryValidationContractAddress string    `mapstructure:"polygon_lottery_validation_contract_address"`
	ScrollLotteryValidationContractAddress  string    `mapstructure:"scroll_lottery_validation_contract_address"`
	ZKSynckLotteryValidationContractAddress string    `mapstructure:"zksync_lottery_validation_contract_address"`
	Networks                                []Network `mapstructure:"networks"`
	Peer                                    string    `mapstructure:"peer"`
	PlateausPrivateKey                      string    `mapstructure:"plateaus_private_key"`
	PlateausRPC                             string    `mapstructure:"plateaus_rpc"`
	PlateausNodeValidatorContractAddress    string    `mapstructure:"plateaus_node_validator_contract_address"`
	PolygonPrivateKey                       string    `mapstructure:"polygon_private_key"`
	PolygonRPC                              string    `mapstructure:"polygon_rpc"`
	ScrollPrivateKey                        string    `mapstructure:"scroll_private_key"`
	ScrollRPC                               string    `mapstructure:"scroll_rpc"`
	ZKSyncPrivateKey                        string    `mapstructure:"zksync_private_key"`
	ZKSyncRPC                               string    `mapstructure:"zksync_rpc"`
	IPFSToken                               string    `mapstructure:"ipfs_token"`
}

type Network struct {
	Name      string  `mapstructure:"name"`
	MinAmount float64 `mapstructure:"min_amount"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("some error while reading config file: %s", err)
	}

	cfg = Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("some error while unmarshal config: %s", err)
	}
}

func GetConfig() *Config {
	return &cfg
}
