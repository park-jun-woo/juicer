//ff:type feature=scan type=model topic=spring
//ff:what Spring Boot 어노테이션 이름 상수
package spring

const (
	AnnRestController = "RestController"
	AnnController     = "Controller"
	AnnResponseBody   = "ResponseBody"
	AnnRequestMapping = "RequestMapping"
	AnnGetMapping     = "GetMapping"
	AnnPostMapping    = "PostMapping"
	AnnPutMapping     = "PutMapping"
	AnnDeleteMapping  = "DeleteMapping"
	AnnPatchMapping   = "PatchMapping"
	AnnPathVariable   = "PathVariable"
	AnnRequestParam   = "RequestParam"
	AnnRequestBody    = "RequestBody"
	AnnRequestPart    = "RequestPart"
	AnnModelAttribute = "ModelAttribute"
	AnnResponseStatus = "ResponseStatus"
	AnnPreAuthorize   = "PreAuthorize"
	AnnSecured        = "Secured"
	AnnRolesAllowed   = "RolesAllowed"
	AnnNotNull        = "NotNull"
	AnnNotBlank       = "NotBlank"
	AnnNotEmpty       = "NotEmpty"
	AnnMin            = "Min"
	AnnMax            = "Max"
	AnnSize           = "Size"
	AnnEmail          = "Email"
	AnnJsonProperty   = "JsonProperty"
	AnnValid          = "Valid"
	AnnValidated      = "Validated"
	AnnRequestHeader  = "RequestHeader"
)

var httpMappingAnnotations = map[string]string{
	AnnGetMapping:    "GET",
	AnnPostMapping:   "POST",
	AnnPutMapping:    "PUT",
	AnnDeleteMapping: "DELETE",
	AnnPatchMapping:  "PATCH",
}

var requestMappingMethods = map[string]string{
	"RequestMethod.GET":    "GET",
	"RequestMethod.POST":   "POST",
	"RequestMethod.PUT":    "PUT",
	"RequestMethod.DELETE": "DELETE",
	"RequestMethod.PATCH":  "PATCH",
	"GET":                  "GET",
	"POST":                 "POST",
	"PUT":                  "PUT",
	"DELETE":               "DELETE",
	"PATCH":                "PATCH",
}

var skipDirs = map[string]bool{
	"build":        true,
	"target":       true,
	".gradle":      true,
	".mvn":         true,
	".git":         true,
	"node_modules": true,
	".idea":        true,
	"out":          true,
}

var httpStatusAnnotations = map[string]string{
	"HttpStatus.OK":                    "200",
	"HttpStatus.CREATED":              "201",
	"HttpStatus.ACCEPTED":            "202",
	"HttpStatus.NO_CONTENT":          "204",
	"HttpStatus.MOVED_PERMANENTLY":   "301",
	"HttpStatus.FOUND":               "302",
	"HttpStatus.BAD_REQUEST":         "400",
	"HttpStatus.UNAUTHORIZED":        "401",
	"HttpStatus.FORBIDDEN":           "403",
	"HttpStatus.NOT_FOUND":           "404",
	"HttpStatus.CONFLICT":            "409",
	"HttpStatus.UNPROCESSABLE_ENTITY": "422",
	"HttpStatus.INTERNAL_SERVER_ERROR": "500",
	"OK":                    "200",
	"CREATED":              "201",
	"ACCEPTED":            "202",
	"NO_CONTENT":          "204",
	"MOVED_PERMANENTLY":   "301",
	"FOUND":               "302",
	"BAD_REQUEST":         "400",
	"UNAUTHORIZED":        "401",
	"FORBIDDEN":           "403",
	"NOT_FOUND":           "404",
	"CONFLICT":            "409",
	"UNPROCESSABLE_ENTITY": "422",
	"INTERNAL_SERVER_ERROR": "500",
}
