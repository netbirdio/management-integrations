package core

import (
	"context"
	"errors"
)

type Module string

const (
	Networks Module = "networks"
	Peers    Module = "peers"
	Groups   Module = "groups"
)

type Operation string

const (
	Read  Operation = "read"
	Write Operation = "write"
)

// @todo move types.UserRole here

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserNotInAccount = errors.New("user does not belong to account")
	ErrUserRoleInvalid  = errors.New("invalid role")

	ErrSettingsFailedToGet = errors.New("failed to get settings")
)

type BaseManager struct {
	module Module
}

func NewBaseManager(module Module) BaseManager {
	return BaseManager{
		module: module,
	}
}

func (m *BaseManager) ValidatePermissions(ctx context.Context, operation Operation) error {
	// auth, err := nbcontext.GetUserAuthFromContext(ctx)
	// if err != nil {
	// 	return status.NewPermissionValidationError(err)
	// }

	// user, err := m.store.GetUserByUserID(ctx, store.LockingStrengthShare, auth.UserId)
	// if err != nil {
	// 	return status.NewPermissionValidationError(err)
	// }

	// if user == nil {
	// 	return status.NewPermissionValidationError(ErrUserNotFound)
	// }

	// if user.AccountID != auth.AccountId {
	// 	return status.NewPermissionValidationError(ErrUserNotInAccount)
	// }

	// if user.Blocked {
	// 	return status.NewPermissionDeniedError()
	// }

	// switch user.Role {
	// case types.UserRoleAdmin, types.UserRoleOwner:
	// 	return nil
	// case types.UserRoleUser:
	// 	return m.validateRegularUserPermissions(ctx, auth.AccountId, operation)
	// case types.UserRoleBillingAdmin:
	// 	return status.NewPermissionDeniedError()
	// default:
	// 	return status.NewPermissionValidationError(ErrUserRoleInvalid)
	// }
	return nil
}

func (m *BaseManager) validateRegularUserPermissions(ctx context.Context, accountId string, operation Operation) error {
	// settings, err := m.store.GetAccountSettings(ctx, store.LockingStrengthShare, accountId)
	// if err != nil {
	// 	return status.NewPermissionValidationError(fmt.Errorf("%w: %w", ErrSettingsFailedToGet, err))
	// }
	// if settings.RegularUsersViewBlocked {
	// 	return status.NewPermissionDeniedError()
	// }

	// if operation == Write {
	// 	return status.NewPermissionDeniedError()
	// }

	// if m.module == Peers {
	// 	return nil
	// }

	// return status.NewPermissionDeniedError()
	return nil
}
