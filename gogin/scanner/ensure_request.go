//ff:func feature=scan type=extract control=sequence
//ff:what ensureRequest 함수
package scanner

func ensureRequest(ep *Endpoint) {
	if ep.Request == nil {
		ep.Request = &Request{}
	}
}
