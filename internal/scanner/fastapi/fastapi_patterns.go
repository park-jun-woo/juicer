//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what FastAPI 데코레이터·클래스·HTTP 메서드 상수 정의
package fastapi

// httpMethods maps decorator attribute names to uppercase HTTP methods.
var httpMethods = map[string]string{
	"get":     "GET",
	"post":    "POST",
	"put":     "PUT",
	"delete":  "DELETE",
	"patch":   "PATCH",
	"options": "OPTIONS",
	"head":    "HEAD",
}

// routerClassNames are the class names that create router instances.
var routerClassNames = map[string]bool{
	"FastAPI":   true,
	"APIRouter": true,
}

// dependsKeywords are substrings in function names that indicate auth middleware.
var dependsKeywords = []string{
	"auth",
	"current_user",
	"get_current",
}

// specialDefaults are FastAPI default function calls that classify parameters.
var specialDefaults = map[string]string{
	"Query":   "query",
	"Body":    "body",
	"File":    "file",
	"Header":  "header",
	"Depends": "depends",
}

// uploadFileTypes are type names that indicate file upload parameters.
var uploadFileTypes = map[string]bool{
	"UploadFile": true,
}

// skipDirs are directory names to skip when walking Python files.
var skipDirs = map[string]bool{
	"venv":        true,
	".venv":       true,
	"__pycache__": true,
	".git":        true,
	"dist":        true,
	"node_modules": true,
	".tox":        true,
	".mypy_cache": true,
	".pytest_cache": true,
}
