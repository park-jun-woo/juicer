//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what @relation의 onDelete/onUpdate를 ON DELETE/UPDATE 절로 변환
package prisma

import "strings"

// referentialActions renders ON DELETE/ON UPDATE clauses from a @relation arg.
func referentialActions(rel string) string {
	var sb strings.Builder
	if act := referentialAction(rel, "onDelete"); act != "" {
		sb.WriteString(" ON DELETE ")
		sb.WriteString(act)
	}
	if act := referentialAction(rel, "onUpdate"); act != "" {
		sb.WriteString(" ON UPDATE ")
		sb.WriteString(act)
	}
	return sb.String()
}
