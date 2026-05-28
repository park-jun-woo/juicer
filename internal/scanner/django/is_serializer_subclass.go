//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 부모 클래스 목록에서 Serializer 서브클래스 여부를 판별한다
package django

// serializerParentNames is the set of known Serializer base classes.
var serializerParentNames = map[string]bool{
	"Serializer":                  true,
	"ModelSerializer":             true,
	"HyperlinkedModelSerializer":  true,
	"ListSerializer":              true,
	"BaseSerializer":              true,
}

// isSerializerSubclass checks if any parent is a known Serializer base class.
func isSerializerSubclass(parents []string) bool {
	for _, p := range parents {
		if serializerParentNames[p] {
			return true
		}
	}
	return false
}
