//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 확장자 없는 모듈 base를 실제 소스 파일로 해석한다 (각 확장자 부착 → index.<ext> → base 그대로 순)
package express

func resolveSourceBase(base string) string {
	if !hasSourceExtension(base) {
		if candidate := firstExistingCandidate(base); candidate != "" {
			return candidate
		}
	}
	if statFile(base) {
		return base
	}
	return ""
}
