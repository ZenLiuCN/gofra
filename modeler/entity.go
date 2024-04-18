package modeler

import (
	"context"
	"database/sql"
	"github.com/ZenLiuCN/fn"
	"github.com/ZenLiuCN/goinfra/utils"
	"log/slog"
	"time"
)

var (
	ByteBuffers = utils.NewByteBufferPool()
)

type Model[ID comparable] interface {
	Model()
}
type Executor interface {
	QueryOne(out any, sql string, args map[string]any) error
	QueryMany(out any, sql string, args map[string]any) error
	Execute(sql string, args map[string]any) (sql.Result, error)
	Close(ctx context.Context) bool
}
type SQLMaker interface {
	Query() string
	Update(fields ...FIELD) string
	Delete() string
	Drop() string
}
type Entity[E Model[ID], ID comparable] interface {
	Save() bool             //update record
	SaveBy(actor ID) bool   //update record
	DeleteBy(actor ID) bool //soft delete record
	Delete() bool           //soft delete record
	Drop() bool             //drop record permanently
	Refresh() bool
	Close(ctx context.Context) bool
}
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

type Configurer int

func (c Configurer) IsModifier() bool {
	return int(c)&int(ConfigurerModifier) > 0
}
func (c Configurer) IsModified() bool {
	return int(c)&int(ConfigurerModified) > 0
}
func (c Configurer) IsSoftRemoved() bool {
	return int(c)&int(ConfigurerSoftRemoved) > 0
}
func (c Configurer) IsVersioned() bool {
	return int(c)&int(ConfigurerVersion) > 0
}
func MakeConfigurer(flags ...ConfigurerFlag) Configurer {
	c := 0
	for _, flag := range flags {
		c |= int(flag)
	}
	return Configurer(c)
}

type ConfigurerFlag int

var (
	ConfigurerAll = Configurer(0b00001111)
)

const (

	// ConfigurerModifier flag that requires record who took action
	ConfigurerModifier ConfigurerFlag = 1 << iota
	// ConfigurerModified flag that requires manually recording update timestamp
	ConfigurerModified
	// ConfigurerSoftRemoved flag that support soft remove via [Model].Delete
	ConfigurerSoftRemoved
	// ConfigurerVersion flag that support optimistic lock with Version field
	ConfigurerVersion
)

type conf struct {
	softRemoved       bool
	useModifiedBy     bool
	manualUpdateTime  bool
	optimisticLocking bool
}

func newConf(config Configurer) conf {
	return conf{
		softRemoved:       config.IsSoftRemoved(),
		useModifiedBy:     config.IsModifier(),
		manualUpdateTime:  config.IsModified(),
		optimisticLocking: config.IsVersioned(),
	}
}

type BaseSQLMaker[ID comparable, E Model[ID]] struct {
	table string
	conf
	fields EntityInfo[ID, E]
}

func NewBaseSQLMaker[ID comparable, E Model[ID]](table string,
	config Configurer,
	fields EntityInfo[ID, E],
) BaseSQLMaker[ID, E] {
	return BaseSQLMaker[ID, E]{
		table:  table,
		conf:   newConf(config),
		fields: fields,
	}
}

func (b BaseSQLMaker[ID, E]) Query() string {
	q := ByteBuffers.Get()
	defer ByteBuffers.Put(q)
	q.WriteString("SELECT * FROM ")
	q.WriteString(b.table)
	q.WriteString(" WHERE ")
	q.WriteString(b.fields.IdName())
	q.WriteString("=:")
	q.WriteString(b.fields.IdName())
	if b.softRemoved {
		q.WriteString(b.fields.RemovedName())
		q.WriteString("= false")
	}
	return q.String()
}

