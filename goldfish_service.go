package goldfish

import (
	"goldfish/proto"
)

type GoldFishService struct {
	server *GoldFishServer
}

func MakeService(port int64) (service *GoldFishService, err error) {
	service = new(GoldFishService)
	service.server, err = MakeServer(port)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (service GoldFishService) Run() {
	go service.server.Run()

	for {
		select {
		case msg := <-service.server.MessageChannel():
			handleMsg(msg)
		}
	}
}

func handleMsg(msg goldfish_message.GoldFishMessage) {

}
