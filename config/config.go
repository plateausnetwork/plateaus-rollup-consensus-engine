package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfg Config

type Config struct {
	LotteryContractAddress            string   `mapstructure:"lottery_contract_address"`
	LotteryValidationContractAddress  string   `mapstructure:"lottery_validation_contract_address"`
	Networks                          []string `mapstructure:"networks"`
	Peer                              string   `mapstructure:"peer"`
	PlateausPrivateKey                string   `mapstructure:"plateaus_private_key"`
	PlateausRPC                       string   `mapstructure:"plateaus_rpc"`
	PlateausValidationContractAddress string   `mapstructure:"plateaus_validation_contract_address"`
	PolygonPrivateKey                 string   `mapstructure:"polygon_private_key"`
	PolygonRPC                        string   `mapstructure:"polygon_rpc"`
	IPFSToken                         string   `mapstructure:"ipfs_token"`
	MinAmount                         float64  `mapstructure:"min_amount,string"`
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
