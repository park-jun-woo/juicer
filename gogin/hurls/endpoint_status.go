//ff:type feature=hurl type=model
//ff:what EndpointStatus 데이터 구조
package hurls

// EndpointStatus tracks one endpoint's ratchet state.
type EndpointStatus struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	TestFile string `json:"test_file,omitempty"`
}
