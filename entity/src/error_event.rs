use crate::error;
use sea_orm::entity::prelude::*;

#[derive(Clone, Debug, PartialEq, DeriveEntityModel)]
#[sea_orm(table_name = "error_events")]
pub struct Model {
  #[sea_orm(primary_key)]
  pub id: i64,
  pub project_id: i64,
  pub error_id: i64,
  pub timestamp: DateTimeWithTimeZone,
  pub stacktrace: String,
  #[sea_orm(column_type = "JsonBinary")]
  pub tags: Json,
  pub payload: Json,
  pub request_info: Json,
}

#[derive(Copy, Clone, Debug, EnumIter)]
pub enum Relation {
  Error,
}

impl RelationTrait for Relation {
  fn def(&self) -> RelationDef {
    match self {
      Self::Error => Entity::belongs_to(error::Entity)
        .from(Column::ErrorId)
        .to(error::Column::Id)
        .into(),
    }
  }
}

impl Related<error::Entity> for Entity {
  fn to() -> RelationDef {
    Relation::Error.def()
  }

  fn via() -> Option<RelationDef> {
    None
  }
}

impl ActiveModelBehavior for ActiveModel {}
