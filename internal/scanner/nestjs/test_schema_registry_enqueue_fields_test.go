//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what enqueueFields Ref 있는 필드만 작업 큐에 추가 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestSchemaRegistryEnqueueFields(t *testing.T) {
	r := &schemaRegistry{}
	fields := []scanner.Field{
		{Name: "a", Type: "string"},
		{Name: "b", Ref: "ChildDto"},
		{Name: "c", Ref: "OtherDto"},
	}
	r.enqueueFields(fields, map[string]string{"ChildDto": "./child"}, "/ref.ts", "/root")
	if len(r.queue) != 2 {
		t.Fatalf("want 2 jobs, got %d", len(r.queue))
	}
	if r.queue[0].typeName != "ChildDto" || r.queue[1].typeName != "OtherDto" {
		t.Errorf("queue: %+v", r.queue)
	}
}
