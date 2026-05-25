//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what 전체 엔드포인트 목록 + 상태 출력
package hurls

import (
	"fmt"
)

// RunList prints all endpoints with their status.
func RunList() error {
	if !SessionExists() {
		fmt.Println("No session found. Run \"huma hurl next --host URL --tests DIR --repo DIR\" first.")
		return nil
	}
	sess, err := LoadSession()
	if err != nil {
		return err
	}

	for _, ep := range sess.Endpoints {
		fmt.Printf("%s  %s", ep.ID, ep.Status)
		if ep.TestFile != "" {
			fmt.Printf("  (%s)", ep.TestFile)
		}
		fmt.Println()
	}
	return nil
}
