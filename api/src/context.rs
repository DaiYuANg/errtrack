use axum_example_service::sea_orm::prelude::Uuid;

pub enum AuthContext {
  User { user_id: Uuid },
  Project { project_id: Uuid },
}
