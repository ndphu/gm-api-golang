package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	MongoDBUri            string
	DBName                string
	MQTTBroker            string
	CrawlerServiceBaseUrl string
	GinDebug              bool
	UseMQTT               bool
	UseCookie             bool
}

type MongoDBCredential struct {
	Uri string
}

type MongoDBConfig struct {
	Name        string
	Credentials MongoDBCredential
}

var conf *Config

func init() {
	conf = &Config{}
	conf.MongoDBUri = os.Getenv("MONGODB_URI")
	conf.DBName = getDBName(conf.MongoDBUri)

	conf.GinDebug = os.Getenv("GIN_DEBUG") == "true"

	conf.UseMQTT = os.Getenv("USE_MQTT") == "true"
	if conf.UseMQTT == true {
		broker := os.Getenv("MQTT_BROKER")
		if broker == "" {
			conf.MQTTBroker = "tcp://iot.eclipse.org:1883"
		} else {
			conf.MQTTBroker = broker
		}
	} else {
		conf.CrawlerServiceBaseUrl = os.Getenv("CRAW_SERVICE_BASE_URL")
		if conf.CrawlerServiceBaseUrl == "" {
			panic("should set CRAW_SERVICE_BASE_URL")
		}
		fmt.Println("using craw service " + conf.CrawlerServiceBaseUrl)
	}

	conf.UseCookie = os.Getenv("USE_COOKIE") == "true"
}

func Get() *Config {
	return conf
}

func getDBName(mongodbUri string) string {
	parsed, e := url.Parse(mongodbUri)
	if e != nil {
		panic(e)
	}
	return strings.Trim(parsed.Path, "/")
}
