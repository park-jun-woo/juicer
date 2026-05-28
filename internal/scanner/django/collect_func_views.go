//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 @api_view 데코레이터 함수 뷰를 수집한다
package django

// collectFuncViews finds all @api_view decorated functions in the parsed files.
func collectFuncViews(files []fileInfo) []funcViewInfo {
	var views []funcViewInfo
	for _, fi := range files {
		views = append(views, collectFuncViewsFromFile(fi)...)
	}
	return views
}
