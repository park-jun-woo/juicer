//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 부모 클래스 목록에서 APIView 서브클래스 여부를 판별한다
package django

// apiViewParentNames is the set of known APIView base classes.
var apiViewParentNames = map[string]bool{
	"APIView":                       true,
	"GenericAPIView":                true,
	"ListAPIView":                   true,
	"CreateAPIView":                 true,
	"RetrieveAPIView":               true,
	"DestroyAPIView":                true,
	"UpdateAPIView":                 true,
	"ListCreateAPIView":             true,
	"RetrieveUpdateAPIView":         true,
	"RetrieveDestroyAPIView":        true,
	"RetrieveUpdateDestroyAPIView":  true,
}

// isAPIViewSubclass checks if any parent is a known APIView base class.
func isAPIViewSubclass(parents []string) bool {
	for _, p := range parents {
		if apiViewParentNames[p] {
			return true
		}
	}
	return false
}
