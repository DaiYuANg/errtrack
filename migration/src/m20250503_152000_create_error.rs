use sea_orm_migration::{prelude::*, schema::*};

#[derive(DeriveMigrationName)]
pub struct Migration;

#[async_trait::async_trait]
impl MigrationTrait for Migration {
  async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
    manager
      .create_table(
        Table::create()
          .table(Project::Table)
          .if_not_exists()
          .col(
            ColumnDef::new(Project::Id)
              .integer()
              .not_null()
              .auto_increment()
              .primary_key(), // 主键 + 自增
          )
          .col(
            ColumnDef::new(Project::Title)
              .string() // VARCHAR 类型
              .not_null(),
          )
          .col(
            ColumnDef::new(Project::Text)
              .text() // TEXT 类型（对应实体中的 column_type = "Text"）
              .not_null(),
          )
          .col(ColumnDef::new(Project::AccessKey).string().not_null())
          .to_owned(),
      )
      .await
  }

  async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
    manager
      .drop_table(Table::drop().table(Project::Table).to_owned())
      .await
  }
}

#[derive(DeriveIden)]
enum Project {
  Table,
  Id,
  Title,
  Text,
  AccessKey,
}
