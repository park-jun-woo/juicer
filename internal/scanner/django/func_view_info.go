//ff:type feature=scan type=model topic=django
//ff:what 함수 기반 뷰 정보 구조체
package django

// funcViewInfo holds information about a DRF function-based view.
type funcViewInfo struct {
	name    string   // function name
	methods []string // HTTP methods from @api_view(["GET", "POST"])
	file    string
	line    int
}
