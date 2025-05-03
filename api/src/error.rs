use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
};
use serde_json::json;
use std::fmt;
use axum_example_service::rdkafka;
use migration::sea_orm;

// 自定义错误类型（覆盖常见错误场景）
#[derive(Debug)]
pub enum AppError {
    ValidationError(String),    // 数据校验失败
    DatabaseError(sea_orm::DbErr),  // 数据库错误
    KafkaError(rdkafka::error::KafkaError), // Kafka 错误
    Unauthorized,               // 认证失败
    InternalServerError,        // 其他内部错误
}

// 错误类型转 HTTP 响应
impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, message) = match self {
            AppError::ValidationError(err) => (StatusCode::BAD_REQUEST, err),
            AppError::DatabaseError(_) => (
                StatusCode::INTERNAL_SERVER_ERROR,
                "Database error".to_string(),
            ),
            AppError::KafkaError(_) => (
                StatusCode::SERVICE_UNAVAILABLE,
                "Event processing backlogged".to_string(),
            ),
            AppError::Unauthorized => (
                StatusCode::UNAUTHORIZED,
                "Invalid DSN key".to_string(),
            ),
            AppError::InternalServerError => (
                StatusCode::INTERNAL_SERVER_ERROR,
                "Internal server error".to_string(),
            ),
        };

        let body = json!({
            "error": message,
            "code": status.as_u16(),
        });

        (status, axum::Json(body)).into_response()
    }
}

// 从其他错误类型自动转换
impl From<validator::ValidationErrors> for AppError {
    fn from(err: validator::ValidationErrors) -> Self {
        AppError::ValidationError(err.to_string())
    }
}

impl From<sea_orm::DbErr> for AppError {
    fn from(err: sea_orm::DbErr) -> Self {
        AppError::DatabaseError(err)
    }
}

impl From<rdkafka::error::KafkaError> for AppError {
    fn from(err: rdkafka::error::KafkaError) -> Self {
        AppError::KafkaError(err)
    }
}

// 其他必要的 From 实现...