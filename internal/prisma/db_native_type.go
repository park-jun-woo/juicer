//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what @db.Xxx native 타입 어노테이션을 소문자 SQL 타입으로 변환
package prisma

import "strings"

// dbNativeType converts a @db.<Type> attribute (e.g. @db.VarChar(255))
// into its lowercase SQL type (varchar(255)).
func dbNativeType(attr string) (string, bool) {
	if !strings.HasPrefix(attr, "@db.") {
		return "", false
	}
	t := strings.TrimSpace(attr[len("@db."):])
	if t == "" {
		return "", false
	}
	if i := strings.IndexByte(t, '('); i >= 0 {
		return strings.ToLower(t[:i]) + t[i:], true
	}
	return strings.ToLower(t), true
}
