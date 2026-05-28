//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 쓰기 메서드에 Serializer body를 추가한다 (APIView용)
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// addWriteMethodBody adds serializer body for write methods (APIView).
func addWriteMethodBody(ep *scanner.Endpoint, method, serializerClass string, serializers map[string]serializerInfo) {
	if serializerClass == "" || !isWriteMethod(method) {
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
