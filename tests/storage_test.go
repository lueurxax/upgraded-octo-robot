package tests

import (
	"testing"

	"github.com/lueurxax/upgraded-octo-robot/pkg/channels"
	"github.com/lueurxax/upgraded-octo-robot/pkg/mutexes"
	"github.com/lueurxax/upgraded-octo-robot/pkg/sync_map"
)

const (
	keya  = "a"
	keyb  = "b"
	value = "123"
)

func Benchmark_SyncMap(b *testing.B) {
	db := sync_map.NewKv()
	// warmup shards
	db.Set(keya, value)
	db.Set(keyb, value)
	for i := 0; i < b.N; i++ {
		go db.Set(keya, value)
		go db.Set(keyb, value)
		db.Get(keyb)
	}
}

func Benchmark_Channels(b *testing.B) {
	db := channels.NewStorage()
	// warmup shards
	db.Set(keya, value)
	db.Set(keyb, value)
	for i := 0; i < b.N; i++ {
		go db.Set(keya, value)
		go db.Set(keyb, value)
		db.Get(keyb)
	}
}

func Benchmark_Mutexes(b *testing.B) {
	db := mutexes.NewStorage()
	// warmup shards
	db.Set(keya, value)
	db.Set(keyb, value)
	for i := 0; i < b.N; i++ {
		go db.Set(keya, value)
		go db.Set(keyb, value)
		db.Get(keyb)
	}
}
