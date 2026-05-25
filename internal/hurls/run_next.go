//ff:func feature=hurl type=parse control=sequence
//ff:what ratchet 한 단계 실행 — 세션 생성 또는 다음 TODO 항목 검증/출력
package hurls

import (
	"fmt"
)

// RunNext executes one ratchet step: create session or verify/present the next TODO item.
func RunNext(host, testsDir, repoDir string) error {
	if !SessionExists() {
		return createSession(host, testsDir, repoDir)
	}

	sess, err := LoadSession()
	if err != nil {
		return err
	}

	if host == "" {
		host = sess.Host
	}
	if testsDir == "" {
		testsDir = sess.TestsDir
	}
	if repoDir == "" {
		repoDir = sess.RepoDir
	}

	idx := firstTODO(sess)
	if idx < 0 {
		total := len(sess.Endpoints)
		fmt.Printf("All tests complete! (%d/%d)\n", total, total)
		return nil
	}

	ep := &sess.Endpoints[idx]
	method, path := parseEndpointID(ep.ID)

	testFile := findTestFile(testsDir, method, path)
	if testFile == "" {
		showTODO(ep, repoDir, testsDir)
		return nil
	}

	fmt.Printf("%s  verifying...\n", ep.ID)
	fmt.Printf("  hurl --test %s --variable host=%s\n", testFile, host)

	passed, stderr := runHurlTest(testFile, host)
	if passed {
		return handlePass(sess, idx, testFile, repoDir, testsDir)
	}
	handleFail(stderr)
	return nil
}
