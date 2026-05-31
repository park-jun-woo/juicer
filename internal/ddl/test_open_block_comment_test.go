//ff:func feature=ddl type=test control=sequence
//ff:what openBlockComment /* 블록주석 인라인 종료/잔여/멀티라인 진입 테스트
package ddl

import "testing"

func TestOpenBlockComment(t *testing.T) {
	// no closing */ on line -> enter block mode, drop line
	lines := []string{"/* start", "more", "*/"}
	out, inBlock := openBlockComment(lines, "/* start")
	if !inBlock || len(out) != 2 || out[0] != "more" {
		t.Errorf("multiline: out=%v inBlock=%v", out, inBlock)
	}
	// inline close with no remainder -> drop line, not in block
	lines = []string{"/* x */", "next"}
	out, inBlock = openBlockComment(lines, "/* x */")
	if inBlock || len(out) != 1 || out[0] != "next" {
		t.Errorf("inline empty: out=%v inBlock=%v", out, inBlock)
	}
	// inline close with remainder -> keep remainder on lines[0]
	lines = []string{"/* x */ rest", "next"}
	out, inBlock = openBlockComment(lines, "/* x */ rest")
	if inBlock || out[0] != "rest" {
		t.Errorf("inline rest: out=%v inBlock=%v", out, inBlock)
	}
}
