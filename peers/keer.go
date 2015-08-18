package peers

import (
	"fmt"
	"net"
	"sync"
)

func NewConnKeeper() *ConnKeeper {
	return &ConnKeeper{
		Conns:   make(map[string]*net.TCPConn),
		RWMutex: &sync.RWMutex{},
	}
}

type ConnKeeper struct {
	Conns map[string]*net.TCPConn
	*sync.RWMutex
}

func (s *ConnKeeper) Login(username string, conn *net.TCPConn) error {
	go s.Set(username, conn)
	//
	return nil
}

func (s *ConnKeeper) Logout(username string) error {
	s.Delete(username)
	//
	return nil
}

func (s *ConnKeeper) SendMsgTo(targetu string, content []byte) error {
	s.Lock()
	defer s.Unlock()
	targetconn, err := s.NoLockerGet(targetu)
	if err != nil {
		return err
	}
	_, err = targetconn.Write(content)
	return err
}

func (s *ConnKeeper) Get(username string) (*net.TCPConn, error) {
	s.RLock()
	defer s.RUnlock()
	conn, ok := s.Conns[username]
	if ok {
		return conn, nil
	}
	return nil, fmt.Errorf("%s is lost.", username)
}

func (s *ConnKeeper) Set(username string, conn *net.TCPConn) {
	s.Lock()
	defer s.Unlock()
	s.Conns[username] = conn
}

func (s *ConnKeeper) Delete(username string) {
	s.Lock()
	defer s.Unlock()
	delete(s.Conns, username)
}

func (s *ConnKeeper) NoLockerGet(username string) (*net.TCPConn, error) {
	conn, ok := s.Conns[username]
	if ok {
		return conn, nil
	}
	return nil, fmt.Errorf("%s is lost.", username)
}

func (s *ConnKeeper) NoLockerSet(username string, conn *net.TCPConn) {
	s.Conns[username] = conn
}
