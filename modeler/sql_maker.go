package modeler

import (
	"strings"
)

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
	return strings.Clone(q.String())
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
	return strings.Clone(q.String())
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
	return strings.Clone(q.String())
}

func (b BaseSQLMaker[ID, E]) Drop() string {
	q := ByteBuffers.Get()
	defer ByteBuffers.Put(q)
	q.WriteString("DELETE FROM ")
	q.WriteString(b.table)
	q.WriteString(" WHERE ")
	q.WriteString(b.fields.IdName())
	q.WriteString("= :")
	q.WriteString(b.fields.IdName())
	return strings.Clone(q.String())
}
