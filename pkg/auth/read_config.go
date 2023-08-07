package auth

import "github.com/mikerybka/webmachine/pkg/util"

func ReadConfig() (config *Config) {
	util.ReadJSON("/data/auth/config.json", config)
	return config
}
