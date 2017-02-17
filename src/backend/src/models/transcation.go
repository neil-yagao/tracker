package models

import (
	"fmt"
	"strings"
	"sync"

	"github.com/astaxie/beego"
)

/**
 * files contains all transaction saving and updating.
 * also provide batch update
 */
type BatchSqls struct {
	Sql    string
	Params []JSONObject
}

type Transaction struct {
	locker     *sync.Mutex
	Executions []BatchSqls
}

func (this *Transaction) Execute() error {
	this.locker = &sync.Mutex{}
	this.locker.Lock()
	defer this.locker.Unlock()
	transaction, err := db.Begin()
	if err != nil {
		return err
	}
	for _, batchSql := range this.Executions {
		for _, param := range batchSql.Params {
			executingSql := QueryBuilder.BuildQuery(batchSql.Sql, param)
			_, err = transaction.Exec(executingSql)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}
	transaction.Commit()
	return nil
}

func (this *BatchSqls) BatchInsert(table string) error {

	var columns []string
	var underscoreColumns []string
	var valueArgs []interface{}
	var valueStrings []string
	var valueString string
	for key, _ := range this.Params[0] {
		if strings.ToLower(key) == "id" {
			continue
		}
		columns = append(columns, key)
		underscoreColumns = append(underscoreColumns, underscore(key))
	}
	valueString = "(" + strings.Repeat("?,", len(columns)) + ")"
	valueString = strings.Replace(valueString, "?,)", "?)", -1)
	for _, param := range this.Params {
		valueStrings = append(valueStrings, valueString)
		for _, column := range columns {
			valueArgs = append(valueArgs, param[column])
		}
	}

	stmt := fmt.Sprintf("INSERT IGNORE INTO  %s ( %s ) VALUES %s", table, strings.Join(underscoreColumns, ", "), strings.Join(valueStrings, ", "))
	beego.Debug(stmt)
	beego.Debug("values:", valueArgs)
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}
	return nil

}
