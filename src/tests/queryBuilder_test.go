package testing

import (
	"models"
	"strings"
	"testing"
)

var equal_statement = "select * from some_table where string = :string and integer = :integer"
var like_statement = "select * from some_table where string like :string "
var in_statement = "select * from some_table where array in :array" // do not support yet

func TestBuildWithoutNil(t *testing.T) {
	var buildedSQL string = ""

	param := make(map[string]interface{})
	param["string"] = "stringValue"
	param["integer"] = 610

	buildedSQL = models.QueryBuilder.Build(equal_statement, param)
	testStringResultMatching("select * from some_table where string = 'stringValue' and integer = 610", buildedSQL, t)
	buildedSQL = models.QueryBuilder.Build(like_statement, param)
	testStringResultMatching("select * from some_table where string like 'stringValue'", buildedSQL, t)
}

type testStruct struct {
	SomeColumn  string
	OtherColumn string
}

func TestInsertSql(t *testing.T) {
	var tst testStruct
	tst.SomeColumn = "something"
	result := models.QueryBuilder.BuildInsert("table", tst)
	t.Log(result)
}

func testStringResultMatching(target, compare string, t *testing.T) {
	if strings.Compare(target, compare) != 0 {
		t.Error("buildedSQL failed, not equal.")
		t.Error("buildedSQL is:" + compare)
	} else {
		t.Log("matched SQL:" + compare)
	}
}
