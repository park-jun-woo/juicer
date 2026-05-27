//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what main.ts를 우선 배치하고 src/ 내 나머지 .ts 파일을 후순위로 모아 반환한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// collectPrefixCandidates returns src/*.ts paths with main.ts first.
// Non-.ts entries and directories are skipped.
func collectPrefixCandidates(root string) []string {
	mainPath := filepath.Join(root, "src", "main.ts")
	candidates := []string{mainPath}

	entries, err := os.ReadDir(filepath.Join(root, "src"))
	if err != nil {
		return candidates
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".ts") || e.Name() == "main.ts" {
			continue
		}
		candidates = append(candidates, filepath.Join(root, "src", e.Name()))
	}
	return candidates
}
