//ff:func feature=scan type=convert control=iteration dimension=1 topic=fastapi
//ff:what 상대 경로 깊이만큼 상위 디렉토리로 이동한다
package fastapi

import "path/filepath"

// navigateUp moves up the directory tree by (dots - 1) levels.
func navigateUp(base string, dots int) string {
	for i := 1; i < dots; i++ {
		base = filepath.Dir(base)
	}
	return base
}
