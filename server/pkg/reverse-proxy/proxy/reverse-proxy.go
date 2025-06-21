package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ReverseProxyService struct {
	// Amount to be initialized
	Amount int `json:"amount"`

	// Environment file path
	EnvironmentPath string `json:"environmentPath"`
}

type ReverseProxyConfig struct {
	Singss  *ReverseProxyService `json:"singss"`
	Singapi *ReverseProxyService `json:"singapi"`
}

func ParseReverseProxyConfig(path string) (*ReverseProxyConfig, error) {
	// Check if folder exists
	if file, err := os.Stat(path); err != nil {
		return nil, err
	} else if !file.IsDir() {
		return nil, errors.New("Invalid path: " + file.Name())
	}
	path = filepath.Join(path, "proxy.json")

	// Check if file is json
	file, err := os.Stat(path)
	if err != nil {
		return nil, err
	} else if !strings.HasSuffix(file.Name(), ".json") {
		return nil, errors.New("Invalid format: " + file.Name() + " must be a json file.")
	}

	var buff []byte = make([]byte, file.Size())

	jsonfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if _, err = jsonfile.Read(buff); err != nil {
		return nil, errors.New("Unable to read file: " + err.Error())
	}

	var config *ReverseProxyConfig = &ReverseProxyConfig{}

	if json.Unmarshal(buff, config) != nil {
		return nil, err
	}

	return config, nil
}

type ReverseProxy struct{}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{}
}

func (p *ReverseProxy) StartWith(c *ReverseProxyConfig) {
	fmt.Println("reverse-proxy.go | StartWith func")
	fmt.Println(c.Singss.Amount, c.Singapi.EnvironmentPath)
	// if ErrorPortAlreadyInUse then Try again in a upper port
}
