//ff:func feature=hurl type=session control=sequence
//ff:what 세션 파일 존재 여부 확인
package hurls

import (
	"os"
)

// SessionExists checks whether a hurl session file exists.
func SessionExists() bool {
	_, err := os.Stat(sessionPath())
	return err == nil
}
