package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConf() {
	dir, _ := os.Getwd()
	confPath := dir + "/conf/"
	fmt.Println(dir)
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath(confPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Conf Err:%s", err.Error()))
	}
}
