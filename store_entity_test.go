package permissionstore

import (
	"context"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

func TestStoreEntityPermissionCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.EntityPermissionCount(context.Background(), NewEntityPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityPermissionCount(context.Background(), NewEntityPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	entityPermission2 := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetPermissionID("PERMISSION_02")

	err = store.EntityPermissionCreate(context.Background(), entityPermission2)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityPermissionCount(context.Background(), NewEntityPermissionQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreEntityPermissionCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreEntityPermissionCreate_Duplicate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err == nil {
		t.Fatal("must return error as duplicated entity to permission relationship")
	}
}

func TestStoreEntityPermissionDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionDelete(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityPermissionFound, err := store.EntityPermissionFindByID(context.Background(), entityPermission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityPermissionFound != nil {
		t.Fatal("EntityPermission MUST be nil")
	}

	entityPermissionFindWithDeleted, err := store.EntityPermissionList(context.Background(), NewEntityPermissionQuery().
		SetID(entityPermission.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityPermissionFindWithDeleted) != 0 {
		t.Fatal("EntityPermission MUST be nil")
	}
}

func TestStoreEntityPermissionDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionDeleteByID(context.Background(), entityPermission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityPermissionFound, err := store.EntityPermissionFindByID(context.Background(), entityPermission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityPermissionFound != nil {
		t.Fatal("EntityPermission MUST be nil")
	}

	entityPermissionFindWithDeleted, err := store.EntityPermissionList(context.Background(), NewEntityPermissionQuery().
		SetID(entityPermission.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityPermissionFindWithDeleted) != 0 {
		t.Fatal("EntityPermission MUST NOT be found")
	}
}

func TestStoreEntityPermissionFindByEntityAndPermission(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = entityPermission.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionCreate(database.Context(context.Background(), store.DB()), entityPermission)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityPermissionFound, errFind := store.EntityPermissionFindByEntityAndPermission(database.Context(context.Background(), store.DB()), entityPermission.EntityType(), entityPermission.EntityID(), entityPermission.PermissionID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityPermissionFound == nil {
		t.Fatal("EntityPermission MUST NOT be nil")
	}

	if entityPermissionFound.ID() != entityPermission.ID() {
		t.Fatal("IDs do not match")
	}

	if entityPermissionFound.EntityID() != entityPermission.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityPermissionFound.EntityType() != entityPermission.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityPermissionFound.PermissionID() != entityPermission.PermissionID() {
		t.Fatal("PermissionIDs do not match")
	}

	if entityPermissionFound.Meta("education_1") != entityPermission.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityPermissionFound.Meta("education_2") != entityPermission.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityPermissionFound.Meta("education_3") != entityPermission.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityPermissionFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = entityPermission.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := database.Context(context.Background(), store.DB())
	err = store.EntityPermissionCreate(ctx, entityPermission)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityPermissionFound, errFind := store.EntityPermissionFindByID(ctx, entityPermission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityPermissionFound == nil {
		t.Fatal("EntityPermission MUST NOT be nil")
	}

	if entityPermissionFound.ID() != entityPermission.ID() {
		t.Fatal("IDs do not match")
	}

	if entityPermissionFound.EntityID() != entityPermission.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityPermissionFound.EntityType() != entityPermission.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityPermissionFound.PermissionID() != entityPermission.PermissionID() {
		t.Fatal("PermissionIDs do not match")
	}

	if entityPermissionFound.Meta("education_1") != entityPermission.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityPermissionFound.Meta("education_2") != entityPermission.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityPermissionFound.Meta("education_3") != entityPermission.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityPermissionList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission1 := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	entityPermission2 := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetPermissionID("PERMISSION_02")

	entityPermissions := []EntityPermissionInterface{
		entityPermission1,
		entityPermission2,
	}

	for _, entityPermission := range entityPermissions {
		err = store.EntityPermissionCreate(context.Background(), entityPermission)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	list1, err := store.EntityPermissionList(context.Background(), NewEntityPermissionQuery().SetPermissionID("PERMISSION_01"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list1) != 1 {
		t.Fatal("unexpected list length:", len(list1))
	}

	list2, err := store.EntityPermissionList(context.Background(), NewEntityPermissionQuery().SetEntityType("USER").SetEntityID("USER_02"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list2) != 1 {
		t.Fatal("unexpected list length:", len(list2))
	}
}

func TestStoreEntityPermissionSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionSoftDelete(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityPermission.SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("EntityPermission MUST be soft deleted")
	}

	entityPermissionFound, errFind := store.EntityPermissionFindByID(context.Background(), entityPermission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityPermissionFound != nil {
		t.Fatal("EntityPermission MUST be soft deleted, so MUST be nil")
	}

	entityPermissionFindWithDeleted, err := store.EntityPermissionList(context.Background(), NewEntityPermissionQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityPermission.ID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityPermissionFindWithDeleted) == 0 {
		t.Fatal("EntityPermission MUST be soft deleted")
	}

	if strings.Contains(entityPermissionFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityPermission MUST be soft deleted", entityPermission.SoftDeletedAt())
	}

	if !entityPermissionFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityPermission MUST be soft deleted")
	}
}

func TestStoreEntityPermissionSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityPermission := NewEntityPermission().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetPermissionID("PERMISSION_01")

	err = store.EntityPermissionCreate(context.Background(), entityPermission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityPermissionSoftDeleteByID(context.Background(), entityPermission.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityPermission.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("EntityPermission MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	entityPermissionFound, errFind := store.EntityPermissionFindByID(context.Background(), entityPermission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityPermissionFound != nil {
		t.Fatal("EntityPermission MUST be nil")
	}
	query := NewEntityPermissionQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityPermission.ID()).
		SetLimit(1)

	entityPermissionFindWithDeleted, err := store.EntityPermissionList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityPermissionFindWithDeleted) == 0 {
		t.Fatal("EntityPermission MUST be soft deleted")
	}

	if strings.Contains(entityPermissionFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityPermission MUST be soft deleted", entityPermission.SoftDeletedAt())
	}

	if !entityPermissionFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityPermission MUST be soft deleted")
	}
}
