//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 router.register() 호출을 수집한다
package django

// extractRouterRegistrationsFromFile finds router.register() calls in a single file.
func extractRouterRegistrationsFromFile(fi fileInfo) []routerRegistration {
	var regs []routerRegistration
	for _, callNode := range findAllByType(fi.root, "call") {
		reg := parseRegisterCall(callNode, fi)
		if reg != nil {
			regs = append(regs, *reg)
		}
	}
	return regs
}
