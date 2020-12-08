package hash

import "encoding/binary"

type Uint128 [2]uint64

func (u *Uint128) setLower64(l uint64)  { u[0] = l }
func (u *Uint128) setHigher64(h uint64) { u[1] = h }
func (u Uint128) Lower64() uint64       { return u[0] }
func (u Uint128) Higher64() uint64      { return u[1] }
func (u Uint128) Bytes() []byte {
	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b, u[0])
	binary.LittleEndian.PutUint64(b[8:], u[1])
	return b
}
