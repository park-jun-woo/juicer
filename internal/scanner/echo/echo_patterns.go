//ff:func feature=scan type=extract control=sequence
//ff:what Echo 패키지 경로, HTTP 메서드명, 라우터 타입명 상수
package echo

// echoPkgPaths lists the import paths of Echo (v4 and v5).
var echoPkgPaths = []string{
	"github.com/labstack/echo/v4",
	"github.com/labstack/echo/v5",
}

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
