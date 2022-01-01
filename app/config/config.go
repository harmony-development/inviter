package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/creasty/defaults"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/ztrue/tracerr"
)

type Config struct {
	Server struct {
		DefaultServer string `hcl:"DefaultServer,optional" default:"https://chat.harmonyapp.io"`
	} `hcl:"Server,block"`
}

// Load reads a config file (JSON format)
func Load() (*Config, error) {
	var config Config
	defaults.MustSet(&config)

	if _, err := os.Stat("config.hcl"); os.IsNotExist(err) {
		file := hclwrite.NewFile()
		gohcl.EncodeIntoBody(&config, file.Body())

		err = ioutil.WriteFile("config.hcl", file.Bytes(), 0755)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("A default configuration has been written to 'config.hcl'. Edit it as appropriate and then restart inviter.")
		os.Exit(0)
	}

	if err := tracerr.Wrap(hclsimple.DecodeFile("config.hcl", nil, &config)); err != nil {
		return nil, err
	}

	return &config, nil
}
