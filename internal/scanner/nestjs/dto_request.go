//ff:type feature=scan type=model topic=nestjs
//ff:what DTO 해석 요청 구조체
package nestjs

// dtoRequest tracks a DTO type that needs resolution.
type dtoRequest struct {
	typeName string
	imports  map[string]string
	referrer string // absolute file path of the referencing file
	epIdx    int    // index into endpoints slice
	isBody   bool   // true=request body, false=response
}
