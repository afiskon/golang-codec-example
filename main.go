package main

import (
	"github.com/ugorji/go/codec"
	. "go-sandbox/codec-examples-gen/types"
	"log"
)

func main() {
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
	log.Printf("bs = %q, len(bs) = %d, cap(bs) = %d", bs, len(bs), cap(bs))

	// Decode bs to v2

	var v2 Hero
	dec := codec.NewDecoderBytes(bs, &cborHandle)
	err = dec.Decode(&v2)
	if err != nil {
		log.Fatalf("dec.Decode() failed, err = %v", err)
	}

	log.Printf("v2 = %v", v2)
	if v2.WariorInfo != nil{
		log.Printf("WariorInfo = %v", *v2.WariorInfo)
	}
	if v2.MageInfo != nil {
		log.Printf("MageInfo = %v", *v2.MageInfo)
	}
}
