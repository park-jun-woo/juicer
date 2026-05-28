//ff:type feature=scan type=model topic=fastify
//ff:what Fastify HTTP 메서드명 상수 및 매핑
package fastify

var httpMethods = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"patch":  "PATCH",
	"delete": "DELETE",
	"all":    "all",
}
