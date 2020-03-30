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

func testStaffGroups(t *testing.T) {
	t.Parallel()

	query := StaffGroups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStaffGroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
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

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStaffGroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StaffGroups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStaffGroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StaffGroupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStaffGroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StaffGroupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if StaffGroup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StaffGroupExists to return true, but got false.")
	}
}

func testStaffGroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	staffGroupFound, err := FindStaffGroup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if staffGroupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStaffGroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StaffGroups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStaffGroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StaffGroups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStaffGroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	staffGroupOne := &StaffGroup{}
	staffGroupTwo := &StaffGroup{}
	if err = randomize.Struct(seed, staffGroupOne, staffGroupDBTypes, false, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}
	if err = randomize.Struct(seed, staffGroupTwo, staffGroupDBTypes, false, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = staffGroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = staffGroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StaffGroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStaffGroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	staffGroupOne := &StaffGroup{}
	staffGroupTwo := &StaffGroup{}
	if err = randomize.Struct(seed, staffGroupOne, staffGroupDBTypes, false, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}
	if err = randomize.Struct(seed, staffGroupTwo, staffGroupDBTypes, false, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = staffGroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = staffGroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func staffGroupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func staffGroupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StaffGroup) error {
	*o = StaffGroup{}
	return nil
}

func testStaffGroupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StaffGroup{}
	o := &StaffGroup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, staffGroupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StaffGroup object: %s", err)
	}

	AddStaffGroupHook(boil.BeforeInsertHook, staffGroupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	staffGroupBeforeInsertHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.AfterInsertHook, staffGroupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	staffGroupAfterInsertHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.AfterSelectHook, staffGroupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	staffGroupAfterSelectHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.BeforeUpdateHook, staffGroupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	staffGroupBeforeUpdateHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.AfterUpdateHook, staffGroupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	staffGroupAfterUpdateHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.BeforeDeleteHook, staffGroupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	staffGroupBeforeDeleteHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.AfterDeleteHook, staffGroupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	staffGroupAfterDeleteHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.BeforeUpsertHook, staffGroupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	staffGroupBeforeUpsertHooks = []StaffGroupHook{}

	AddStaffGroupHook(boil.AfterUpsertHook, staffGroupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	staffGroupAfterUpsertHooks = []StaffGroupHook{}
}

func testStaffGroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStaffGroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(staffGroupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStaffGroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
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

func testStaffGroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StaffGroupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStaffGroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StaffGroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	staffGroupDBTypes = map[string]string{`ID`: `text`, `Del`: `boolean`, `CreatedAt`: `timestamp without time zone`, `CreStaffID`: `text`, `UpdatedAt`: `timestamp without time zone`, `UpdateStaffID`: `text`, `Name`: `text`}
	_                 = bytes.MinRead
)

func testStaffGroupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(staffGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(staffGroupAllColumns) == len(staffGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStaffGroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(staffGroupAllColumns) == len(staffGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StaffGroup{}
	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, staffGroupDBTypes, true, staffGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(staffGroupAllColumns, staffGroupPrimaryKeyColumns) {
		fields = staffGroupAllColumns
	} else {
		fields = strmangle.SetComplement(
			staffGroupAllColumns,
			staffGroupPrimaryKeyColumns,
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

	slice := StaffGroupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStaffGroupsUpsert(t *testing.T) {
	t.Parallel()

	if len(staffGroupAllColumns) == len(staffGroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StaffGroup{}
	if err = randomize.Struct(seed, &o, staffGroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StaffGroup: %s", err)
	}

	count, err := StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, staffGroupDBTypes, false, staffGroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StaffGroup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StaffGroup: %s", err)
	}

	count, err = StaffGroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
