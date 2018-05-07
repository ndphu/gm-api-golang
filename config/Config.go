package config

import (
	"os"
	"encoding/json"
	"net/url"
	"strings"
)

type Config struct {
	IsLocal bool
	MongoDBUri string
	DBName string
	MQTTBroker string
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
		c.IsLocal = true
		c.MongoDBUri = os.Getenv("MONGODB_URI")
	} else {
		c.IsLocal = false
		c.MongoDBUri = getMongoDBUri(vcapServices)
	}
	c.DBName = getDBName(c.MongoDBUri)

	broker:= os.Getenv("MQTT_BROKER")
	if broker == "" {
		c.MQTTBroker = "tcp://iot.eclipse.org:1883"
	} else {
		c.MQTTBroker = broker
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
