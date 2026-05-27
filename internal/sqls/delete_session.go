//ff:func feature=sql type=parse control=sequence
//ff:what .codist 세션 파일 삭제
package sqls

import (
	"errors"
	"os"
)

// DeleteSession removes the sql session file.
//
func DeleteSession() error {
	err := os.Remove(sessionPath())
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

