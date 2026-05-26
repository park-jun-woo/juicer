//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what requirements.txt 또는 pyproject.toml에서 fastapi 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectFastAPI(root string) bool {
	for _, name := range []string{"requirements.txt", "pyproject.toml"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		if strings.Contains(string(data), "fastapi") {
			return true
		}
	}
	return false
}
