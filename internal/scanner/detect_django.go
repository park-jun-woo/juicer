//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what requirements.txt, setup.py, pyproject.toml에서 django 또는 djangorestframework 의존을 확인한다 (flask, fastapi 미포함 시에만 true)
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectDjango(root string) bool {
	for _, name := range []string{"requirements.txt", "setup.py", "pyproject.toml"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		content := strings.ToLower(string(data))
		hasDjango := strings.Contains(content, "django") || strings.Contains(content, "djangorestframework")
		hasFlask := strings.Contains(content, "flask")
		hasFastAPI := strings.Contains(content, "fastapi")
		if hasDjango && !hasFlask && !hasFastAPI {
			return true
		}
	}
	return false
}
