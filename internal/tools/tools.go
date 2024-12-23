package tools

import (
	"strconv"
	"strings"
)

func getTokens(uri string) []string {
	tokens := strings.Split(uri, "/")
	result := make([]string, 0)
	for _, token := range tokens {
		result = append(result, strings.Split(token, "?")[0])
	}
	return result
	// return strings.Split(uri, "/")
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
		return 0, nil
	}
	return strconv.Atoi(tokens[2])
}

func GetParams(query string) map[string]string {
	raw := strings.Split(query, "&")
	params := make(map[string]string, 0)
	for _, keyValue := range raw {
		kv := strings.Split(keyValue, "=")
		params[kv[0]] = kv[1]
	}
	return params
}
