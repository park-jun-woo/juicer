//ff:func feature=scan type=model topic=actix
//ff:what Actix-web 패턴 상수
package actix

var httpMacros = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"delete": "DELETE",
	"patch":  "PATCH",
	"head":   "HEAD",
}

var webMethodBuilders = map[string]string{
	"get":    "GET",
	"post":   "POST",
	"put":    "PUT",
	"delete": "DELETE",
	"patch":  "PATCH",
	"head":   "HEAD",
}

var httpResponseStatuses = map[string]string{
	"Ok":                  "200",
	"Created":             "201",
	"Accepted":            "202",
	"NoContent":           "204",
	"MovedPermanently":    "301",
	"Found":               "302",
	"BadRequest":          "400",
	"Unauthorized":        "401",
	"Forbidden":           "403",
	"NotFound":            "404",
	"Conflict":            "409",
	"UnprocessableEntity": "422",
	"InternalServerError": "500",
}

var skipDirs = map[string]bool{
	"target":       true,
	".git":         true,
	"node_modules": true,
	".idea":        true,
}
