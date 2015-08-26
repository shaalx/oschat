package msg

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

var msg Msg

var osmsg OSMsg

type M interface {
	Reset()
	String() string
	ProtoMessage()
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println("error:", err)
		return true
	}
	return false
}
func TestMsg(t *testing.T) {
	recMsg := &OSMsg{Fromu: proto.String("jack"), Tou: proto.String("tom"), Content: proto.String("first")}
	b, err := proto.Marshal(recMsg)
	if checkerr(err) {
		return
	}

	buf := proto.NewBuffer(b)
	err = proto.Unmarshal(b, &msg)
	if checkerr(err) {
		return
	}
	fmt.Println(msg)
	fmt.Println(msg.String())

	err = proto.Unmarshal(buf.Bytes(), &osmsg)
	if checkerr(err) {
		return
	}

	fmt.Println(osmsg)
	fmt.Println(osmsg.String())

	a := []int{1, 2}
	a1 := a[:1]
	fmt.Println(a, a1)
	a1 = append(a1, 3)
	fmt.Println(a, a1)
	a1 = append(a1, 4)
	fmt.Println(a, a1)
	a1[1] = 10
	fmt.Println(a, a1)
}
