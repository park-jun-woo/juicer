//ff:type feature=hurl type=session
//ff:what Hurl ratchet 세션 상태 컨테이너
package hurls

// Session holds ratchet progress for Hurl endpoint testing.
type Session struct {
	Host      string           `json:"host"`
	TestsDir  string           `json:"tests_dir"`
	RepoDir   string           `json:"repo_dir"`
	Endpoints []EndpointStatus `json:"endpoints"`
}
