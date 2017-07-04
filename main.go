package redis_go

import (
	"errors"
	"fmt"
	"net"
)

type NRedis struct {
	*net.TCPConn
}

func main() {
	fmt.Println("redis_go")
}

func Conn(hostname string, port string) (*NRedis, error) {
	raddr, err := net.ResolveTCPAddr("tcp4", hostname+":"+port)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		return nil, errors.New("connection refused")
	}
	//defer c.Close()
	return &NRedis{conn}, nil
}
