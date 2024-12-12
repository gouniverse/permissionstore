package permissionstore

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/utils"
	_ "modernc.org/sqlite"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && utils.FileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			return nil, err
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(filepath string) (StoreInterface, error) {
	db, err := initDB(filepath)

	if err != nil {
		return nil, err
	}

	store, err := NewStore(NewStoreOptions{
		DB:                        db,
		PermissionTableName:       "permissions_permission_table",
		EntityPermissionTableName: "permissions_entity_permission_table",
		AutomigrateEnabled:        true,
		DebugEnabled:              true,
		SqlLogger:                 slog.New(slog.NewTextHandler(os.Stdout, nil)),
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}

func TestStoreWithTx(t *testing.T) {
	store, err := initStore("test_store_with_tx.db")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	db := store.DB()

	if db == nil {
		t.Fatal("unexpected nil db")
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	tx, err := db.Begin()

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tx == nil {
		t.Fatal("unexpected nil tx")
	}

	txCtx := database.Context(context.Background(), tx)

	// create permission
	permission := NewPermission().
		SetStatus(PERMISSION_STATUS_ACTIVE).
		SetHandle("PERMISSION_HANDLE").
		SetTitle("PERMISSION_TITLE")

	err = store.PermissionCreate(txCtx, permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// update permission
	permission.SetTitle("PERMISSION_TITLE_2")
	err = store.PermissionUpdate(txCtx, permission)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check permission
	permissionFound, errFind := store.PermissionFindByID(database.Context(context.Background(), db), permission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound != nil {
		t.Fatal("Permission MUST be nil, as transaction not committed")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check permission
	permissionFound, errFind = store.PermissionFindByID(database.Context(context.Background(), db), permission.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if permissionFound == nil {
		t.Fatal("Permission MUST be not nil, as transaction committed")
	}

	if permissionFound.Title() != "PERMISSION_TITLE_2" {
		t.Fatal("Permission MUST be PERMISSION_TITLE_2, as transaction committed")
	}
}
