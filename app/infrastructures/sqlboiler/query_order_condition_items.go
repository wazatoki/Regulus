// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// QueryOrderConditionItem is an object representing the database table.
type QueryOrderConditionItem struct {
	ID                string      `db:"id" boil:"id" json:"id" toml:"id" yaml:"id"`
	Del               bool        `db:"del" boil:"del" json:"del" toml:"del" yaml:"del"`
	CreatedAt         null.Time   `db:"created_at" boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	CreStaffID        null.String `db:"cre_staff_id" boil:"cre_staff_id" json:"cre_staff_id,omitempty" toml:"cre_staff_id" yaml:"cre_staff_id,omitempty"`
	UpdatedAt         null.Time   `db:"updated_at" boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UpdateStaffID     null.String `db:"update_staff_id" boil:"update_staff_id" json:"update_staff_id,omitempty" toml:"update_staff_id" yaml:"update_staff_id,omitempty"`
	QueryConditionsID string      `db:"query_conditions_id" boil:"query_conditions_id" json:"query_conditions_id" toml:"query_conditions_id" yaml:"query_conditions_id"`
	OrderFieldID      string      `db:"order_field_id" boil:"order_field_id" json:"order_field_id" toml:"order_field_id" yaml:"order_field_id"`
	OrderFieldKeyWord string      `db:"order_field_key_word" boil:"order_field_key_word" json:"order_field_key_word" toml:"order_field_key_word" yaml:"order_field_key_word"`
	RowOrder          int         `db:"row_order" boil:"row_order" json:"row_order" toml:"row_order" yaml:"row_order"`

	R *queryOrderConditionItemR `db:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L queryOrderConditionItemL  `db:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QueryOrderConditionItemColumns = struct {
	ID                string
	Del               string
	CreatedAt         string
	CreStaffID        string
	UpdatedAt         string
	UpdateStaffID     string
	QueryConditionsID string
	OrderFieldID      string
	OrderFieldKeyWord string
	RowOrder          string
}{
	ID:                "id",
	Del:               "del",
	CreatedAt:         "created_at",
	CreStaffID:        "cre_staff_id",
	UpdatedAt:         "updated_at",
	UpdateStaffID:     "update_staff_id",
	QueryConditionsID: "query_conditions_id",
	OrderFieldID:      "order_field_id",
	OrderFieldKeyWord: "order_field_key_word",
	RowOrder:          "row_order",
}

var QueryOrderConditionItemTableColumns = struct {
	ID                string
	Del               string
	CreatedAt         string
	CreStaffID        string
	UpdatedAt         string
	UpdateStaffID     string
	QueryConditionsID string
	OrderFieldID      string
	OrderFieldKeyWord string
	RowOrder          string
}{
	ID:                "query_order_condition_items.id",
	Del:               "query_order_condition_items.del",
	CreatedAt:         "query_order_condition_items.created_at",
	CreStaffID:        "query_order_condition_items.cre_staff_id",
	UpdatedAt:         "query_order_condition_items.updated_at",
	UpdateStaffID:     "query_order_condition_items.update_staff_id",
	QueryConditionsID: "query_order_condition_items.query_conditions_id",
	OrderFieldID:      "query_order_condition_items.order_field_id",
	OrderFieldKeyWord: "query_order_condition_items.order_field_key_word",
	RowOrder:          "query_order_condition_items.row_order",
}

// Generated where

var QueryOrderConditionItemWhere = struct {
	ID                whereHelperstring
	Del               whereHelperbool
	CreatedAt         whereHelpernull_Time
	CreStaffID        whereHelpernull_String
	UpdatedAt         whereHelpernull_Time
	UpdateStaffID     whereHelpernull_String
	QueryConditionsID whereHelperstring
	OrderFieldID      whereHelperstring
	OrderFieldKeyWord whereHelperstring
	RowOrder          whereHelperint
}{
	ID:                whereHelperstring{field: "\"query_order_condition_items\".\"id\""},
	Del:               whereHelperbool{field: "\"query_order_condition_items\".\"del\""},
	CreatedAt:         whereHelpernull_Time{field: "\"query_order_condition_items\".\"created_at\""},
	CreStaffID:        whereHelpernull_String{field: "\"query_order_condition_items\".\"cre_staff_id\""},
	UpdatedAt:         whereHelpernull_Time{field: "\"query_order_condition_items\".\"updated_at\""},
	UpdateStaffID:     whereHelpernull_String{field: "\"query_order_condition_items\".\"update_staff_id\""},
	QueryConditionsID: whereHelperstring{field: "\"query_order_condition_items\".\"query_conditions_id\""},
	OrderFieldID:      whereHelperstring{field: "\"query_order_condition_items\".\"order_field_id\""},
	OrderFieldKeyWord: whereHelperstring{field: "\"query_order_condition_items\".\"order_field_key_word\""},
	RowOrder:          whereHelperint{field: "\"query_order_condition_items\".\"row_order\""},
}

// QueryOrderConditionItemRels is where relationship names are stored.
var QueryOrderConditionItemRels = struct {
	QueryCondition string
}{
	QueryCondition: "QueryCondition",
}

// queryOrderConditionItemR is where relationships are stored.
type queryOrderConditionItemR struct {
	QueryCondition *QueryCondition `db:"QueryCondition" boil:"QueryCondition" json:"QueryCondition" toml:"QueryCondition" yaml:"QueryCondition"`
}

// NewStruct creates a new relationship struct
func (*queryOrderConditionItemR) NewStruct() *queryOrderConditionItemR {
	return &queryOrderConditionItemR{}
}

// queryOrderConditionItemL is where Load methods for each relationship are stored.
type queryOrderConditionItemL struct{}

var (
	queryOrderConditionItemAllColumns            = []string{"id", "del", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "query_conditions_id", "order_field_id", "order_field_key_word", "row_order"}
	queryOrderConditionItemColumnsWithoutDefault = []string{"id", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "query_conditions_id", "order_field_id", "order_field_key_word", "row_order"}
	queryOrderConditionItemColumnsWithDefault    = []string{"del"}
	queryOrderConditionItemPrimaryKeyColumns     = []string{"id"}
)

type (
	// QueryOrderConditionItemSlice is an alias for a slice of pointers to QueryOrderConditionItem.
	// This should almost always be used instead of []QueryOrderConditionItem.
	QueryOrderConditionItemSlice []*QueryOrderConditionItem
	// QueryOrderConditionItemHook is the signature for custom QueryOrderConditionItem hook methods
	QueryOrderConditionItemHook func(context.Context, boil.ContextExecutor, *QueryOrderConditionItem) error

	queryOrderConditionItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	queryOrderConditionItemType                 = reflect.TypeOf(&QueryOrderConditionItem{})
	queryOrderConditionItemMapping              = queries.MakeStructMapping(queryOrderConditionItemType)
	queryOrderConditionItemPrimaryKeyMapping, _ = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, queryOrderConditionItemPrimaryKeyColumns)
	queryOrderConditionItemInsertCacheMut       sync.RWMutex
	queryOrderConditionItemInsertCache          = make(map[string]insertCache)
	queryOrderConditionItemUpdateCacheMut       sync.RWMutex
	queryOrderConditionItemUpdateCache          = make(map[string]updateCache)
	queryOrderConditionItemUpsertCacheMut       sync.RWMutex
	queryOrderConditionItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var queryOrderConditionItemBeforeInsertHooks []QueryOrderConditionItemHook
var queryOrderConditionItemBeforeUpdateHooks []QueryOrderConditionItemHook
var queryOrderConditionItemBeforeDeleteHooks []QueryOrderConditionItemHook
var queryOrderConditionItemBeforeUpsertHooks []QueryOrderConditionItemHook

var queryOrderConditionItemAfterInsertHooks []QueryOrderConditionItemHook
var queryOrderConditionItemAfterSelectHooks []QueryOrderConditionItemHook
var queryOrderConditionItemAfterUpdateHooks []QueryOrderConditionItemHook
var queryOrderConditionItemAfterDeleteHooks []QueryOrderConditionItemHook
var queryOrderConditionItemAfterUpsertHooks []QueryOrderConditionItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QueryOrderConditionItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QueryOrderConditionItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QueryOrderConditionItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QueryOrderConditionItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QueryOrderConditionItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QueryOrderConditionItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QueryOrderConditionItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QueryOrderConditionItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QueryOrderConditionItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range queryOrderConditionItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQueryOrderConditionItemHook registers your hook function for all future operations.
func AddQueryOrderConditionItemHook(hookPoint boil.HookPoint, queryOrderConditionItemHook QueryOrderConditionItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		queryOrderConditionItemBeforeInsertHooks = append(queryOrderConditionItemBeforeInsertHooks, queryOrderConditionItemHook)
	case boil.BeforeUpdateHook:
		queryOrderConditionItemBeforeUpdateHooks = append(queryOrderConditionItemBeforeUpdateHooks, queryOrderConditionItemHook)
	case boil.BeforeDeleteHook:
		queryOrderConditionItemBeforeDeleteHooks = append(queryOrderConditionItemBeforeDeleteHooks, queryOrderConditionItemHook)
	case boil.BeforeUpsertHook:
		queryOrderConditionItemBeforeUpsertHooks = append(queryOrderConditionItemBeforeUpsertHooks, queryOrderConditionItemHook)
	case boil.AfterInsertHook:
		queryOrderConditionItemAfterInsertHooks = append(queryOrderConditionItemAfterInsertHooks, queryOrderConditionItemHook)
	case boil.AfterSelectHook:
		queryOrderConditionItemAfterSelectHooks = append(queryOrderConditionItemAfterSelectHooks, queryOrderConditionItemHook)
	case boil.AfterUpdateHook:
		queryOrderConditionItemAfterUpdateHooks = append(queryOrderConditionItemAfterUpdateHooks, queryOrderConditionItemHook)
	case boil.AfterDeleteHook:
		queryOrderConditionItemAfterDeleteHooks = append(queryOrderConditionItemAfterDeleteHooks, queryOrderConditionItemHook)
	case boil.AfterUpsertHook:
		queryOrderConditionItemAfterUpsertHooks = append(queryOrderConditionItemAfterUpsertHooks, queryOrderConditionItemHook)
	}
}

