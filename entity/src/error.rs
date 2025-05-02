use crate::project;
use sea_orm::entity::prelude::*;

#[derive(Clone, Debug, PartialEq, DeriveEntityModel)]
#[sea_orm(table_name = "errors")]
pub struct Model {
  #[sea_orm(primary_key)]
  pub id: i32,
  pub project_id: i32,
  pub fingerprint: String,
  pub message: String,
  pub culprit: Option<String>,
  pub level: String,
  pub environment: String,
  pub first_seen: DateTimeUtc,
  pub last_seen: DateTimeUtc,
  pub event_count: i32,
}

#[derive(Copy, Clone, Debug, EnumIter)]
pub enum Relation {
  Project,
  Events,
}

impl RelationTrait for Relation {
  fn def(&self) -> RelationDef {
    match self {
      Self::Project => Entity::belongs_to(project::Entity)
        .from(Column::ProjectId)
        .to(project::Column::Id)
        .into(),
      Self::Events => Entity::has_many(super::error_event::Entity).into(),
    }
  }
}

impl ActiveModelBehavior for ActiveModel {}
