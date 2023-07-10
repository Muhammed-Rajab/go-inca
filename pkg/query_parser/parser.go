package main

import "fmt"

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

func Parse(query string) (*ParsedQuery, error) {
	/*
		Task
		-----
		1. Get the command type
			- First item
		2. Get all the args as a slice considering that there's gonna be strings and all type of shit. (I really regret doing this bruh, I've no clue:)
	*/
	/*
		GET key
		SET key val
		SET key val ttl
		DELETE key
		CLEAR

		TTL key
		EXPIRE key ttl
		KEYS
	*/
	parsed := &ParsedQuery{}

	// Case 1: Empty query
	if len(query) == 0 {
		return parsed, fmt.Errorf("query can't be empty")
	}

	// Case 2: Valid Command parse using regex
	return parsed, nil
}

func main() {
	Parse("GET key")
	Parse("SET name \"Rajab\"")
	Parse("SET name \"Rajab\" 10000")
	Parse("DELETE name")
	Parse("CLEAR")
	Parse("TTL name")
	Parse("EXPIRE name 800")
	Parse("KEYS")
}
