//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Cargo.toml에서 actix-web 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectActix(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "Cargo.toml"))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "actix-web")
}
