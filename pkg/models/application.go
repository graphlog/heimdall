package models

type AuthorizationKeys struct {
	APIKey    string
	APISecret string
}

type Application struct {
	Name      string `db:"name"`
	APIKey    string `db:"api_key"`
	APISecret string `db:"api_secret"`
}
