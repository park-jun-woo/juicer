//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractRecordParams 테스트
package dotnet

import "testing"

func TestExtractRecordParams(t *testing.T) {
	root, src := parseCS(t, `public record CreateUserDto(string Name, int Age);`)
	recs := findAllByType(root, "record_declaration")
	if len(recs) == 0 {
		t.Skip("no record")
	}
	fields := extractRecordParams(recs[0], src)
	if len(fields) != 2 || fields[0].Name != "Name" {
		t.Fatalf("got %+v", fields)
	}
	if fields[0].Type != "string" || fields[1].Type != "integer" {
		t.Fatalf("types: %+v", fields)
	}
}
