//ff:type feature=scan type=model
//ff:what caller 인자 해석 결과
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

type callerArgResult struct {
	status     string
	typeName   string
	fields     []scanner.Field
	confidence string
	skip       bool
}
