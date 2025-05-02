mod flash;
mod context;

use axum::{
  extract::{Form, Path, State},
  http::StatusCode,
  routing::post,
  Router,
};
use axum_example_service::{
  sea_orm::{Database, DatabaseConnection},
  Mutation as MutationCore,
};
use entity::post;
use flash::{post_response, PostResponse};
use migration::{Migrator, MigratorTrait};
use serde::{Deserialize, Serialize};
use std::env;
use tower_cookies::{CookieManagerLayer, Cookies};
use utoipa::{Modify, OpenApi, ToSchema};
use utoipa::openapi::security::{ApiKey, ApiKeyValue, SecurityScheme};

const TODO_TAG: &str = "todo";
#[derive(OpenApi)]
#[openapi(
  modifiers(&SecurityAddon),
  tags(
            (name = TODO_TAG, description = "Todo items management API")
  )
)]
struct ApiDoc;

struct SecurityAddon;

impl Modify for SecurityAddon {
  fn modify(&self, openapi: &mut utoipa::openapi::OpenApi) {
    if let Some(components) = openapi.components.as_mut() {
      components.add_security_scheme(
        "api_key",
        SecurityScheme::ApiKey(ApiKey::Header(ApiKeyValue::new("todo_apikey"))),
      )
    }
  }
}

#[tokio::main]
async fn start() -> anyhow::Result<()> {
  env::set_var("RUST_LOG", "debug");
  tracing_subscriber::fmt()
      .with_max_level(tracing::Level::DEBUG)
      .with_test_writer()
      .init();

  dotenvy::dotenv().ok();
  let db_url = env::var("DATABASE_URL").expect("DATABASE_URL is not set in .env file");
  let host = env::var("HOST").expect("HOST is not set in .env file");
  let port = env::var("PORT").expect("PORT is not set in .env file");
  let server_url = format!("{host}:{port}");

  let conn = Database::connect(db_url)
    .await
    .expect("Database connection failed");
  Migrator::up(&conn, None).await?;

  let state = AppState { conn };

  let app = Router::new()
    .route("/", post(create_post))
    .route("/{id}", post(update_post))
    .route("/delete/{id}", post(delete_post))
    .layer(CookieManagerLayer::new())
    .with_state(state);

  let listener = tokio::net::TcpListener::bind(&server_url).await?;
  axum::serve(listener, app).await?;

  Ok(())
}

#[derive(Clone)]
struct AppState {
  conn: DatabaseConnection,
}

#[derive(Deserialize)]
struct Params {
  page: Option<u64>,
  posts_per_page: Option<u64>,
}

#[derive(Deserialize, Serialize, Debug, Clone, ToSchema)]
struct FlashData {
  kind: String,
  message: String,
}

async fn create_post(
  state: State<AppState>,
  mut cookies: Cookies,
  form: Form<post::Model>,
) -> Result<PostResponse, (StatusCode, &'static str)> {
  let form = form.0;

  MutationCore::create_post(&state.conn, form)
    .await
    .expect("could not insert post");

  let data = FlashData {
    kind: "success".to_owned(),
    message: "Post successfully added".to_owned(),
  };

  Ok(post_response(&mut cookies, data))
}

async fn update_post(
  state: State<AppState>,
  Path(id): Path<i32>,
  mut cookies: Cookies,
  form: Form<post::Model>,
) -> Result<PostResponse, (StatusCode, String)> {
  let form = form.0;

  MutationCore::update_post_by_id(&state.conn, id, form)
    .await
    .expect("could not edit post");

  let data = FlashData {
    kind: "success".to_owned(),
    message: "Post successfully updated".to_owned(),
  };

  Ok(post_response(&mut cookies, data))
}

async fn delete_post(
  state: State<AppState>,
  Path(id): Path<i32>,
  mut cookies: Cookies,
) -> Result<PostResponse, (StatusCode, &'static str)> {
  MutationCore::delete_post(&state.conn, id)
    .await
    .expect("could not delete post");

  let data = FlashData {
    kind: "success".to_owned(),
    message: "Post successfully deleted".to_owned(),
  };

  Ok(post_response(&mut cookies, data))
}

pub fn main() {
  let result = start();

  if let Some(err) = result.err() {
    println!("Error: {err}");
  }
}
