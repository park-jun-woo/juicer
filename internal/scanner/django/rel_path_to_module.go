//ff:func feature=scan type=convert control=sequence topic=django
//ff:what 파일 상대경로를 dotted Python 모듈 경로로 변환한다
package django

import (
	"path/filepath"
	"strings"
)

// relPathToModule converts a file relative path to a dotted Python module path.
// e.g. "blog/urls.py" -> "blog.urls", "pkg/__init__.py" -> "pkg".
func relPathToModule(relPath string) string {
	p := filepath.ToSlash(relPath)
	p = strings.TrimSuffix(p, ".py")
	parts := strings.Split(p, "/")
	if n := len(parts); n > 0 && parts[n-1] == "__init__" {
		parts = parts[:n-1]
	}
	return strings.Join(parts, ".")
}
