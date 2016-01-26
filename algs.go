package main

import (
	"crypto/rand"
	"encoding/binary"
	"io"

	mt "github.com/EricLagergren/go-prng/mersenne_twister_64"
	xs "github.com/EricLagergren/go-prng/xorshift"
)

func randAlg(b64 bool) string {
	buf := make([]byte, *length)
	n, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		exit(err.Error())
	}

	if n != *length {
		exit("io.ReadFull did not read %d bytes, instead read %d", *length, n)
	}

	return format(buf, b64)
}

// round rounds up to the nearest multiple of a number to prevent
// PutUvarint from panicking.
func round(have, want int) int {
	return 1 + ((have + want - 1) & (^(want - 1)))
}

func xorshiftAlg(b64 bool) string {
	r := &xs.Shift4096Star{}
	r.Seed()
	return doAlg(r.Next, b64)
}

func mersenneAlg(b64 bool) string {
	return doAlg(mt.NewMersennePrime().Int64, b64)
}

func doAlg(fn func() uint64, b64 bool) string {
	buf := make([]byte, round(*length, binary.MaxVarintLen64))
	for i := 0; i < len(buf)-1; i += binary.MaxVarintLen64 {
		binary.PutUvarint(buf[i:], fn())
	}
	return format(buf, b64)
}
