//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_FallbackEnv 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_FallbackEnv(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix(configService.getOrThrow('app.apiPrefix', { infer: true }));
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	writeFile(t, dir, ".env.example", "API_PREFIX=api\nDATABASE_HOST=localhost\n")
	prefix := detectGlobalPrefix(dir)
	if prefix != "api" {
		t.Fatalf("expected api, got %q", prefix)
	}
}
