//ff:func feature=scan type=extract control=sequence
//ff:what package.json에서 fastify 의존을 확인한다 (express, @nestjs/core 미포함 시에만 true)
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectFastify(root string) bool {
	data, err := os.ReadFile(filepath.Join(root, "package.json"))
	if err != nil {
		return false
	}
	content := string(data)
	if strings.Contains(content, "@nestjs/core") {
		return false
	}
	if strings.Contains(content, "\"express\"") {
		return false
	}
	return strings.Contains(content, "\"fastify\"")
}
