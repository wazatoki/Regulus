// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Maker is an object representing the database table.
type Maker struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Del           bool        `boil:"del" json:"del" toml:"del" yaml:"del"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	CreStaffID    null.String `boil:"cre_staff_id" json:"cre_staff_id,omitempty" toml:"cre_staff_id" yaml:"cre_staff_id,omitempty"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UpdateStaffID null.String `boil:"update_staff_id" json:"update_staff_id,omitempty" toml:"update_staff_id" yaml:"update_staff_id,omitempty"`
	Name          string      `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *makerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L makerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MakerColumns = struct {
	ID            string
	Del           string
	CreatedAt     string
	CreStaffID    string
	UpdatedAt     string
	UpdateStaffID string
	Name          string
}{
	ID:            "id",
	Del:           "del",
	CreatedAt:     "created_at",
	CreStaffID:    "cre_staff_id",
	UpdatedAt:     "updated_at",
	UpdateStaffID: "update_staff_id",
	Name:          "name",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MakerWhere = struct {
	ID            whereHelperstring
	Del           whereHelperbool
	CreatedAt     whereHelpernull_Time
	CreStaffID    whereHelpernull_String
	UpdatedAt     whereHelpernull_Time
	UpdateStaffID whereHelpernull_String
	Name          whereHelperstring
}{
	ID:            whereHelperstring{field: "\"makers\".\"id\""},
	Del:           whereHelperbool{field: "\"makers\".\"del\""},
	CreatedAt:     whereHelpernull_Time{field: "\"makers\".\"created_at\""},
	CreStaffID:    whereHelpernull_String{field: "\"makers\".\"cre_staff_id\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"makers\".\"updated_at\""},
	UpdateStaffID: whereHelpernull_String{field: "\"makers\".\"update_staff_id\""},
	Name:          whereHelperstring{field: "\"makers\".\"name\""},
}

// MakerRels is where relationship names are stored.
var MakerRels = struct {
}{}

// makerR is where relationships are stored.
type makerR struct {
}

// NewStruct creates a new relationship struct
func (*makerR) NewStruct() *makerR {
	return &makerR{}
}

// makerL is where Load methods for each relationship are stored.
type makerL struct{}

var (
	makerAllColumns            = []string{"id", "del", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "name"}
	makerColumnsWithoutDefault = []string{"id", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "name"}
	makerColumnsWithDefault    = []string{"del"}
	makerPrimaryKeyColumns     = []string{"id"}
)

type (
	// MakerSlice is an alias for a slice of pointers to Maker.
	// This should generally be used opposed to []Maker.
	MakerSlice []*Maker
	// MakerHook is the signature for custom Maker hook methods
	MakerHook func(context.Context, boil.ContextExecutor, *Maker) error

	makerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	makerType                 = reflect.TypeOf(&Maker{})
	makerMapping              = queries.MakeStructMapping(makerType)
	makerPrimaryKeyMapping, _ = queries.BindMapping(makerType, makerMapping, makerPrimaryKeyColumns)
	makerInsertCacheMut       sync.RWMutex
	makerInsertCache          = make(map[string]insertCache)
	makerUpdateCacheMut       sync.RWMutex
	makerUpdateCache          = make(map[string]updateCache)
	makerUpsertCacheMut       sync.RWMutex
	makerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var makerBeforeInsertHooks []MakerHook
var makerBeforeUpdateHooks []MakerHook
var makerBeforeDeleteHooks []MakerHook
var makerBeforeUpsertHooks []MakerHook

var makerAfterInsertHooks []MakerHook
var makerAfterSelectHooks []MakerHook
var makerAfterUpdateHooks []MakerHook
var makerAfterDeleteHooks []MakerHook
var makerAfterUpsertHooks []MakerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Maker) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Maker) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Maker) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Maker) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Maker) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Maker) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Maker) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Maker) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Maker) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range makerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMakerHook registers your hook function for all future operations.
func AddMakerHook(hookPoint boil.HookPoint, makerHook MakerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		makerBeforeInsertHooks = append(makerBeforeInsertHooks, makerHook)
	case boil.BeforeUpdateHook:
		makerBeforeUpdateHooks = append(makerBeforeUpdateHooks, makerHook)
	case boil.BeforeDeleteHook:
		makerBeforeDeleteHooks = append(makerBeforeDeleteHooks, makerHook)
	case boil.BeforeUpsertHook:
		makerBeforeUpsertHooks = append(makerBeforeUpsertHooks, makerHook)
	case boil.AfterInsertHook:
		makerAfterInsertHooks = append(makerAfterInsertHooks, makerHook)
	case boil.AfterSelectHook:
		makerAfterSelectHooks = append(makerAfterSelectHooks, makerHook)
	case boil.AfterUpdateHook:
		makerAfterUpdateHooks = append(makerAfterUpdateHooks, makerHook)
	case boil.AfterDeleteHook:
		makerAfterDeleteHooks = append(makerAfterDeleteHooks, makerHook)
	case boil.AfterUpsertHook:
		makerAfterUpsertHooks = append(makerAfterUpsertHooks, makerHook)
	}
}

// One returns a single maker record from the query.
func (q makerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Maker, error) {
	o := &Maker{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for makers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Maker records from the query.
func (q makerQuery) All(ctx context.Context, exec boil.ContextExecutor) (MakerSlice, error) {
	var o []*Maker

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to Maker slice")
	}

	if len(makerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Maker records in the query.
func (q makerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count makers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q makerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if makers exists")
	}

	return count > 0, nil
}

// Makers retrieves all the records using an executor.
func Makers(mods ...qm.QueryMod) makerQuery {
	mods = append(mods, qm.From("\"makers\""))
	return makerQuery{NewQuery(mods...)}
}

// FindMaker retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMaker(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Maker, error) {
	makerObj := &Maker{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"makers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, makerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from makers")
	}

	return makerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Maker) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no makers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(makerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	makerInsertCacheMut.RLock()
	cache, cached := makerInsertCache[key]
	makerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			makerAllColumns,
			makerColumnsWithDefault,
			makerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(makerType, makerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(makerType, makerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"makers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"makers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into makers")
	}

	if !cached {
		makerInsertCacheMut.Lock()
		makerInsertCache[key] = cache
		makerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Maker.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Maker) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	makerUpdateCacheMut.RLock()
	cache, cached := makerUpdateCache[key]
	makerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			makerAllColumns,
			makerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboiler: unable to update makers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"makers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, makerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(makerType, makerMapping, append(wl, makerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update makers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by update for makers")
	}

	if !cached {
		makerUpdateCacheMut.Lock()
		makerUpdateCache[key] = cache
		makerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q makerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all for makers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected for makers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MakerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlboiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), makerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"makers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, makerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all in maker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected all in update all maker")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Maker) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no makers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(makerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	makerUpsertCacheMut.RLock()
	cache, cached := makerUpsertCache[key]
	makerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			makerAllColumns,
			makerColumnsWithDefault,
			makerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			makerAllColumns,
			makerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert makers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(makerPrimaryKeyColumns))
			copy(conflict, makerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"makers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(makerType, makerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(makerType, makerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to upsert makers")
	}

	if !cached {
		makerUpsertCacheMut.Lock()
		makerUpsertCache[key] = cache
		makerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Maker record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Maker) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboiler: no Maker provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), makerPrimaryKeyMapping)
	sql := "DELETE FROM \"makers\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete from makers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by delete for makers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q makerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboiler: no makerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from makers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for makers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MakerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(makerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), makerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"makers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, makerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from maker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for makers")
	}

	if len(makerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Maker) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMaker(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MakerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MakerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), makerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"makers\".* FROM \"makers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, makerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in MakerSlice")
	}

	*o = slice

	return nil
}

// MakerExists checks if the Maker row exists.
func MakerExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"makers\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if makers exists")
	}

	return exists, nil
}
