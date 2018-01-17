package goldfish

import (
	"goldfish/proto"

	"github.com/golang/protobuf/proto"
)

type GoldFishServer struct {
	udpServer *Server
	msgChan   chan goldfish_message.GoldFishMessage
}

func MakeServer(port int64) (server *GoldFishServer, err error) {
	server = new(GoldFishServer)
	server.udpServer, err = MakeUDPServer(port)
	server.msgChan = make(chan goldfish_message.GoldFishMessage)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (server GoldFishServer) Run() (err error) {
	go server.udpServer.Run()

	for {
		select {
		case data := <-server.udpServer.ReadChannel():
			msg := &goldfish_message.GoldFishMessage{}
			err = proto.Unmarshal(data.data, msg)
			if err != nil {
				return err
			}
			server.msgChan <- *msg
		}
	}
}

func (server GoldFishServer) MessageChannel() chan goldfish_message.GoldFishMessage {
	return server.msgChan
}

func (server GoldFishServer) Stop() {
	server.udpServer.Stop()
}
