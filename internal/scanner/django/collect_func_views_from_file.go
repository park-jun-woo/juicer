//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 @api_view 데코레이터 함수 뷰를 수집한다
package django

// collectFuncViewsFromFile finds @api_view decorated functions in a single file.
func collectFuncViewsFromFile(fi fileInfo) []funcViewInfo {
	var views []funcViewInfo
	for _, funcDef := range findAllByType(fi.root, "function_definition") {
		fv := parseFuncView(funcDef, fi)
		if fv != nil {
			views = append(views, *fv)
		}
	}
	return views
}
