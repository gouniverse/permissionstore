package permissionstore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

type StoreInterface interface {
	// AutoMigrate auto migrates the database schema
	AutoMigrate() error

	// EnableDebug enables or disables the debug mode
	EnableDebug(debug bool)

	// DB returns the underlying database connection
	DB() *sql.DB

	// == Permission Methods =======================================================//

	// PermissionCount returns the number of permissions based on the given query options
	PermissionCount(ctx context.Context, options PermissionQueryInterface) (int64, error)

	// PermissionCreate creates a new permission
	PermissionCreate(ctx context.Context, permission PermissionInterface) error

	// PermissionDelete deletes a permission
	PermissionDelete(ctx context.Context, permission PermissionInterface) error

	// PermissionDeleteByID deletes a permission by its ID
	PermissionDeleteByID(ctx context.Context, id string) error

	// PermissionFindByHandle returns a permission by its handle
	PermissionFindByHandle(ctx context.Context, handle string) (PermissionInterface, error)

	// PermissionFindByID returns a permission by its ID
	PermissionFindByID(ctx context.Context, id string) (PermissionInterface, error)

	// PermissionList returns a list of permissions based on the given query options
	PermissionList(ctx context.Context, query PermissionQueryInterface) ([]PermissionInterface, error)

	// PermissionSoftDelete soft deletes a permission
	PermissionSoftDelete(ctx context.Context, permission PermissionInterface) error

	// PermissionSoftDeleteByID soft deletes a permission by its ID
	PermissionSoftDeleteByID(ctx context.Context, id string) error

	// PermissionUpdate updates a permission
	PermissionUpdate(ctx context.Context, permission PermissionInterface) error

	// == EntityPermission Methods =================================================//

	// EntityPermissionCount returns the number of permission entities mappings based on the given query options
	EntityPermissionCount(ctx context.Context, options EntityPermissionQueryInterface) (int64, error)

	// EntityPermissionCreate creates a new permission entity mapping
	EntityPermissionCreate(ctx context.Context, entityPermission EntityPermissionInterface) error

	// EntityPermissionDelete deletes a permission entity mapping
	EntityPermissionDelete(ctx context.Context, entityPermission EntityPermissionInterface) error

	// EntityPermissionDeleteByID deletes a permission entity mapping by its ID
	EntityPermissionDeleteByID(ctx context.Context, id string) error

	// EntityPermissionFindByEntityAndPermission returns a permission entity mapping by its entity type, entity ID and permission ID
	EntityPermissionFindByEntityAndPermission(ctx context.Context, entityType string, entityID string, permissionID string) (EntityPermissionInterface, error)

	// EntityPermissionFindByID returns a permission entity mapping by its ID
	EntityPermissionFindByID(ctx context.Context, id string) (EntityPermissionInterface, error)

	// EntityPermissionList returns a list of permission entity mappings based on the given query options
	EntityPermissionList(ctx context.Context, query EntityPermissionQueryInterface) ([]EntityPermissionInterface, error)

	// EntityPermissionSoftDelete soft deletes a permission entity mapping
	EntityPermissionSoftDelete(ctx context.Context, entityPermission EntityPermissionInterface) error

	// EntityPermissionSoftDeleteByID soft deletes a permission entity mapping by its ID
	EntityPermissionSoftDeleteByID(ctx context.Context, id string) error

	// EntityPermissionUpdate updates a permission entity mapping
	EntityPermissionUpdate(ctx context.Context, entityPermission EntityPermissionInterface) error
}

type PermissionInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) PermissionInterface

	Handle() string
	SetHandle(handle string) PermissionInterface

	ID() string
	SetID(id string) PermissionInterface

	Memo() string
	SetMemo(memo string) PermissionInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	Status() string
	SetStatus(status string) PermissionInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) PermissionInterface

	Title() string
	SetTitle(title string) PermissionInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) PermissionInterface
}

type EntityPermissionInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) EntityPermissionInterface

	EntityType() string
	SetEntityType(entityType string) EntityPermissionInterface

	EntityID() string
	SetEntityID(entityID string) EntityPermissionInterface

	ID() string
	SetID(id string) EntityPermissionInterface

	Memo() string
	SetMemo(memo string) EntityPermissionInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	PermissionID() string
	SetPermissionID(permissionID string) EntityPermissionInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) EntityPermissionInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) EntityPermissionInterface
}

type UserInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	Get(columnName string) string
	Set(columnName string, value string)

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool
	IsUnverified() bool

	IsAdministrator() bool
	IsManager() bool
	IsSuperuser() bool

	IsRegistrationCompleted() bool

	// setters and getters

	BusinessName() string
	SetBusinessName(businessName string) UserInterface

	Country() string
	SetCountry(country string) UserInterface

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) UserInterface

	Email() string
	SetEmail(email string) UserInterface

	ID() string
	SetID(id string) UserInterface

	FirstName() string
	SetFirstName(firstName string) UserInterface

	LastName() string
	SetLastName(lastName string) UserInterface

	Memo() string
	SetMemo(memo string) UserInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error
	UpsertMetas(metas map[string]string) error

	MiddleNames() string
	SetMiddleNames(middleNames string) UserInterface

	Password() string
	SetPassword(password string) UserInterface

	Phone() string
	SetPhone(phone string) UserInterface

	ProfileImageUrl() string
	SetProfileImageUrl(profileImageUrl string) UserInterface

	Permission() string
	SetPermission(permission string) UserInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	Timezone() string
	SetTimezone(timezone string) UserInterface

	Status() string
	SetStatus(status string) UserInterface

	PasswordCompare(password string) bool

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
