//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what pom.xml 또는 build.gradle에서 spring-boot-starter-web 의존을 확인한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectSpring(root string) bool {
	for _, name := range []string{"pom.xml", "build.gradle", "build.gradle.kts"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			continue
		}
		if strings.Contains(string(data), "spring-boot-starter-web") {
			return true
		}
	}
	return false
}
