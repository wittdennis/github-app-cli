package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ApplicationId int64
	PrivateKey    []byte
}

func GetConfig() Config {
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}

	if !viper.IsSet("appId") {
		cobra.CheckErr(fmt.Errorf("missing required 'appId' in config"))
	}

	appId, err := strconv.ParseInt(viper.GetString("appId"), 10, 64)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("can't parse 'appId' into int64 id"))
	}

	if !viper.IsSet("privateKey") {
		cobra.CheckErr(fmt.Errorf("missing required 'privateKey' in config"))
	}
	privateKey := viper.GetString("privateKey")

	return Config{
		ApplicationId: appId,
		PrivateKey:    []byte(privateKey),
	}
}
