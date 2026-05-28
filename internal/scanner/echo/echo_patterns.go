//ff:func feature=scan type=extract control=sequence
//ff:what Echo 패키지 경로, HTTP 메서드명, 라우터 타입명 상수
package echo

// echoPkgPath is the import path of Echo v4.
const echoPkgPath = "github.com/labstack/echo/v4"

// echoMethods maps HTTP method registration names to true.
var echoMethods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true,
	"HEAD": true, "OPTIONS": true, "Any": true,
}

// echoRouterTypes lists Echo types that can register routes.
var echoRouterTypes = map[string]bool{
	"Echo":  true,
	"Group": true,
}
