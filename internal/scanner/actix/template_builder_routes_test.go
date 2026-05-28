//ff:type feature=scan type=test topic=actix
//ff:what 빌더 패턴 라우트 테스트용 소스 템플릿
package actix

var builderRoutesSource = `
use actix_web::{web, App, HttpServer, HttpResponse};

async fn list_users() -> HttpResponse {
    HttpResponse::Ok().json(users)
}

async fn get_user(path: web::Path<i64>) -> HttpResponse {
    HttpResponse::Ok().json(user)
}

async fn create_user(body: web::Json<CreateUserRequest>) -> HttpResponse {
    HttpResponse::Created().json(user)
}

fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/api")
            .service(
                web::resource("/users")
                    .route(web::get().to(list_users))
                    .route(web::post().to(create_user))
            )
            .service(
                web::resource("/users/{id}")
                    .route(web::get().to(get_user))
            )
    );
}
`

var builderCreateUserStruct = `
use serde::Deserialize;

#[derive(Deserialize)]
struct CreateUserRequest {
    name: String,
    email: String,
}
`
