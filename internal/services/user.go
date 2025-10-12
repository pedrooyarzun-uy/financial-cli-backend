package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/helpers"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type UserService interface {
	SignUp(usr dto.SignUpReq) error
	SignIn(usr dto.SignInReq) error
}

type userService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		ur,
	}
}

func (s *userService) SignUp(req dto.SignUpReq) error {

	//Check if user is on DB
	res := s.ur.GetByEmail(req.Email)

	//User exists
	if res.Name != "" {

		if !res.Verified {

			if res.VerificationTokenExpiresAt.After(time.Now()) {
				newToken := uuid.New().String()
				err := s.ur.UpdateToken(res.Id, newToken, time.Now().Add(5*time.Hour))

				if err != nil {
					return ErrUserCreationFailed
				}

				err = helpers.SendMail(req.Email, req.Name, "Example", "Example")

				if err != nil {
					return ErrUserCreationFailed
				}

				return nil
			}

			//Send email again with same token
			err := helpers.SendMail(req.Email, req.Name, "Example", "Example")

			if err != nil {
				return ErrUserAlreadyExists
			}

			return nil
		}

		return ErrUserAlreadyExists
	}

	password, _ := helpers.HashPassword(req.Password)

	//Auth token
	token := uuid.New().String()

	usr := domain.User{
		Name:                       req.Name,
		Email:                      req.Email,
		Password:                   password,
		VerificationToken:          token,
		VerificationTokenExpiresAt: time.Now().Add(5 * time.Hour),
		Verified:                   false,
		Deleted:                    false,
	}

	//Save user on db
	err := s.ur.Create(usr)

	if err != nil {
		return ErrUserCreationFailed
	}

	//Send email to user
	err = helpers.SendMail(req.Email, req.Name, "Example", "Example")

	if err != nil {
		return ErrUserCreationFailed
	}

	return nil
}

func (s *userService) SignIn(dto.SignInReq) error {
	return nil
}
