package main

import (
	"fmt"
	"github.com/ninjacn/redis_go"
	"log"
)

func main() {
	c, err := redis_go.Conn("127.0.0.1", "6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	//ping, err := redis_go.Info()
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Printf("%q", ping)

	//msg, err := redis_go.Set("name", "\"ming")
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(msg)

	//g, err := redis_go.Get("name")
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(g)

	g, err := redis_go.Incr("int")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)

}
