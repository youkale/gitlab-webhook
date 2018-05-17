package config

import (
	"testing"
	"path/filepath"
	"fmt"
)

var c = Config{
	Path:        "/webhook",
	Address:     "0.0.0.0:8964",
	SecretToken: "123",
}

func TestLoadFile(t *testing.T) {
	srcPath, _ := filepath.Abs("")
	defaultConfigFile := filepath.Join(srcPath, "config.json")
	config := loadConfig(defaultConfigFile)
	fmt.Printf("%v", config)
}
