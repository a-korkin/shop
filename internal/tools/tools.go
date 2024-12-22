package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func getTokens(uri string) []string {
	return strings.Split(uri, "/")
}

func GetResource(uri string) string {
	tokens := getTokens(uri)
	if len(tokens) < 2 {
		return ""
	}
	return tokens[1]
}

func GetId(uri string) (int, error) {
	tokens := getTokens(uri)
	if len(tokens) < 3 {
		return 0, fmt.Errorf("uri tokens too short")
	}
	return strconv.Atoi(tokens[2])
}
