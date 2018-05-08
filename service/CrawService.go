package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/ndphu/gm-api-golang/config"
	"io/ioutil"
	"net/http"
	"time"
)

type CrawRequest struct {
	ResponseTo string        `json:"responseTo"`
	Items      []RequestItem `json:"items"`
}

type RequestItem struct {
	Id       string `json:"id"`
	Input    string `json:"input"`
	Result   string `json:"result"`
	SubTitle string `json:"subTitle"`
	Error    string `json:"error"`
}

func CrawVideoSource(playUrl string) (videoSource string, srt string, err error) {
	if config.Get().UseMQTT == true {
		return CrawVideoSourceMQTT(playUrl)
	} else {
		return CrawServiceHttp(playUrl)
	}
}

func CrawVideoSourceMQTT(playUrl string) (videoSource string, srt string, err error) {
	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
	req := CrawRequest{
		ResponseTo: reqId,
		Items: []RequestItem{{
			Id:    reqId,
			Input: playUrl,
		}},
	}
	var res []RequestItem
	syncChan := make(chan int, 0)

	jsonBuffer, err := json.Marshal(req)
	if err != nil {
		return "", "", err
	}

	fmt.Println("will publish message " + string(jsonBuffer))

	fmt.Println("connecting to mqtt broker")
	opts := MQTT.NewClientOptions().AddBroker(config.Get().MQTTBroker)
	opts.SetClientID("gm-api-golang-client-" + reqId)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, message MQTT.Message) {
		fmt.Println("message received from " + message.Topic())
		fmt.Println("payload = " + string(message.Payload()))
		err = json.Unmarshal(message.Payload(), &res)
		syncChan <- 0
	})
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("mqtt connect error " + err.Error())
		return "", "", err
	}
	fmt.Println("mqtt connected")
	defer client.Disconnect(250)

	resChannel := "/topic/crawler/response/" + reqId
	if token := client.Subscribe(resChannel, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("mqtt subscribe error " + token.Error().Error())
	}
	fmt.Println("subcribed to reponse channel " + resChannel)

	reqTopic := "/topic/crawler/request"
	publishMessage(reqTopic, client, jsonBuffer)
	go func() {
		time.Sleep(2 * time.Minute)
		err = errors.New("timeout waiting for message on channel " + resChannel)
		syncChan <- 1
	}()

	<-syncChan

	if err != nil {
		fmt.Println("failed to craw video source via mqtt. err = " + err.Error())
		return "", "", err
	} else {
		return res[0].Result, res[0].SubTitle, nil
	}
}

func CrawServiceHttp(playUrl string) (videoSource string, srt string, err error) {
	reqId := fmt.Sprintf("%d", time.Now().UnixNano())
	items := []RequestItem{{
		Id:    reqId,
		Input: playUrl,
	}}
	jsonBuffer, err := json.Marshal(items)
	if err != nil {
		return "", "", err
	}
	crawServiceUrl := config.Get().CrawlerServiceBaseUrl + "/api/craw"
	fmt.Println("calling craw service at " + crawServiceUrl)
	fmt.Println("body = " + string(jsonBuffer))
	resp, err := http.Post(crawServiceUrl,
		"application/json",
		bytes.NewBuffer(jsonBuffer))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var respItems []RequestItem
	err = json.Unmarshal(body, &respItems)
	if err != nil {
		return "", "", err
	}

	return respItems[0].Result, respItems[0].SubTitle, nil
}

func publishMessage(reqTopic string, client MQTT.Client, jsonBuffer []byte) {
	fmt.Println("publising message to " + reqTopic)
	token := client.Publish(reqTopic, 0, false, string(jsonBuffer))
	token.Wait()
	fmt.Println("message published")
}
