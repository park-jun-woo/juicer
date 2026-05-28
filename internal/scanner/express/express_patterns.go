//ff:type feature=scan type=model topic=express
//ff:what Express HTTP 메서드명 상수 및 매핑
package express

var httpMethods = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"patch":  "PATCH",
	"delete": "DELETE",
	"all":    "all",
}
