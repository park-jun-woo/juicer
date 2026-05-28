//ff:type feature=scan type=model topic=django
//ff:what @action 데코레이터 정보 구조체
package django

// actionInfo holds information about a DRF @action decorator.
type actionInfo struct {
	name    string   // method name
	methods []string // HTTP methods
	detail  bool     // True = /{pk}/action_name/, False = /action_name/
	urlPath string   // custom url_path if specified
	line    int
}
