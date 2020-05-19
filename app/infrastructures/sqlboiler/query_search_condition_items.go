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

// QuerySearchConditionItem is an object representing the database table.
type QuerySearchConditionItem struct {
	ID                string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Del               bool        `boil:"del" json:"del" toml:"del" yaml:"del"`
	CreatedAt         null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	CreStaffID        null.String `boil:"cre_staff_id" json:"cre_staff_id,omitempty" toml:"cre_staff_id" yaml:"cre_staff_id,omitempty"`
	UpdatedAt         null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UpdateStaffID     null.String `boil:"update_staff_id" json:"update_staff_id,omitempty" toml:"update_staff_id" yaml:"update_staff_id,omitempty"`
	QueryConditionsID string      `boil:"query_conditions_id" json:"query_conditions_id" toml:"query_conditions_id" yaml:"query_conditions_id"`
	SearchFieldID     string      `boil:"search_field_id" json:"search_field_id" toml:"search_field_id" yaml:"search_field_id"`
	ConditionValue    string      `boil:"condition_value" json:"condition_value" toml:"condition_value" yaml:"condition_value"`
	MatchType         string      `boil:"match_type" json:"match_type" toml:"match_type" yaml:"match_type"`
	Operator          string      `boil:"operator" json:"operator" toml:"operator" yaml:"operator"`
	RowOrder          int         `boil:"row_order" json:"row_order" toml:"row_order" yaml:"row_order"`

	R *querySearchConditionItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L querySearchConditionItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuerySearchConditionItemColumns = struct {
	ID                string
	Del               string
	CreatedAt         string
	CreStaffID        string
	UpdatedAt         string
	UpdateStaffID     string
	QueryConditionsID string
	SearchFieldID     string
	ConditionValue    string
	MatchType         string
	Operator          string
	RowOrder          string
}{
	ID:                "id",
	Del:               "del",
	CreatedAt:         "created_at",
	CreStaffID:        "cre_staff_id",
	UpdatedAt:         "updated_at",
	UpdateStaffID:     "update_staff_id",
	QueryConditionsID: "query_conditions_id",
	SearchFieldID:     "search_field_id",
	ConditionValue:    "condition_value",
	MatchType:         "match_type",
	Operator:          "operator",
	RowOrder:          "row_order",
}

// Generated where

var QuerySearchConditionItemWhere = struct {
	ID                whereHelperstring
	Del               whereHelperbool
	CreatedAt         whereHelpernull_Time
	CreStaffID        whereHelpernull_String
	UpdatedAt         whereHelpernull_Time
	UpdateStaffID     whereHelpernull_String
	QueryConditionsID whereHelperstring
	SearchFieldID     whereHelperstring
	ConditionValue    whereHelperstring
	MatchType         whereHelperstring
	Operator          whereHelperstring
	RowOrder          whereHelperint
}{
	ID:                whereHelperstring{field: "\"query_search_condition_items\".\"id\""},
	Del:               whereHelperbool{field: "\"query_search_condition_items\".\"del\""},
	CreatedAt:         whereHelpernull_Time{field: "\"query_search_condition_items\".\"created_at\""},
	CreStaffID:        whereHelpernull_String{field: "\"query_search_condition_items\".\"cre_staff_id\""},
	UpdatedAt:         whereHelpernull_Time{field: "\"query_search_condition_items\".\"updated_at\""},
	UpdateStaffID:     whereHelpernull_String{field: "\"query_search_condition_items\".\"update_staff_id\""},
	QueryConditionsID: whereHelperstring{field: "\"query_search_condition_items\".\"query_conditions_id\""},
	SearchFieldID:     whereHelperstring{field: "\"query_search_condition_items\".\"search_field_id\""},
	ConditionValue:    whereHelperstring{field: "\"query_search_condition_items\".\"condition_value\""},
	MatchType:         whereHelperstring{field: "\"query_search_condition_items\".\"match_type\""},
	Operator:          whereHelperstring{field: "\"query_search_condition_items\".\"operator\""},
	RowOrder:          whereHelperint{field: "\"query_search_condition_items\".\"row_order\""},
}

// QuerySearchConditionItemRels is where relationship names are stored.
var QuerySearchConditionItemRels = struct {
	QueryCondition string
}{
	QueryCondition: "QueryCondition",
}

// querySearchConditionItemR is where relationships are stored.
type querySearchConditionItemR struct {
	QueryCondition *QueryCondition
}

// NewStruct creates a new relationship struct
func (*querySearchConditionItemR) NewStruct() *querySearchConditionItemR {
	return &querySearchConditionItemR{}
}

// querySearchConditionItemL is where Load methods for each relationship are stored.
type querySearchConditionItemL struct{}

var (
	querySearchConditionItemAllColumns            = []string{"id", "del", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "query_conditions_id", "search_field_id", "condition_value", "match_type", "operator", "row_order"}
	querySearchConditionItemColumnsWithoutDefault = []string{"id", "created_at", "cre_staff_id", "updated_at", "update_staff_id", "query_conditions_id", "search_field_id", "condition_value", "match_type", "operator", "row_order"}
	querySearchConditionItemColumnsWithDefault    = []string{"del"}
	querySearchConditionItemPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuerySearchConditionItemSlice is an alias for a slice of pointers to QuerySearchConditionItem.
	// This should generally be used opposed to []QuerySearchConditionItem.
	QuerySearchConditionItemSlice []*QuerySearchConditionItem
	// QuerySearchConditionItemHook is the signature for custom QuerySearchConditionItem hook methods
	QuerySearchConditionItemHook func(context.Context, boil.ContextExecutor, *QuerySearchConditionItem) error

	querySearchConditionItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	querySearchConditionItemType                 = reflect.TypeOf(&QuerySearchConditionItem{})
	querySearchConditionItemMapping              = queries.MakeStructMapping(querySearchConditionItemType)
	querySearchConditionItemPrimaryKeyMapping, _ = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, querySearchConditionItemPrimaryKeyColumns)
	querySearchConditionItemInsertCacheMut       sync.RWMutex
	querySearchConditionItemInsertCache          = make(map[string]insertCache)
	querySearchConditionItemUpdateCacheMut       sync.RWMutex
	querySearchConditionItemUpdateCache          = make(map[string]updateCache)
	querySearchConditionItemUpsertCacheMut       sync.RWMutex
	querySearchConditionItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var querySearchConditionItemBeforeInsertHooks []QuerySearchConditionItemHook
var querySearchConditionItemBeforeUpdateHooks []QuerySearchConditionItemHook
var querySearchConditionItemBeforeDeleteHooks []QuerySearchConditionItemHook
var querySearchConditionItemBeforeUpsertHooks []QuerySearchConditionItemHook

var querySearchConditionItemAfterInsertHooks []QuerySearchConditionItemHook
var querySearchConditionItemAfterSelectHooks []QuerySearchConditionItemHook
var querySearchConditionItemAfterUpdateHooks []QuerySearchConditionItemHook
var querySearchConditionItemAfterDeleteHooks []QuerySearchConditionItemHook
var querySearchConditionItemAfterUpsertHooks []QuerySearchConditionItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QuerySearchConditionItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QuerySearchConditionItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QuerySearchConditionItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QuerySearchConditionItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QuerySearchConditionItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QuerySearchConditionItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QuerySearchConditionItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QuerySearchConditionItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QuerySearchConditionItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range querySearchConditionItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuerySearchConditionItemHook registers your hook function for all future operations.
func AddQuerySearchConditionItemHook(hookPoint boil.HookPoint, querySearchConditionItemHook QuerySearchConditionItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		querySearchConditionItemBeforeInsertHooks = append(querySearchConditionItemBeforeInsertHooks, querySearchConditionItemHook)
	case boil.BeforeUpdateHook:
		querySearchConditionItemBeforeUpdateHooks = append(querySearchConditionItemBeforeUpdateHooks, querySearchConditionItemHook)
	case boil.BeforeDeleteHook:
		querySearchConditionItemBeforeDeleteHooks = append(querySearchConditionItemBeforeDeleteHooks, querySearchConditionItemHook)
	case boil.BeforeUpsertHook:
		querySearchConditionItemBeforeUpsertHooks = append(querySearchConditionItemBeforeUpsertHooks, querySearchConditionItemHook)
	case boil.AfterInsertHook:
		querySearchConditionItemAfterInsertHooks = append(querySearchConditionItemAfterInsertHooks, querySearchConditionItemHook)
	case boil.AfterSelectHook:
		querySearchConditionItemAfterSelectHooks = append(querySearchConditionItemAfterSelectHooks, querySearchConditionItemHook)
	case boil.AfterUpdateHook:
		querySearchConditionItemAfterUpdateHooks = append(querySearchConditionItemAfterUpdateHooks, querySearchConditionItemHook)
	case boil.AfterDeleteHook:
		querySearchConditionItemAfterDeleteHooks = append(querySearchConditionItemAfterDeleteHooks, querySearchConditionItemHook)
	case boil.AfterUpsertHook:
		querySearchConditionItemAfterUpsertHooks = append(querySearchConditionItemAfterUpsertHooks, querySearchConditionItemHook)
	}
}

