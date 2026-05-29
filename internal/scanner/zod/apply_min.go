//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what Zod .min(N) 메서드를 Field에 반영한다 (string→minLength, number→minimum)
package zod

import "github.com/park-jun-woo/codistill/internal/scanner"

// ApplyMin — .min(N)을 Field에 반영
func ApplyMin(f *scanner.Field, m ChainMethod) {
	if len(m.Args) == 0 {
		return
	}
	n := parseIntArg(m.Args[0])
	if n == nil {
		return
	}
	if f.Type == "string" {
		f.MinLength = n
	} else {
		f.Minimum = n
	}
}
