//ff:type feature=scan type=test topic=express
//ff:what 테스트용 엔드포인트 비교 구조(method/path)
package express

type endpointLike struct{ method, path string }
