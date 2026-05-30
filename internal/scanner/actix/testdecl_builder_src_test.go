//ff:type feature=scan type=test topic=actix
//ff:what builderSrc 테스트 보조 선언
package actix

const builderSrc = `
pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::resource("/users")
            .route(web::get().to(list_users))
            .route(web::post().to(create_user)),
    );
}
`
