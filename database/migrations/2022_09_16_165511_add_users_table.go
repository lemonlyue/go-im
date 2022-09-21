package migrations

import (
	"database/sql"
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Nickname  string `gorm:"type:varchar(255);not null;index"`
		AvatarUrl string `gorm:"type:varchar(255);not null"`
		Email     string `gorm:"type:varchar(255);not null;uniqueIndex"`
		Password  string `gorm:"type:varchar(255);not null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_09_16_165511_add_users_table", up, down)
}
