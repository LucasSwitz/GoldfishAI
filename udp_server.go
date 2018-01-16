package goldfish

import (
	"fmt"
	"net"
	"strconv"
)

const (
	STOP_MESSAGE = iota
)

type incommingData struct {
	addr *net.UDPAddr
	data []byte
}

type Server struct {
	port        int64
	addr        *net.UDPAddr
	running     bool
	controlChan chan int
	connection  *net.UDPConn
	readChan    chan incommingData
}

/*Make Makes new server and returns*/
func MakeUDPServer(port int64) (s *Server, err error) {
	s = new(Server)
	s.port = port
	ps := strconv.FormatInt(port, 10)
	s.addr, err = net.ResolveUDPAddr("udp", ":"+ps)
	if err != nil {
		return nil, err
	}
	s.controlChan = make(chan int)
	s.connection, err = net.ListenUDP("udp", s.addr)

	if err != nil {
		return nil, err
	}

	s.readChan = make(chan incommingData)
	fmt.Printf("Made server on port: %d \n", port)
	return s, err
}

func (server Server) Run() (err error) {
	go server.read()
	fmt.Printf("Starting server: %d \n", server.port)
	for {
		select {
		case sig := <-server.controlChan:
			fmt.Printf("Recieved signal: %d", sig)
			switch sig {
			case STOP_MESSAGE:
				fmt.Printf("Stopping server....")
				server.connection.Close()
				return nil
			default:
				fmt.Printf("Invalid Control Message: %d", sig)
			}
		}
	}
}

func (server Server) ReadChannel() chan incommingData {
	return server.readChan
}

func (server Server) Write(addr *net.UDPAddr, data []byte) (err error) {
	_, err = server.connection.WriteToUDP(data, addr)
	return err
}

func (server Server) read() (err error) {
	fmt.Printf("Starting read: %d \n", server.port)
	buffer := make([]byte, 1024)
	for {
		n, addr, err := server.connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Errorf("Server read error", err)
			return err
		}
		data := []byte(buffer[0:n])
		server.readChan <- incommingData{addr, data}
	}
}

func (server *Server) Stop() {
	server.controlChan <- STOP_MESSAGE
}
