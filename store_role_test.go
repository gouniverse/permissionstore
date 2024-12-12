package permissionstore

import (
	"context"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

func TestStorePermissionCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.PermissionCount(context.Background(), NewPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")
	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.PermissionCount(context.Background(), NewPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	err = store.PermissionCreate(context.Background(), NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.PermissionCount(context.Background(), NewPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStorePermissionCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStorePermissionDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.PermissionDelete(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	permissionFound, err := store.PermissionFindByID(context.Background(), permission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if permissionFound != nil {
		t.Fatal("Permission MUST be nil")
	}

	permissionFindWithDeleted, err := store.PermissionList(context.Background(), NewPermissionQuery().
		SetID(permission.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(permissionFindWithDeleted) != 0 {
		t.Fatal("Permission MUST be nil")
	}
}

func TestStorePermissionDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.PermissionDeleteByID(context.Background(), permission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	permissionFound, err := store.PermissionFindByID(context.Background(), permission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if permissionFound != nil {
		t.Fatal("Permission MUST be nil")
	}

	permissionFindWithDeleted, err := store.PermissionList(context.Background(), NewPermissionQuery().
		SetID(permission.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(permissionFindWithDeleted) != 0 {
		t.Fatal("Permission MUST NOT be found")
	}
}

func TestStorePermissionFindByHandle(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = permission.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.PermissionCreate(database.Context(context.Background(), store.DB()), permission)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	permissionFound, errFind := store.PermissionFindByHandle(database.Context(context.Background(), store.DB()), permission.Handle())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound == nil {
		t.Fatal("Permission MUST NOT be nil")
	}

	if permissionFound.ID() != permission.ID() {
		t.Fatal("IDs do not match")
	}

	if permissionFound.Handle() != permission.Handle() {
		t.Fatal("Handles do not match")
	}

	if permissionFound.Title() != permission.Title() {
		t.Fatal("Titles do not match")
	}

	if permissionFound.Status() != permission.Status() {
		t.Fatal("Statuses do not match")
	}

	if permissionFound.Meta("education_1") != permission.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if permissionFound.Meta("education_2") != permission.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if permissionFound.Meta("education_3") != permission.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStorePermissionFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = permission.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := database.Context(context.Background(), store.DB())
	err = store.PermissionCreate(ctx, permission)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	permissionFound, errFind := store.PermissionFindByID(ctx, permission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound == nil {
		t.Fatal("Permission MUST NOT be nil")
	}

	if permissionFound.ID() != permission.ID() {
		t.Fatal("IDs do not match")
	}

	if permissionFound.Handle() != permission.Handle() {
		t.Fatal("Handles do not match")
	}

	if permissionFound.Title() != permission.Title() {
		t.Fatal("Titles do not match")
	}

	if permissionFound.Status() != permission.Status() {
		t.Fatal("Statuses do not match")
	}

	if permissionFound.Meta("education_1") != permission.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if permissionFound.Meta("education_2") != permission.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if permissionFound.Meta("education_3") != permission.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStorePermissionList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission1 := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE_1").
		SetTitle("PERMISSION_TITLE_1")

	permission2 := NewPermission().
		SetStatus(PERMISSION_STATUS_INACTIVE).
		SetHandle("PERMISSION_HANDLE_2").
		SetTitle("PERMISSION_TITLE_2")

	permissions := []PermissionInterface{
		permission1,
		permission2,
	}

	for _, permission := range permissions {
		err = store.PermissionCreate(context.Background(), permission)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listActive, err := store.PermissionList(context.Background(), NewPermissionQuery().SetStatus(PERMISSION_STATUS_ACTIVE))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listActive) != 1 {
		t.Fatal("unexpected list length:", len(listActive))
	}

	listEmail, err := store.PermissionList(context.Background(), NewPermissionQuery().SetHandle("PERMISSION_HANDLE_2"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listEmail) != 1 {
		t.Fatal("unexpected list length:", len(listEmail))
	}
}

func TestStorePermissionSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.PermissionSoftDelete(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if permission.SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Permission MUST be soft deleted")
	}

	permissionFound, errFind := store.PermissionFindByID(context.Background(), permission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound != nil {
		t.Fatal("Permission MUST be soft deleted, so MUST be nil")
	}

	permissionFindWithDeleted, err := store.PermissionList(context.Background(), NewPermissionQuery().
		SetSoftDeletedIncluded(true).
		SetID(permission.ID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(permissionFindWithDeleted) == 0 {
		t.Fatal("Permission MUST be soft deleted")
	}

	if strings.Contains(permissionFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("Permission MUST be soft deleted", permission.SoftDeletedAt())
	}

	if !permissionFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Permission MUST be soft deleted")
	}
}

func TestStorePermissionSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(context.Background(), permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.PermissionSoftDeleteByID(context.Background(), permission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if permission.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("Permission MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	permissionFound, errFind := store.PermissionFindByID(context.Background(), permission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound != nil {
		t.Fatal("Permission MUST be nil")
	}
	query := NewPermissionQuery().
		SetSoftDeletedIncluded(true).
		SetID(permission.ID()).
		SetLimit(1)

	permissionFindWithDeleted, err := store.PermissionList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(permissionFindWithDeleted) == 0 {
		t.Fatal("Permission MUST be soft deleted")
	}

	if strings.Contains(permissionFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("Permission MUST be soft deleted", permission.SoftDeletedAt())
	}

	if !permissionFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Permission MUST be soft deleted")
	}
}
