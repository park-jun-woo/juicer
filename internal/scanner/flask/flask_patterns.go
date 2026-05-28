//ff:func feature=scan type=extract control=selection topic=flask
//ff:what Flask 데코레이터·HTTP 메서드·Blueprint 상수 정의
package flask

// httpMethods maps decorator attribute names to uppercase HTTP methods.
var httpMethods = map[string]string{
	"get":     "GET",
	"post":    "POST",
	"put":     "PUT",
	"delete":  "DELETE",
	"patch":   "PATCH",
	"options": "OPTIONS",
	"head":    "HEAD",
	"route":   "", // handled separately — requires methods arg
}

// shortcutMethods are decorator names that imply a single HTTP method (Flask 2.0+).
var shortcutMethods = map[string]string{
	"get":     "GET",
	"post":    "POST",
	"put":     "PUT",
	"delete":  "DELETE",
	"patch":   "PATCH",
	"options": "OPTIONS",
	"head":    "HEAD",
}

// skipDirs are directory names to skip when walking Python files.
var skipDirs = map[string]bool{
	"venv":          true,
	".venv":         true,
	"__pycache__":   true,
	".git":          true,
	"dist":          true,
	"node_modules":  true,
	".tox":          true,
	".mypy_cache":   true,
	".pytest_cache": true,
}
