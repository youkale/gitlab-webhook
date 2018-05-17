package config

import (
	"path/filepath"
	"os"
	"flag"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"errors"
)

type Setting struct {
	Event       string   `json:"event"`
	BuildName   string   `json:"build_name"`
	BuildStage  string   `json:"build_stage"`
	BuildStatus string   `json:"build_status"`
	Command     []string `json:"command"`
	ProjectName string   `json:"project_name"`
	Ref         string   `json:"ref"`
}

type Config struct {
	Address     string `json:"address"`
	Path        string `json:"path"`
	SecretToken string `json:"secret_token"`
	Settings    []Setting
}


func GetConfig() *Config {
	defaultConfigFile := ""
	workingDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		defaultConfigFile = filepath.Join(workingDir, "config.json")
	}
	if !fileExists(defaultConfigFile) {
		srcPath, _ := filepath.Abs("")
		defaultConfigFile = filepath.Join(srcPath, "config.json")
	}
	var configFile string
	flag.StringVar(&configFile, "config", defaultConfigFile, "Config file for this Point server.")
	return loadConfig(configFile)

}

func loadConfig(configFile string) *Config {
	if fileExists(configFile) {
		if byt, err := ioutil.ReadFile(configFile); nil == err {
			conf := Config{}
			if e := json.Unmarshal(byt, &conf); nil == e {
				return &conf
			} else {
				panic(errors.New(fmt.Sprintf("config file unmarsha has error %v", configFile)))
			}
		} else {
			panic(errors.New(fmt.Sprintf("read config file has error %v", err.Error())))
		}
	} else {
		panic(errors.New(fmt.Sprintf("config file is do not exists %v", configFile)))
	}
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}
