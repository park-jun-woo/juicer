//ff:type feature=scan type=model topic=fastapi
//ff:what include_router 호출 정보 구조체
package fastapi

// includeCall represents an app.include_router(router, prefix=...) call.
type includeCall struct {
	parentVar   string // e.g., "app"
	childVar    string // e.g., "router"
	extraPrefix string // prefix kwarg from include_router call
}
