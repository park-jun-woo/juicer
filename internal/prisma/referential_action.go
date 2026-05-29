//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what onDelete/onUpdate 값(Cascade 등)을 SQL 참조 동작 문자열로 변환
package prisma

import "strings"

// prismaActions maps Prisma referential actions to SQL clauses.
var prismaActions = map[string]string{
	"Cascade":    "CASCADE",
	"Restrict":   "RESTRICT",
	"NoAction":   "NO ACTION",
	"SetNull":    "SET NULL",
	"SetDefault": "SET DEFAULT",
}

// referentialAction extracts the value of key (onDelete/onUpdate) from a
// @relation argument and maps it to its SQL clause, or "" if absent.
func referentialAction(rel, key string) string {
	idx := strings.Index(rel, key+":")
	if idx < 0 {
		return ""
	}
	rest := strings.TrimSpace(rel[idx+len(key)+1:])
	rest, _ = firstWord(rest)
	return prismaActions[rest]
}
