use std::collections::HashMap;
use serde::Deserialize;
use validator::Validate;

#[derive(Deserialize, Validate)] // 使用 validator 做基础校验
pub struct EventPayload {
    #[validate(length(min = 1, max = 64))]
    event_id: String,
    #[validate(length(min = 1))]
    pub(crate) message: String,
    #[serde(default)]
    pub(crate) tags: HashMap<String, String>,
    pub(crate) platform: String,
    // timestamp: Option<DateTime<Utc>>,
}