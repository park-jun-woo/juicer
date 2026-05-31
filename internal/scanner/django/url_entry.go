//ff:type feature=scan type=model topic=django
//ff:what URL 패턴 엔트리 구조체
package django

// urlEntry represents a single path() entry from urls.py.
type urlEntry struct {
	pattern       string // URL pattern e.g. "api/users/<int:pk>/"
	viewName      string // view reference e.g. "UserViewSet" or "views.health_check"
	isInclude     bool   // whether the second arg is include(...)
	includeModule string // module path for include("app.urls")
	// methodActions maps an HTTP method (lowercase, e.g. "get") to its ViewSet
	// action (e.g. "list") as declared in as_view({"get": "list", ...}).
	methodActions map[string]string
}
