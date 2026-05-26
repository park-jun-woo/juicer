//ff:func feature=scan type=extract control=sequence
//ff:what EnsureRequest 함수
package scanner

func EnsureRequest(ep *Endpoint) {
	if ep.Request == nil {
		ep.Request = &Request{}
	}
}
