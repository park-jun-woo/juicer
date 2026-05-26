//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_Found 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_Found(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
import { NestFactory } from '@nestjs/core';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('api');
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	prefix := detectGlobalPrefix(dir)
	if prefix != "api" {
		t.Fatalf("expected api, got %q", prefix)
	}
}
