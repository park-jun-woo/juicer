//ff:type feature=scan type=model topic=hono
//ff:what Hono HTTP 메서드명 상수 및 매핑
package hono

var httpMethods = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"patch":  "PATCH",
	"delete": "DELETE",
	"all":    "all",
}
