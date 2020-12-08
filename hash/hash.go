package hash

import (
	"encoding/binary"
	"hash"
)

// Hash128 is the common interface implemented by all 128-bit hash functions.
type Hash128 interface {
	hash.Hash
	Sum128() Uint128
}

type (
	City32  []byte
	City64  []byte
	City128 []byte
)

func New32() hash.Hash32 { return &City32{} }
func New64() hash.Hash64 { return &City64{} }
func New128() Hash128    { return &City128{} }

func (c *City32) Sum(b []byte) []byte {
	b2 := make([]byte, 4)
	binary.BigEndian.PutUint32(b2, c.Sum32())
	return append(b, b2...)
}
func (c *City32) Sum32() uint32  { return CityHash32(*c, uint32(len(*c))) }
func (c *City32) Reset()         { *c = (*c)[0:0] }
func (c *City32) BlockSize() int { return 1 }
func (c *City32) Size() int      { return 4 }
func (c *City32) Write(s []byte) (n int, err error) {
	*c = append(*c, s...)
	return len(s), nil
}

func (c *City64) Sum(b []byte) []byte {
	b2 := make([]byte, 8)
	binary.BigEndian.PutUint64(b2, c.Sum64())
	return append(b, b2...)
}
func (c *City64) Sum64() uint64  { return CityHash64(*c, uint32(len(*c))) }
func (c *City64) Reset()         { *c = (*c)[0:0] }
func (c *City64) BlockSize() int { return 1 }
func (c *City64) Size() int      { return 8 }
func (c *City64) Write(s []byte) (n int, err error) {
	*c = append(*c, s...)
	return len(s), nil
}

func (c *City128) Sum(b []byte) []byte { return append(b, c.Sum128().Bytes()...) }
func (c *City128) Sum128() Uint128     { return CityHash128(*c, uint32(len(*c))) }
func (c *City128) Reset()              { *c = (*c)[0:0] }
func (c *City128) BlockSize() int      { return 1 }
func (c *City128) Size() int           { return 16 }
func (c *City128) Write(s []byte) (n int, err error) {
	*c = append(*c, s...)
	return len(s), nil
}
