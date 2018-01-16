package goldfish

import (
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
	err := client.Connect()

	if err != nil {
		t.Error(err)
		return
	}
}

func TestGoldfishPing(t *testing.T) {
	var waitGroup sync.WaitGroup
	server, _ := MakeUDPServer(4443)
	client, _ := MakeClient("127.0.0.1", 4443)
	go server.Run()

	err := client.Connect()

	if err != nil {
		t.Error(err)
		return
	}

	msg := []byte{30}

	err = client.Send(msg, 1)

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

func SigOrTimeout(t *testing.T, sig *chan struct{}, timeout time.Duration) {
	select {
	case <-*sig:
		return
	case <-time.After(timeout):
		t.Error("Timeout")
		return
	}
}
