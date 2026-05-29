//ff:func feature=scan type=extract control=selection topic=actix
//ff:what scoped 타입 식별자를 extractor 종류(path/json/query/form)로 분류한다
package actix

func classifyExtractor(scopedType string) string {
	switch scopedType {
	case "web::Path":
		return "path"
	case "web::Json":
		return "json"
	case "web::Query":
		return "query"
	case "web::Form":
		return "form"
	}
	return ""
}
