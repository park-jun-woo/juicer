//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what base에 각 소스 확장자(base+ext, 이어서 index.<ext>)를 붙여 처음 존재하는 파일 경로를 반환한다
package express

import "path/filepath"

func firstExistingCandidate(base string) string {
	for _, ext := range sourceExtensions {
		if candidate := base + ext; statFile(candidate) {
			return candidate
		}
	}
	for _, ext := range sourceExtensions {
		if candidate := filepath.Join(base, "index"+ext); statFile(candidate) {
			return candidate
		}
	}
	return ""
}