// One returns a single querySearchConditionItem record from the query.
func (q querySearchConditionItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*QuerySearchConditionItem, error) {
	o := &QuerySearchConditionItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for query_search_condition_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QuerySearchConditionItem records from the query.
func (q querySearchConditionItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuerySearchConditionItemSlice, error) {
	var o []*QuerySearchConditionItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to QuerySearchConditionItem slice")
	}

	if len(querySearchConditionItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QuerySearchConditionItem records in the query.
func (q querySearchConditionItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count query_search_condition_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q querySearchConditionItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if query_search_condition_items exists")
	}

	return count > 0, nil
}

// QueryCondition pointed to by the foreign key.
func (o *QuerySearchConditionItem) QueryCondition(mods ...qm.QueryMod) queryConditionQuery {
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
func (querySearchConditionItemL) LoadQueryCondition(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuerySearchConditionItem interface{}, mods queries.Applicator) error {
	var slice []*QuerySearchConditionItem
	var object *QuerySearchConditionItem

	if singular {
		object = maybeQuerySearchConditionItem.(*QuerySearchConditionItem)
	} else {
		slice = *maybeQuerySearchConditionItem.(*[]*QuerySearchConditionItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &querySearchConditionItemR{}
		}
		args = append(args, object.QueryConditionsID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &querySearchConditionItemR{}
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

	query := NewQuery(qm.From(`query_conditions`), qm.WhereIn(`query_conditions.id in ?`, args...))
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

	if len(querySearchConditionItemAfterSelectHooks) != 0 {
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
		foreign.R.QuerySearchConditionItems = append(foreign.R.QuerySearchConditionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QueryConditionsID == foreign.ID {
				local.R.QueryCondition = foreign
				if foreign.R == nil {
					foreign.R = &queryConditionR{}
				}
				foreign.R.QuerySearchConditionItems = append(foreign.R.QuerySearchConditionItems, local)
				break
			}
		}
	}

	return nil
}

// SetQueryCondition of the querySearchConditionItem to the related item.
// Sets o.R.QueryCondition to related.
// Adds o to related.R.QuerySearchConditionItems.
func (o *QuerySearchConditionItem) SetQueryCondition(ctx context.Context, exec boil.ContextExecutor, insert bool, related *QueryCondition) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"query_search_condition_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"query_conditions_id"}),
		strmangle.WhereClause("\"", "\"", 2, querySearchConditionItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.QueryConditionsID = related.ID
	if o.R == nil {
		o.R = &querySearchConditionItemR{
			QueryCondition: related,
		}
	} else {
		o.R.QueryCondition = related
	}

	if related.R == nil {
		related.R = &queryConditionR{
			QuerySearchConditionItems: QuerySearchConditionItemSlice{o},
		}
	} else {
		related.R.QuerySearchConditionItems = append(related.R.QuerySearchConditionItems, o)
	}

	return nil
}

// QuerySearchConditionItems retrieves all the records using an executor.
func QuerySearchConditionItems(mods ...qm.QueryMod) querySearchConditionItemQuery {
	mods = append(mods, qm.From("\"query_search_condition_items\""))
	return querySearchConditionItemQuery{NewQuery(mods...)}
}

// FindQuerySearchConditionItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuerySearchConditionItem(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*QuerySearchConditionItem, error) {
	querySearchConditionItemObj := &QuerySearchConditionItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"query_search_condition_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, querySearchConditionItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from query_search_condition_items")
	}

	return querySearchConditionItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QuerySearchConditionItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no query_search_condition_items provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(querySearchConditionItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	querySearchConditionItemInsertCacheMut.RLock()
	cache, cached := querySearchConditionItemInsertCache[key]
	querySearchConditionItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			querySearchConditionItemAllColumns,
			querySearchConditionItemColumnsWithDefault,
			querySearchConditionItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"query_search_condition_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"query_search_condition_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sqlboiler: unable to insert into query_search_condition_items")
	}

	if !cached {
		querySearchConditionItemInsertCacheMut.Lock()
		querySearchConditionItemInsertCache[key] = cache
		querySearchConditionItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the QuerySearchConditionItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QuerySearchConditionItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	querySearchConditionItemUpdateCacheMut.RLock()
	cache, cached := querySearchConditionItemUpdateCache[key]
	querySearchConditionItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			querySearchConditionItemAllColumns,
			querySearchConditionItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboiler: unable to update query_search_condition_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"query_search_condition_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, querySearchConditionItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, append(wl, querySearchConditionItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "sqlboiler: unable to update query_search_condition_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by update for query_search_condition_items")
	}

	if !cached {
		querySearchConditionItemUpdateCacheMut.Lock()
		querySearchConditionItemUpdateCache[key] = cache
		querySearchConditionItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q querySearchConditionItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all for query_search_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected for query_search_condition_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuerySearchConditionItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), querySearchConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"query_search_condition_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, querySearchConditionItemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all in querySearchConditionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected all in update all querySearchConditionItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuerySearchConditionItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no query_search_condition_items provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(querySearchConditionItemColumnsWithDefault, o)

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

	querySearchConditionItemUpsertCacheMut.RLock()
	cache, cached := querySearchConditionItemUpsertCache[key]
	querySearchConditionItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			querySearchConditionItemAllColumns,
			querySearchConditionItemColumnsWithDefault,
			querySearchConditionItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			querySearchConditionItemAllColumns,
			querySearchConditionItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert query_search_condition_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(querySearchConditionItemPrimaryKeyColumns))
			copy(conflict, querySearchConditionItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"query_search_condition_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(querySearchConditionItemType, querySearchConditionItemMapping, ret)
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
		return errors.Wrap(err, "sqlboiler: unable to upsert query_search_condition_items")
	}

	if !cached {
		querySearchConditionItemUpsertCacheMut.Lock()
		querySearchConditionItemUpsertCache[key] = cache
		querySearchConditionItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single QuerySearchConditionItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QuerySearchConditionItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboiler: no QuerySearchConditionItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), querySearchConditionItemPrimaryKeyMapping)
	sql := "DELETE FROM \"query_search_condition_items\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete from query_search_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by delete for query_search_condition_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q querySearchConditionItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboiler: no querySearchConditionItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from query_search_condition_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for query_search_condition_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuerySearchConditionItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(querySearchConditionItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), querySearchConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"query_search_condition_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, querySearchConditionItemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from querySearchConditionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for query_search_condition_items")
	}

	if len(querySearchConditionItemAfterDeleteHooks) != 0 {
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
func (o *QuerySearchConditionItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuerySearchConditionItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuerySearchConditionItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuerySearchConditionItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), querySearchConditionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"query_search_condition_items\".* FROM \"query_search_condition_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, querySearchConditionItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in QuerySearchConditionItemSlice")
	}

	*o = slice

	return nil
}

// QuerySearchConditionItemExists checks if the QuerySearchConditionItem row exists.
func QuerySearchConditionItemExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"query_search_condition_items\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if query_search_condition_items exists")
	}

	return exists, nil
}
