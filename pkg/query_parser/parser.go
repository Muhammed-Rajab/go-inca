package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	type ParsedQuery struct {
		CommandType string;
		Args []string;
	}

	Function Definitions:
	func Parse(query string) ParsedQuery
*/

type ParsedQuery struct {
	CommandType string
	Args        []string
}

func createParsedQuery(cmdType string, args []string) *ParsedQuery {
	return &ParsedQuery{cmdType, args}
}

func cleanKeys(key string) string {
	return strings.TrimSpace(key)
}

func Parse(query string) (*ParsedQuery, error) {

	parsed := &ParsedQuery{}

	// Case 1: Empty query
	if len(query) == 0 {
		return parsed, fmt.Errorf("parse error: query empty")
	}

	// Case 2: Valid Command parse using regex
	re := regexp.MustCompile(`^(?i)(get|set|delete|clear|ttl|expire|keys)(\s+.+){0,3}\s*$`)

	if !re.MatchString(query) {
		return parsed, fmt.Errorf("parse error: invalid query")
	}

	// Case 3: Return ParsedQuery based on the command
	pattern := `^(?i)(GET|SET|DELETE|CLEAR|TTL|EXPIRE|KEYS){1}([ ]+[^\s]*)?([ ]\"(?:.*?)*\")?([ ]+\d*)?$`
	re = regexp.MustCompile(pattern)
	splitted := re.FindStringSubmatch(query)[1:]
	cmd := splitted[0]

	// fmt.Printf("%q\n", splitted)

	if cmd == "GET" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("GET", []string{key}), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	} else if cmd == "DELETE" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("DELETE", []string{key}), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	} else if cmd == "CLEAR" {
		return createParsedQuery("DELETE", []string{}), nil
	} else if cmd == "KEYS" {
		return createParsedQuery("KEYS", []string{}), nil
	} else if cmd == "TTL" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("TTL", []string{key}), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	}

	return parsed, nil
}

func main() {
	queries := []string{
		`GET`,
		`GET key`,
		`GET 12key`,
		`GET @#!$$_12key`,
		`GET "key"`,
		`SET name "Rajab is good"`,
		`SET name "Rajab" 10000`,
		"DELETE name",
		"CLEAR",
		"TTL name",
		"EXPIRE name 800",
		"KEYS",
	}

	for index, query := range queries {
		parsed, err := Parse(query)
		if err != nil {
			fmt.Printf("%d -> %s -> ERROR: %s\n", index, query, err.Error())
			continue
		}
		fmt.Printf("%d -> %s -> %v\n", index, query, parsed)
	}
}
