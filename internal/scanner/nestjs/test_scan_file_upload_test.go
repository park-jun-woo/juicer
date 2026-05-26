//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 파일 업로드 스캔 테스트
package nestjs

import "testing"

func TestScan_FileUpload(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Post, UploadedFile } from '@nestjs/common';

@Controller('upload')
export class UploadController {
  @Post()
  uploadFile(@UploadedFile() file: any) {}
}
`
	writeFile(t, dir, "src/upload.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil || len(ep.Request.Files) != 1 {
		t.Fatalf("expected 1 file param, got %v", ep.Request)
	}
	if ep.Request.Files[0].Name != "file" {
		t.Errorf("file param name: want file, got %s", ep.Request.Files[0].Name)
	}
}
