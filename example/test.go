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

	//d, err := redis_go.Select(2)
	//if err != nil {
	//log.Fatal(err)
	//}

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

	age, err := redis_go.Incr("age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(age)

}
