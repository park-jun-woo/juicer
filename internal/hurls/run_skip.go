//ff:func feature=hurl type=parse control=sequence
//ff:what 현재 TODO 항목을 SKIP으로 마킹
package hurls

import (
	"fmt"
)

// RunSkip marks the current TODO item as SKIP.
func RunSkip() error {
	if !SessionExists() {
		fmt.Println("No session found. Run \"juicer hurl next --host URL --tests DIR --repo DIR\" first.")
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

	ep := &sess.Endpoints[idx]
	ep.Status = "SKIP"
	if err := SaveSession(sess); err != nil {
		return err
	}
	fmt.Printf("Skipped: %s\n", ep.ID)
	return nil
}
