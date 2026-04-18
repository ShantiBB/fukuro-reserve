package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func ValidateRegisterRequest(req dto.RegisterRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(consts.ErrPasswordRequired)
	}
	return nil
}

func ValidateLoginRequest(req dto.LoginRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(consts.ErrPasswordRequired)
	}
	return nil
}

func ValidateRefreshTokenRequest(req dto.RefreshTokenRequest) error {
	if strings.TrimSpace(req.RefreshToken) == "" {
		return errors.New(consts.ErrRefreshRequired)
	}
	return nil
}

func ValidateCreateUserRequest(req dto.CreateUserRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(consts.ErrPasswordRequired)
	}
	return nil
}

func ValidateUpdateUserRequest(req dto.UpdateUserRequest) error {
	if strings.TrimSpace(req.Email) == "" && strings.TrimSpace(req.Username) == "" {
		return errors.New(consts.ErrEmailOrUserReq)
	}
	if strings.TrimSpace(req.Email) != "" {
		if err := validateEmail(req.Email); err != nil {
			return fmt.Errorf(consts.ErrWrapWithField, consts.FieldEmail, err)
		}
	}
	return nil
}

func ValidateUpdateUserRoleRequest(req dto.UpdateUserRoleRequest) error {
	if strings.TrimSpace(req.Role) == "" {
		return errors.New(consts.ErrRoleRequired)
	}
	return nil
}
