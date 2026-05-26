//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what env 파일에서 API_PREFIX 값을 읽는다
package nestjs

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// readEnvPrefix reads API_PREFIX from an env file (e.g. .env.example).
func readEnvPrefix(root, name string) string {
	f, err := os.Open(filepath.Join(root, name))
	if err != nil {
		return ""
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		k, v, _ := strings.Cut(line, "=")
		if strings.TrimSpace(k) == "API_PREFIX" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
