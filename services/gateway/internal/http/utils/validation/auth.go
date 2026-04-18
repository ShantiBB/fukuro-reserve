package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	valconst "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation/constants"
)

func ValidateRegisterRequest(req dto.RegisterRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(valconst.ErrPasswordRequired)
	}
	return nil
}

func ValidateLoginRequest(req dto.LoginRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(valconst.ErrPasswordRequired)
	}
	return nil
}

func ValidateRefreshTokenRequest(req dto.RefreshTokenRequest) error {
	if strings.TrimSpace(req.RefreshToken) == "" {
		return errors.New(valconst.ErrRefreshRequired)
	}
	return nil
}

func ValidateCreateUserRequest(req dto.CreateUserRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldEmail, err)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New(valconst.ErrPasswordRequired)
	}
	return nil
}

func ValidateUpdateUserRequest(req dto.UpdateUserRequest) error {
	if strings.TrimSpace(req.Email) == "" && strings.TrimSpace(req.Username) == "" {
		return errors.New(valconst.ErrEmailOrUserReq)
	}
	if strings.TrimSpace(req.Email) != "" {
		if err := validateEmail(req.Email); err != nil {
			return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldEmail, err)
		}
	}
	return nil
}

func ValidateUpdateUserRoleRequest(req dto.UpdateUserRoleRequest) error {
	if strings.TrimSpace(req.Role) == "" {
		return errors.New(valconst.ErrRoleRequired)
	}
	return nil
}
