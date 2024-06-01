package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"sso/internal/domain/models"
	"sso/internal/lib/logger/sl"
	"sso/internal/lib/logger/sl/jwt"
	"sso/internal/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserSaver interface {
	SaveUser(
			ctx context.Context,
			email string,
			passHash []byte,
	) (user models.User, err error)
}

type UserProvider interface {  
	User(ctx context.Context, email string) (models.User, error)  
	IsAdmin(ctx context.Context, userID uint) (bool, error)
}

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	tokenTTL    time.Duration
}

func New(  
	log *slog.Logger,  
	userSaver UserSaver,  
	userProvider UserProvider,  
	tokenTTL time.Duration,  
) *Auth {  
	return &Auth{  
		 usrSaver:    userSaver,  
		 usrProvider: userProvider,  
		 log:         log,  
		 tokenTTL:    tokenTTL,  // Время жизни возвращаемых токенов
	}  
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (string, error) {
	// op (operation) - имя текущей функции и пакета. Такую метку удобно
	// добавлять в логи и в текст ошибок, чтобы легче было искать хвосты
	// в случае поломок.
	const op = "Auth.RegisterNewUser"

	// Создаём локальный объект логгера с доп. полями, содержащими полезную инфу
	// о текущем вызове функции
	log := a.log.With(
			slog.String("op", op),
			slog.String("email", email),
	)

	log.Info("registering user")

	// Генерируем хэш и соль для пароля.
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
			log.Error("failed to generate password hash", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
	}

	// Сохраняем пользователя в БД
	user, err := a.usrSaver.SaveUser(ctx, email, passHash)
	if err != nil {
			log.Error("failed to save user", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
	}

	return CreateToken(op, user, a)
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string, // пароль в чистом виде, аккуратней с логами!
) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
			slog.String("op", op),
			slog.String("username", email),
			// password либо не логируем, либо логируем в замаскированном виде
	)

	log.Info("attempting to login user")

	// Достаём пользователя из БД
	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
			if errors.Is(err, storage.ErrUserNotFound) {
					a.log.Warn("user not found", sl.Err(err))

					return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
			}

			a.log.Error("failed to get user", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем корректность полученного пароля
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
			a.log.Info("invalid credentials", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	// Создаём токен авторизации
	return CreateToken(op, user, a)
}

func CreateToken(op string, user models.User, a *Auth) (string, error){
	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
			a.log.Error("failed to generate token", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID uint) (bool, error) {
	const op = "Auth.IsAdmin"

	log := a.log.With(
			slog.String("op", op),
			slog.Uint64("user_id", uint64(userID)),
	)

	log.Info("checking if user is admin")

	isAdmin, err := a.usrProvider.IsAdmin(ctx, userID)
	if err != nil {
			return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))

	return isAdmin, nil
}
