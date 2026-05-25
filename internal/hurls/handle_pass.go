//ff:func feature=hurl type=parse control=sequence
//ff:what hurl 테스트 통과 시 세션 업데이트 및 다음 TODO 출력
package hurls

import (
	"fmt"
)

// handlePass marks the endpoint as DONE and shows next TODO.
func handlePass(sess *Session, idx int, testFile, repoDir, testsDir string) error {
	ep := &sess.Endpoints[idx]
	ep.Status = "DONE"
	ep.TestFile = testFile
	if err := SaveSession(sess); err != nil {
		return err
	}
	done := countStatus(sess, "DONE")
	total := len(sess.Endpoints)
	fmt.Printf("  PASS  (%d/%d)\n", done, total)

	nextIdx := firstTODO(sess)
	if nextIdx >= 0 {
		fmt.Println()
		showTODO(&sess.Endpoints[nextIdx], repoDir, testsDir)
	}
	return nil
}
