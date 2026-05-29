//ff:func feature=scan type=extract control=selection topic=express
//ff:what 마운트 엔트리 하나를 인라인/크로스파일/모호 케이스로 분기해 엣지를 추가한다
package express

func addMountEdges(g *mountGraph, m mountEntry, allRouters map[string]map[string]bool) {
	parent := routerKey{m.sourceFile, m.sourceRouter}
	switch {
	case m.filePath == "":
		if m.varName != "" {
			graphAddEdge(g, parent, routerKey{m.sourceFile, m.varName}, m.prefix)
		}
	case resolveChildVar(m.filePath, m.varName, allRouters) != "":
		cv := resolveChildVar(m.filePath, m.varName, allRouters)
		graphAddEdge(g, parent, routerKey{m.filePath, cv}, m.prefix)
	default:
		addAmbiguousMountEdges(g, parent, m, allRouters)
	}
}
