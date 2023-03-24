package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfg Config

type Config struct {
	PlateausLotteryContractAddress       string    `mapstructure:"plateaus_lottery_contract_address"`
	Networks                             []Network `mapstructure:"networks"`
	Peer                                 string    `mapstructure:"peer"`
	PlateausPrivateKey                   string    `mapstructure:"plateaus_private_key"`
	PlateausRPC                          string    `mapstructure:"plateaus_rpc"`
	PlateausNodeValidatorContractAddress string    `mapstructure:"plateaus_node_validator_contract_address"`
	IPFSToken                            string    `mapstructure:"ipfs_token"`
}

type Network struct {
	Name                             string  `mapstructure:"name"`
	MinAmount                        float64 `mapstructure:"min_amount"`
	RPC                              string  `mapstructure:"rpc"`
	LotteryValidationContractAddress string  `mapstructure:"lottery_validation_contract_address"`
	PrivateKey                       string  `mapstructure:"private_key"`
}

func (n Network) GetLotteryValidationContractAddress() string {
	return n.LotteryValidationContractAddress
}

func (n Network) GetPrivateKey() string {
	return n.PrivateKey
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
