//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_FallbackConfig 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_FallbackConfig(t *testing.T) {
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
	configTs := `
export default registerAs('app', () => ({
  apiPrefix: process.env.API_PREFIX || 'api',
  port: parseInt(process.env.APP_PORT, 10) || 3000,
}));
`
	writeFile(t, dir, "src/config/app.config.ts", configTs)
	prefix := detectGlobalPrefix(dir)
	if prefix != "api" {
		t.Fatalf("expected api, got %q", prefix)
	}
}
