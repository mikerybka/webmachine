package email

import "github.com/mikerybka/webmachine/pkg/util"

var config *Config

type Config struct {
	Sender         string
	SendgridAPIKey string
}

func init() {
	util.ReadJSON("/data/email/config.json", config)
}
