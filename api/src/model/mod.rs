pub mod event_payload;

use serde::{Deserialize, Serialize};
use utoipa::ToSchema;

#[derive(Deserialize)]
pub struct Params {
  page: Option<u64>,
  posts_per_page: Option<u64>,
}

#[derive(Deserialize, Serialize, Debug, Clone, ToSchema)]
pub struct FlashData {
  pub kind: String,
  pub message: String,
}
