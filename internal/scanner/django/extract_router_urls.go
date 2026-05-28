//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 DefaultRouter.register() 호출을 수집한다
package django

// extractRouterRegistrations finds DefaultRouter().register() calls in all files.
func extractRouterRegistrations(files []fileInfo) []routerRegistration {
	var regs []routerRegistration
	for _, fi := range files {
		regs = append(regs, extractRouterRegistrationsFromFile(fi)...)
	}
	return regs
}
