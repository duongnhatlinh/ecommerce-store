package database

import (
	"ecommercestore/internal/conf"
	"fmt"
	"github.com/rs/zerolog/log"
)

func ResetTestDatabase() {
	ConnectionDatabase(conf.NewTestConfig())

	tables := []string{
		"users",
		"user_addresses",
		"products",
		"product_likes",
		"order_items",
		"orders",
		"inventories",
		"discounts",
		"categories",
		"carts",
	}

	for _, table := range tables {
		result := DB.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if result.Error != nil {
			log.Panic().Err(result.Error).Str("table", table).Msg("Error clearing test database")
		}

		result = DB.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table))
		if result.Error != nil {
			log.Panic().Err(result.Error).Str("table", table).Msg("Error resetting auto increment")
		}
	}
}
