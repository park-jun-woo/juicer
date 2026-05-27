//ff:func feature=sql type=parse control=sequence
//ff:what 현재 TODO 항목을 SKIP으로 마킹
package sqls

import (
	"fmt"
)

// RunSkip marks the current TODO item as SKIP.
//
func RunSkip() error {
	if !SessionExists() {
		fmt.Println("No session found. Run \"codist sql next --repo DIR --queries DIR\" first.")
		return nil
	}
	sess, err := LoadSession()
	if err != nil {
		return err
	}

	idx := firstTODO(sess)
	if idx < 0 {
		fmt.Println("No TODO items to skip.")
		return nil
	}

	m := &sess.Methods[idx]
	m.Status = "SKIP"
	if err := SaveSession(sess); err != nil {
		return err
	}
	fmt.Printf("Skipped: %s\n", m.ID)
	return nil
}

