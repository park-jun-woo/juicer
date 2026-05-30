//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractUsings 테스트
package dotnet

import "testing"

func TestExtractUsings(t *testing.T) {
	root, src := parseCS(t, `using System;
using Microsoft.AspNetCore.Mvc;
namespace App {}`)
	usings := extractUsings(root, src)
	if len(usings) != 2 || usings[0] != "System" {
		t.Fatalf("got %v", usings)
	}
}
