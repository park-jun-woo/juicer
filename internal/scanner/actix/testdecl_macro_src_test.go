//ff:type feature=scan type=test topic=actix
//ff:what macroSrc 테스트 보조 선언
package actix

const macroSrc = `
#[get("/health")]
async fn health() -> impl Responder { "" }
`
