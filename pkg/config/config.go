package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	RestfulAddr string `json:"restful_addr" yaml:"RestfulAddr"`
	GRPCAddr    string `json:"grpc_addr" yaml:"GRPCAddr"`
	Debug       bool   `json:"debug" yaml:"Debug"`

	DBPath string `json:"db_path" yaml:"DBPath"` // https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.db
}

var BaseConfig = initBase()

func initBase() *config {
	file, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		if err := ioutil.WriteFile("./configs/config.yaml", []byte(cfp), 00666); err != nil {
			log.Fatalln(err)
		} else {
			log.Fatalln("Please fill out the profile")
		}
	}
	var cfg config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalln(err)
	}
	return &cfg
}

const cfp = `
RestfulAddr: "0.0.0.0:8086"
GRPCAddr: "0.0.0.0:8087"
Debug: true
DBPath: "./ip2region.db" # https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.db
`
