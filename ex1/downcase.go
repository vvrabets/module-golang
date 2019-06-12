package downcase

func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func Downcase(s string) (string, error) {
	str := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			str[i] = string(s[i] + 32)
		} else {
			str[i] = string(s[i])
		}
	}
	result := Join(str, "")
	return result, nil
}
