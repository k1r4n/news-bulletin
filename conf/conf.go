package conf

import (
	"log"
	"os"

	"github.com/hifx/envconfig"
)

// Vars is the object struct of config variables
type Vars struct {
	APIKey            string `envconfig:"APIKEY"`
	Index             string `envconfig:"INDEX"`
	ChannelTypeTop    string `envconfig:"CHANNEL_TYPE_TOP"`
	ChannelTypeLatest string `envconfig:"CHANNEL_TYPE_LATEST"`
	ChannelEndPoint   string `envconfig:"CHANNEL_ENDPOINT"`
	ArticleTypeTop    string `envconfig:"ARTICLE_TYPE_TOP"`
	ArticleTypeLatest string `envconfig:"ARTICLE_TYPE_LATEST"`
	ArticleEndPoint   string `envconfig:"ARTICLE_ENDPOINT"`
	ErrorLog          string `envconfig:"ERROR_LOG"`
}

// Read is used to create and object of the config values
func Read() *Vars {
	vars := &Vars{}
	if err := envconfig.InitWithPrefix(vars, "NBC"); err != nil {
		log.Printf("error reading configuration: %s\n", err)
		os.Exit(1)
	}
	return vars
}
