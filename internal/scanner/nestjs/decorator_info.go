//ff:type feature=scan type=model topic=nestjs
//ff:what 데코레이터 파싱 결과 구조체
package nestjs

// decoratorInfo holds a parsed decorator name and optional argument.
type decoratorInfo struct {
	name string
	arg  string
}
