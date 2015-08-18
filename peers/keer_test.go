package peers

import (
	"testing"
)

var Conn_Keeper *ConnKeeper

func init() {
	Conn_Keeper = NewConnKeeper()
	Conn_Keeper.Set("jack", nil)
}

func TestING(t *testing.T) {
	conn, err := Conn_Keeper.Get("jack")
	t.Log(conn, err)
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
