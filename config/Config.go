package config

import (
	"encoding/json"
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

type VcapServices struct {
	Mlab []MongoDBConfig
}

type MongoDBCredential struct {
	Uri string
}

type MongoDBConfig struct {
	Name        string
	Credentials MongoDBCredential
}

const VCAPSERVICES = "VCAP_SERVICES"

var conf *Config

func init() {
	conf = &Config{}
	vcapServices := os.Getenv(VCAPSERVICES)
	if vcapServices == "" {
		conf.MongoDBUri = os.Getenv("MONGODB_URI")
	} else {
		conf.MongoDBUri = getMongoDBUri(vcapServices)
	}
	conf.DBName = getDBName(conf.MongoDBUri)

	conf.GinDebug = os.Getenv("GIN_DEBUG") == "true"

	conf.UseMQTT = os.Getenv("USE_MQTT") == "true"
	if conf.UseMQTT {
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

func getMongoDBUri(vcapServicesEnv string) string {
	vs := VcapServices{}
	err := json.Unmarshal([]byte(vcapServicesEnv), &vs)
	if err != nil {
		panic(err)
	}
	return vs.Mlab[0].Credentials.Uri
}

func getDBName(mongodbUri string) string {
	parsed, e := url.Parse(mongodbUri)
	if e != nil {
		panic(e)
	}
	return strings.Trim(parsed.Path, "/")
}
