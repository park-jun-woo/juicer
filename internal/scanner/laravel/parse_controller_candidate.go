//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 후보 경로가 존재하고 파싱 가능하면 fileInfo를 반환한다
package laravel

import (
	"os"
)

func parseControllerCandidate(absRoot, candidate string) *fileInfo {
	if _, err := os.Stat(candidate); err != nil {
		return nil
	}
	fi, err := parseFile(absRoot, candidate)
	if err != nil {
		return nil
	}
	return fi
}
