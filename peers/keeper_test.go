package peers

import (
	"testing"
	"time"
)

var Conn_Keeper *ConnKeeper

func init() {
	Conn_Keeper = NewConnKeeper()
	Conn_Keeper.Set("jack", nil)
}

func TestING(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Conn_Keeper.Get("jack")
	}
	go Conn_Keeper.Set("tom", nil)
	for i := 0; i < 10; i++ {
		go Conn_Keeper.Get("tom")
	}
	time.Sleep(2e9)

}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conn_Keeper.Get("jack")
	}
}

func BenchmarkNoLockerGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conn_Keeper.NoLockerGet("jack")
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conn_Keeper.Set("jack", nil)
	}
}

func BenchmarkNoLockerSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conn_Keeper.NoLockerSet("jack", nil)
	}
}
