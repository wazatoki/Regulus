// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testQueryOrderConditionItems(t *testing.T) {
	t.Parallel()

	query := QueryOrderConditionItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testQueryOrderConditionItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQueryOrderConditionItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := QueryOrderConditionItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQueryOrderConditionItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QueryOrderConditionItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testQueryOrderConditionItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := QueryOrderConditionItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if QueryOrderConditionItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected QueryOrderConditionItemExists to return true, but got false.")
	}
}

func testQueryOrderConditionItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	queryOrderConditionItemFound, err := FindQueryOrderConditionItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if queryOrderConditionItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testQueryOrderConditionItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = QueryOrderConditionItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testQueryOrderConditionItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := QueryOrderConditionItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testQueryOrderConditionItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	queryOrderConditionItemOne := &QueryOrderConditionItem{}
	queryOrderConditionItemTwo := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, queryOrderConditionItemOne, queryOrderConditionItemDBTypes, false, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}
	if err = randomize.Struct(seed, queryOrderConditionItemTwo, queryOrderConditionItemDBTypes, false, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = queryOrderConditionItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = queryOrderConditionItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QueryOrderConditionItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testQueryOrderConditionItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	queryOrderConditionItemOne := &QueryOrderConditionItem{}
	queryOrderConditionItemTwo := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, queryOrderConditionItemOne, queryOrderConditionItemDBTypes, false, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}
	if err = randomize.Struct(seed, queryOrderConditionItemTwo, queryOrderConditionItemDBTypes, false, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = queryOrderConditionItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = queryOrderConditionItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func queryOrderConditionItemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func queryOrderConditionItemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *QueryOrderConditionItem) error {
	*o = QueryOrderConditionItem{}
	return nil
}

func testQueryOrderConditionItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &QueryOrderConditionItem{}
	o := &QueryOrderConditionItem{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem object: %s", err)
	}

	AddQueryOrderConditionItemHook(boil.BeforeInsertHook, queryOrderConditionItemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemBeforeInsertHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.AfterInsertHook, queryOrderConditionItemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemAfterInsertHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.AfterSelectHook, queryOrderConditionItemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemAfterSelectHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.BeforeUpdateHook, queryOrderConditionItemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemBeforeUpdateHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.AfterUpdateHook, queryOrderConditionItemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemAfterUpdateHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.BeforeDeleteHook, queryOrderConditionItemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemBeforeDeleteHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.AfterDeleteHook, queryOrderConditionItemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemAfterDeleteHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.BeforeUpsertHook, queryOrderConditionItemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemBeforeUpsertHooks = []QueryOrderConditionItemHook{}

	AddQueryOrderConditionItemHook(boil.AfterUpsertHook, queryOrderConditionItemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	queryOrderConditionItemAfterUpsertHooks = []QueryOrderConditionItemHook{}
}

func testQueryOrderConditionItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQueryOrderConditionItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(queryOrderConditionItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testQueryOrderConditionItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testQueryOrderConditionItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := QueryOrderConditionItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testQueryOrderConditionItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := QueryOrderConditionItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	queryOrderConditionItemDBTypes = map[string]string{`ID`: `text`, `Del`: `boolean`, `CreatedAt`: `timestamp without time zone`, `CreStaffID`: `text`, `UpdatedAt`: `timestamp without time zone`, `UpdateStaffID`: `text`, `QueryConditionID`: `text`, `OrderFieldID`: `text`, `OrderFieldKeyWord`: `text`, `RowOrder`: `integer`}
	_                              = bytes.MinRead
)

func testQueryOrderConditionItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(queryOrderConditionItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(queryOrderConditionItemAllColumns) == len(queryOrderConditionItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testQueryOrderConditionItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(queryOrderConditionItemAllColumns) == len(queryOrderConditionItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &QueryOrderConditionItem{}
	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, queryOrderConditionItemDBTypes, true, queryOrderConditionItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(queryOrderConditionItemAllColumns, queryOrderConditionItemPrimaryKeyColumns) {
		fields = queryOrderConditionItemAllColumns
	} else {
		fields = strmangle.SetComplement(
			queryOrderConditionItemAllColumns,
			queryOrderConditionItemPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := QueryOrderConditionItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testQueryOrderConditionItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(queryOrderConditionItemAllColumns) == len(queryOrderConditionItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := QueryOrderConditionItem{}
	if err = randomize.Struct(seed, &o, queryOrderConditionItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QueryOrderConditionItem: %s", err)
	}

	count, err := QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, queryOrderConditionItemDBTypes, false, queryOrderConditionItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize QueryOrderConditionItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert QueryOrderConditionItem: %s", err)
	}

	count, err = QueryOrderConditionItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
