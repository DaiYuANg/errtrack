mod api_doc;
mod app_state;
mod config;
mod context;
mod error;
mod event;
mod event_handler;
mod flash;
mod model;
mod post_handler;

use crate::app_state::AppState;
use crate::config::ErrTrackConfig;
use crate::event_handler::ingest_event;
use crate::post_handler::{create_post, delete_post, update_post};
use axum::{routing::post, Router};
use axum_example_service::sea_orm::Database;
use migration::{Migrator, MigratorTrait};
use std::env;
use tower_cookies::CookieManagerLayer;

async fn init() -> ErrTrackConfig {
  dotenvy::dotenv().ok();

  tracing_subscriber::fmt()
    .with_max_level(tracing::Level::DEBUG)
    .with_test_writer()
    .init();

  config::load().expect("config load error")
}

async fn configure(err_track_config: ErrTrackConfig) -> AppState {
  let db_url = err_track_config.clone().database.url;

  let conn = Database::connect(db_url)
    .await
    .expect("Database connection failed");

  Migrator::up(&conn, None).await.expect("migrate error");

  AppState {
    conn,
    err_track_config,
  }
}

#[tokio::main]
async fn start_api() -> anyhow::Result<()> {
  let config = init().await;
  let app_state = configure(config.clone()).await;

  let app = Router::new()
    .route("/", post(create_post))
    .route("/{id}", post(update_post))
    .route("/delete/{id}", post(delete_post))
    .route("/event", post(ingest_event))
    .layer(CookieManagerLayer::new())
    .with_state(app_state);

  let listener = tokio::net::TcpListener::bind(config.clone().server.server_url()).await?;
  axum::serve(listener, app).await?;

  Ok(())
}

pub fn main() {
  let result = start_api();

  if let Some(err) = result.err() {
    println!("Error: {err}");
  }
}
