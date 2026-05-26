//ff:type feature=scan type=model topic=nestjs
//ff:what NestJS 데코레이터 이름 상수
package nestjs

// HTTP method decorators
const (
	DecController   = "Controller"
	DecGet          = "Get"
	DecPost         = "Post"
	DecPut          = "Put"
	DecPatch        = "Patch"
	DecDelete       = "Delete"
)

// Parameter decorators
const (
	DecParam        = "Param"
	DecQuery        = "Query"
	DecBody         = "Body"
	DecUploadedFile = "UploadedFile"
)

// Other decorators
const (
	DecHttpCode  = "HttpCode"
	DecUseGuards = "UseGuards"
)

// httpMethods maps HTTP decorator names to HTTP method strings.
var httpMethods = map[string]string{
	DecGet:    "GET",
	DecPost:   "POST",
	DecPut:    "PUT",
	DecPatch:  "PATCH",
	DecDelete: "DELETE",
}
