//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractLambdaParams_And_Request_Round5 테스트
package dotnet

import "testing"

func TestExtractLambdaParams_And_Request_Round5(t *testing.T) {
	fi := csFileInfo(t, `app.MapPost("/u/{id}", (int id, [FromBody] CreateDto dto) => Results.Ok());`)
	inv := findAllByType(fi.root, "invocation_expression")[0]
	args := findChildByType(inv, "argument_list")
	argNodes := childrenOfType(args, "argument")
	req := extractLambdaRequest(argNodes, fi.src, "/u/{id}")
	if req == nil {
		t.Fatal("expected request from lambda")
	}
	if req.Body == nil || req.Body.TypeName != "CreateDto" {
		t.Fatalf("body: %+v", req.Body)
	}
	if len(req.PathParams) != 1 || req.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", req.PathParams)
	}

	if got := extractLambdaRequest(argNodes[:1], fi.src, "/u"); got != nil {
		t.Fatalf("expected nil for single arg, got %+v", got)
	}

	lambda := findChildByType(argNodes[1], "lambda_expression")
	if lambda == nil {
		t.Fatal("no lambda expression")
	}
	lreq := extractLambdaParams(lambda, fi.src, "/u/{id}")
	if lreq == nil || lreq.Body == nil || lreq.Body.TypeName != "CreateDto" {
		t.Fatalf("extractLambdaParams: %+v", lreq)
	}
}
