//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what NestJS globalPrefix 없는 컨트롤러 스캔 테스트
package nestjs

import "testing"

func TestScan_NoPrefixController(t *testing.T) {
	dir := t.TempDir()

	mainTS := `
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTS)

	ctrlRoot := `
import { Controller, Get } from '@nestjs/common';

@Controller('/')
export class HomeController {
  @Get()
  root() {}
}
`
	writeFile(t, dir, "src/home.controller.ts", ctrlRoot)

	ctrlHealth := `
import { Controller, Get } from '@nestjs/common';

@Controller('health')
export class HealthController {
  @Get()
  check() {}
}
`
	writeFile(t, dir, "src/health.controller.ts", ctrlHealth)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	paths := make(map[string]bool)
	for _, ep := range result.Endpoints {
		paths[ep.Path] = true
	}
	if !paths["/"] {
		t.Error("expected / endpoint")
	}
	if !paths["/health"] {
		t.Error("expected /health endpoint")
	}
}
