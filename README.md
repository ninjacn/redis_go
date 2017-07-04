# redis_go

```
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
```
