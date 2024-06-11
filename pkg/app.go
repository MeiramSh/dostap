package pkg

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	DB *pgxpool.Pool
}

func App(DATABASE_URL string) (Application, error) {
	app := &Application{}
	conn, err := NewConn(DATABASE_URL)
	if err != nil {
		return Application{}, err
	}
	app.DB = conn

	return *app, nil
}

func (app *Application) CloseDBConnection() {
	app.DB.Close()
}
