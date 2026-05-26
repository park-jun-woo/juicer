//ff:type feature=scan type=model
//ff:what caller 인자 해석 결과
package gogin

import (
	"github.com/park-jun-woo/juicer/internal/scanner"
)

type callerArgResult struct {
	status     string
	typeName   string
	fields     []scanner.Field
	confidence string
	skip       bool
}
