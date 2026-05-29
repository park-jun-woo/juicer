//ff:func feature=ddl type=render control=sequence
//ff:what 테이블 수에 맞춘 zero-pad 파일 prefix 자리수 계산 (최소 2자리)
package ddl

import "strconv"

// prefixWidth returns the digit width needed to zero-pad numeric file prefixes
// for n tables (prefixes range 0..n). Minimum width is 2 (e.g. 00_, 01_) so the
// enum file 00_enums.sql always sorts before table files.
func prefixWidth(n int) int {
	width := len(strconv.Itoa(n))
	if width < 2 {
		width = 2
	}
	return width
}
