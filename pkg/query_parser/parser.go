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

func cleanTTL(ttl string) (string, error) {
	_, err := strconv.ParseFloat(strings.TrimSpace(ttl), 32)
	if err != nil {
		return "-1", fmt.Errorf("parse error: invalid ttl provided")
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
	pattern := `^(?i)(GET|SET|DELETE|CLEAR|TTL|EXPIRE|KEYS){1}([ ]+[^\s]*)?([ ]+\"(?:.*?)*\")?([ ]+\d*)?$`
	re := regexp.MustCompile(pattern)
	// re := regexp.MustCompile(`^(?i)(get|set|delete|clear|ttl|expire|keys)(\s+.+){0,3}\s*$`)

	if !re.MatchString(query) {
		return parsed, fmt.Errorf("parse error: invalid query")
	}

	// Case 3: Return ParsedQuery based on the command
	splitted := re.FindStringSubmatch(query)[1:]
	cmd := splitted[0]

	// fmt.Printf("%q\n", splitted)

	if cmd == "GET" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("GET", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	} else if cmd == "DELETE" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("DELETE", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	} else if cmd == "CLEAR" {
		return createParsedQuery("DELETE", "", "", ""), nil
	} else if cmd == "KEYS" {
		return createParsedQuery("KEYS", "", "", ""), nil
	} else if cmd == "TTL" {
		if splitted[1] != "" {
			key := cleanKeys(splitted[1])
			return createParsedQuery("TTL", key, "", ""), nil
		}
		return parsed, fmt.Errorf("parse error: invalid key provided")
	} else if cmd == "SET" {
		args := splitted[1:]
		length := len(args)
		if length < 2 {
			return parsed, fmt.Errorf("parse error: key/value not provided")
		}
		key := cleanKeys(args[0])
		value := cleanKeys(args[1])
		ttl := "-1"
		if length >= 3 && args[2] != "" {
			val, err := cleanTTL(args[2])
			if err != nil {
				return parsed, err
			}
			ttl = val
		}
		return createParsedQuery("SET", key, value, ttl), nil
	}

	return parsed, nil
}

func main() {
	queries := []string{
		`GET`,
		`GET key`,
		`GET          key    `,
		`GET 12key`,
		`GET @#!$$_12key`,
		`GET "key"`,
		`SET name "Rajab is good"`,
		`SET name       "Rajab is good"  `,
		`SET name "adasdasdasd  \n\nRajab is good"  `,
		`SET name "adasdasdasd  \n\nRajab is good`,
		`SET name "adasdasdasd  \n\nRajab is good" sfsfsd fsdfs fsd fd`,
		`SET name "Rajab" 10000`,
		`SET name "भारत" 10000`,
		`SET name "Rajab" 10000   `,
		`SET name`,
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
