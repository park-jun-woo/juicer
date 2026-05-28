//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::prefix()->group() / Route::middleware()->group() 체인에서 그룹 라우트를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractRouteGroups extracts routes from Route::prefix()->group() and
// Route::middleware()->group() chain calls.
func extractRouteGroups(fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	var routes []routeInfo
	memberCalls := findAllByType(fi.root, "member_call_expression")
	for _, mc := range memberCalls {
		rs := extractOneGroup(mc, fi, outerPrefix, outerMiddleware)
		routes = append(routes, rs...)
	}
	return routes
}

// extractOneGroup handles a single Route::prefix('x')->group(fn) or
// Route::middleware(['auth'])->group(fn) call.
func extractOneGroup(mc *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	// Must have ->group(...) at the end
	groupMethodName := ""
	for i := 0; i < int(mc.ChildCount()); i++ {
		child := mc.Child(i)
		if child.Type() == "name" {
			groupMethodName = nodeText(child, fi.src)
		}
	}
	if groupMethodName != "group" {
		return nil
	}

	// The object part is a scoped_call_expression: Route::prefix('x') or Route::middleware([...])
	scopedCall := findChildByType(mc, "scoped_call_expression")
	if scopedCall == nil {
		// Could be chained: Route::prefix('x')->middleware([...])->group(fn)
		innerMC := findChildByType(mc, "member_call_expression")
		if innerMC != nil {
			return extractChainedGroup(mc, innerMC, fi, outerPrefix, outerMiddleware)
		}
		return nil
	}

	prefix, mw := extractGroupModifier(scopedCall, fi)
	combinedPrefix := joinGroupPrefix(outerPrefix, prefix)
	combinedMW := mergeMiddleware(outerMiddleware, mw)

	// Get the closure body from group(function() { ... })
	groupArgs := findChildByType(mc, "arguments")
	if groupArgs == nil {
		return nil
	}
	closureBody := extractClosureBody(groupArgs, fi)
	if closureBody == nil {
		return nil
	}

	return collectRoutesFromBody(closureBody, fi, combinedPrefix, combinedMW)
}

// extractChainedGroup handles Route::prefix('x')->middleware(['y'])->group(fn) chains.
func extractChainedGroup(outerMC, innerMC *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	prefix := outerPrefix
	mw := make([]string, len(outerMiddleware))
	copy(mw, outerMiddleware)

	// Walk the chain to collect prefix and middleware
	walkChain(innerMC, fi, &prefix, &mw)

	groupArgs := findChildByType(outerMC, "arguments")
	if groupArgs == nil {
		return nil
	}
	closureBody := extractClosureBody(groupArgs, fi)
	if closureBody == nil {
		return nil
	}
	return collectRoutesFromBody(closureBody, fi, prefix, mw)
}

// walkChain walks a member_call_expression or scoped_call_expression chain
// to accumulate prefix and middleware.
func walkChain(node *sitter.Node, fi fileInfo, prefix *string, mw *[]string) {
	if node.Type() == "scoped_call_expression" {
		p, m := extractGroupModifier(node, fi)
		*prefix = joinGroupPrefix(*prefix, p)
		*mw = mergeMiddleware(*mw, m)
		return
	}
	if node.Type() == "member_call_expression" {
		// Get the method name of this call
		methodName := ""
		for i := 0; i < int(node.ChildCount()); i++ {
			child := node.Child(i)
			if child.Type() == "name" {
				methodName = nodeText(child, fi.src)
			}
		}
		// Extract args for this method call
		args := findChildByType(node, "arguments")
		if args != nil && methodName == "middleware" {
			m := extractMiddlewareArgs(args, fi)
			*mw = mergeMiddleware(*mw, m)
		} else if args != nil && methodName == "prefix" {
			argNodes := childrenOfType(args, "argument")
			if len(argNodes) > 0 {
				p := extractStringContent(argNodes[0], fi.src)
				*prefix = joinGroupPrefix(*prefix, p)
			}
		}
		// Walk deeper
		inner := findChildByType(node, "scoped_call_expression")
		if inner != nil {
			walkChain(inner, fi, prefix, mw)
		}
		innerMC := findChildByType(node, "member_call_expression")
		if innerMC != nil {
			walkChain(innerMC, fi, prefix, mw)
		}
	}
}

