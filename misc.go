package bdd

import "unicode/utf8"

var space = map[int]string{}

func runeWidth(r rune) int {
	switch {
	case r == utf8.RuneError || r < '\x20':
		return 0
	case '\x20' <= r && r < '\u2000':
		return 1
	case '\u2000' <= r && r < '\uFF61':
		return 2
	case '\uFF61' <= r && r < '\uFFA0':
		return 1
	case '\uFFA0' <= r:
		return 2
	}

	return 0
}

func strLen(str string) (i int) {
	for _, v := range str {
		i += runeWidth(v)
	}
	return
}

func spac(depth int) string {
	b := []byte{}
	if depth > 0 {
		for i := 0; i != depth; i++ {
			b = append(b, Space)
		}
	}
	return string(b)
}

func spacing(depth int) string {
	v, ok := space[depth]
	if ok {
		return v
	}
	v = "\n" + spac(depth-1)
	space[depth] = v
	return v
}
