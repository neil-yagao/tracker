package models

import "testing"

func TestBatchInsert(t *testing.T) {
	trans := new(Transaction)
	batchSql := new(BatchSqls)
	batchSql.Params = append(batchSql.Params, map[string]interface{}{"param1": 1})
	batchSql.Params = append(batchSql.Params, map[string]interface{}{"param2": "2"})
	batchSql.Params = append(batchSql.Params, map[string]interface{}{"param3": 3.0})
	trans.Executions = append(trans.Executions, *batchSql)
	trans.BatchInsert("table")
}
