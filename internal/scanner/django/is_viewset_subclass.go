//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 상속체인을 전이적으로 walk하여 ViewSet 서브클래스 여부를 판별한다
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

// isViewSetSubclass reports whether any ancestor (transitively, via the class
// index) is a known DRF ViewSet base class or mixin. A nil index degrades to a
// direct-parent check.
func isViewSetSubclass(parents []string, idx classIndex) bool {
	return hasAncestorIn(parents, viewsetParentNames, idx)
}
