//ff:func feature=scan type=extract control=sequence
//ff:what go.mod에서 gin-gonic/gin 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectGoGin(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "github.com/gin-gonic/gin")
}
