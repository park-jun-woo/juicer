//ff:func feature=sql type=parse control=sequence
//ff:what .huma 디렉토리 삭제
package sqls

import (
	"errors"
	"os"
)

// DeleteSession removes the .huma directory.
//
func DeleteSession() error {
	if _, err := os.Stat(sessionDirName); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return os.RemoveAll(sessionDirName)
}

