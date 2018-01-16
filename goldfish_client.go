package goldfish

import (
	"fmt"
	"net"
	"strconv"
)

type Client struct {
	port        int64
	addrS       string
	addr        *net.UDPAddr
	conn        *net.UDPConn
	readChan    chan []byte
	controlChan chan int
}

func MakeClient(ip string, port int64) (c Client, err error) {
	conAddr := ip + ":" + strconv.FormatInt(port, 10)
	c.addr, err = net.ResolveUDPAddr("udp", conAddr)
	c.readChan = make(chan []byte)
	return c, err
}

func (c Client) Send(msg []byte, size int64) error {
	n, err := c.conn.Write(msg)
	fmt.Printf("Client sent %d bytes", n)
	return err
}

func (c *Client) Connect() (err error) {
	c.conn, err = net.DialUDP("udp", nil, c.addr)
	go c.Listen()
	return err
}

func (c Client) Read() []byte {
	return <-c.readChan
}

func (c Client) Listen() error {
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