func (b BaseSQLMaker[ID, E]) Update(fields ...FIELD) string {
	if len(fields) == 0 {
		panic("no field provided")
	}
	q := ByteBuffers.Get()
	defer ByteBuffers.Put(q)
	q.WriteString("UPDATE ")
	q.WriteString(b.table)
	q.WriteString(" SET ")
	for i, field := range fields {
		if i > 0 {
			q.WriteByte(',')
		}
		q.WriteString(b.fields[field].Name)
		q.WriteString("= :")
		q.WriteString(b.fields[field].Name)
	}
	if b.useModifiedBy {
		q.WriteByte(',')
		q.WriteString(b.fields.ModifiedByName())
		q.WriteString("= :")
		q.WriteString(b.fields.ModifiedByName())
	}
	if b.manualUpdateTime {
		q.WriteByte(',')
		q.WriteString(b.fields.ModifiedAtName())
		q.WriteString("= CURRENT_TIMESTAMP")
	}
	if b.optimisticLocking {
		q.WriteByte(',')
		q.WriteString(b.fields.VersionName())
		q.WriteString("=")
		q.WriteString(b.fields.VersionName())
		q.WriteString("+1")
	}
	q.WriteString(" WHERE ")
	q.WriteString(b.fields.IdName())
	q.WriteString("= :")
	q.WriteString(b.fields.IdName())
	if b.softRemoved {
		q.WriteString(" AND ")
		q.WriteString(b.fields.RemovedName())
		q.WriteString("= false")
	}
	if b.optimisticLocking {
		q.WriteString(" AND ")
		q.WriteString(b.fields.VersionName())
		q.WriteString("= :")
		q.WriteString(b.fields.VersionName())
	}
	return q.String()
}

func (b BaseSQLMaker[ID, E]) Delete() string {
	q := ByteBuffers.Get()
	defer ByteBuffers.Put(q)
	q.WriteString("UPDATE ")
	q.WriteString(b.table)
	q.WriteString(" SET ")
	q.WriteString(b.fields.RemovedName())
	q.WriteString("= true")
	if b.useModifiedBy {
		q.WriteByte(',')
		q.WriteString(b.fields.ModifiedByName())
		q.WriteString("= :")
		q.WriteString(b.fields.ModifiedByName())
	}
	if b.manualUpdateTime {
		q.WriteByte(',')
		q.WriteString(b.fields.ModifiedAtName())
		q.WriteString("= CURRENT_TIMESTAMP")
	}
	if b.optimisticLocking {
		q.WriteByte(',')
		q.WriteString(b.fields.VersionName())
		q.WriteString("=")
		q.WriteString(b.fields.VersionName())
		q.WriteString("+1")
	}
	q.WriteString(" WHERE ")
	q.WriteString(b.fields.IdName())
	q.WriteString("= :")
	q.WriteString(b.fields.IdName())
	if b.softRemoved {
		q.WriteString(" AND ")
		q.WriteString(b.fields.RemovedName())
		q.WriteString("= false")
	}
	if b.optimisticLocking {
		q.WriteString(" AND ")
		q.WriteString(b.fields.VersionName())
		q.WriteString("= :")
		q.WriteString(b.fields.VersionName())
	}
	return q.String()
}

func (b BaseSQLMaker[ID, E]) Drop() string {
	q := ByteBuffers.Get()
	defer ByteBuffers.Put(q)
	q.WriteString("DELETE ")
	q.WriteString(b.table)
	q.WriteString(" WHERE ")
	q.WriteString(b.fields.IdName())
	q.WriteString("= :")
	q.WriteString(b.fields.IdName())
	return q.String()
}

type BaseEntity[ID comparable, E Model[ID]] struct {
	conf     conf
	pointer  *E
	fields   EntityInfo[ID, E]
	modified map[FIELD]any
	executor Executor
	dml      SQLMaker
	closer   func(ctx context.Context) bool
	invalid  bool
}

