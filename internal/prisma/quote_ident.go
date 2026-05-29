//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what SQL 식별자를 큰따옴표로 인용하고 내부 따옴표를 ""로 이스케이프
package prisma

import "strings"

// quoteIdent wraps a SQL identifier in double quotes, escaping any embedded
// double quote as "". Always quoting matches Prisma's policy and stays safe
// regardless of reserved words or case folding.
func quoteIdent(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}
