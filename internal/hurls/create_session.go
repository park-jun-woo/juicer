//ff:func feature=hurl type=session control=sequence
//ff:what scanner.Scan 호출 후 새 hurl 세션 생성
package hurls

import (
	"fmt"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// createSession runs scanner.Scan and initializes a new hurl session.
func createSession(host, testsDir, repoDir string) error {
	if host == "" {
		return fmt.Errorf("--host is required for first run")
	}
	if testsDir == "" {
		return fmt.Errorf("--tests is required for first run")
	}
	if repoDir == "" {
		return fmt.Errorf("--repo is required for first run")
	}

	fmt.Println("No session found. Scanning Go source for endpoints...")

	result, err := scanner.Scan(repoDir)
	if err != nil {
		return err
	}

	sorted := sortEndpoints(result.Endpoints)
	endpoints := buildEndpointStatuses(sorted)

	sess := &Session{
		Host:      host,
		TestsDir:  testsDir,
		RepoDir:   repoDir,
		Endpoints: endpoints,
	}

	if err := SaveSession(sess); err != nil {
		return err
	}

	fmt.Printf("Found %d endpoints.\nSession created.\n\n", len(endpoints))

	if len(sorted) > 0 {
		printSkeleton(&sorted[0], testsDir)
	}
	return nil
}
