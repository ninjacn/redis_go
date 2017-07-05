# redis_go
[![Build Status](https://travis-ci.org/ninjacn/redis_go.svg?branch=master)](https://travis-ci.org/ninjacn/redis_go)
<pre><code>
package main

import (
	"fmt"
	"github.com/ninjacn/redis_go"
	"log"
)

func main() {
	redis_go, err := redis_go.Conn("127.0.0.1", "6379")
	if err != nil {
		log.Fatal(err)
	}
	defer redis_go.Close()

	age, err := redis_go.Incr("age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(age)
    }
</code></pre>
