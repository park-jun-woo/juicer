//ff:func feature=sql type=parse control=sequence
//ff:what 세션 파일 존재 여부 확인
package sqls

import (
	"os"
)

// SessionExists checks whether a session file exists.
//
func SessionExists() bool {
	_, err := os.Stat(sessionPath())
	return err == nil
}

