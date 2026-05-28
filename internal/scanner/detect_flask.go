//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what requirements.txt, setup.py, pyproject.toml에서 flask 의존을 확인한다 (fastapi 미포함 시에만 true)
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectFlask(root string) bool {
	for _, name := range []string{"requirements.txt", "setup.py", "pyproject.toml"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		content := strings.ToLower(string(data))
		if strings.Contains(content, "flask") && !strings.Contains(content, "fastapi") {
			return true
		}
	}
	return false
}
