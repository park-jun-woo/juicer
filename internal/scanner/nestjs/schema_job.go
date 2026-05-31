//ff:type feature=scan type=model topic=nestjs
//ff:what 중첩 스키마 해석 작업 큐 항목 구조체
package nestjs

// schemaJob is a single named-type resolution task in the nested-schema BFS.
type schemaJob struct {
	typeName    string
	imports     map[string]string
	referrer    string
	projectRoot string
}
