//ff:type feature=scan type=model topic=django
//ff:what Serializer 정보 구조체
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// serializerInfo holds information about a DRF Serializer class.
type serializerInfo struct {
	name   string
	fields []scanner.Field
}
