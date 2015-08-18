package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/shaalx/oschat/msg"
	"github.com/shaalx/oschat/peers"
	"net"
	"sync"
)

func main() {
	newServer().Serve()
}

type Server struct {
	keeper *peers.ConnKeeper
	ip     string
	port   int
}

func newServer() *Server {
	return &Server{
		keeper: peers.NewConnKeeper(),
		ip:     "",
		port:   8080,
	}
}

func (s *Server) Serve() {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(s.ip), s.port, ""})
	if err != nil {
		panic(err)
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	for {
		conn, err := listener.AcceptTCP()
		if checkerr(err) {
			continue
		}
		go s.Accept(conn)
	}
}
func (s *Server) Accept(conn *net.TCPConn) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
	bs := make([]byte, 1024)
	var msg_tmp msg.OSMsg
	var login_user string
	once := sync.Once{}
	for {
		n, err := conn.Read(bs)
		if checkerr(err) {
			break
		}
		err = proto.Unmarshal(bs[:n], &msg_tmp)
		if checkerr(err) {
			continue
		}
		once.Do(func() {
			login_user = msg_tmp.GetFromu()
			err := s.keeper.Login(login_user, conn)
			if checkerr(err) {
				panic(err)
			}
		})
		err = s.keeper.SendMsgTo(msg_tmp.GetTou(), bs[:n])
		if checkerr(err) {
			continue
		}
		msg_tmp.Reset()
	}
	defer func() {
		s.keeper.Logout(login_user)
	}()

}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
