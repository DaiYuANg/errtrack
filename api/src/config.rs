use figment::providers::{Env, Format, Json, Toml};
use figment::Figment;
use serde::Deserialize;

#[derive(Deserialize, Clone)]
pub struct ErrTrackConfig {
  pub server: ServerConfig,
  pub database: DatabaseConfig,
}

#[derive(Deserialize, Clone)]
pub struct ServerConfig {
  pub port: u16,
  pub host: String,
}

impl ServerConfig {
  pub fn server_url(&self) -> String {
    format!("{}:{}", self.host, self.port)
  }
}

#[derive(Deserialize, Clone)]
pub struct DatabaseConfig {
  pub url: String,
}

pub fn load() -> figment::Result<ErrTrackConfig> {
  Figment::new()
    .merge(Toml::file("errtrack.toml"))
    .merge(Env::prefixed("ERR_TRACK_"))
    .extract_lossy()
}
