package goldfish

import (
	"fmt"
	"goldfish/proto"

	"github.com/golang/protobuf/proto"
)

type GoldFishClient struct {
	udpClient  *UDPClient
	msgChannel chan goldfish_message.GoldFishMessage
}

func MakeClient(ip string, port int64) (c *GoldFishClient, err error) {
	c = new(GoldFishClient)
	c.udpClient, err = MakeUDPClient(ip, port)
	c.msgChannel = make(chan goldfish_message.GoldFishMessage)
	return c, err
}

func (c GoldFishClient) Send(msg goldfish_message.GoldFishMessage) error {
	data, err := proto.Marshal(&msg)

	if err != nil {
		return err
	}

	err = c.udpClient.Write(data)
	return err
}

func (c GoldFishClient) Dial() (err error) {
	err = c.udpClient.Dial()

	if err != nil {
		return err
	}

	go c.listen()
	return nil
}

func (c GoldFishClient) MessageChannel() chan goldfish_message.GoldFishMessage {
	return c.msgChannel
}

func (c GoldFishClient) listen() error {
	for {
		select {
		case data := <-c.udpClient.readChan:
			msg := &goldfish_message.GoldFishMessage{}
			err := proto.Unmarshal(data, msg)
			if err != nil {
				fmt.Printf("%s \n", err.Error())
			} else {
				c.msgChannel <- *msg
			}

		}
	}
}
