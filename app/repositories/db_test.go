package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupTestData() {
	boil.DebugMode = true
	db := createDB()
	con, _ := db.Open()
	defer con.DB.Close()

	tearDownTestData(con)

	insertTestDataToStaffs(con)
	insertTestDataToStaffGroups(con)
	insertTestDataToJoinStaffsStaffGroups(con)
	insertTestDataToQueryCondition(con)
	insertTestDataToQuerySearchConditionItems(con)
	inserTestDataToJoinQueryConditionsStaffGroups(con)
}

func insertTestDataToQueryCondition(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into query_conditions "
	sql += "(id, pattern_name, category_name, is_disclose, owner_id) "
	sql += "values"
	sql += "('queryConditionid0', 'patternName0', 'staff', false, 'staffid1'),"
	sql += "('queryConditionid1', 'patternName1', 'staff', false, 'staffid1'),"
	sql += "('queryConditionid2', 'patternName2', 'staff', false, 'staffid1'),"
	sql += "('queryConditionid3', 'patternName3', 'staff', false, 'staffid1'),"
	sql += "('queryConditionid4', 'patternName4', 'staff', false, 'staffid1'),"
	sql += "('queryConditionid5', 'patternName5', 'staff-group', false, 'staffid2'),"
	sql += "('queryConditionid6', 'patternName6', 'staff-group', false, 'staffid2'),"
	sql += "('queryConditionid7', 'patternName7', 'staff-group', false, 'staffid2'),"
	sql += "('queryConditionid8', 'patternName8', 'staff-group', false, 'staffid2'),"
	sql += "('queryConditionid9', 'patternName9', 'staff-group', false, 'staffid2')"

	_, err = con.Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Fatal("insert into query_conditions " + err.Error())
	}
}

func insertTestDataToQuerySearchConditionItems(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into query_search_condition_items "
	sql += "(id, query_conditions_id, search_field_id, condition_value, match_type, operator, row_order) "
	sql += "values"
	sql += "('searchConditionItem0', 'queryConditionid0', 'account-id', '1', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem1', 'queryConditionid1', 'account-id', '2', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem2', 'queryConditionid2', 'account-id', '3', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem3', 'queryConditionid3', 'account-id', '4', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem4', 'queryConditionid4', 'account-id', '5', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem5', 'queryConditionid5', 'name', '1', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem6', 'queryConditionid6', 'name', '2', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem7', 'queryConditionid7', 'name', '3', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem8', 'queryConditionid8', 'name', '4', 'pertialmatch', 'and', 1),"
	sql += "('searchConditionItem9', 'queryConditionid9', 'name', '5', 'pertialmatch', 'and', 1)"
	_, err = con.Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Fatal("insert into query_search_condition_items " + err.Error())
	}
}

func inserTestDataToJoinQueryConditionsStaffGroups(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into join_query_conditions_staff_groups "
	sql += "(query_conditions_id, staff_groups_id) "
	sql += "values"
	sql += "('queryConditionid0', 'staffgroupid1'),"
	sql += "('queryConditionid0', 'staffgroupid2'),"
	sql += "('queryConditionid1', 'staffgroupid1'),"
	sql += "('queryConditionid1', 'staffgroupid2'),"
	sql += "('queryConditionid2', 'staffgroupid1'),"
	sql += "('queryConditionid2', 'staffgroupid2'),"
	sql += "('queryConditionid3', 'staffgroupid1'),"
	sql += "('queryConditionid3', 'staffgroupid2'),"
	sql += "('queryConditionid4', 'staffgroupid1'),"
	sql += "('queryConditionid4', 'staffgroupid2'),"
	sql += "('queryConditionid5', 'staffgroupid1'),"
	sql += "('queryConditionid5', 'staffgroupid2'),"
	sql += "('queryConditionid6', 'staffgroupid1'),"
	sql += "('queryConditionid6', 'staffgroupid2'),"
	sql += "('queryConditionid7', 'staffgroupid1'),"
	sql += "('queryConditionid7', 'staffgroupid2'),"
	sql += "('queryConditionid8', 'staffgroupid1'),"
	sql += "('queryConditionid8', 'staffgroupid2'),"
	sql += "('queryConditionid9', 'staffgroupid1'),"
	sql += "('queryConditionid9', 'staffgroupid2')"

	_, err = con.Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Fatal("insert into join_query_conditions_staff_groups " + err.Error())
	}
}

func insertTestDataToJoinStaffsStaffGroups(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into join_staffs_staff_groups "
	sql += "(staffs_id, staff_groups_id) "
	sql += "values"
	sql += "('staffid1', 'staffgroupid1'), "
	sql += "('staffid1', 'staffgroupid2'), "
	sql += "('staffid1', 'staffgroupid4'), "
	sql += "('staffid2', 'staffgroupid1'), "
	sql += "('staffid3', 'staffgroupid2'), "
	sql += "('staffid4', 'staffgroupid3'), "
	sql += "('staffid5', 'staffgroupid1')"
	_, err = con.Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Fatal("insert into join_staffs_staff_groups " + err.Error())
	}
}

func insertTestDataToStaffGroups(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into staff_groups "
	sql += "(id, name, del) "
	sql += "values"
	sql += "('staffgroupid1', 'staff group name 1', false), "
	sql += "('staffgroupid2', 'staff group name 2', false),"
	sql += "('staffgroupid3', 'staff group name 3', false),"
	sql += "('staffgroupid4', 'staff group name 4', true)"
	_, err = con.Exec(sql)
	if err != nil {
		log.Fatal("insertStaffGroupTestData " + err.Error())
	}
}

func insertTestDataToStaffs(con *sqlx.DB) {
	var err error
	var sql string
	sql = "insert into staffs (id, account_id, name, password) "
	sql += "values"
	sql += "('staffid1', '12345', 'name 1', 'password 1'), "
	sql += "('staffid2', '22345', 'name 2', 'password 2'), "
	sql += "('staffid3', '32345', 'name 3', 'password 3'), "
	sql += "('staffid4', '42345', 'name 4', 'password 4'), "
	sql += "('staffid5', '52345', 'name 5', 'password 5')"
	_, err = con.Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Fatal("insert into staffs " + err.Error())
	}
}

func tearDownTestData(con *sqlx.DB) {
	var err error

	_, err = con.Exec("delete from join_query_conditions_staff_groups")
	if err != nil {
		log.Fatal("delete from join_query_conditions_staff_groups " + err.Error())
	}
	_, err = con.Exec("delete from query_search_condition_items")
	if err != nil {
		log.Fatal("delete from query_search_condition_items " + err.Error())
	}
	_, err = con.Exec("delete from query_conditions")
	if err != nil {
		log.Fatal("delete from query_conditions " + err.Error())
	}

	_, err = con.Exec("delete from join_staffs_staff_groups")
	if err != nil {
		log.Fatal("delete from join_staffs_staff_groups " + err.Error())
	}
	_, err = con.Exec("delete from staffs")
	if err != nil {
		log.Fatal("delete from staffs " + err.Error())
	}
	_, err = con.Exec("delete from staff_groups")
	if err != nil {
		log.Fatal("delete from staff_groups " + err.Error())
	}
}
