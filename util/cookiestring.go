package util

import (
	"fmt"
	"strings"
)

func ToCookiestring(cookie map[string]string) string {
	rt := ""
	for k, v := range cookie {
		rt += fmt.Sprintf("%s=%s; ", k, v)
	}
	return strings.Trim(rt, "; ")
}

func FromCookiestring(s string) map[string]string {
	tokenstrings := strings.Split(s, ";")
	tokens := map[string]string{}
	for _, v := range tokenstrings {
		trim := strings.Trim(v, " ")
		split := strings.Split(trim, "=")
		tokens[split[0]] = split[1]
	}
	return tokens
}
