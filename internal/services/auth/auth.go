package auth

import (
	"cmd/sso/main.go/internal/domain/models"
	"context"
	"log/slog"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	AppProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (models.User, error)
}
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID string) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int32) (models.App, error)
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		AppProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email, password string, appID int32) (string, error) {
	panic("not implemented")
}

func (a *Auth) RegisterNewUser(ctx context.Context, email, password string) (string, error) {
	const op = "auth.RegisterNewUser"
	panic("not implemented")
}

func (a *Auth) IsAdmin(ctx context.Context, email, password string, appID int32) (string, error) {
	panic("not implemented")
}
