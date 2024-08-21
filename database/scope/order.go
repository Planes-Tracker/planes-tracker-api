package scope

import (
	"github.com/LockBlock-dev/planes-tracker-api/constant"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Order(order string, sortBy string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sortBy == "" {
			sortBy = clause.PrimaryKey
		}

		return db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: sortBy},
			Desc:   order == constant.ORDERING_DESCENDING,
		})
	}
}
