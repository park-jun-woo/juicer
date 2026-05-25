//ff:func feature=hurl type=session control=sequence
//ff:what hurl 세션 파일 삭제
package hurls

import (
	"os"
)

// DeleteSession removes the hurl session file.
func DeleteSession() error {
	return os.Remove(sessionPath())
}
