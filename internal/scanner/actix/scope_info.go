//ff:type feature=scan type=model topic=actix
//ff:what web::scope 정보(prefix와 등록된 핸들러 이름 목록)
package actix

type scopeInfo struct {
	prefix   string
	handlers []string // handler names registered via .service(handler_name)
}
