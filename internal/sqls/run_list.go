//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 전체 메서드 목록 + 상태 출력
package sqls

import (
	"fmt"
)

// RunList prints all methods with their status.
//
func RunList() error {
	if !SessionExists() {
		fmt.Println("No session found. Run \"codist sql next --repo DIR --queries DIR\" first.")
		return nil
	}
	sess, err := LoadSession()
	if err != nil {
		return err
	}

	for _, m := range sess.Methods {
		fmt.Printf("%s  %s", m.ID, m.Status)
		if m.QueryName != "" {
			fmt.Printf("  (%s)", m.QueryName)
		}
		fmt.Println()
	}
	return nil
}

