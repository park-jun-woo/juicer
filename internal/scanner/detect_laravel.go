//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what composer.json에서 laravel/framework 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectLaravel(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "composer.json"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "laravel/framework")
}
