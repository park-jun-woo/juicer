//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 세션의 idx 번째 메서드 스켈레톤 정보 출력
package sqls

import (
	"fmt"
)

// printSkeleton outputs the skeleton info for the method at index idx.
// methods is the pre-extracted list of MethodSkeleton from Extract.
//
func printSkeleton(sess *Session, idx int, methods []MethodSkeleton) {
	m := sess.Methods[idx]

	// Find matching skeleton
	var sk *MethodSkeleton
	for i, s := range methods {
		if s.Repo+"."+s.Method == m.ID {
			sk = &methods[i]
			break
		}
	}

	if sk == nil {
		fmt.Printf("%s  TODO\n", m.ID)
		return
	}

	queryName := toQueryName(m.ID)
	hint := sqlcHint(sk)

	crud := sk.CRUD
	if sk.Dynamic {
		crud += " (dynamic)"
	}

	fmt.Printf("%s  TODO\n", m.ID)
	fmt.Printf("  crud: %s\n", crud)
	fmt.Printf("  tables: %s\n", formatSlice(sk.Tables))
	if len(sk.Params) > 0 {
		fmt.Printf("  params: %s\n", formatSlice(sk.Params))
	}
	if len(sk.Returns) > 0 {
		fmt.Printf("  returns: %s\n", formatSlice(sk.Returns))
	}
	if len(sk.SQLFragments) > 0 {
		fmt.Printf("  sql: %s\n", sk.SQLFragments[0])
	}
	fmt.Printf("  → Write query as \"-- name: %s %s\"\n", queryName, hint)
}

