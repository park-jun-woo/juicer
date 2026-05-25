//ff:type feature=sql type=session
//ff:what SQL ratchet 세션 상태 컨테이너
package sqls

// Session holds ratchet progress for SQL query migration.
type Session struct {
	RepoDir    string         `json:"repo_dir"`
	QueriesDir string         `json:"queries_dir"`
	Methods    []MethodStatus `json:"methods"`
}
