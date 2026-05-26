//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_NoPrefix 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_NoPrefix(t *testing.T) {
	dir := t.TempDir()
	mainTs := `
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTs)
	prefix := detectGlobalPrefix(dir)
	if prefix != "" {
		t.Fatalf("expected empty, got %q", prefix)
	}
}
