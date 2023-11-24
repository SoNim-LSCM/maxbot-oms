package mqtt

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	errorHandler "github.com/SoNim-LSCM/maxbot_oms/errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %+v", err)
}

func MqttSetup() {
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	broker := os.Getenv("MQTT_BROKER")
	clientId := "go_" + strconv.Itoa(rand.Intn(100))
	port := 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	fmt.Printf("tcp://%s:%d %s (%s) %s\n", broker, port, username, password, clientId)
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	token := client.Connect()
	// token.Wait()
	// token.WaitTimeout(5 * time.Second)
	if token.Error() != nil {
		fmt.Println(token.Error().Error())
		errorHandler.CheckError(token.Error(), "MQTT FAIL")
	}
	fmt.Printf("tcp://%s:%d %s (%s) %s\n", broker, port, username, password, clientId)
	sub(client)

}

func sub(client mqtt.Client) {
	// Subscribe to the LWT connection status
	topic := "/maxbot/SONIM_TEST1/status"
	token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println(string(msg.Payload()))
	})
	token.Wait()
	fmt.Printf("Subscribed to LWT %s\n", topic)
}
