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
	DecRoles     = "Roles"
)

// authLevelPatterns maps decorator name keywords to auth levels.
// Checked in order: first match wins.
// "Public" before "Auth" so @ApiPublic → "public", not "auth_required".
// "AuthOptional" before "Auth" so @AuthOptional → "auth_optional", not "auth_required".
var authLevelPatterns = []struct {
	contains string
	level    string
}{
	{"Public", "public"},
	{"AuthOptional", "auth_optional"},
	{"Auth", "auth_required"},
}

// httpMethods maps HTTP decorator names to HTTP method strings.
var httpMethods = map[string]string{
	DecGet:    "GET",
	DecPost:   "POST",
	DecPut:    "PUT",
	DecPatch:  "PATCH",
	DecDelete: "DELETE",
}

// httpStatusMap maps NestJS HttpStatus enum members to numeric codes.
var httpStatusMap = map[string]int{
	"CONTINUE":                        100,
	"SWITCHING_PROTOCOLS":             101,
	"PROCESSING":                      102,
	"OK":                              200,
	"CREATED":                         201,
	"ACCEPTED":                        202,
	"NON_AUTHORITATIVE_INFORMATION":   203,
	"NO_CONTENT":                      204,
	"RESET_CONTENT":                   205,
	"PARTIAL_CONTENT":                 206,
	"MOVED_PERMANENTLY":               301,
	"FOUND":                           302,
	"SEE_OTHER":                       303,
	"NOT_MODIFIED":                    304,
	"TEMPORARY_REDIRECT":              307,
	"PERMANENT_REDIRECT":              308,
	"BAD_REQUEST":                     400,
	"UNAUTHORIZED":                    401,
	"FORBIDDEN":                       403,
	"NOT_FOUND":                       404,
	"METHOD_NOT_ALLOWED":              405,
	"NOT_ACCEPTABLE":                  406,
	"REQUEST_TIMEOUT":                 408,
	"CONFLICT":                        409,
	"GONE":                            410,
	"UNPROCESSABLE_ENTITY":            422,
	"TOO_MANY_REQUESTS":               429,
	"INTERNAL_SERVER_ERROR":           500,
	"NOT_IMPLEMENTED":                 501,
	"BAD_GATEWAY":                     502,
	"SERVICE_UNAVAILABLE":             503,
	"GATEWAY_TIMEOUT":                 504,
}
