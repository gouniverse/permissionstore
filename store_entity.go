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

func (store *store) EntityPermissionCount(ctx context.Context, options EntityPermissionQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.entityPermissionSelectQuery(options)

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

func (store *store) EntityPermissionCreate(ctx context.Context, entityPermission EntityPermissionInterface) error {
	if entityPermission == nil {
		return errors.New("permissionstore > EntityPermissionCreate. entityPermission is nil")
	}

	if entityPermission.PermissionID() == "" {
		return errors.New("permissionstore > EntityPermissionCreate. entityPermission permissionID is empty")
	}

	if entityPermission.EntityID() == "" {
		return errors.New("permissionstore > EntityPermissionCreate. entityPermission entityID is empty")
	}

	if entityPermission.EntityType() == "" {
		return errors.New("permissionstore > EntityPermissionCreate. entityPermission entityType is empty")
	}

	entityPermissionExists, err := store.EntityPermissionFindByEntityAndPermission(
		ctx,
		entityPermission.EntityType(),
		entityPermission.EntityID(),
		entityPermission.PermissionID(),
	)

	if err != nil {
		return err
	}

	if entityPermissionExists != nil {
		return errors.New("permissionstore > EntityPermissionCreate. entityPermission with the same entityType-entityID-permissionID combination already exists")
	}

	entityPermission.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	entityPermission.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := entityPermission.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.entityPermissionTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityPermissionstore: database is nil")
	}

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	entityPermission.MarkAsNotDirty()

	return nil
}

func (store *store) EntityPermissionDelete(ctx context.Context, entityPermission EntityPermissionInterface) error {
	if entityPermission == nil {
		return errors.New("entityPermission is nil")
	}

	return store.EntityPermissionDeleteByID(ctx, entityPermission.ID())
}

func (store *store) EntityPermissionDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("entityPermission id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.entityPermissionTableName).
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

func (store *store) EntityPermissionFindByEntityAndPermission(
	ctx context.Context,
	entityType string,
	entityID string,
	permissionID string,
) (entityPermission EntityPermissionInterface, err error) {
	if entityType == "" {
		return nil, errors.New("EntityPermissionFindByEntityAndPermission entityType is empty")
	}

	if entityID == "" {
		return nil, errors.New("EntityPermissionFindByEntityAndPermission entityID is empty")
	}

	if permissionID == "" {
		return nil, errors.New("EntityPermissionFindByEntityAndPermission permissionID is empty")
	}

	query := NewEntityPermissionQuery().
		SetEntityType(entityType).
		SetEntityID(entityID).
		SetPermissionID(permissionID).
		SetLimit(1)

	list, err := store.EntityPermissionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityPermissionFindByID(ctx context.Context, id string) (entityPermission EntityPermissionInterface, err error) {
	if id == "" {
		return nil, errors.New("entityPermission id is empty")
	}

	query := NewEntityPermissionQuery().SetID(id).SetLimit(1)

	list, err := store.EntityPermissionList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityPermissionList(ctx context.Context, query EntityPermissionQueryInterface) ([]EntityPermissionInterface, error) {
	if query == nil {
		return []EntityPermissionInterface{}, errors.New("at entityPermission list > entityPermission query is nil")
	}

	q, columns, err := store.entityPermissionSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []EntityPermissionInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []EntityPermissionInterface{}, errors.New("entityPermissionstore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []EntityPermissionInterface{}, err
	}

	list := []EntityPermissionInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewEntityPermissionFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) EntityPermissionSoftDelete(ctx context.Context, entityPermission EntityPermissionInterface) error {
	if entityPermission == nil {
		return errors.New("at entityPermission soft delete > entityPermission is nil")
	}

	entityPermission.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.EntityPermissionUpdate(ctx, entityPermission)
}

func (store *store) EntityPermissionSoftDeleteByID(ctx context.Context, id string) error {
	entityPermission, err := store.EntityPermissionFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.EntityPermissionSoftDelete(ctx, entityPermission)
}

func (store *store) EntityPermissionUpdate(ctx context.Context, entityPermission EntityPermissionInterface) error {
	if entityPermission == nil {
		return errors.New("at entityPermission update > entityPermission is nil")
	}

	entityPermission.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := entityPermission.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.entityPermissionTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(entityPermission.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityPermissionstore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	entityPermission.MarkAsNotDirty()

	return err
}

func (store *store) entityPermissionSelectQuery(options EntityPermissionQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("entityPermission options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.entityPermissionTableName)

	if options.HasEntityID() {
		q = q.Where(goqu.C(COLUMN_ENTITY_ID).Eq(options.EntityID()))
	}

	if options.HasEntityType() {
		q = q.Where(goqu.C(COLUMN_ENTITY_TYPE).Eq(options.EntityType()))
	}

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasPermissionID() {
		q = q.Where(goqu.C(COLUMN_PERMISSION_ID).Eq(options.PermissionID()))
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
		return q, columns, nil // soft deleted entityPermissions requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
