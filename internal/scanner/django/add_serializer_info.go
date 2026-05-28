//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 쓰기 메서드에 Serializer 필드를 엔드포인트 body에 추가한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// addSerializerInfo adds serializer fields to the endpoint for write methods.
func addSerializerInfo(ep *scanner.Endpoint, am actionMethod, serializerClass string, serializers map[string]serializerInfo) {
	if serializerClass == "" || !isWriteMethod(am.method) {
		return
	}
	si, ok := serializers[serializerClass]
	if !ok {
		return
	}
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
	ep.Request.Body = &scanner.Body{
		TypeName: si.name,
		Fields:   si.fields,
	}
}
