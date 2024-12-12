package permissionstore

import "errors"

type PermissionQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) PermissionQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) PermissionQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) PermissionQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) PermissionQueryInterface

	HasHandle() bool
	Handle() string
	SetHandle(handle string) PermissionQueryInterface

	HasID() bool
	ID() string
	SetID(id string) PermissionQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) PermissionQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) PermissionQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) PermissionQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) PermissionQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) PermissionQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) PermissionQueryInterface

	HasStatus() bool
	Status() string
	SetStatus(status string) PermissionQueryInterface

	HasStatusIn() bool
	StatusIn() []string
	SetStatusIn(statusIn []string) PermissionQueryInterface

	HasTitleLike() bool
	TitleLike() string
	SetTitleLike(titleLike string) PermissionQueryInterface

	hasProperty(name string) bool
}

func NewPermissionQuery() PermissionQueryInterface {
	return &permissionQueryImplementation{
		properties: make(map[string]any),
	}
}

type permissionQueryImplementation struct {
	properties map[string]any
}

func (c *permissionQueryImplementation) Validate() error {
	if c.HasID() && c.ID() == "" {
		return errors.New("permission query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("permission query. id_in cannot be empty")
	}

	if c.HasStatus() && c.Status() == "" {
		return errors.New("permission query. status cannot be empty")
	}

	if c.HasTitleLike() && c.TitleLike() == "" {
		return errors.New("permission query. title_like cannot be empty")
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

func (c *permissionQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *permissionQueryImplementation) SetColumns(columns []string) PermissionQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *permissionQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *permissionQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *permissionQueryImplementation) SetCountOnly(countOnly bool) PermissionQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *permissionQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *permissionQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *permissionQueryImplementation) SetCreatedAtGte(createdAtGte string) PermissionQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *permissionQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *permissionQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *permissionQueryImplementation) SetCreatedAtLte(createdAtLte string) PermissionQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *permissionQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *permissionQueryImplementation) HasHandle() bool {
	return c.hasProperty("handle")
}

func (c *permissionQueryImplementation) Handle() string {
	if !c.HasHandle() {
		return ""
	}

	return c.properties["handle"].(string)
}

func (c *permissionQueryImplementation) SetHandle(handle string) PermissionQueryInterface {
	c.properties["handle"] = handle

	return c
}

func (c *permissionQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *permissionQueryImplementation) SetID(id string) PermissionQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *permissionQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *permissionQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *permissionQueryImplementation) SetIDIn(idIn []string) PermissionQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *permissionQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *permissionQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *permissionQueryImplementation) SetLimit(limit int) PermissionQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *permissionQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *permissionQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *permissionQueryImplementation) SetOffset(offset int) PermissionQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *permissionQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *permissionQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *permissionQueryImplementation) SetOrderBy(orderBy string) PermissionQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *permissionQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *permissionQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *permissionQueryImplementation) SetSortDirection(sortDirection string) PermissionQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *permissionQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *permissionQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *permissionQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) PermissionQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *permissionQueryImplementation) HasStatus() bool {
	return c.hasProperty("status")
}

func (c *permissionQueryImplementation) Status() string {
	if !c.HasStatus() {
		return ""
	}

	return c.properties["status"].(string)
}

func (c *permissionQueryImplementation) SetStatus(status string) PermissionQueryInterface {
	c.properties["status"] = status

	return c
}

func (c *permissionQueryImplementation) HasStatusIn() bool {
	return c.hasProperty("status_in")
}

func (c *permissionQueryImplementation) StatusIn() []string {
	if !c.HasStatusIn() {
		return []string{}
	}

	return c.properties["status_in"].([]string)
}

func (c *permissionQueryImplementation) SetStatusIn(statusIn []string) PermissionQueryInterface {
	c.properties["status_in"] = statusIn

	return c
}

func (c *permissionQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *permissionQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *permissionQueryImplementation) SetTitleLike(titleLike string) PermissionQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *permissionQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
