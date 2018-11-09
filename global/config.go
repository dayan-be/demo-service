package global

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Configs struct {
	Srv struct {
		SrvName string `yaml:"srvName"`
		SrvId   uint32 `yaml:"srvId"`
		Addr    string `yaml:"addr"`
		Version string `yaml:"version"`
	}

	Log struct {
		LogLevel    int  `yaml:"logLevel"`
		LogFileSize int  `yaml:"logFileSize"`
		JsonFile    bool `yaml:"jsonFile"`
	}
}

var cfgIns *Configs

func Config() *Configs {
	return cfgIns
}

func LoadConfig(filename string){

	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		goto FAILED
	}

	err = yaml.Unmarshal(buff, cfgIns)
	if err != nil {
		goto FAILED
	}

FAILED:
	fmt.Printf("failed:%v",err)
	os.Exit(1)
}

func init() {
	cfgIns = &Configs{}
}