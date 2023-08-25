package config

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLoadConfig(t *testing.T) {
	var (
		basePath = "/tmp/vtun"
		jsonPath = path.Join(basePath, "config.json")
		yamlPath = path.Join(basePath, "config.yaml")
	)

	_ = os.MkdirAll(basePath, os.ModePerm)
	defer os.RemoveAll(basePath)

	// write test config
	cfg := Config{
		DeviceName: "eth0",
		ServerMode: true,
	}
	jsonOut, err := json.Marshal(cfg)
	assert.Nil(t, err)
	os.WriteFile(jsonPath, jsonOut, os.ModePerm)

	yamlOut, err := yaml.Marshal(cfg)
	assert.Nil(t, err)
	os.WriteFile(yamlPath, yamlOut, os.ModePerm)

	// test yaml
	ncfg := Config{}
	err = ncfg.LoadConfig(yamlPath)
	assert.Nil(t, err)
	assert.Equal(t, cfg.DeviceName, ncfg.DeviceName)
	assert.Equal(t, true, ncfg.ServerMode)

	// test json
	ncfg = Config{}
	err = ncfg.LoadConfig(jsonPath)
	assert.Nil(t, err)
	assert.Equal(t, cfg.DeviceName, ncfg.DeviceName)
	assert.Equal(t, true, ncfg.ServerMode)
}
