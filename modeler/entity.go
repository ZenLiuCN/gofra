package modeler

import (
	"context"
	"database/sql"
	"github.com/ZenLiuCN/fn"
	cfg "github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/gofra/units"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	ByteBuffers = units.NewByteBufferPool()
)

type Model[ID comparable] interface {
	Model()
}
type Executor interface {
	QueryOne(ctx context.Context, out any, sql string, args map[string]any) error
	Execute(ctx context.Context, sql string, args map[string]any) (sql.Result, error)
	Close(ctx context.Context) bool
}
type SQLMaker interface {
	Query() string
	Update(fields ...FIELD) string
	Delete() string
	Drop() string
}
type Entity[E Model[ID], ID comparable] interface {
	Save(ctx context.Context) bool               //update record
	SaveBy(ctx context.Context, actor ID) bool   //update record
	DeleteBy(ctx context.Context, actor ID) bool //soft delete record
	Delete(ctx context.Context) bool             //soft delete record
	Drop(ctx context.Context) bool               //drop record permanently
	Refresh(ctx context.Context) bool
	Close(ctx context.Context) bool
}

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
func (s *BaseEntity[ID, E]) Refresh(ctx context.Context) bool {
	if s.invalid {
		return false
	}
	q := s.dml.Query()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	err := s.executor.QueryOne(ctx, s.pointer, q, m)
	if err == nil {
		return true
	}
	cfg.Internal().Error("refresh entity", "error", err, "query", q, "parameter", m)
	return false
}
func (s *BaseEntity[ID, E]) DoWrite(f FIELD, v any) bool {
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
func (s *BaseEntity[ID, E]) DoRead(f FIELD) any {
	if v, ok := s.modified[f]; ok {
		return v
	}
	return s.fields.Get(s.pointer, f)
}
func (s *BaseEntity[ID, E]) IsModified() bool {
	return len(s.modified) > 0
}
func (s *BaseEntity[ID, E]) IsInvalid() bool {
	return s.invalid
}
func (s *BaseEntity[ID, E]) Save(ctx context.Context) bool {
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
	r, err := s.executor.Execute(ctx, q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			for i := range s.modified {
				s.fields.Set(s.pointer, i, s.modified[i])
				delete(s.modified, i)
			}
			return true
		} else if n != 1 {
			cfg.Internal().Error("delete entity not effect one record", "query", q, "parameter", m)
			return false
		}

	}
	cfg.Internal().Error("save modification", "error", err, "query", q, "parameter", m)
	return false
}
func (s *BaseEntity[ID, E]) SaveBy(ctx context.Context, actor ID) bool {
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
	r, err := s.executor.Execute(ctx, q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			for i := range s.modified {
				delete(s.modified, i)
			}
			return true
		} else if n != 1 {
			cfg.Internal().Error("delete entity not effect one record", "query", q, "parameter", m)
			return false
		}

	}
	cfg.Internal().Error("save modification", "error", err, "query", q, "parameter", m)
	return false
}
func (s *BaseEntity[ID, E]) Delete(ctx context.Context) bool {
	if s.invalid {
		return false
	}
	q := s.dml.Delete()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	if s.conf.optimisticLocking {
		s.fields.Entry(s.pointer, FIELD_VERSION, m)
	}
	r, err := s.executor.Execute(ctx, q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			cfg.Internal().Error("delete entity not effect one record", "query", q, "parameter", m)
			return false
		}
	}
	cfg.Internal().Error("delete entity", "error", err, "query", q, "parameter", m)
	return false
}
func (s *BaseEntity[ID, E]) DeleteBy(ctx context.Context, actor ID) bool {
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
	r, err := s.executor.Execute(ctx, q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			cfg.Internal().Error("delete entity not effect one record", "query", q, "parameter", m)
			return false
		}
	}
	cfg.Internal().Error("delete entity", "error", err, "query", q, "parameter", m)
	return false
}
func (s *BaseEntity[ID, E]) Drop(ctx context.Context) bool {
	if s.invalid {
		return false
	}
	q := s.dml.Drop()
	m := map[string]any{s.fields.IdName(): s.fields[FIELD_ID].Getter(s.pointer)}
	r, err := s.executor.Execute(ctx, q, m)
	if err == nil {
		var n int64
		if n, err = r.RowsAffected(); err == nil && n == 1 {
			s.invalid = true
			return true
		} else if n != 1 {
			cfg.Internal().Error("delete entity not effect one record", "query", q, "parameter", m)
			return false
		}
	}
	cfg.Internal().Error("drop entity", "query", q, "parameter", m, "error", err)
	return false
}
func (s *BaseEntity[ID, E]) Close(ctx context.Context) bool {
	if len(s.modified) > 0 {
		s.Save(ctx)
	}
	return s.closer(ctx)
}

type FullModelEntity[ID comparable] struct {
	Identified[ID]
	Auditable[ID]
	Versioned
	SoftRemoved
}

func (f FullModelEntity[ID]) Model() {

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

type SqlxExecutor struct {
	*sqlx.DB
}

func (s SqlxExecutor) QueryOne(ctx context.Context, out any, sql string, args map[string]any) (err error) {
	var r *sqlx.Rows
	if len(args) == 0 {
		r, err = s.QueryxContext(ctx, sql)
	} else {
		r, err = s.NamedQueryContext(ctx, sql, args)
	}
	if err != nil {
		return err
	}
	if r.Next() {
		err = r.StructScan(out)
	}
	return err
}

func (s SqlxExecutor) Execute(ctx context.Context, q string, args map[string]any) (r sql.Result, err error) {
	if len(args) == 0 {
		r, err = s.ExecContext(ctx, q)
	} else {
		r, err = s.NamedExecContext(ctx, q, args)
	}
	return r, err
}

func (s SqlxExecutor) Close(ctx context.Context) bool {
	err := s.DB.Close()
	if err == nil {
		return true
	}
	return false
}
