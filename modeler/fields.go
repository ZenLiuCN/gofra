package modeler

import "time"

type FIELD int
type FieldInfo[ID comparable, E Model[ID]] struct {
	Field  FIELD
	Name   string
	Getter func(*E) any
	Setter func(*E, any)
}
type EntityInfo[ID comparable, E Model[ID]] map[FIELD]FieldInfo[ID, E]

func (c EntityInfo[ID, E]) Entry(e *E, f FIELD, m map[string]any) {
	m[c[f].Name] = c[f].Getter(e)
}
func (c EntityInfo[ID, E]) Get(e *E, f FIELD) any {
	return c[f].Getter(e)
}
func (c EntityInfo[ID, E]) Set(e *E, f FIELD, v any) {
	c[f].Setter(e, v)
}
func (c EntityInfo[ID, E]) Name(f FIELD) string {
	return c[f].Name
}
func (c EntityInfo[ID, E]) IdName() string {
	return c[FIELD_ID].Name
}
func (c EntityInfo[ID, E]) GetId(e *E) ID {
	return c[FIELD_ID].Getter(e).(ID)
}

func (c EntityInfo[ID, E]) VersionName() string {
	return c[FIELD_VERSION].Name
}
func (c EntityInfo[ID, E]) GetVersion(e *E) int {
	v, ok := c[FIELD_VERSION].Getter(e).(int)
	if ok {
		return v
	}
	return -1
}
func (c EntityInfo[ID, E]) RemovedName() string {
	return c[FIELD_REMOVED].Name
}
func (c EntityInfo[ID, E]) GetRemoved(e *E) bool {
	v, ok := c[FIELD_REMOVED].Getter(e).(bool)
	if ok {
		return v
	}
	return true
}
func (c EntityInfo[ID, E]) ModifiedByName() string {
	return c[FIELD_MODIFIED_BY].Name
}
func (c EntityInfo[ID, E]) GetModifiedBy(e *E) ID {
	return c[FIELD_MODIFIED_BY].Getter(e).(ID)
}
func (c EntityInfo[ID, E]) ModifiedAtName() string {
	return c[FIELD_MODIFIED_AT].Name
}
func (c EntityInfo[ID, E]) GetModifiedAt(e *E) time.Time {
	v, ok := c[FIELD_MODIFIED_AT].Getter(e).(time.Time)
	if ok {
		return v
	}
	return time.Time{}
}
func (c EntityInfo[ID, E]) CreateByName() string {
	return c[FIELD_CREATE_BY].Name
}
func (c EntityInfo[ID, E]) GetCreateBy(e *E) ID {
	return c[FIELD_CREATE_BY].Getter(e).(ID)
}
func (c EntityInfo[ID, E]) CreateAtName() string {
	return c[FIELD_CREATE_AT].Name
}
func (c EntityInfo[ID, E]) GetCreateAt(e *E) time.Time {
	v, ok := c[FIELD_CREATE_AT].Getter(e).(time.Time)
	if ok {
		return v
	}
	return time.Time{}
}

var (
	BaseAllFields = map[FIELD]string{
		FIELD_ID:          "id",
		FIELD_CREATE_AT:   "create_at",
		FIELD_MODIFIED_AT: "modified_at",
		FIELD_REMOVED:     "removed",
		FIELD_VERSION:     "version",
		FIELD_CREATE_BY:   "create_by",
		FIELD_MODIFIED_BY: "modified_by",
	}
)

func EntityInfoBuilder[ID comparable, E Model[ID]](
	getter func(*E, FIELD) any,
	setter func(*E, FIELD, any),
	fields map[FIELD]string,
) (e EntityInfo[ID, E]) {
	e = EntityInfo[ID, E]{}
	for field, s := range fields {
		f := field
		e[f] = FieldInfo[ID, E]{
			Name:  s,
			Field: f,
			Getter: func(e *E) any {
				return getter(e, f)
			},
			Setter: func(e *E, v any) {
				setter(e, f, v)
			},
		}
	}
	return
}

const (
	FIELD_ALL FIELD = iota - 1
	FIELD_NONE
	FIELD_ID
	FIELD_CREATE_AT
	FIELD_MODIFIED_AT
	FIELD_REMOVED
	FIELD_VERSION
	FIELD_CREATE_BY
	FIELD_MODIFIED_BY
	FIELD_BUILTIN_MAX
)
