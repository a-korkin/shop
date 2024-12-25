package tools

import (
	pb "github.com/a-korkin/shop/internal/common"
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
}

func getParams(query string) map[string]string {
	raw := strings.Split(query, "&")
	params := make(map[string]string, 0)
	for _, keyValue := range raw {
		kv := strings.Split(keyValue, "=")
		params[kv[0]] = kv[1]
	}
	return params
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
		return -1, nil
	}
	return strconv.Atoi(tokens[2])
}

func GetPageParams(rawQuery string) *pb.PageParams {
	pageParams := pb.PageParams{
		Page:  1,
		Limit: 20,
	}
	if rawQuery == "" {
		return &pageParams
	}
	queryParams := getParams(rawQuery)
	if p, ok := queryParams["page"]; ok {
		if page, err := strconv.Atoi(p); err == nil {
			pageParams.Page = max(int32(page), 1)
		}
	}
	if l, ok := queryParams["limit"]; ok {
		if limit, err := strconv.Atoi(l); err == nil {
			pageParams.Limit = int32(limit)
		}
	}
	return &pageParams
}
