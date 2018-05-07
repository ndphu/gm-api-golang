package config

import (
	"os"
	"encoding/json"
	"net/url"
	"strings"
	"fmt"
)

type Config struct {
	MongoDBUri            string
	DBName                string
	MQTTBroker            string
	CrawlerServiceBaseUrl string
	GinDebug              bool
	UseMQTT bool
}

type VcapServices struct {
	Mlab []MongoDBConfig
}

type MongoDBCredential struct {
	Uri string
}

type MongoDBConfig struct {
	Name string
	Credentials MongoDBCredential
}

const VCAPSERVICES = "VCAP_SERVICES"
var conf *Config

func init() {
	c := Config{}
	vcapServices := os.Getenv(VCAPSERVICES)
	if vcapServices == "" {
		c.MongoDBUri = os.Getenv("MONGODB_URI")
	} else {
		c.MongoDBUri = getMongoDBUri(vcapServices)
	}
	c.DBName = getDBName(c.MongoDBUri)

	ginDebug := os.Getenv("GIN_DEBUG")
	if ginDebug == "true" {
		c.GinDebug = true
	} else {
		c.GinDebug = false
	}

	useMQTT := os.Getenv("USE_MQTT")
	if useMQTT == "true" {
		c.UseMQTT = true
		broker:= os.Getenv("MQTT_BROKER")
		if broker == "" {
			c.MQTTBroker = "tcp://iot.eclipse.org:1883"
		} else {
			c.MQTTBroker = broker
		}
	} else {
		c.UseMQTT = false
		c.CrawlerServiceBaseUrl = os.Getenv("CRAW_SERVICE_BASE_URL")
		if c.CrawlerServiceBaseUrl == "" {
			panic("should set CRAW_SERVICE_BASE_URL")
		}
		fmt.Println("using craw service " + c.CrawlerServiceBaseUrl)
	}
	conf = &c
}

func Get() (*Config) {
	return conf
}

func getMongoDBUri(vcapServicesEnv string) string {
	vs := VcapServices{}
	err:=json.Unmarshal([]byte(vcapServicesEnv), &vs)
	if err!=nil {
		panic(err)
	}
	return vs.Mlab[0].Credentials.Uri
}

func getDBName(mongodbUri string) string {
	parsed, e := url.Parse(mongodbUri)
	if e != nil {
		panic(e)
	}
	return strings.Trim(parsed.Path, "/");
}
