package schedule

import (
	"context"
	"log"

	"github.com/cooking-club/recipes/internal/db"
)

func GetSchedule(gid int, pos int, limit int) ([]db.Record, error) {
	ctx := context.TODO()

	var records []db.Record
	err := db.Select(&records).
		Relation("Professor").
		Relation("Room").
		Join("JOIN group_records AS gr ON gr.record_id = record.id").
		Where("gr.group_id = ?", gid).
		Where("record.position >= ?", pos).
		Where("record.position < ?", limit).
		Order("record.position ASC").
		Scan(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return records, nil
}