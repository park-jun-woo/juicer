//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what Zod .max(N) 메서드를 Field에 반영한다 (string→maxLength, number→maximum)
package zod

import "github.com/park-jun-woo/codistill/internal/scanner"

// ApplyMax — .max(N)을 Field에 반영
func ApplyMax(f *scanner.Field, m ChainMethod) {
	if len(m.Args) == 0 {
		return
	}
	n := parseIntArg(m.Args[0])
	if n == nil {
		return
	}
	if f.Type == "string" {
		f.MaxLength = n
	} else {
		f.Maximum = n
	}
}
