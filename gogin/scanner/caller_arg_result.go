//ff:type feature=scan type=model
//ff:what caller 인자 해석 결과
package scanner

type callerArgResult struct {
	status     string
	typeName   string
	fields     []Field
	confidence string
	skip       bool
}
