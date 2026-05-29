//ff:type feature=scan type=model topic=laravel
//ff:what apiResource CRUD 액션 정의(method/suffix/action/hasParam)
package laravel

// apiResourceAction pairs a controller method with its HTTP method and detail flag.
type apiResourceAction struct {
	method   string
	suffix   string // URL suffix: "" for collection, "/{resource}" for detail
	action   string // controller method name
	hasParam bool   // true if the route has a path parameter
}
