//ff:func feature=hurl type=render control=sequence
//ff:what TODO 엔드포인트의 스켈레톤 정보 출력
package hurls

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/juicer/scanner"
)

// printSkeleton outputs skeleton info for an endpoint from scan data.
func printSkeleton(ep *scanner.Endpoint, testsDir string) {
	fmt.Printf("%s %s  TODO\n", ep.Method, ep.Path)

	// Response status codes
	if len(ep.Responses) > 0 {
		var codes []string
		for _, r := range ep.Responses {
			codes = append(codes, r.Status)
		}
		fmt.Printf("  responses: %s\n", formatSlice(codes))
	}

	// Auth
	if len(ep.Middleware) == 0 {
		fmt.Println("  auth: none")
	} else {
		fmt.Printf("  auth: required\n")
		fmt.Printf("  middleware: %s\n", formatSlice(ep.Middleware))
	}

	// Request body
	if ep.Request != nil && ep.Request.Body != nil {
		body := ep.Request.Body
		typeName := body.TypeName
		if typeName == "" {
			typeName = "object"
		}
		var fieldNames []string
		for _, f := range body.Fields {
			fieldNames = append(fieldNames, f.JSON)
		}
		if len(fieldNames) > 0 {
			fmt.Printf("  request_body: %s {%s}\n", typeName, strings.Join(fieldNames, ", "))
		} else {
			fmt.Printf("  request_body: %s\n", typeName)
		}
	}

	// Query params
	if ep.Request != nil && len(ep.Request.Query) > 0 {
		var names []string
		for _, q := range ep.Request.Query {
			names = append(names, q.Name)
		}
		fmt.Printf("  query_params: %s\n", formatSlice(names))
	}

	// Suggested filename
	suggested := suggestFilename(ep.Path)
	fmt.Printf("  -> Write test to %s%s\n", testsDir, suggested)

	// Warn about write endpoints needing cleanup
	if ep.Method == "POST" {
		fmt.Printf("  ! Write endpoint -- pair with DELETE for cleanup\n")
	}
}
