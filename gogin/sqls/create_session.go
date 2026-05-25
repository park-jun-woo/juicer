//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what Extract 호출 후 새 세션 생성
package sqls

import (
	"fmt"
)

// createSession runs Extract and initializes a new session.
//
func createSession(repoDir, queriesDir string) error {
	if repoDir == "" {
		return fmt.Errorf("--repo is required for first run")
	}
	if queriesDir == "" {
		return fmt.Errorf("--queries is required for first run")
	}

	fmt.Println("No session found. Analyzing...")

	result, err := Extract(repoDir)
	if err != nil {
		return err
	}

	methods := make([]MethodStatus, len(result.Methods))
	for i, m := range result.Methods {
		methods[i] = MethodStatus{
			ID:     m.Repo + "." + m.Method,
			Status: "TODO",
		}
	}

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: queriesDir,
		Methods:    methods,
	}

	if err := SaveSession(sess); err != nil {
		return err
	}

	fmt.Printf("Found %d methods.\nSession created.\n\n", len(methods))

	if len(methods) > 0 {
		printSkeleton(sess, 0)
	}
	return nil
}

