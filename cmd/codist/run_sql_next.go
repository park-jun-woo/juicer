//ff:func feature=ratchet type=command control=sequence
//ff:what sql next 실행 — ratchet 세션의 다음 메서드로 진행한다
package main

import (
	"github.com/park-jun-woo/codistill/internal/sqls"
)

func runSQLNext(repoDir string, queriesDir string) error {
	return sqls.RunNext(repoDir, queriesDir)
}
