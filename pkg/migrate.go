package pkg

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(p *pgxpool.Pool) error {
	log.Println("Migrating...")
	if err := drop(p); err != nil {
		log.Println("Error while dropping tables")
		return err
	}
	if err := tables(p); err != nil {
		log.Println("Error while creating tables")
		return err
	}
	// _, err := utils.CreateAdmin(p)
	// if err != nil {
	// 	return err
	// }
	if err := mock(p); err != nil {
		log.Println("Error while writing mock data")
		return err
	}
	log.Println("Migration completed")
	return nil
}

func drop(p *pgxpool.Pool) error {
	migrationSQL, err := os.ReadFile("migrations/drop.sql")
	if err != nil {
		return err
	}
	_, err = p.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		return err
	}
	return nil
}

func tables(p *pgxpool.Pool) error {
	migrationSQL, err := os.ReadFile("migrations/tables.sql")
	if err != nil {
		return err
	}
	_, err = p.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		return err
	}
	return nil
}

func mock(p *pgxpool.Pool) error {
	migrationSQL, err := os.ReadFile("migrations/mock.sql")
	if err != nil {
		return err
	}
	_, err = p.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		return err
	}
	return nil
}
