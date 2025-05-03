use serde::{Deserialize, Serialize};
use chrono::{DateTime, Utc};
use uuid::Uuid;

// 用于 Kafka 消息传递的事件结构
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Event {
    pub event_id: String,
    pub project_id: i32,
    pub message: String,
    pub platform: String,
    pub tags: std::collections::HashMap<String, String>,
    pub timestamp: DateTime<Utc>, // 注意与数据库模型的时间类型匹配
}
