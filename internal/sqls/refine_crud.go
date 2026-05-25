//ff:func feature=sql type=parse control=sequence
//ff:what SQL 조각 내용으로 INSERT/UPDATE/DELETE 판별
package sqls

import (
	"strings"
)

// refineCRUD determines INSERT/UPDATE/DELETE from SQL fragment content.
//
func refineCRUD(fragments []string) string {
	joined := strings.ToUpper(strings.Join(fragments, " "))
	if strings.Contains(joined, "INSERT") {
		return "INSERT"
	}
	if strings.Contains(joined, "UPDATE") && strings.Contains(joined, "SET") {
		return "UPDATE"
	}
	if strings.Contains(joined, "DELETE") {
		return "DELETE"
	}
	return "EXEC"
}

