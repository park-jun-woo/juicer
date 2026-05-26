//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 전역 접두사 스캔 테스트
package nestjs

import "testing"

func TestScan_GlobalPrefix(t *testing.T) {
	dir := t.TempDir()

	mainTS := `
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('api/v1');
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTS)

	ctrl := `
import { Controller, Get } from '@nestjs/common';

@Controller('health')
export class HealthController {
  @Get()
  check() {}
}
`
	writeFile(t, dir, "src/health.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	if result.Endpoints[0].Path != "/api/v1/health" {
		t.Errorf("path: want /api/v1/health, got %s", result.Endpoints[0].Path)
	}
}
