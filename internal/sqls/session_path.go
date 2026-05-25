//ff:func feature=sql type=parse control=sequence
//ff:what 세션 파일 전체 경로 반환
package sqls

import (
	"path/filepath"
)

// sessionPath returns the full path to the session file.
//
func sessionPath() string {
	return filepath.Join(sessionDirName, sessionFileName)
}

