package redis_go

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func resp() ([][]byte, error) {
	str := make([]byte, 4096)
	_, err := conn.Read(str)
	if err != nil {
		return make([][]byte, 0), err
	}
	str = bytes.Trim(str, "\x00") //移除null
	fmt.Printf("%q\n", str)
	firstChar := string(str[:1])
	fmt.Printf("%q\n", firstChar)
	strSlice := bytes.Split(str[1:], []byte("\r\n"))
	//fmt.Printf("%q\n", strSlice)
	//fmt.Printf("%s", len(strSlice))
	switch firstChar {
	case "+":
		return strSlice, nil
	case "$":
		return strSlice, nil
	case ":":
		return strSlice, nil
	case "-":
		return strSlice, errors.New(string(strSlice[0]))
	default:
		return strSlice, errors.New("First char not match.")
	}
	return strSlice, nil
}

func Info() (map[string]string, error) {
	command := "*1\r\n$4\r\nINFO\r\n"
	_, err := conn.Write([]byte(command))
	m := make(map[string]string)
	if err != nil {
		return m, err
	}
	msg, err := resp()
	for i := range msg {
		v := string(msg[i])
		if strings.Contains(v, ":") {
			s := strings.Split(v, ":")
			m[s[0]] = s[1]
		}
	}
	if err != nil {
		return m, err
	}
	return m, nil
}

func Ping() (string, error) {
	command := "*1\r\n$4\r\nPING\r\n"
	_, err := conn.Write([]byte(command))
	if err != nil {
		return "", err
	}
	msg, err := resp()
	if err != nil {
		return "", err
	}
	return string(msg[0]), nil
}

func Get(key string) (string, error) {
	command := "*2\r\n$3\r\nGET\r\n$" + strconv.Itoa(len(key)) + "\r\n" + key + "\r\n"
	_, err := conn.Write([]byte(command))
	if err != nil {
		return "", err
	}
	msg, err := resp()
	//fmt.Printf("%q\n", msg)
	if err != nil {
		return "", err
	}
	return string(msg[1]), nil
}

func Set(key string, value string) (string, error) {
	command := "*3\r\n$3\r\nSET\r\n$" + strconv.Itoa(len(key)) + "\r\n" + key + "\r\n$" + strconv.Itoa(len(value)) + "\r\n" + value + "\r\n"
	_, err := conn.Write([]byte(command))
	if err != nil {
		return "", err
	}
	msg, err := resp()
	//fmt.Printf("%q\n", msg)
	if err != nil {
		return "", err
	}
	return string(msg[0]), nil
}

func Incr(key string) (int, error) {
	command := "*2\r\n$4\r\nINCR\r\n$" + strconv.Itoa(len(key)) + "\r\n" + key + "\r\n"
	_, err := conn.Write([]byte(command))
	if err != nil {
		return 0, err
	}
	msg, err := resp()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(string(msg[0]))
	if err != nil {
		return 0, err
	}
	return num, nil
}
