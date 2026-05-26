//ff:func feature=scan type=extract control=sequence
//ff:what 프로젝트 루트의 설정 파일로 프레임워크를 감지한다
package scanner

func DetectFramework(root string) string {
	if detectGoGin(root) {
		return "gogin"
	}
	if detectNestJS(root) {
		return "nestjs"
	}
	if detectFastAPI(root) {
		return "fastapi"
	}
	return ""
}
