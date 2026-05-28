//ff:func feature=scan type=test control=sequence
//ff:what package.json 없을 때 fastify 감지 거부 테스트
package scanner

import "testing"

func TestDetectFastify_NoPkg(t *testing.T) {
	dir := t.TempDir()
	if detectFastify(dir) {
		t.Error("expected false without package.json")
	}
}
