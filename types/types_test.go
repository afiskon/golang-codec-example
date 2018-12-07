package types

import (
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
//  . "github.com/afiskon/golang-codec-example/types"
	"log"
	"testing"
)

type TypeVer1 struct {
	Field1 int
	Field2 string
	Field3 int
}

type TypeVer2 struct {
	Field3 int
	Field4 string
	Field5 bool
}

func implementsSelferInterface(obj codec.Selfer) bool {
	return true
}

func encodeDecode() {
	var (
		cborHandle codec.CborHandle
		err error
	)

	//v1 := Hero{ "Alex", 123, 456, &WariorInfo{ BOW, 10 }, nil}
	v1 := Hero{ "Bob", 234, 567, nil, &MageInfo{ []Spell{FIREBALL, THUNDERBOLT}, 42 }}

	var bs []byte
	enc := codec.NewEncoderBytes(&bs, &cborHandle)
	err = enc.Encode(v1)
	if err != nil {
		log.Fatalf("enc.Encode() failed, err = %v", err)
	}
	// log.Printf("bs = %X, len(bs) = %d, cap(bs) = %d", bs, len(bs), cap(bs))

	// Decode bs to v2

	var v2 Hero
	dec := codec.NewDecoderBytes(bs, &cborHandle)
	err = dec.Decode(&v2)
	if err != nil {
		log.Fatalf("dec.Decode() failed, err = %v", err)
	}
}

// make sure user didn't forget to run `go generate ./...` according to README.md
func TestSerialization(t *testing.T) {
	hero := Hero{ "Alex", 123, 456, &WariorInfo{ BOW, 10 }, nil}
	res := implementsSelferInterface(&hero)
	if !res {
		t.FailNow()
	}
}

func TestMigration(t *testing.T) {
	var (
		cborHandle codec.CborHandle
		err error
	)

	v1 := TypeVer1{ 123, "Field2", 456}
	var bs []byte
	enc := codec.NewEncoderBytes(&bs, &cborHandle)
	err = enc.Encode(v1)
	if err != nil {
		log.Fatalf("enc.Encode() failed, err = %v", err)
	}

	var v2 TypeVer2
	dec := codec.NewDecoderBytes(bs, &cborHandle)
	err = dec.Decode(&v2)
	if err != nil {
		log.Fatalf("dec.Decode() failed, err = %v", err)
	}

	assert.Equal(t, v2.Field3, 456)
	assert.Equal(t, v2.Field4, "")
	assert.Equal(t, v2.Field5, false)
}

// to execute benchmarks, use `go test -bench=. ./tests/...` command
func BenchmarkSerialization(t *testing.B) {
	for i := 0; i < 1000000; i++ {
		encodeDecode()
	}
}
