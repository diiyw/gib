package mysql

import "strings"

const (
	SelectTyp      = "SELECT "
	InsertTyp      = "INSERT INTO "
	ReplaceIntoTyp = "REPLACE INTO "
	DeleteTyp      = "DELETE FROM "
	UpdateType     = "UPDATE SET "
)

type Builder struct {
	table      *table
	cond       []*cond
	typ        *typ
	join       []*join
	columns    []string
	orderBy    []string
	groupBy    []string
	bindValues []interface{}
}

func (b *Builder) make() string {

	var sql strings.Builder
	sql.WriteString(b.typ.name)
	switch b.typ.name {
	case SelectTyp:

		sql.WriteString(strings.Join(b.typ.columns, ","))
		sql.WriteString(" FROM ")

		b.tableAlias(&sql, b.table)

		sql.WriteString(" ")
		for _, join := range b.join {

			sql.WriteString(join.typ)
			sql.WriteString(" ")

			b.tableAlias(&sql, join.table)

			sql.WriteString(" ON ")
			sql.WriteString(join.on.column)
			sql.WriteString(join.on.op)
			sql.WriteString(join.on.value.(string))
		}

	case InsertTyp, ReplaceIntoTyp:

		sql.WriteString(b.table.name)
		sql.WriteString("(")
		sql.WriteString(strings.Join(b.typ.columns, ","))
		sql.WriteString(")VALUES(")

		if len(b.typ.values) != 0 {

			sql.WriteString("?")
			if len(b.bindValues) != 1 {
				for i := 0; i < len(b.bindValues)-1; i++ {
					sql.WriteString(",?")
				}
			}
		}

		sql.WriteString(")")
	case DeleteTyp:

		b.tableAlias(&sql, b.table)

	case UpdateType:

		sets := strings.Join(b.typ.columns, "=? ,")
		sql.WriteString(strings.Trim(sets, ","))
		b.bindValues = append(b.bindValues, b.typ.values)
	}

	if len(b.cond) != 0 {
		sql.WriteString(" WHERE ")

		sql.WriteString(b.cond[0].column)
		sql.WriteString(b.cond[0].op)
		sql.WriteString(" ? ")

		b.bindValues = append(b.bindValues, b.cond[0].value)

		for _, cond := range b.cond[1:] {

			sql.WriteString(cond.typ)
			sql.WriteString(cond.column)
			sql.WriteString(cond.op)
			sql.WriteString(" ? ")

			b.bindValues = append(b.bindValues, cond.value)

			if cond.op == " BETWEEN " {
				sql.WriteString(" AND ? ")
				b.bindValues = append(b.bindValues, cond.value)
			}
		}
	}
	return sql.String()
}

func (b *Builder) tableAlias(sql *strings.Builder, t *table) {
	sql.WriteString(t.name)
	if t.alias != "" {
		sql.WriteString(" AS " + t.alias)
	}
}

type table struct {
	name, alias string
}

type join struct {
	typ   string
	table *table
	on    *cond
}

type typ struct {
	name    string
	columns []string
	values  []interface{}
}

type cond struct {
	op     string
	column string
	value  interface{}
	typ    string
}

func Table(name, as string) *table {
	t := new(table)
	t.name = name
	if len(name) == 2 {
		t.alias = as
	}
	return t
}

func And(column, op string, value interface{}) *cond {
	return &cond{column: column, op: op, value: value, typ: " AND "}
}

func Or(column, op string, value interface{}) *cond {
	return &cond{column: column, op: op, value: value, typ: " OR "}
}

func AndBetween(column string, value interface{}) *cond {
	return &cond{column: column, op: " BETWEEN ", value: value, typ: " AND "}
}

func OrBetween(column string, value interface{}) *cond {
	return &cond{column: column, op: " BETWEEN ", value: value, typ: " OR "}
}

func On(column, op, nextColumn string) *cond {
	return &cond{column: column, op: op, value: nextColumn, typ: " ON "}
}

func Select(columns ...string) *typ {
	return &typ{name: SelectTyp, columns: columns}
}

func Insert(values map[string]interface{}) *typ {
	return newType(InsertTyp, values)
}

func ReplaceInto(values map[string]interface{}) *typ {
	return newType(ReplaceIntoTyp, values)
}

func Delete() *typ {
	return &typ{name: DeleteTyp}
}

func Update(sets map[string]interface{}) *typ {

	return newType(UpdateType, sets)
}

func newType(name string, sets map[string]interface{}) *typ {
	keys := make([]string, 0)
	values := make([]interface{}, 0)
	for key, value := range sets {
		keys = append(keys, key)
		values = append(values, value)
	}
	return &typ{name: name, columns: keys, values: values}
}

func Join(table *table, c *cond, typ string) *join {
	return &join{typ: typ, table: table, on: c}
}

func LeftJoin(table *table, c *cond) *join {
	return Join(table, c, "LEFT")
}
