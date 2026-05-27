//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what resolveEnumTypeNames 다른 파일 enum 해석 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveEnumTypeNames_CrossFile(t *testing.T) {
	tmp := t.TempDir()
	writeFile(t, tmp, "entities/task.entity.ts", `
export enum TaskStatus { OPEN = 'open', IN_PROGRESS = 'in_progress', DONE = 'done' }
`)
	dtoSrc := []byte(`
import { TaskStatus } from '../entities/task.entity';
class CreateTaskDto {
  @IsEnum(TaskStatus)
  status: TaskStatus;
}
`)
	root, err := parseTypeScript(dtoSrc)
	if err != nil {
		t.Fatal(err)
	}
	dtoPath := filepath.Join(tmp, "dto", "create-task.dto.ts")
	imports := map[string]string{
		"TaskStatus": "../entities/task.entity",
	}
	fields := []dtoField{
		{name: "status", enumTypeName: "TaskStatus"},
	}
	resolveEnumTypeNames(fields, root, dtoSrc, dtoPath, imports, tmp)
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
