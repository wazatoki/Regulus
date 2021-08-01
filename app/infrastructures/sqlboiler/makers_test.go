// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMakers(t *testing.T) {
	t.Parallel()

	query := Makers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMakersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
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

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMakersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Makers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMakersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MakerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMakersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MakerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Maker exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MakerExists to return true, but got false.")
	}
}

func testMakersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	makerFound, err := FindMaker(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if makerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMakersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Makers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMakersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Makers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMakersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	makerOne := &Maker{}
	makerTwo := &Maker{}
	if err = randomize.Struct(seed, makerOne, makerDBTypes, false, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}
	if err = randomize.Struct(seed, makerTwo, makerDBTypes, false, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = makerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = makerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Makers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMakersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	makerOne := &Maker{}
	makerTwo := &Maker{}
	if err = randomize.Struct(seed, makerOne, makerDBTypes, false, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}
	if err = randomize.Struct(seed, makerTwo, makerDBTypes, false, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = makerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = makerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func makerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func makerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Maker) error {
	*o = Maker{}
	return nil
}

func testMakersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Maker{}
	o := &Maker{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, makerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Maker object: %s", err)
	}

	AddMakerHook(boil.BeforeInsertHook, makerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	makerBeforeInsertHooks = []MakerHook{}

	AddMakerHook(boil.AfterInsertHook, makerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	makerAfterInsertHooks = []MakerHook{}

	AddMakerHook(boil.AfterSelectHook, makerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	makerAfterSelectHooks = []MakerHook{}

	AddMakerHook(boil.BeforeUpdateHook, makerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	makerBeforeUpdateHooks = []MakerHook{}

	AddMakerHook(boil.AfterUpdateHook, makerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	makerAfterUpdateHooks = []MakerHook{}

	AddMakerHook(boil.BeforeDeleteHook, makerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	makerBeforeDeleteHooks = []MakerHook{}

	AddMakerHook(boil.AfterDeleteHook, makerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	makerAfterDeleteHooks = []MakerHook{}

	AddMakerHook(boil.BeforeUpsertHook, makerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	makerBeforeUpsertHooks = []MakerHook{}

	AddMakerHook(boil.AfterUpsertHook, makerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	makerAfterUpsertHooks = []MakerHook{}
}

func testMakersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMakersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(makerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMakersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
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

func testMakersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MakerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMakersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Makers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	makerDBTypes = map[string]string{`ID`: `text`, `Del`: `boolean`, `CreatedAt`: `timestamp without time zone`, `CreStaffID`: `text`, `UpdatedAt`: `timestamp without time zone`, `UpdateStaffID`: `text`, `Name`: `text`}
	_            = bytes.MinRead
)

func testMakersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(makerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(makerAllColumns) == len(makerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, makerDBTypes, true, makerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMakersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(makerAllColumns) == len(makerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Maker{}
	if err = randomize.Struct(seed, o, makerDBTypes, true, makerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, makerDBTypes, true, makerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(makerAllColumns, makerPrimaryKeyColumns) {
		fields = makerAllColumns
	} else {
		fields = strmangle.SetComplement(
			makerAllColumns,
			makerPrimaryKeyColumns,
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

	slice := MakerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMakersUpsert(t *testing.T) {
	t.Parallel()

	if len(makerAllColumns) == len(makerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Maker{}
	if err = randomize.Struct(seed, &o, makerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Maker: %s", err)
	}

	count, err := Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, makerDBTypes, false, makerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Maker struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Maker: %s", err)
	}

	count, err = Makers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
