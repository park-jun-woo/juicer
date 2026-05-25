//ff:func feature=sql type=parse control=sequence
//ff:what 세션 삭제
package sqls

import (
	"fmt"
)

// RunReset deletes the session.
//
func RunReset() error {
	if !SessionExists() {
		fmt.Println("No session found.")
		return nil
	}
	if err := DeleteSession(); err != nil {
		return err
	}
	fmt.Println("Session deleted.")
	return nil
}

