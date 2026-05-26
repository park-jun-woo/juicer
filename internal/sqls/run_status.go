//ff:func feature=sql type=parse control=sequence
//ff:what 세션 진행 요약 출력 (TODO/DONE/SKIP 수)
package sqls

import (
	"fmt"
)

// RunStatus prints session progress summary.
//
func RunStatus() error {
	if !SessionExists() {
		fmt.Println("No session found. Run \"juicer sql next --repo DIR --queries DIR\" first.")
		return nil
	}
	sess, err := LoadSession()
	if err != nil {
		return err
	}

	todo := countStatus(sess, "TODO")
	done := countStatus(sess, "DONE")
	skip := countStatus(sess, "SKIP")
	fmt.Printf("TODO: %d, DONE: %d, SKIP: %d\n", todo, done, skip)
	return nil
}

