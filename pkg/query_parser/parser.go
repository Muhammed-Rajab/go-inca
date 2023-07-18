package queryparser

import (
	"fmt"
	"regexp"
	"strconv"
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
type args struct {
	Key string
	Val string
	TTL string
}

type ParsedQuery struct {
	CommandType string
	Args        args
}

func createParsedQuery(cmdType string, key, val, ttl string) *ParsedQuery {
	return &ParsedQuery{cmdType, args{key, val, ttl}}
}

func cleanKeys(key string) string {
	return strings.TrimSpace(key)
}

func cleanString(s string) string {
	return strings.TrimSpace(s)
}

func cleanTTL(ttl string) (string, error) {
	v, err := strconv.ParseFloat(strings.TrimSpace(ttl), 32)
	if err != nil {
		return "-1", fmt.Errorf("parse error: invalid ttl provided")
	}
	if v <= 0 {
		return "-1", nil
	}
	return ttl, nil
}

func Parse(query string) (*ParsedQuery, error) {

	query = strings.TrimSpace(query)

	parsed := &ParsedQuery{}

	// Case 1: Empty query
	if len(query) == 0 {
		return parsed, fmt.Errorf("parse error: query empty")
	}

	// Case 2: Valid Command parse using regex
	// CMD KEY VAL TTL
	pattern := `^(?i)(GET|SET|DELETE|CLEAR|TTL|EXPIRE|KEYS){1}([ ]+[^\s]*)?([ ]+\"(?:.*?)*\")?([ ]+-?\d*)?$`
	re := regexp.MustCompile(pattern)

	// Case 3: Query doesn't match the pattern
	if !re.MatchString(query) {
		return parsed, fmt.Errorf("parse error: invalid query")
	}

	// Case 4: Return ParsedQuery based on the command
	splitted := re.FindStringSubmatch(query)[1:]
	args := splitted[1:]
	cmd := splitted[0]

	if cmd == "GET" {
		key := cleanKeys(splitted[1])
		if key != "" {
			return createParsedQuery("GET", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided %s", key)
	} else if cmd == "DELETE" {
		key := cleanKeys(splitted[1])
		if key != "" {
			return createParsedQuery("DELETE", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided %s", key)
	} else if cmd == "CLEAR" {
		return createParsedQuery("CLEAR", "", "", ""), nil
	} else if cmd == "KEYS" {
		return createParsedQuery("KEYS", "", "", ""), nil
	} else if cmd == "TTL" {
		key := cleanKeys(splitted[1])
		if key != "" {
			return createParsedQuery("TTL", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided %s", key)
	} else if cmd == "SET" {
		key := cleanKeys(args[0])
		value := cleanString(args[1])
		if key == "" {
			return parsed, fmt.Errorf("parse error: key not provided")
		}

		ttl := "-1"

		// Check if ttl is present in query
		if args[2] != "" {
			val, err := cleanTTL(args[2])
			if err != nil {
				return parsed, err
			}
			ttl = cleanString(val)
		}

		// Remove this clean string method and handle it properly later
		return createParsedQuery("SET", key, value, ttl), nil
	} else if cmd == "EXPIRE" {
		key := cleanKeys(args[0])
		ttl := cleanString(args[2])

		if key == "" || ttl == "" {
			return parsed, fmt.Errorf("parse error: key/duration value not provided")
		}

		ttl, err := cleanTTL(ttl)

		if err != nil {
			return parsed, err
		}

		return createParsedQuery("EXPIRE", cleanString(key), "", cleanString(ttl)), nil
	}

	return parsed, fmt.Errorf("parser error: command not found")
}
