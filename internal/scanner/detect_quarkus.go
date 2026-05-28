//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what pom.xml 또는 build.gradle에서 quarkus-rest 또는 quarkus-resteasy 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectQuarkus(root string) bool {
	for _, name := range []string{"pom.xml", "build.gradle", "build.gradle.kts"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		content := string(data)
		if strings.Contains(content, "spring-boot") {
			continue
		}
		if strings.Contains(content, "quarkus-rest") || strings.Contains(content, "quarkus-resteasy") {
			return true
		}
	}
	return false
}
