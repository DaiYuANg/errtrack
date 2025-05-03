use crate::app_state::AppState;
use crate::error::AppError;
use crate::event::Event;
use crate::model::event_payload::EventPayload;
use axum::extract::State;
use axum::Json;
use serde::Serialize;
use uuid::Uuid;
use validator::Validate;

#[derive(Serialize)]
pub struct EventResponse {
  pub event_id: String,
  // 可选扩展字段
  // pub message: String,
}

pub async fn ingest_event(
  state: State<AppState>,
  Json(payload): Json<EventPayload>,
) -> Result<Json<EventResponse>, AppError> {
  // 数据校验
  payload.validate()?;

  // 构造事件
  let event = Event {
    event_id: Uuid::new_v4().to_string(),
    project_id: 1, // 简化项目ID获取
    message: payload.message,
    platform: payload.platform,
    tags: payload.tags,
    timestamp: chrono::Utc::now().into(),
  };

  // 发送到Kafka
  // let json = serde_json::to_vec(&event)?;
  // state.producer.send(
  //     FutureRecord::to("events")
  //         .payload(&json)
  //         .key(&event.project_id.to_string()),
  //     Duration::from_secs(3),
  // ).await?;

  Ok(Json(EventResponse {
    event_id: event.event_id,
  }))
}
