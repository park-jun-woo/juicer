//ff:func feature=scan type=command control=sequence
//ff:what TestRunSQL_YAML_Stdout 테스트
package main

import (
	"testing"
)

func TestRunSQL_YAML_Stdout(t *testing.T) {
	// Empty dir — no *_repo.go files, returns empty result
	dir := t.TempDir()
	runSQL([]string{dir})
}