// One returns a single queryOrderConditionItem record from the query.
func (q queryOrderConditionItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*QueryOrderConditionItem, error) {
	o := &QueryOrderConditionItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for query_order_condition_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QueryOrderConditionItem records from the query.
func (q queryOrderConditionItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (QueryOrderConditionItemSlice, error) {
	var o []*QueryOrderConditionItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to QueryOrderConditionItem slice")
	}

	if len(queryOrderConditionItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QueryOrderConditionItem records in the query.
func (q queryOrderConditionItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count query_order_condition_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q queryOrderConditionItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if query_order_condition_items exists")
	}

	return count > 0, nil
}

// QueryCondition pointed to by the foreign key.
func (o *QueryOrderConditionItem) QueryCondition(mods ...qm.QueryMod) queryConditionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.QueryConditionsID),
	}

	queryMods = append(queryMods, mods...)

	query := QueryConditions(queryMods...)
	queries.SetFrom(query.Query, "\"query_conditions\"")

	return query
}

// LoadQueryCondition allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (queryOrderConditionItemL) LoadQueryCondition(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQueryOrderConditionItem interface{}, mods queries.Applicator) error {
	var slice []*QueryOrderConditionItem
	var object *QueryOrderConditionItem

	if singular {
		object = maybeQueryOrderConditionItem.(*QueryOrderConditionItem)
	} else {
		slice = *maybeQueryOrderConditionItem.(*[]*QueryOrderConditionItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &queryOrderConditionItemR{}
		}
		args = append(args, object.QueryConditionsID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &queryOrderConditionItemR{}
			}

			for _, a := range args {
				if a == obj.QueryConditionsID {
					continue Outer
				}
			}

			args = append(args, obj.QueryConditionsID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`query_conditions`),
		qm.WhereIn(`query_conditions.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load QueryCondition")
	}

	var resultSlice []*QueryCondition
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice QueryCondition")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for query_conditions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for query_conditions")
	}

	if len(queryOrderConditionItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.QueryCondition = foreign
		if foreign.R == nil {
			foreign.R = &queryConditionR{}
		}
		foreign.R.QueryOrderConditionItems = append(foreign.R.QueryOrderConditionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QueryConditionsID == foreign.ID {
				local.R.QueryCondition = foreign
				if foreign.R == nil {
					foreign.R = &queryConditionR{}
				}
				foreign.R.QueryOrderConditionItems = append(foreign.R.QueryOrderConditionItems, local)
				break
			}
		}
	}

	return nil
}

// SetQueryCondition of the queryOrderConditionItem to the related item.
// Sets o.R.QueryCondition to related.
// Adds o to related.R.QueryOrderConditionItems.
func (o *QueryOrderConditionItem) SetQueryCondition(ctx context.Context, exec boil.ContextExecutor, insert bool, related *QueryCondition) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"query_order_condition_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"query_conditions_id"}),
		strmangle.WhereClause("\"", "\"", 2, queryOrderConditionItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.QueryConditionsID = related.ID
	if o.R == nil {
		o.R = &queryOrderConditionItemR{
			QueryCondition: related,
		}
	} else {
		o.R.QueryCondition = related
	}

	if related.R == nil {
		related.R = &queryConditionR{
			QueryOrderConditionItems: QueryOrderConditionItemSlice{o},
		}
	} else {
		related.R.QueryOrderConditionItems = append(related.R.QueryOrderConditionItems, o)
	}

	return nil
}

// QueryOrderConditionItems retrieves all the records using an executor.
func QueryOrderConditionItems(mods ...qm.QueryMod) queryOrderConditionItemQuery {
	mods = append(mods, qm.From("\"query_order_condition_items\""))
	return queryOrderConditionItemQuery{NewQuery(mods...)}
}

// FindQueryOrderConditionItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQueryOrderConditionItem(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*QueryOrderConditionItem, error) {
	queryOrderConditionItemObj := &QueryOrderConditionItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"query_order_condition_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, queryOrderConditionItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from query_order_condition_items")
	}

	if err = queryOrderConditionItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return queryOrderConditionItemObj, err
	}

	return queryOrderConditionItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QueryOrderConditionItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no query_order_condition_items provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(queryOrderConditionItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	queryOrderConditionItemInsertCacheMut.RLock()
	cache, cached := queryOrderConditionItemInsertCache[key]
	queryOrderConditionItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			queryOrderConditionItemAllColumns,
			queryOrderConditionItemColumnsWithDefault,
			queryOrderConditionItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"query_order_condition_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"query_order_condition_items\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into query_order_condition_items")
	}

	if !cached {
		queryOrderConditionItemInsertCacheMut.Lock()
		queryOrderConditionItemInsertCache[key] = cache
		queryOrderConditionItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the QueryOrderConditionItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QueryOrderConditionItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	queryOrderConditionItemUpdateCacheMut.RLock()
	cache, cached := queryOrderConditionItemUpdateCache[key]
	queryOrderConditionItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			queryOrderConditionItemAllColumns,
			queryOrderConditionItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboiler: unable to update query_order_condition_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"query_order_condition_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, queryOrderConditionItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, append(wl, queryOrderConditionItemPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update query_order_condition_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by update for query_order_condition_items")
	}

	if !cached {
		queryOrderConditionItemUpdateCacheMut.Lock()
		queryOrderConditionItemUpdateCache[key] = cache
		queryOrderConditionItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q queryOrderConditionItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all for query_order_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected for query_order_condition_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QueryOrderConditionItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), queryOrderConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"query_order_condition_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, queryOrderConditionItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all in queryOrderConditionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected all in update all queryOrderConditionItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QueryOrderConditionItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no query_order_condition_items provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(queryOrderConditionItemColumnsWithDefault, o)

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

	queryOrderConditionItemUpsertCacheMut.RLock()
	cache, cached := queryOrderConditionItemUpsertCache[key]
	queryOrderConditionItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			queryOrderConditionItemAllColumns,
			queryOrderConditionItemColumnsWithDefault,
			queryOrderConditionItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			queryOrderConditionItemAllColumns,
			queryOrderConditionItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert query_order_condition_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(queryOrderConditionItemPrimaryKeyColumns))
			copy(conflict, queryOrderConditionItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"query_order_condition_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(queryOrderConditionItemType, queryOrderConditionItemMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
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
		return errors.Wrap(err, "sqlboiler: unable to upsert query_order_condition_items")
	}

	if !cached {
		queryOrderConditionItemUpsertCacheMut.Lock()
		queryOrderConditionItemUpsertCache[key] = cache
		queryOrderConditionItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single QueryOrderConditionItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QueryOrderConditionItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboiler: no QueryOrderConditionItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), queryOrderConditionItemPrimaryKeyMapping)
	sql := "DELETE FROM \"query_order_condition_items\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete from query_order_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by delete for query_order_condition_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q queryOrderConditionItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboiler: no queryOrderConditionItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from query_order_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for query_order_condition_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QueryOrderConditionItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(queryOrderConditionItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), queryOrderConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"query_order_condition_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, queryOrderConditionItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from queryOrderConditionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for query_order_condition_items")
	}

	if len(queryOrderConditionItemAfterDeleteHooks) != 0 {
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
func (o *QueryOrderConditionItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQueryOrderConditionItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QueryOrderConditionItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QueryOrderConditionItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), queryOrderConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"query_order_condition_items\".* FROM \"query_order_condition_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, queryOrderConditionItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in QueryOrderConditionItemSlice")
	}

	*o = slice

	return nil
}

// QueryOrderConditionItemExists checks if the QueryOrderConditionItem row exists.
func QueryOrderConditionItemExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"query_order_condition_items\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if query_order_condition_items exists")
	}

	return exists, nil
}
