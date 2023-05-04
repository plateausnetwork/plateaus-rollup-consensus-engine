package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
)

// homeDir the default value is the current PATH project
var homeDir string = "."
var cfg Config

type Config struct {
	PlateausLotteryContractAddress       string    `mapstructure:"plateaus_lottery_contract_address" validate:"required"`
	Networks                             []Network `mapstructure:"networks" validate:"dive"`
	Peer                                 string    `mapstructure:"peer" validate:"required"`
	PlateausPrivateKey                   string    `mapstructure:"plateaus_private_key" validate:"required"`
	PlateausRPC                          string    `mapstructure:"plateaus_rpc" validate:"required,url"`
	PlateausRest                         string    `mapstructure:"plateaus_rest" validate:"required,url"`
	PlateausNodeValidatorContractAddress string    `mapstructure:"plateaus_node_validator_contract_address" validate:"required"`
	IPFSToken                            string    `mapstructure:"ipfs_token" validate:"required"`
}

type Network struct {
	Name                             string  `mapstructure:"name" validate:"required"`
	MinAmount                        float64 `mapstructure:"min_amount" validate:"required,gte=0.1"`
	RPC                              string  `mapstructure:"rpc" validate:"required,url"`
	LotteryValidationContractAddress string  `mapstructure:"lottery_validation_contract_address" validate:"required"`
	PrivateKey                       string  `mapstructure:"private_key" validate:"required"`
}

func (n Network) GetLotteryValidationContractAddress() string {
	return n.LotteryValidationContractAddress
}

func (n Network) GetPrivateKey() string {
	return n.PrivateKey
}

func checkMinAmount(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	return value > 0.1
}

func init() {
	pathFile := GetAbsolutePath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath(pathFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("some error while reading config file: %s", err)
	}

	cfg = Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("some error while unmarshal config: %s", err)
	}

	val := validator.New()
	val.RegisterValidation("checkMinAmount", checkMinAmount)

	if err := val.Struct(cfg); err != nil {
		log.Panicf("error to validate config.yml: %s", err)
	}
}

func GetConfig() *Config {
	return &cfg
}

// GetAbsolutePath method to return the absolute path based on homeDir variable
func GetAbsolutePath(file string) string {
	return fmt.Sprintf("%s/%s", homeDir, file)
}
