package bdd

import (
	"strconv"
	"strings"
)

func Flatten(nested map[string]interface{}) (flatmap map[string]interface{}) {
	if nested != nil {
		var prefixes []string
		flatten(0, flatmap, nested, prefixes)
		return
	}
	return nested
}

func flatten(deep int, flatMap map[string]interface{}, nested interface{}, prefixes []string) {
	switch t := nested.(type) {
	case map[string]interface{}:
		for k, v := range t {
			flatten(deep+1, flatMap, v, append(prefixes, k))
		}
	case []interface{}:
		for i, v := range t {
			flatten(deep+1, flatMap, v, append(prefixes, strconv.Itoa(i)))
		}
	default:
		key := strings.Join(prefixes, ".")
		flatMap[key] = t
	}
}
