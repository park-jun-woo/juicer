//ff:type feature=scan type=model topic=express
//ff:what 라우터 인스턴스 식별자 (파일 경로 + 로컬 변수명)
package express

// routerKey는 라우터 인스턴스 하나를 식별한다.
// 같은 파일이 여러 prefix로 마운트되거나, 한 파일에 라우터가 여럿 있어도
// (file, varName) 쌍으로 구분되므로 prefix가 붕괴하지 않는다.
type routerKey struct {
	file    string
	varName string
}
