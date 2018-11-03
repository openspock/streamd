// Package conf provides interfaces and functions to access openspock
// config.
package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Base is the base config for streamd
var Base Config

func init() {
	if _, err := toml.DecodeFile("/etc/streamd.toml", &Base); err != nil {
		log.Fatal(err)
	}
	log.Println(Base)
}

// Streamd is the hostname and port config
type Streamd struct {
	Port int64
}

// Config is the base config for streamd
type Config struct {
	Streamd Streamd
}
