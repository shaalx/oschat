package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	. "github.com/toukii/oschat/msg"
	"testing"

	"reflect"
	"unsafe"
)

func checkerr(err error) bool {
	if nil != err {
		fmt.Println("error:", err)
		return true
	}
	return false
}
func t1() {
	msg := &OSMsg{Fromu: proto.String("jack"), Tou: proto.String("tom"), Content: proto.String("first")}
	mset := &proto.MessageSet{}
	proto.RegisterMessageSetType(msg, 1, "osmsg")
	b, err := proto.Marshal(msg)
	if checkerr(err) {
		return
	}
	// m := make(map[int32]proto.Extension)
	has := mset.Has(msg)
	fmt.Println(has)
	err = proto.Unmarshal(b, mset)
	if checkerr(err) {
		return
	}
	fmt.Println(mset)
}

// func TestT1(t *testing.T) {
// 	t1()
// }

func t2() {
	msg := &OSMsg{Fromu: proto.String("jack"), Tou: proto.String("tom"), Content: proto.String("first")}
	fmt.Println(msg)
	b, err := proto.Marshal(msg)
	if checkerr(err) {
		return
	}
	var desc OneofDescriptorProto

	err = proto.Unmarshal(b, &desc)
	if checkerr(err) {
		return
	}
	fmt.Println(desc)
	fmt.Println(desc.GetName())
}

// func TestT2(t *testing.T) {
// 	t2()
// }

func t3() {
	msg := &OSMsg{Fromu: proto.String("jack"), Tou: proto.String("tom"), Content: proto.String("first")}
	fmt.Println(msg)
	b, err := proto.Marshal(msg)
	if checkerr(err) {
		return
	}
	typ := reflect.TypeOf(msg)
	v := reflect.New(typ)
	val := (*OSMsg)(unsafe.Pointer(v.Pointer()))
	err = proto.Unmarshal(b, val)
	if checkerr(err) {
		return
	}
	fmt.Println(val)
}

func TestT3(t *testing.T) {
	t3()
}
