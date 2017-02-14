package testing

import (
	"models"
	"strings"
	"testing"
)

var equal_statement = "select * from some_table where string = :string and integer = :integer"
var like_statement = "select * from some_table where string like :string "
var in_statement = "select * from some_table where array in :array" // only support []string

func TestBuildWithoutNil(t *testing.T) {
	var buildedSQL string = ""

	param := make(map[string]interface{})
	param["string"] = "stringValue"
	param["integer"] = 610

	buildedSQL = models.QueryBuilder.BuildQuery(equal_statement, param)
	testStringResultMatching("select * from some_table where string = 'stringValue' and integer = 610", buildedSQL, t)
	buildedSQL = models.QueryBuilder.BuildQuery(like_statement, param)
	testStringResultMatching("select * from some_table where string like 'stringValue'", buildedSQL, t)
	param = make(map[string]interface{})
	arrayParam := make([]string, 2)
	arrayParam[0] = "str1"
	arrayParam[1] = "str2"
	param["array"] = arrayParam
	buildedSQL = models.QueryBuilder.BuildQuery(in_statement, param)
	testStringResultMatching("select * from some_table where array in ('str1','str2')", buildedSQL, t)

}

type testStruct struct {
	SomeColumn  string
	OtherColumn string
}

/*
 * deprecated due to query builder should only building sql
 * without any sql operation
func TestInsertSql(t *testing.T) {
	var tst testStruct
	tst.SomeColumn = "something"
	result := models.QueryBuilder.BuildAndInsert("table", tst)
	t.Log(result)
}*/

func testStringResultMatching(target, compare string, t *testing.T) {
	t.Log("expecting:" + target)
	if strings.Compare(strings.Replace(target, " ", "", -1), strings.Replace(compare, " ", "", -1)) != 0 {
		t.Error("buildedSQL failed, not equal.")
		t.Error("buildedSQL is:" + compare)
	} else {
		t.Log("matched SQL:" + compare)
	}
}
