//ff:func feature=scan type=extract control=sequence
//ff:what package.json에서 @nestjs/core 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectNestJS(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "package.json"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "@nestjs/core")
}
