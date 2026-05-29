//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_NestJSFramework NestJS 프레임워크 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_NestJSFramework(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{"@nestjs/core":"*"}}`), 0o644)
	srcDir := filepath.Join(dir, "src")
	os.MkdirAll(srcDir, 0o755)
	os.WriteFile(filepath.Join(srcDir, "app.controller.ts"), []byte(`import { Controller, Get } from '@nestjs/common';
@Controller('test')
export class AppController {
  @Get()
  getHello(): string { return 'hello'; }
}
`), 0o644)
	execScan([]string{"--framework", "nestjs", dir})
}
