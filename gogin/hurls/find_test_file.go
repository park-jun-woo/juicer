//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what testsDir에서 엔드포인트에 매칭되는 .hurl 파일 탐색
package hurls

import (
	"os"
	"path/filepath"
	"strings"
)

// findTestFile searches testsDir for a .hurl file that matches the endpoint.
// Matching: file contains "METHOD <scheme>{{host}}/path" or "METHOD {{host}}/path".
func findTestFile(testsDir, method, path string) string {
	entries, err := os.ReadDir(testsDir)
	if err != nil {
		return ""
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".hurl") {
			continue
		}
		fp := filepath.Join(testsDir, e.Name())
		data, err := os.ReadFile(fp)
		if err != nil {
			continue
		}
		if matchesEndpoint(string(data), method, path) {
			return fp
		}
	}
	return ""
}
