//ff:func feature=prisma type=convert control=selection topic=prisma
//ff:what Prisma @default 식을 SQL DEFAULT 식으로 변환 (now/uuid/리터럴/문자열)
package prisma

import "strings"

// mapDefaultExpr converts a Prisma default expression to a SQL default.
// autoincrement()/cuid() return "" (no DEFAULT; handled elsewhere).
func mapDefaultExpr(inner string) string {
	switch inner {
	case "autoincrement()", "cuid()", "cuid(2)":
		return ""
	case "now()", "true", "false":
		return inner
	case "uuid()", "uuid(4)", "uuid(7)":
		return "gen_random_uuid()"
	}
	if strings.HasPrefix(inner, `"`) && strings.HasSuffix(inner, `"`) && len(inner) >= 2 {
		return "'" + inner[1:len(inner)-1] + "'"
	}
	return inner
}
