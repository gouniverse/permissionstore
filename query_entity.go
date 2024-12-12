package permissionstore

import "errors"

type EntityPermissionQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) EntityPermissionQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) EntityPermissionQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) EntityPermissionQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) EntityPermissionQueryInterface

	HasEntityID() bool
	EntityID() string
	SetEntityID(entityID string) EntityPermissionQueryInterface

	HasEntityType() bool
	EntityType() string
	SetEntityType(entityType string) EntityPermissionQueryInterface

	HasID() bool
	ID() string
	SetID(id string) EntityPermissionQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) EntityPermissionQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) EntityPermissionQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) EntityPermissionQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) EntityPermissionQueryInterface

	HasPermissionID() bool
	PermissionID() string
	SetPermissionID(permissionID string) EntityPermissionQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) EntityPermissionQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) EntityPermissionQueryInterface

	hasProperty(name string) bool
}

func NewEntityPermissionQuery() EntityPermissionQueryInterface {
	return &permissionEntityQueryImplementation{
		properties: make(map[string]any),
	}
}

type permissionEntityQueryImplementation struct {
	properties map[string]any
}

func (c *permissionEntityQueryImplementation) Validate() error {
	if c.HasCreatedAtGte() && c.CreatedAtGte() == "" {
		return errors.New("permission query. created_at_gte cannot be empty")
	}

	if c.HasCreatedAtLte() && c.CreatedAtLte() == "" {
		return errors.New("permission query. created_at_lte cannot be empty")
	}

	if c.HasEntityID() && c.EntityID() == "" {
		return errors.New("permission query. entity_id cannot be empty")
	}

	if c.HasEntityType() && c.EntityType() == "" {
		return errors.New("permission query. entity_type cannot be empty")
	}

	if c.HasID() && c.ID() == "" {
		return errors.New("permission query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("permission query. id_in cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("permission query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("permission query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("permission query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("permission query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *permissionEntityQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *permissionEntityQueryImplementation) SetColumns(columns []string) EntityPermissionQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *permissionEntityQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *permissionEntityQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *permissionEntityQueryImplementation) SetCountOnly(countOnly bool) EntityPermissionQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *permissionEntityQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *permissionEntityQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *permissionEntityQueryImplementation) SetCreatedAtGte(createdAtGte string) EntityPermissionQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *permissionEntityQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *permissionEntityQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *permissionEntityQueryImplementation) SetCreatedAtLte(createdAtLte string) EntityPermissionQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *permissionEntityQueryImplementation) HasEntityType() bool {
	return c.hasProperty("entity_type")
}

func (c *permissionEntityQueryImplementation) EntityType() string {
	if !c.HasEntityType() {
		return ""
	}

	return c.properties["entity_type"].(string)
}

func (c *permissionEntityQueryImplementation) SetEntityType(entityType string) EntityPermissionQueryInterface {
	c.properties["entity_type"] = entityType

	return c
}

func (c *permissionEntityQueryImplementation) HasEntityID() bool {
	return c.hasProperty("entity_id")
}

func (c *permissionEntityQueryImplementation) EntityID() string {
	if !c.HasEntityID() {
		return ""
	}

	return c.properties["entity_id"].(string)
}

func (c *permissionEntityQueryImplementation) SetEntityID(entityID string) EntityPermissionQueryInterface {
	c.properties["entity_id"] = entityID

	return c
}

func (c *permissionEntityQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *permissionEntityQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *permissionEntityQueryImplementation) SetID(id string) EntityPermissionQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *permissionEntityQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *permissionEntityQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *permissionEntityQueryImplementation) SetIDIn(idIn []string) EntityPermissionQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *permissionEntityQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *permissionEntityQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *permissionEntityQueryImplementation) SetLimit(limit int) EntityPermissionQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *permissionEntityQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *permissionEntityQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *permissionEntityQueryImplementation) SetOffset(offset int) EntityPermissionQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *permissionEntityQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *permissionEntityQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *permissionEntityQueryImplementation) SetOrderBy(orderBy string) EntityPermissionQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *permissionEntityQueryImplementation) HasPermissionID() bool {
	return c.hasProperty("permission_id")
}

func (c *permissionEntityQueryImplementation) PermissionID() string {
	if !c.HasPermissionID() {
		return ""
	}

	return c.properties["permission_id"].(string)
}

func (c *permissionEntityQueryImplementation) SetPermissionID(permissionID string) EntityPermissionQueryInterface {
	c.properties["permission_id"] = permissionID

	return c
}

func (c *permissionEntityQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *permissionEntityQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *permissionEntityQueryImplementation) SetSortDirection(sortDirection string) EntityPermissionQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *permissionEntityQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *permissionEntityQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *permissionEntityQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) EntityPermissionQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *permissionEntityQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *permissionEntityQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *permissionEntityQueryImplementation) SetTitleLike(titleLike string) EntityPermissionQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *permissionEntityQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
