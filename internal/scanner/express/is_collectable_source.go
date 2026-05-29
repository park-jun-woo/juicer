//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 파일명이 수집 대상 소스인지 판정한다 (.d.ts 선언 파일 제외 + 알려진 소스 확장자)
package express

import "strings"

// isCollectableSource reports whether a file name is a source file the scanner
// should collect: a known source extension that is not a `.d.ts` declaration.
func isCollectableSource(name string) bool {
	if strings.HasSuffix(name, ".d.ts") {
		return false
	}
	return hasSourceExtension(name)
}
