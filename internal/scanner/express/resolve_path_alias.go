//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what path alias를 실제 절대 경로로 해석한다 (@/api/users → src/api/users.ts)
package express

import (
	"path/filepath"
	"strings"
)

func resolvePathAlias(absRoot, importPath string, aliases map[string]string) string {
	for prefix, replacement := range aliases {
		if !strings.HasPrefix(importPath, prefix) {
			continue
		}
		rest := importPath[len(prefix):]
		base := filepath.Join(absRoot, replacement, rest)
		resolved := resolveExtension(base)
		if resolved != "" {
			return resolved
		}
	}
	return ""
}
