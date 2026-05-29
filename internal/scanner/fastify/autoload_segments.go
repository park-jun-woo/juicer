//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 베이스 디렉터리 기준 상대 경로를 autoload prefix 세그먼트 목록으로 변환한다
package fastify

import (
	"path/filepath"
	"strings"
)

func autoloadSegments(rel string) []string {
	dir := filepath.Dir(rel)
	var segs []string
	if dir != "." && dir != "" {
		segs = append(segs, strings.Split(filepath.ToSlash(dir), "/")...)
	}
	base := filepath.Base(rel)
	stem := strings.TrimSuffix(base, filepath.Ext(base))
	if stem != "index" {
		segs = append(segs, stem)
	}
	return segs
}
