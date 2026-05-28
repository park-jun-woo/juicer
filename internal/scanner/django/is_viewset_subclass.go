//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 부모 클래스 목록에서 ViewSet 서브클래스 여부를 판별한다
package django

// viewsetParentNames is the set of known ViewSet base classes and mixins.
var viewsetParentNames = map[string]bool{
	"ModelViewSet":         true,
	"ReadOnlyModelViewSet": true,
	"ViewSet":              true,
	"GenericViewSet":       true,
	"ViewSetMixin":         true,
	"CreateModelMixin":     true,
	"ListModelMixin":       true,
	"RetrieveModelMixin":   true,
	"UpdateModelMixin":     true,
	"DestroyModelMixin":    true,
}

// isViewSetSubclass checks if any parent is a known ViewSet class or mixin.
func isViewSetSubclass(parents []string) bool {
	for _, p := range parents {
		if viewsetParentNames[p] {
			return true
		}
	}
	return false
}
