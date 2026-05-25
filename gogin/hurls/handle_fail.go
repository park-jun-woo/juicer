//ff:func feature=hurl type=render control=sequence
//ff:what hurl 테스트 실패 시 에러 출력
package hurls

import (
	"fmt"
)

// handleFail prints the failure message and stderr output.
func handleFail(stderr string) {
	fmt.Printf("  FAIL\n")
	if stderr != "" {
		fmt.Printf("  %s\n", stderr)
	}
	fmt.Printf("  > Fix the test. Next \"huma hurl next\" will re-verify.\n")
}
