//ff:func feature=sql type=parse control=sequence
//ff:what ratchet 한 단계 실행 — 세션 생성 또는 다음 TODO 항목 검증/출력
package sqls

import (
	"fmt"
)

// RunNext executes one ratchet step: create session or verify/present the next TODO item.
//
func RunNext(repoDir, queriesDir string) error {
	if !SessionExists() {
		return createSession(repoDir, queriesDir)
	}

	sess, err := LoadSession()
	if err != nil {
		return err
	}

	// Use session values if flags not provided
	if repoDir == "" {
		repoDir = sess.RepoDir
	}
	if queriesDir == "" {
		queriesDir = sess.QueriesDir
	}

	idx := firstTODO(sess)
	if idx < 0 {
		total := len(sess.Methods)
		fmt.Printf("All queries complete! (%d/%d)\n", total, total)
		return nil
	}

	// Extract once for all printSkeleton calls in this invocation
	result, extractErr := Extract(repoDir)
	var methods []MethodSkeleton
	if extractErr == nil {
		methods = result.Methods
	}

	m := &sess.Methods[idx]
	queryName := toQueryName(m.ID)

	// Check if query exists in queriesDir
	found := queryExists(queriesDir, queryName)
	if !found {
		printSkeleton(sess, idx, methods)
		return nil
	}

	// Query found — run sqlc generate
	fmt.Printf("%s  verifying...\n", m.ID)

	passed, stderr := runSqlcGenerate()
	if passed {
		m.Status = "DONE"
		m.QueryName = queryName
		if err := SaveSession(sess); err != nil {
			return err
		}
		done := countStatus(sess, "DONE")
		total := len(sess.Methods)
		fmt.Printf("  sqlc generate: PASS\n")
		fmt.Printf("DONE ✓  (%d/%d)\n", done, total)

		// Show next TODO if any
		nextIdx := firstTODO(sess)
		if nextIdx >= 0 {
			fmt.Println()
			printSkeleton(sess, nextIdx, methods)
		}
	} else {
		fmt.Printf("  sqlc generate: FAIL\n")
		if stderr != "" {
			fmt.Printf("  %s\n", stderr)
		}
		fmt.Printf("  ▶ Fix the query. Next \"juicer sql next\" will re-verify.\n")
	}
	return nil
}