func NewBaseEntity[ID comparable, E Model[ID]](
	conf Configurer,
	pointer *E,
	fields EntityInfo[ID, E],
	executor Executor,
	dml SQLMaker,
	closer func(ctx context.Context) bool,
) BaseEntity[ID, E] {
	return BaseEntity[ID, E]{
		conf:     newConf(conf),
		pointer:  pointer,
		fields:   fields,
		executor: executor,
		dml:      dml,
		closer:   closer,
	}
}
func (s *BaseEntity[ID, E]) Refresh() bool {
	if s.invalid {
		return false
	}
	q := s.dml.Query()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	err := s.executor.QueryOne(s.pointer, q, m)
	if err == nil {
		return true
	}
	slog.With("refresh", slog.String("query", q), slog.Any("parameter", m)).Error("refresh entity", err)
	return false
}
func (s *BaseEntity[ID, E]) DoModify(f FIELD, v any) bool {
	if s.invalid {
		return false
	}
	if f == FIELD_ID {
		return false
	}
	if s.modified == nil {
		s.modified = make(map[FIELD]any, len(s.fields))
	}
	s.modified[f] = v
	return false
}
func (s *BaseEntity[ID, E]) IsInvalid() bool {
	return s.invalid
}
func (s *BaseEntity[ID, E]) Save() bool {
	if s.invalid {
		return false
	}
	if len(s.modified) == 0 {
		return false
	}
	q := s.dml.Update(fn.MapKeys(s.modified)...)
	m := make(map[string]any, len(s.modified))
	for i, a := range s.modified {
		m[s.fields[i].Name] = a
	}
	if s.conf.optimisticLocking {
		s.fields.Entry(s.pointer, FIELD_VERSION, m)
	}
	r, err := s.executor.Execute(q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			for i := range s.modified {
				s.fields.Set(s.pointer, i, s.modified[i])
				delete(s.modified, i)
			}
			return true
		} else if n != 1 {
			slog.With("save", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity not effect one record")
			return false
		}

	}
	slog.With("save", slog.String("query", q), slog.Any("parameter", m)).Error("save modification", err)
	return false
}
func (s *BaseEntity[ID, E]) SaveBy(actor ID) bool {
	if s.invalid {
		return false
	}
	if len(s.modified) == 0 {
		return false
	}
	q := s.dml.Update(fn.MapKeys(s.modified)...)
	m := make(map[string]any, len(s.modified))
	for i, a := range s.modified {
		m[s.fields[i].Name] = a
	}
	if s.conf.optimisticLocking {
		s.fields.Entry(s.pointer, FIELD_VERSION, m)
	}
	if s.conf.useModifiedBy {
		m[s.fields.ModifiedByName()] = actor
	}
	r, err := s.executor.Execute(q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			for i := range s.modified {
				delete(s.modified, i)
			}
			return true
		} else if n != 1 {
			slog.With("save", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity not effect one record")
			return false
		}

	}
	slog.With("save", slog.String("query", q), slog.Any("parameter", m)).Error("save modification", err)
	return false
}
func (s *BaseEntity[ID, E]) Delete() bool {
	if s.invalid {
		return false
	}
	q := s.dml.Delete()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	if s.conf.optimisticLocking {
		s.fields.Entry(s.pointer, FIELD_VERSION, m)
	}
	r, err := s.executor.Execute(q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			slog.With("delete", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity not effect one record")
			return false
		}
	}
	slog.With("delete", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity", err)
	return false
}
func (s *BaseEntity[ID, E]) DeleteBy(actor ID) bool {
	if s.invalid {
		return false
	}
	q := s.dml.Delete()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	if s.conf.optimisticLocking {
		s.fields.Entry(s.pointer, FIELD_VERSION, m)
	}
	if s.conf.useModifiedBy {
		m[s.fields.ModifiedByName()] = actor
	}
	r, err := s.executor.Execute(q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			slog.With("delete", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity not effect one record")
			return false
		}
	}
	slog.With("delete", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity", err)
	return false
}
func (s *BaseEntity[ID, E]) Drop() bool {
	if s.invalid {
		return false
	}
	q := s.dml.Drop()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	r, err := s.executor.Execute(q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			slog.With("drop", slog.String("query", q), slog.Any("parameter", m)).Error("delete entity not effect one record")
			return false
		}
	}
	slog.With("drop", slog.String("query", q), slog.Any("parameter", m)).Error("drop entity", err)
	return false
}
func (s *BaseEntity[ID, E]) Close(ctx context.Context) bool {
	if len(s.modified) > 0 {
		s.Save()
	}
	return s.closer(ctx)
}

type FullModel[ID comparable] struct {
	Identified[ID]
	Auditable[ID]
	Versioned
	SoftRemoved
}

func (f FullModel[ID]) Model() {

}

type Identified[ID comparable] struct {
	Id ID `db:"id"`
}
type Traceable struct {
	CreateAt   time.Time `db:"create_at"`
	ModifiedAt time.Time `db:"modified_at"`
}
type Auditable[ID comparable] struct {
	Traceable
	CreateBy   ID `db:"create_by"`
	ModifiedBy ID `db:"modified_by"`
}
type Versioned struct {
	Version int `db:"version"`
}
type SoftRemoved struct {
	Removed bool `db:"removed"`
}
