package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/shaalx/oschat/msg"
	"net"
	"time"

	"os"
)

const (
	addr = "127.0.0.1:8080"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	Client(conn)
}

func Client(conn net.Conn) {
	go func() {
		for {
			receiveMsg(conn)
		}
	}()

	endsend := sendMsg(conn)
	select {
	case <-endsend:
		os.Exit(-1)
	}
}

func sendMsg(conn net.Conn) chan bool {
	// _msg := msg.OSMsg{Fromu: proto.String("jack"), Tou: proto.String("tom"), Content: proto.String("hello,tom")}
	_msg := msg.OSMsg{Fromu: proto.String("tom"), Tou: proto.String("jack"), Content: proto.String("hello,jack")}
	end := make(chan bool)
	for {
		// send
		data, err := proto.Marshal(&_msg)
		if checkerr(err) {
			end <- true
			break
		}
		conn.Write(data)

		proto.Unmarshal(data, &_msg)
		fmt.Printf("[send -----> message]%s\n", _msg.String())
		time.Sleep(4e9)
	}
	return end
}

func receiveMsg(conn net.Conn) {
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if checkerr(err) {
		os.Exit(-1)
		return
	}
	// fmt.Printf("\nread (%d) byte from %v :\n%v\n", n, conn.RemoteAddr(), buf[:n])

	protobuf := proto.NewBuffer(buf[:n])
	var _msg msg.OSMsg
	pumerr := protobuf.Unmarshal(&_msg)
	if checkerr(pumerr) {
		return
	}
	fmt.Printf("[received <----- message]%v\n", _msg.String())
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
