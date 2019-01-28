package config

import (
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ConfigDir string

const ServiceConfigFile = "service_info.yml"

var ServiceConf ServiceConfig

type ServiceConfig struct {
	ServiceName string `yaml:"ServiceName"`
	ServicePort string `yaml:"ServicePort"`
}

func Init() {
	ConfigDir = os.Getenv("CONF_DIR")
	loadServiceConfig()
}

func loadServiceConfig() {
	configFile, err := ioutil.ReadFile(filepath.Join(ConfigDir, ServiceConfigFile))
	if err != nil {
		grpclog.Fatalf("ReadFile, err= %v", err)
		return
	}
	yaml.Unmarshal(configFile, &ServiceConf)
	grpclog.Infof("Conf=%v", ServiceConf)
}
