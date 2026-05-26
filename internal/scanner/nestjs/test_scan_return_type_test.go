//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 반환 타입 스캔 테스트
package nestjs

import "testing"

func TestScan_ReturnType(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Get } from '@nestjs/common';

@Controller('data')
export class DataController {
  @Get()
  getData(): Promise<string> {}
}
`
	writeFile(t, dir, "src/data.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if len(ep.Responses) == 0 {
		t.Fatal("expected responses")
	}
	if ep.Responses[0].TypeName != "string" {
		t.Errorf("response type: want string, got %s", ep.Responses[0].TypeName)
	}
}
