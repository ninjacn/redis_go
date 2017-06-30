package redis_go

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	fmt.Println("redis_go")
}

var conn *net.TCPConn

func Conn(hostname string, port string) (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp4", hostname+":"+port)
	if err != nil {
		return nil, err
	}
	c, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		return nil, errors.New("connection refused")
	}
	//defer c.Close()
	conn = c
	return c, nil
}
