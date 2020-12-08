package hash

var (
	Uint32Base62  _uint32
	Uint64Base62  _uint64
	Uint128Base62 _uint128
)

type (
	_uint32  struct{}
	_uint64  struct{}
	_uint128 struct{}
)

func (_uint32) Encoding(i uint32) (s string) {
	for ; i != 0; i /= 62 {
		s = base62Set[i%62] + s
	}
	return
}
func (_uint32) Decoding(s string) (i uint32) {
	for _, v := range s {
		i = i*62 + uint32(base62To(string(v)))
	}
	return
}

var base62Set = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z",
}

func base62To(s string) int {
	for i, v := range base62Set {
		if v == s {
			return i
		}
	}
	return 0
}

func (_uint64) Encoding(i uint64) (s string) {
	for ; i != 0; i /= 62 {
		s = base62Set[i%62] + s
	}
	return
}
func (_uint64) Decoding(s string) (i uint64) {
	for _, v := range s {
		i = i*62 + uint64(base62To(string(v)))
	}
	return
}

func (_uint128) Encoding(i Uint128) (s string) {
	for ; i[1] != 0; i[1] /= 62 {
		s = base62Set[i[1]%62] + s
	}
	for ; i[0] != 0; i[0] /= 62 {
		s = base62Set[i[0]%62] + s
	}
	return
}
func (_uint128) Decoding(s string) (i Uint128) {
	for _, v := range s[:11] {
		i[0] = i[0]*62 + uint64(base62To(string(v)))
	}
	for _, v := range s[11:] {
		i[1] = i[1]*62 + uint64(base62To(string(v)))
	}
	return
}
