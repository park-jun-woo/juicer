//ff:type feature=scan type=model topic=flask
//ff:what Blueprint 인스턴스 정보 구조체
package flask

// blueprintInfo holds information about a Blueprint instance.
type blueprintInfo struct {
	varName   string // variable name, e.g. "users_bp"
	name      string // blueprint name, e.g. "users"
	urlPrefix string // url_prefix from constructor
}
