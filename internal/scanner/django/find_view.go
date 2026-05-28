//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 이름으로 ViewSet, APIView, FuncView를 검색한다
package django

// findViewSet finds a ViewSet by name.
func findViewSet(viewsets []viewsetInfo, name string) *viewsetInfo {
	for i := range viewsets {
		if viewsets[i].name == name {
			return &viewsets[i]
		}
	}
	return nil
}
