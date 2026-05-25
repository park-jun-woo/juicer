package hurls

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestPrintSkeleton_Basic(t *testing.T) {
	ep := &scanner.Endpoint{
		Method: "GET",
		Path:   "/api/health",
	}
	printSkeleton(ep, "/tmp/tests")
}

func TestPrintSkeleton_WithResponses(t *testing.T) {
	ep := &scanner.Endpoint{
		Method:    "POST",
		Path:      "/api/users",
		Responses: []scanner.Response{{Status: "200"}, {Status: "400"}},
		Request: &scanner.Request{
			Body: &scanner.Body{TypeName: "CreateUser", Fields: []scanner.Field{{JSON: "name"}, {JSON: "email"}}},
			Query: []scanner.Param{{Name: "page"}},
		},
		Middleware: []string{"auth"},
	}
	printSkeleton(ep, "/tmp/tests")
}

func TestPrintSkeleton_NoMiddleware(t *testing.T) {
	ep := &scanner.Endpoint{
		Method: "GET",
		Path:   "/api/health",
	}
	printSkeleton(ep, "/tmp/tests")
}
