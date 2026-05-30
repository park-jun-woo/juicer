//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what csFileInfo 테스트 헬퍼
package dotnet

import "testing"

func csFileInfo(t *testing.T, src string) *fileInfo {
	t.Helper()
	root, b := parseCS(t, src)
	return &fileInfo{
		absPath:     "/abs/C.cs",
		relPath:     "C.cs",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		usings:      extractUsings(root, b),
	}
}
