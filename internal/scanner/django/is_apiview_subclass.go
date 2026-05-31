//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 상속체인을 전이적으로 walk하여 APIView 서브클래스 여부를 판별한다
package django

// apiViewParentNames is the set of known APIView base classes.
var apiViewParentNames = map[string]bool{
	"APIView":                      true,
	"GenericAPIView":               true,
	"ListAPIView":                  true,
	"CreateAPIView":                true,
	"RetrieveAPIView":              true,
	"DestroyAPIView":               true,
	"UpdateAPIView":                true,
	"ListCreateAPIView":            true,
	"RetrieveUpdateAPIView":        true,
	"RetrieveDestroyAPIView":       true,
	"RetrieveUpdateDestroyAPIView": true,
}

// isAPIViewSubclass reports whether any ancestor (transitively, via the class
// index) is a known DRF APIView base class. A nil index degrades to a
// direct-parent check.
func isAPIViewSubclass(parents []string, idx classIndex) bool {
	return hasAncestorIn(parents, apiViewParentNames, idx)
}
