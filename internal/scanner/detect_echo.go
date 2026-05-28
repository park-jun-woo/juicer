//ff:func feature=scan type=extract control=sequence
//ff:what go.mod에서 labstack/echo 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectEcho(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "github.com/labstack/echo")
}
