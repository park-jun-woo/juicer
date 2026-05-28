//ff:type feature=scan type=model topic=dotnet
//ff:what ASP.NET Core 어트리뷰트 이름 상수
package dotnet

const (
	AttrApiController = "ApiController"
	AttrRoute         = "Route"
	AttrHttpGet       = "HttpGet"
	AttrHttpPost      = "HttpPost"
	AttrHttpPut       = "HttpPut"
	AttrHttpDelete    = "HttpDelete"
	AttrHttpPatch     = "HttpPatch"
	AttrFromBody      = "FromBody"
	AttrFromQuery     = "FromQuery"
	AttrFromRoute     = "FromRoute"
	AttrFromHeader    = "FromHeader"
	AttrFromForm      = "FromForm"
	AttrAuthorize     = "Authorize"
	AttrAllowAnonymous = "AllowAnonymous"
	AttrRequired      = "Required"
	AttrStringLength  = "StringLength"
	AttrMaxLength     = "MaxLength"
	AttrMinLength     = "MinLength"
	AttrRange         = "Range"
	AttrEmailAddress  = "EmailAddress"
	AttrProducesResponseType = "ProducesResponseType"
)

var httpMethodAttributes = map[string]string{
	AttrHttpGet:    "GET",
	AttrHttpPost:   "POST",
	AttrHttpPut:    "PUT",
	AttrHttpDelete: "DELETE",
	AttrHttpPatch:  "PATCH",
}

var mapMethods = map[string]string{
	"MapGet":    "GET",
	"MapPost":   "POST",
	"MapPut":    "PUT",
	"MapDelete": "DELETE",
	"MapPatch":  "PATCH",
}

var skipDirs = map[string]bool{
	"bin":          true,
	"obj":          true,
	".git":         true,
	"node_modules": true,
	"Migrations":   true,
	".vs":          true,
	"packages":     true,
	"TestResults":  true,
}

var diTypes = map[string]bool{
	"ILogger":              true,
	"IConfiguration":       true,
	"IWebHostEnvironment":  true,
	"IHostEnvironment":     true,
	"IServiceProvider":     true,
	"IMemoryCache":         true,
	"IDistributedCache":    true,
	"HttpContext":          true,
	"CancellationToken":   true,
	"ClaimsPrincipal":     true,
	"LinkGenerator":       true,
}

var resultsStatusMethods = map[string]string{
	"Ok":           "200",
	"Created":      "201",
	"Accepted":     "202",
	"NoContent":    "204",
	"BadRequest":   "400",
	"Unauthorized": "401",
	"Forbid":       "403",
	"NotFound":     "404",
	"Conflict":     "409",
	"Problem":      "500",
}
