package permissionstore

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (store *store) PermissionCount(ctx context.Context, options PermissionQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.permissionSelectQuery(options)

	sqlStr, params, errSql := q.Prepared(true).
		Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	store.logSql("select", sqlStr, params...)

	mapped, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, params...)
	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}

func (store *store) PermissionCreate(ctx context.Context, permission PermissionInterface) error {
	if permission == nil {
		return errors.New("permission is nil")
	}

	permission.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	permission.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := permission.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.permissionTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("permissionstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	permission.MarkAsNotDirty()

	return nil
}

func (store *store) PermissionDelete(ctx context.Context, permission PermissionInterface) error {
	if permission == nil {
		return errors.New("permission is nil")
	}

	return store.PermissionDeleteByID(ctx, permission.ID())
}

func (store *store) PermissionDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("permission id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.permissionTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_ID).Eq(id)).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("delete", sqlStr, params...)

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	return err
}

func (store *store) PermissionFindByHandle(ctx context.Context, handle string) (permission PermissionInterface, err error) {
	if handle == "" {
		return nil, errors.New("permission handle is empty")
	}

	query := NewPermissionQuery().SetHandle(handle).SetLimit(1)

	list, err := store.PermissionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) PermissionFindByID(ctx context.Context, id string) (permission PermissionInterface, err error) {
	if id == "" {
		return nil, errors.New("permission id is empty")
	}

	query := NewPermissionQuery().SetID(id).SetLimit(1)

	list, err := store.PermissionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) PermissionList(ctx context.Context, query PermissionQueryInterface) ([]PermissionInterface, error) {
	if query == nil {
		return []PermissionInterface{}, errors.New("at permission list > permission query is nil")
	}

	q, columns, err := store.permissionSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []PermissionInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []PermissionInterface{}, errors.New("permissionstore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []PermissionInterface{}, err
	}

	list := []PermissionInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewPermissionFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) PermissionSoftDelete(ctx context.Context, permission PermissionInterface) error {
	if permission == nil {
		return errors.New("at permission soft delete > permission is nil")
	}

	permission.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.PermissionUpdate(ctx, permission)
}

func (store *store) PermissionSoftDeleteByID(ctx context.Context, id string) error {
	permission, err := store.PermissionFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.PermissionSoftDelete(ctx, permission)
}

func (store *store) PermissionUpdate(ctx context.Context, permission PermissionInterface) error {
	if permission == nil {
		return errors.New("at permission update > permission is nil")
	}

	permission.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := permission.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.permissionTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(permission.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("permissionstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	permission.MarkAsNotDirty()

	return err
}

func (store *store) permissionSelectQuery(options PermissionQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("permission options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.permissionTableName)

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasStatus() {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status()))
	}

	if options.HasStatusIn() {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn()))
	}

	if options.HasHandle() {
		q = q.Where(goqu.C(COLUMN_HANDLE).Eq(options.Handle()))
	}

	if options.HasTitleLike() {
		q = q.Where(goqu.C(COLUMN_TITLE).ILike(`%` + options.TitleLike() + `%`))
	}

	if options.HasCreatedAtGte() && options.HasCreatedAtLte() {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()),
		)
	} else if options.HasCreatedAtGte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()))
	} else if options.HasCreatedAtLte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()))
	}

	if !options.IsCountOnly() {
		if options.HasLimit() {
			q = q.Limit(cast.ToUint(options.Limit()))
		}

		if options.HasOffset() {
			q = q.Offset(cast.ToUint(options.Offset()))
		}
	}

	if options.HasOrderBy() {
		sort := lo.Ternary(options.HasSortDirection(), options.SortDirection(), sb.DESC)
		if strings.EqualFold(sort, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy()).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy()).Desc())
		}
	}

	columns = []any{}

	for _, column := range options.Columns() {
		columns = append(columns, column)
	}

	if options.SoftDeletedIncluded() {
		return q, columns, nil // soft deleted permissions requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
