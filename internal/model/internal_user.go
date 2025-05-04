package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
	email         string
}
