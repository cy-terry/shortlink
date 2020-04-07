package main

import (
	"ShortLink/config"
	"ShortLink/schema"
	"ShortLink/service"
	"ShortLink/util"
	"ShortLink/web"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "./config/app.yml", "yml config path")
	flag.Parse()
	data, _ := ioutil.ReadFile(confPath)
	if data == nil {
		log.Fatalln("Configuration does not exist")
	}
	err := yaml.Unmarshal(data, &config.Conf)
	if err != nil {
		log.Fatal(err)
	}

	db := schema.InitDB(&config.Conf.Serve.Db)
	s := &service.Service{
		Converter: &util.HashURLConverter{},
		UniqueID:  &util.UUIDUniqueID{},
		DB:        db,
	}
	_ = web.RunServer(&config.Conf.Serve.Web, s, util.UnAuth{})
}
