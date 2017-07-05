package redis_go

import (
	"log"
	"testing"
)

func TestCommandIncr(t *testing.T) {
	redis_go, err := Conn("127.0.0.1", "6379")
	if err != nil {
		log.Fatal(err)
	}
	defer redis_go.Close()

	age, err := redis_go.Incr("age")
	if err != nil {
		log.Fatal(err)
	}
	if age <= 0 {
		t.Fail()
	}
}
