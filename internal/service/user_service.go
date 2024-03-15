package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/sagikazarmark/slog-shim"
	"golang-marketplace/internal/entity"
	helpers "golang-marketplace/internal/helper"
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

func (s *UserService) Register(ctx context.Context, request *model.RegisterRequest) (*model.LoginRegisterResponse, error) {

	// handle request
	err := helpers.ValidationError(s.Validate, request)
	if err != nil {
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	user, err := s.getUsername(request.Username)
	if user != nil {
		return nil, &fiber.Error{
			Code:    409,
			Message: "Username already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashedPassword")
	}

	user = &entity.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hashedPassword),
	}

	err = s.Repository.Create(user)
	if err != nil {
		log.Println("Gagal menyimpan user", err)
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

	resp := &model.LoginRegisterResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: t,
	}

	return resp, nil
}

func (s *UserService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginRegisterResponse, error) {

	// handle request
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	user, err := s.getUsername(request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("Error Password is Wrong")
		return nil, &fiber.Error{
			Code:    400,
			Message: "Password is wrong",
		}
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

	resp := &model.LoginRegisterResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: t,
	}

	return resp, nil
}

func (s *UserService) getUsername(username string) (*entity.User, error) {
	user := entity.User{Username: username}
	err := s.Repository.GetByUsername(&user)
	if err != nil {
		log.Println("Error Get User by Username")
		return nil, &fiber.Error{
			Code:    404,
			Message: "User NotFound",
		}
	}
	return &user, nil
}
