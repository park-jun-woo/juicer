//ff:func feature=hurl type=session control=sequence
//ff:what 세션 파일 전체 경로 반환
package hurls

import (
	"path/filepath"
)

// sessionPath returns the full path to the hurl session file.
func sessionPath() string {
	return filepath.Join(sessionDirName, sessionFileName)
}
