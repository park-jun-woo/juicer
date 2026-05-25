package sqls

import (
	"regexp"
)

var (
	reFrom   = regexp.MustCompile(`(?i)\bFROM\s+(\w+)`)
	reInsert = regexp.MustCompile(`(?i)\bINSERT\s+INTO\s+(\w+)`)
	reUpdate = regexp.MustCompile(`(?i)\bUPDATE\s+(\w+)\s+SET\b`)
	reDelete = regexp.MustCompile(`(?i)\bDELETE\s+FROM\s+(\w+)`)
	reJoin   = regexp.MustCompile(`(?i)\bJOIN\s+(\w+)`)

	sqlKeywords = regexp.MustCompile(`(?i)\b(SELECT|INSERT|UPDATE|DELETE|FROM|WHERE|JOIN)\b`)

	// SQL reserved words that should not be treated as table names
	reservedWords = map[string]bool{
		"lateral": true, "select": true, "insert": true, "update": true,
		"delete": true, "from": true, "where": true, "join": true,
		"inner": true, "outer": true, "left": true, "right": true,
		"cross": true, "full": true, "on": true, "and": true,
		"or": true, "not": true, "in": true, "exists": true,
		"set": true, "values": true, "into": true, "as": true,
		"order": true, "by": true, "group": true, "having": true,
		"limit": true, "offset": true, "union": true, "all": true,
		"distinct": true, "case": true, "when": true, "then": true,
		"else": true, "end": true, "null": true, "true": true,
		"false": true, "is": true, "like": true, "ilike": true,
		"between": true, "asc": true, "desc": true, "cascade": true,
		"restrict": true, "returning": true, "coalesce": true,
		"count": true, "sum": true, "avg": true, "min": true, "max": true,
	}
)
