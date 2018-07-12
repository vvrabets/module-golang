package cipher

//import "fmt"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyCaesar struct{}

type MyShift struct {
	shift int
}

type MyVigenere struct {
	key string
}

func NewCaesar() Cipher {
	return MyCaesar{}
}

func NewShift(shift int) Cipher {
	if shift > 25 || shift < -25 || shift == 0 {
		return MyShift{0}
	}
	return MyShift{shift}
}

func (c MyCaesar) Encode(s string) string {
	str := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			tmp := s[i] + 32 + 3
			if tmp > 122 {
				tmp = tmp - 122 + 96
			}
			str[i] = tmp
		} else if s[i] > 96 && s[i] < 123 {
			tmp := s[i] + 3
			if tmp > 122 {
				tmp = tmp - 122 + 96
			}
			str[i] = tmp
		} else {
			continue
		}
	}
	return string(str)
}

func (c MyCaesar) Decode(s string) string {
	str := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != 0 {
			tmp := s[i] - 3
			if tmp < 97 {
				tmp = 123 - (97 - tmp)
			}
			str[i] = tmp
		}
	}
	return string(str)
}

func (shift MyShift) Encode(s string) string {
	str := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			tmp := s[i] + 32 + byte(shift.shift)
			if tmp > 122 {
				tmp = tmp - 122 + 96
			}
			str[i] = tmp
		} else if s[i] > 96 && s[i] < 123 {
			tmp := s[i] + byte(shift.shift)
			if tmp > 122 {
				tmp = tmp - 122 + 96
			}
			str[i] = tmp
		} else {
			continue
		}
	}
	return string(str)
}

func (shift MyShift) Decode(s string) string {
	str := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != 0 {
			tmp := s[i] - byte(shift.shift)
			if tmp < 97 {
				tmp = 123 - (97 - tmp)
			}
			str[i] = tmp
		}
	}
	return string(str)
}

func (v MyVigenere) Encode(p string) string {
	s := ReadyToEncode(p)
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] + uint8(v.key[i%len(v.key)]) - uint8('a')

		if tmp > 122 {
			tmp -= 26
		}
		str[i] = tmp
	}

	return string(str)
}

func (v MyVigenere) Decode(s string) string {
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] - uint8(v.key[i%len(v.key)]) + uint8('a')

		if tmp < 97 {
			tmp += 26
		}

		str[i] = tmp
	}

	return string(str)
}

func NewVigenere(key string) Cipher {
	if !invalid_key(key) {
		return nil
	}

	return MyVigenere{key}
}

func ReadyToEncode(s string) string {
	count := 0
	j := 0

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}

	str := make([]byte, count)

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			if s[i] >= 'A' && s[i] <= 'Z' {
				str[j] = s[i] - 'A' + 'a'
			} else {
				str[j] = s[i]
			}
			j++
		}
	}
	return string(str)
}

func invalid_key(s string) bool {
	flag := false

	if len(s) < 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || (s[i] > 40 && s[i] < 91) {
			return false
		}

		if s[i] != 'a' {
			flag = true
		}
	}
	return flag
}

/*func main() {
	s := "todayisholiday"
	var c Cipher = NewCaesar()
	//	shift := NewShift(3)
	fmt.Println(c.Encode(s))
	fmt.Println(c.Decode(c.Encode(s)))
}*/
