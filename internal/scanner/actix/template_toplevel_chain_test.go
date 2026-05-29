//ff:type feature=scan type=test topic=actix
//ff:what 최상위 .service() 체인 + .to()-only + scoped/generic/closure 핸들러 테스트용 소스 템플릿
package actix

var topLevelChainSource = `
use actix_web::{web, App, HttpResponse};

async fn alpha() -> HttpResponse { HttpResponse::Ok().finish() }
async fn beta() -> HttpResponse { HttpResponse::Ok().finish() }
async fn gamma() -> HttpResponse { HttpResponse::Ok().finish() }
async fn delta() -> HttpResponse { HttpResponse::Ok().finish() }
async fn epsilon() -> HttpResponse { HttpResponse::Ok().finish() }

fn main() {
    HttpServer::new(|| {
        App::new()
            .service(web::resource("/delta").route(web::get().to(delta)))
            .service(web::resource("/epsilon").route(web::post().to(epsilon)))
    });
}

fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(web::resource("/alpha").route(web::get().to(alpha)))
       .service(web::resource("/beta").route(web::post().to(beta)))
       .service(web::resource("/gamma").route(web::get().to(gamma)));
    cfg.service(web::resource("/about").to(view::about));
    cfg.service(web::resource("/search").route(web::get().to(api::search::<String>)));
    cfg.service(web::resource("/inline").route(web::get().to(|| async { "x" })));
}
`
