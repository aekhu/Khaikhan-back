package util

import (
	"flag"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	buildType := flag.String("env", "prod", "Environment type [dev,prod]")
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	flag.Parse()
	viper.AddConfigPath(path)
	if *buildType != "dev" && *buildType != "prod" {
		log.Fatalln("Please enter valid environment type.")
	}
	viper.SetConfigName(*buildType)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	spew.Dump(config)
	return config, nil
}