// extractGroupModifier extracts prefix or middleware from Route::prefix('x') or Route::middleware([...]).
func extractGroupModifier(scopedCall *sitter.Node, fi fileInfo) (string, []string) {
	methodName := ""
	foundScope := false
	for i := 0; i < int(scopedCall.ChildCount()); i++ {
		child := scopedCall.Child(i)
		if child.Type() == "name" {
			if !foundScope {
				foundScope = true
				continue
			}
			methodName = nodeText(child, fi.src)
			break
		}
	}
	args := findChildByType(scopedCall, "arguments")
	if args == nil {
		return "", nil
	}
	switch methodName {
	case "prefix":
		argNodes := childrenOfType(args, "argument")
		if len(argNodes) > 0 {
			return extractStringContent(argNodes[0], fi.src), nil
		}
	case "middleware":
		return "", extractMiddlewareArgs(args, fi)
	}
	return "", nil
}

// extractMiddlewareArgs extracts middleware names from argument list.
func extractMiddlewareArgs(args *sitter.Node, fi fileInfo) []string {
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return nil
	}
	// Could be a single string or an array
	arr := findAllByType(argNodes[0], "array_creation_expression")
	if len(arr) > 0 {
		return extractStringArray(arr[0], fi.src)
	}
	// Single string
	s := extractStringContent(argNodes[0], fi.src)
	if s != "" {
		return []string{s}
	}
	return nil
}

// extractStringArray extracts string values from an array literal.
func extractStringArray(arr *sitter.Node, src []byte) []string {
	var result []string
	elems := childrenOfType(arr, "array_element_initializer")
	for _, elem := range elems {
		s := extractStringContent(elem, src)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// extractClosureBody extracts the compound_statement (body) from a closure argument.
func extractClosureBody(groupArgs *sitter.Node, fi fileInfo) *sitter.Node {
	closures := findAllByType(groupArgs, "anonymous_function_creation_expression")
	if len(closures) == 0 {
		// Try arrow function
		closures = findAllByType(groupArgs, "arrow_function")
	}
	if len(closures) == 0 {
		return nil
	}
	body := findChildByType(closures[0], "compound_statement")
	return body
}

// collectRoutesFromBody extracts routes from inside a closure body.
func collectRoutesFromBody(body *sitter.Node, fi fileInfo, prefix string, middleware []string) []routeInfo {
	var routes []routeInfo
	// Create a temporary fileInfo with the body as root for targeted search
	bodyFI := fileInfo{
		absPath: fi.absPath,
		relPath: fi.relPath,
		src:     fi.src,
		root:    body,
	}
	routes = append(routes, collectRoutes(bodyFI, prefix, middleware)...)
	routes = append(routes, collectAPIResource(bodyFI, prefix, middleware)...)

	// Recursively handle nested groups
	memberCalls := findAllByType(body, "member_call_expression")
	for _, mc := range memberCalls {
		rs := extractOneGroup(mc, bodyFI, prefix, middleware)
		routes = append(routes, rs...)
	}

	return routes
}

// joinGroupPrefix joins outer prefix with inner prefix.
func joinGroupPrefix(outer, inner string) string {
	if outer == "" {
		return inner
	}
	if inner == "" {
		return outer
	}
	return outer + "/" + inner
}

// mergeMiddleware merges two middleware slices, avoiding duplicates.
func mergeMiddleware(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	seen := make(map[string]bool, len(a))
	for _, m := range a {
		seen[m] = true
	}
	result := make([]string, len(a))
	copy(result, a)
	for _, m := range b {
		if !seen[m] {
			result = append(result, m)
		}
	}
	return result
}
