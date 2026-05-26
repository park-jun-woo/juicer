//ff:type feature=scan type=model topic=fastapi
//ff:what Pydantic 모델 해석 요청 구조체
package fastapi

// modelRequest tracks a Pydantic model type that needs resolution.
type modelRequest struct {
	typeName string
	imports  []importInfo
	referrer string // absolute file path of the referencing file
	epIdx    int    // index into endpoints slice
	isBody   bool   // true=request body, false=response
}
