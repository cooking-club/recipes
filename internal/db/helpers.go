package db

import "github.com/uptrace/bun"

func Insert(model any) *bun.InsertQuery {
	return db.NewInsert().Model(model)
}

func Select(model any) *bun.SelectQuery {
	return db.NewSelect().Model(model)
}

func Update(model any) *bun.UpdateQuery {
	return db.NewUpdate().Model(model)
}

func Delete(model any) *bun.DeleteQuery {
	return db.NewDelete().Model(model)
}