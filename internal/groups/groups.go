package groups

import "github.com/cooking-club/recipes/internal/db"
import "context"

func GetGroups() ([]db.Group, error) {
	ctx := context.Background()

	var groups []db.Group
	err := db.Select(&groups).Relation("Department").Scan(ctx)
	return groups, err
}
