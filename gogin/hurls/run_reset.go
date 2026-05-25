//ff:func feature=hurl type=parse control=sequence
//ff:what hurl 세션 삭제
package hurls

import (
	"fmt"
)

// RunReset deletes the hurl session.
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
