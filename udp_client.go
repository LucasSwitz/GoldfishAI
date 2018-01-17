package goldfish

import (
	"fmt"
	"net"
	"strconv"
)

type UDPClient struct {
	port        int64
	addrS       string
	addr        *net.UDPAddr
	conn        *net.UDPConn
	readChan    chan []byte
	controlChan chan int
}

func MakeUDPClient(ip string, port int64) (c *UDPClient, err error) {
	c = new(UDPClient)
	conAddr := ip + ":" + strconv.FormatInt(port, 10)
	c.addr, err = net.ResolveUDPAddr("udp", conAddr)
	c.readChan = make(chan []byte)
	return c, err
}

func (c UDPClient) Write(msg []byte) error {
	_, err := c.conn.Write(msg)
	return err
}

func (c *UDPClient) Dial() (err error) {
	c.conn, err = net.DialUDP("udp", nil, c.addr)
	go c.listen()
	return err
}

func (c UDPClient) ReadChannel() chan []byte {
	return c.readChan
}

func (c UDPClient) listen() error {
	fmt.Printf("Starting client listening...")
	for {
		buffer := make([]byte, 1024)
		n, _, err := c.conn.ReadFromUDP(buffer)
		if err != nil {
			return err
		}

		if n == 0 {
			fmt.Printf("No bytes read...")
			return nil
		}

		c.readChan <- buffer
	}
}
