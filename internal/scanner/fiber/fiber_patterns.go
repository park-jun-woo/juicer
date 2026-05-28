//ff:type feature=scan type=model
//ff:what Fiber 패키지 경로, HTTP 메서드명, 라우터 타입, 핸들러 변수 상수
package fiber

// fiberPkgPath is the import path of the Fiber v2 package.
const fiberPkgPath = "github.com/gofiber/fiber/v2"

// fiberMethods maps Fiber HTTP method names to true.
// Fiber uses capitalized first letter: Get, Post, Put, Patch, Delete, Head, Options.
var fiberMethods = map[string]bool{
	"Get": true, "Post": true, "Put": true, "Patch": true, "Delete": true,
	"Head": true, "Options": true, "All": true,
}

// fiberMethodToHTTP maps Fiber method names to uppercase HTTP method names.
var fiberMethodToHTTP = map[string]string{
	"Get": "GET", "Post": "POST", "Put": "PUT", "Patch": "PATCH", "Delete": "DELETE",
	"Head": "HEAD", "Options": "OPTIONS", "All": "ALL",
}

// fiberRouterTypes maps Fiber router type names.
var fiberRouterTypes = map[string]bool{
	"App":    true,
	"Group":  true,
	"Router": true,
}

// bindMethods maps Fiber body binding methods.
var bindMethods = map[string]bool{
	"BodyParser": true,
}

// queryMethods maps Fiber query parameter methods.
var queryMethods = map[string]bool{
	"Query": true,
}

// paramMethods maps Fiber path parameter methods.
var paramMethods = map[string]bool{
	"Params": true,
}

// formMethods maps Fiber form value methods.
var formMethods = map[string]bool{
	"FormValue": true,
}

// fileMethods maps Fiber file upload methods.
var fileMethods = map[string]bool{
	"FormFile": true,
}

// rawBodyMethods maps Fiber raw body methods.
var rawBodyMethods = map[string]bool{
	"Body": true,
}

// responseMethods maps Fiber response methods to their kind.
var responseMethods = map[string]string{
	"JSON":       "json",
	"SendString": "string",
	"Send":       "data",
	"SendFile":   "file",
	"Redirect":   "redirect",
	"Status":     "status",
	"SendStatus": "status",
}
