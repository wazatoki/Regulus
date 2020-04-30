// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("Makers", testMakers)
	t.Run("QueryConditions", testQueryConditions)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroups)
	t.Run("QueryDisplayItems", testQueryDisplayItems)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItems)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItems)
	t.Run("Staffs", testStaffs)
	t.Run("StaffGroups", testStaffGroups)
}

func TestDelete(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("Makers", testMakersDelete)
	t.Run("QueryConditions", testQueryConditionsDelete)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsDelete)
	t.Run("QueryDisplayItems", testQueryDisplayItemsDelete)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsDelete)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsDelete)
	t.Run("Staffs", testStaffsDelete)
	t.Run("StaffGroups", testStaffGroupsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("Makers", testMakersQueryDeleteAll)
	t.Run("QueryConditions", testQueryConditionsQueryDeleteAll)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsQueryDeleteAll)
	t.Run("QueryDisplayItems", testQueryDisplayItemsQueryDeleteAll)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsQueryDeleteAll)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsQueryDeleteAll)
	t.Run("Staffs", testStaffsQueryDeleteAll)
	t.Run("StaffGroups", testStaffGroupsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("Makers", testMakersSliceDeleteAll)
	t.Run("QueryConditions", testQueryConditionsSliceDeleteAll)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsSliceDeleteAll)
	t.Run("QueryDisplayItems", testQueryDisplayItemsSliceDeleteAll)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsSliceDeleteAll)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsSliceDeleteAll)
	t.Run("Staffs", testStaffsSliceDeleteAll)
	t.Run("StaffGroups", testStaffGroupsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("Makers", testMakersExists)
	t.Run("QueryConditions", testQueryConditionsExists)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsExists)
	t.Run("QueryDisplayItems", testQueryDisplayItemsExists)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsExists)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsExists)
	t.Run("Staffs", testStaffsExists)
	t.Run("StaffGroups", testStaffGroupsExists)
}

func TestFind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("Makers", testMakersFind)
	t.Run("QueryConditions", testQueryConditionsFind)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsFind)
	t.Run("QueryDisplayItems", testQueryDisplayItemsFind)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsFind)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsFind)
	t.Run("Staffs", testStaffsFind)
	t.Run("StaffGroups", testStaffGroupsFind)
}

func TestBind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("Makers", testMakersBind)
	t.Run("QueryConditions", testQueryConditionsBind)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsBind)
	t.Run("QueryDisplayItems", testQueryDisplayItemsBind)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsBind)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsBind)
	t.Run("Staffs", testStaffsBind)
	t.Run("StaffGroups", testStaffGroupsBind)
}

func TestOne(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("Makers", testMakersOne)
	t.Run("QueryConditions", testQueryConditionsOne)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsOne)
	t.Run("QueryDisplayItems", testQueryDisplayItemsOne)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsOne)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsOne)
	t.Run("Staffs", testStaffsOne)
	t.Run("StaffGroups", testStaffGroupsOne)
}

func TestAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("Makers", testMakersAll)
	t.Run("QueryConditions", testQueryConditionsAll)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsAll)
	t.Run("QueryDisplayItems", testQueryDisplayItemsAll)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsAll)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsAll)
	t.Run("Staffs", testStaffsAll)
	t.Run("StaffGroups", testStaffGroupsAll)
}

func TestCount(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("Makers", testMakersCount)
	t.Run("QueryConditions", testQueryConditionsCount)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsCount)
	t.Run("QueryDisplayItems", testQueryDisplayItemsCount)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsCount)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsCount)
	t.Run("Staffs", testStaffsCount)
	t.Run("StaffGroups", testStaffGroupsCount)
}

func TestHooks(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsHooks)
	t.Run("Makers", testMakersHooks)
	t.Run("QueryConditions", testQueryConditionsHooks)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsHooks)
	t.Run("QueryDisplayItems", testQueryDisplayItemsHooks)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsHooks)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsHooks)
	t.Run("Staffs", testStaffsHooks)
	t.Run("StaffGroups", testStaffGroupsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("Makers", testMakersInsert)
	t.Run("Makers", testMakersInsertWhitelist)
	t.Run("QueryConditions", testQueryConditionsInsert)
	t.Run("QueryConditions", testQueryConditionsInsertWhitelist)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsInsert)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsInsertWhitelist)
	t.Run("QueryDisplayItems", testQueryDisplayItemsInsert)
	t.Run("QueryDisplayItems", testQueryDisplayItemsInsertWhitelist)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsInsert)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsInsertWhitelist)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsInsert)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsInsertWhitelist)
	t.Run("Staffs", testStaffsInsert)
	t.Run("Staffs", testStaffsInsertWhitelist)
	t.Run("StaffGroups", testStaffGroupsInsert)
	t.Run("StaffGroups", testStaffGroupsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("Makers", testMakersReload)
	t.Run("QueryConditions", testQueryConditionsReload)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsReload)
	t.Run("QueryDisplayItems", testQueryDisplayItemsReload)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsReload)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsReload)
	t.Run("Staffs", testStaffsReload)
	t.Run("StaffGroups", testStaffGroupsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("Makers", testMakersReloadAll)
	t.Run("QueryConditions", testQueryConditionsReloadAll)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsReloadAll)
	t.Run("QueryDisplayItems", testQueryDisplayItemsReloadAll)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsReloadAll)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsReloadAll)
	t.Run("Staffs", testStaffsReloadAll)
	t.Run("StaffGroups", testStaffGroupsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("Makers", testMakersSelect)
	t.Run("QueryConditions", testQueryConditionsSelect)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsSelect)
	t.Run("QueryDisplayItems", testQueryDisplayItemsSelect)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsSelect)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsSelect)
	t.Run("Staffs", testStaffsSelect)
	t.Run("StaffGroups", testStaffGroupsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("Makers", testMakersUpdate)
	t.Run("QueryConditions", testQueryConditionsUpdate)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsUpdate)
	t.Run("QueryDisplayItems", testQueryDisplayItemsUpdate)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsUpdate)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsUpdate)
	t.Run("Staffs", testStaffsUpdate)
	t.Run("StaffGroups", testStaffGroupsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("Makers", testMakersSliceUpdateAll)
	t.Run("QueryConditions", testQueryConditionsSliceUpdateAll)
	t.Run("QueryConditionDiscloseGroups", testQueryConditionDiscloseGroupsSliceUpdateAll)
	t.Run("QueryDisplayItems", testQueryDisplayItemsSliceUpdateAll)
	t.Run("QueryOrderConditionItems", testQueryOrderConditionItemsSliceUpdateAll)
	t.Run("QuerySearchConditionItems", testQuerySearchConditionItemsSliceUpdateAll)
	t.Run("Staffs", testStaffsSliceUpdateAll)
	t.Run("StaffGroups", testStaffGroupsSliceUpdateAll)
}
