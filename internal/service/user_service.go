package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/sagikazarmark/slog-shim"
	"golang-marketplace/internal/entity"
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserService struct {
	Repository *repository.UserRepository
	Validate   *validator.Validate
	Log        *slog.Logger
}

func NewUserService(r *repository.UserRepository, validate *validator.Validate, log *slog.Logger) *UserService {
	return &UserService{Repository: r, Validate: validate, Log: log}
}

func (s *UserService) Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashedPassword")
	}

	user := &entity.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hashedPassword),
	}

	err = s.Repository.Create(user)
	if err != nil {
		log.Fatal("Gagal menyimpan user", err)
	}

	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	//TODO :: config secret jwt
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	resp := &model.RegisterResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: t,
	}

	return resp, nil
}
