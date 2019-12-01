package services

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/graphlog/heimdall/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ApplicationService struct {
	TableName string
	DB        *sqlx.DB
}

func NewApplicationService(db *sqlx.DB) *ApplicationService {
	return &ApplicationService{
		TableName: "applications",
		DB:        db,
	}
}

func (a *ApplicationService) Validate(auth *models.AuthorizationKeys) bool {
	var app []models.Application

	sql, args, err := sq.Select("api_key", "api_secret").From(a.TableName).Where(sq.Eq{"api_key": auth.APIKey, "enabled": true}).Limit(1).ToSql()

	if err != nil {
		logrus.Error(err)
		return false
	}

	if err := a.DB.Select(&app, sql, args...); err != nil {
		logrus.Error(err)
		return false
	}

	if len(app) == 1 {
		currApp := app[0]

		if err := bcrypt.CompareHashAndPassword([]byte(currApp.APISecret), []byte(auth.APISecret)); err != nil {
			return false
		}

		return true
	}

	return false

}
