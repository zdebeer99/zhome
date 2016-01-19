package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Load() *Config {
	fpath, err := filepath.Abs("./config.yaml")
	if err != nil {
		log.Fatalln("Get 'config.yaml' path failed.", err)
	}
	log.Println("Loading config.", fpath)

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("Failed to load '%s' file. exciting! error %s", fpath, err)
	}

	var cf Config
	err = yaml.Unmarshal(data, &cf)
	if err != nil {
		log.Fatalf("Failed to parse config.yaml file. exciting! parse error %s", err)
	}
	if cf.Database == "" {
		cf.Database = "localhost/zhome"
	}
	if cf.BindAddress == "" {
		cf.BindAddress = ":3001"
	}
	return &cf
}

func Save(cf *Config) {

}
