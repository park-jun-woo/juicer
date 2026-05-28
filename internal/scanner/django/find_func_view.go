//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 이름으로 함수 뷰를 검색한다
package django

// findFuncView finds a function view by name.
func findFuncView(funcViews []funcViewInfo, name string) *funcViewInfo {
	for i := range funcViews {
		if funcViews[i].name == name {
			return &funcViews[i]
		}
	}
	return nil
}
