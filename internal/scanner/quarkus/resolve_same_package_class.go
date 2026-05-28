//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 같은 패키지에서 클래스 파일 경로를 해석한다
package quarkus

import (
	"os"
	"path/filepath"
)

func resolveSamePackageClass(referrerPath, className string) string {
	dir := filepath.Dir(referrerPath)
	candidate := filepath.Join(dir, className+".java")
	if _, err := os.Stat(candidate); err == nil {
		return candidate
	}
	return ""
}
