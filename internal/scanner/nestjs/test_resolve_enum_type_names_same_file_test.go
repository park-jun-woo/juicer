//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what resolveEnumTypeNames 같은 파일 내 enum 해석 테스트
package nestjs

import "testing"

func TestResolveEnumTypeNames_SameFile(t *testing.T) {
	src := []byte(`
enum TaskStatus { OPEN = 'open', IN_PROGRESS = 'in_progress', DONE = 'done' }
class CreateTaskDto {
  @IsEnum(TaskStatus)
  status: TaskStatus;
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := []dtoField{
		{name: "status", enumTypeName: "TaskStatus"},
	}
	resolveEnumTypeNames(fields, root, src, "/fake/dto.ts", nil, "")
	if len(fields[0].enum) != 3 {
		t.Fatalf("expected 3 enum values, got %d: %v", len(fields[0].enum), fields[0].enum)
	}
	tests := []struct {
		index int
		want  string
	}{
		{0, "open"},
		{1, "in_progress"},
		{2, "done"},
	}
	for _, tt := range tests {
		if fields[0].enum[tt.index] != tt.want {
			t.Errorf("enum[%d] = %q, want %q", tt.index, fields[0].enum[tt.index], tt.want)
		}
	}
}
