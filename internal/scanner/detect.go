//ff:func feature=scan type=extract control=sequence
//ff:what 프로젝트 루트의 설정 파일로 프레임워크를 감지한다
package scanner

func DetectFramework(root string) string {
	if detectGoGin(root) {
		return "gogin"
	}
	if detectFiber(root) {
		return "fiber"
	}
	if detectEcho(root) {
		return "echo"
	}
	if detectNestJS(root) {
		return "nestjs"
	}
	if detectFastify(root) {
		return "fastify"
	}
	if detectHono(root) {
		return "hono"
	}
	if detectSupaFunc(root) {
		return "supafunc"
	}
	if detectFastAPI(root) {
		return "fastapi"
	}
	if detectFlask(root) {
		return "flask"
	}
	if detectDjango(root) {
		return "django"
	}
	if detectExpress(root) {
		return "express"
	}
	if detectSpring(root) {
		return "spring"
	}
	if detectQuarkus(root) {
		return "quarkus"
	}
	if detectLaravel(root) {
		return "laravel"
	}
	if detectDotnet(root) {
		return "dotnet"
	}
	if detectActix(root) {
		return "actix"
	}
	return ""
}
