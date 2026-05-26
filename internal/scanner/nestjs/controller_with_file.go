//ff:type feature=scan type=model topic=nestjs
//ff:what 컨트롤러 정보와 파일 경로 쌍 구조체
package nestjs

// controllerWithFile pairs a controllerInfo with its absolute file path.
type controllerWithFile struct {
	info    controllerInfo
	absFile string
}
