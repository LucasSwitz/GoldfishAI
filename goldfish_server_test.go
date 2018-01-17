package goldfish

import (
	"goldfish/proto"
	"sync"
	"testing"
	"time"
)

func TestUDPServerControl(t *testing.T) {
	server, err := MakeUDPServer(4443)
	var waitGroup sync.WaitGroup
	if err != nil {
		t.Error(err)
		return
	}

	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		err = server.Run()

		if err != nil {
			t.Error(err)
			return
		}
	}()

	c := make(chan struct{})

	server.Stop()

	go func() {
		defer close(c)
		waitGroup.Wait()
	}()

	SigOrTimeout(t, &c, time.Millisecond*50)
}

func TestGoldfishClientMake(t *testing.T) {
	_, err := MakeClient("127.0.0.1", 4443)

	if err != nil {
		t.Error(err)
	}
}

func TestGoldfishClientConnect(t *testing.T) {
	client, _ := MakeClient("127.0.0.1", 4443)
	err := client.Dial()

	if err != nil {
		t.Error(err)
		return
	}
}

func TestUDPPing(t *testing.T) {
	var waitGroup sync.WaitGroup
	server, _ := MakeUDPServer(4443)
	client, _ := MakeUDPClient("127.0.0.1", 4443)
	go server.Run()

	err := client.Dial()

	if err != nil {
		t.Error(err)
		return
	}

	msg := []byte{30}

	err = client.Write(msg)

	if err != nil {
		t.Error(err)
		return
	}

	c := make(chan struct{})

	waitGroup.Add(1)
	go func() {
		<-server.ReadChannel()
		waitGroup.Done()
	}()

	go func() {
		defer close(c)
		waitGroup.Wait()
	}()

	SigOrTimeout(t, &c, time.Millisecond*10)
}

func TestGoldFishPing(t *testing.T) {
	var waitGroup sync.WaitGroup
	server, _ := MakeServer(4444)
	client, _ := MakeClient("127.0.0.1", 4444)

	go server.Run()

	err := client.Dial()

	if err != nil {
		t.Error(err)
		return
	}

	msg := goldfish_message.GoldFishMessage{goldfish_message.GoldFishMessage_MANAGMENT, nil, nil, nil}

	err = client.Send(msg)

	if err != nil {
		t.Error(err)
		return
	}

	c := make(chan struct{})

	waitGroup.Add(1)
	go func() {
		msg := <-server.MessageChannel()

		if msg.Type != goldfish_message.GoldFishMessage_MANAGMENT {
			t.Error("Corrupt Message.")
		}

		waitGroup.Done()
	}()

	go func() {
		defer close(c)
		waitGroup.Wait()
	}()

	SigOrTimeout(t, &c, time.Millisecond*10)

}

func SigOrTimeout(t *testing.T, sig *chan struct{}, timeout time.Duration) {
	select {
	case <-*sig:
		return
	case <-time.After(timeout):
		t.Error("Timeout")
		return
	}
}
