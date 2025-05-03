use axum_example_service::sea_orm::DatabaseConnection;
use crate::config::ErrTrackConfig;

#[derive(Clone)]
pub struct AppState {
  pub conn: DatabaseConnection,
  pub err_track_config: ErrTrackConfig
}
