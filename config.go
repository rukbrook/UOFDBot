package main

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"log"
	"time"
)

const ConfigFile string = `./config.yml`
const GameInterval = int64(24*time.Hour) / int64(time.Millisecond)
const ConnectionTypeChannel string = `CHANNEL`
const ConnectionTypeWebhook string = `WEBHOOK`

var config struct {
	BotName         string `yaml:"bot_name"`
	BotToken        string `yaml:"bot_token"`
	BotTimeout      int    `yaml:"bot_timeout"`
	ConnectionType  string `yaml:"connection_type"`
	DbPath          string `yaml:"db_path"`
	WebhookHost     string `yaml:"webhook_host"`
	WebhookPort     string `yaml:"webhook_port"`
	WebhookCertPath string `yaml:"webhook_cert_path"`
	WebhookKeyPath  string `yaml:"webhook_key_path"`
	ServerAddress   string `yaml:"server_adress"`
}

func LoadConfig() {
	file, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(`Config successfully loaded!`)
}