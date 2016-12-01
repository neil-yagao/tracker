package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type queryBuilder struct {
}

type JSONObject map[string]interface{}

type CRUD interface {
	Query(sql string) *sql.Rows
}

type basicCRUD struct {
}

var BasicCRUD basicCRUD

var QueryBuilder *queryBuilder

//sequence is important
var OPERATION = "<>|=|!=|like|( not +in )|( in )|<=|>=|<|>"

var OPERATION_PATTERN = regexp.MustCompile(OPERATION)
var WHERE_CLAUSE_MATCHING_PATTERN = regexp.MustCompile("\\S+\\s*(" + OPERATION + ")\\s*:\\S+")

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "powerlift:password@tcp(127.0.0.1:3306)/powerlift")
	checkErr(err)
}

//todo return a slice of map[string]interface
//based on the struct passing in
func (b *basicCRUD) Query(sql string) *sql.Rows {
	beego.Debug("execute sql:" + sql)
	rows, err := db.Query(sql)
	checkErr(err)
	return rows
}

func (b *basicCRUD) BuildAndQuery(sql string, param map[string]interface{}) *sql.Rows {
	querySql := QueryBuilder.BuildQuery(sql, param)
	return b.Query(querySql)
}

func (b *basicCRUD) InsertOne(sql string) int64 {
	return insertOne(sql)
}

func insertOne(sql string) int64 {
	result, err := db.Exec(sql)
	checkErr(err)
	if lastInsert, err := result.LastInsertId(); err == nil {
		return lastInsert
	} else {
		beego.Error(err)
		panic(err)
	}
}

func (s *basicCRUD) Save(table string, value interface{}) int64 {
	insertSql := buildInsert(table, value)
	beego.Debug("saving sql:" + insertSql)
	return insertOne(insertSql)
}

func checkErr(err error) {
	if err != nil {
		beego.Error(err)
	}
}
func (query *queryBuilder) BuildQuery(sql string, param map[string]interface{}) string {
	return build(sql, param)
}

func (b *basicCRUD) BuildAndUpdate(sql string, params []map[string]interface{}) {
	update(sql, params)
}

func (b *basicCRUD) Update(sql string) {
	rs, err := db.Exec(sql)
	beego.Debug("executing sql:", sql)
	checkErr(err)
	if rows, _ := rs.RowsAffected(); rows > 1 {
		beego.Warning("update sql:%v has affected more than one row.", sql)
	}
}

func update(sql string, params []map[string]interface{}) {
	for _, param := range params {
		str := build(sql, param)
		beego.Debug("update:", str)
		db.Exec(str)
	}

	checkErr(err)
}

func build(sql string, param map[string]interface{}) string {
	var buildSQL, replaceCriteria string
	buildSQL = sql
	for _, criteria := range WHERE_CLAUSE_MATCHING_PATTERN.FindAllString(sql, -1) {
		column := OPERATION_PATTERN.Split(criteria, 2)[0]
		variable := OPERATION_PATTERN.Split(criteria, 2)[1]
		variable = strings.TrimSpace(variable)[1:]
		operator := OPERATION_PATTERN.FindString(criteria)
		operator = strings.TrimSpace(operator)
		paramValue := param[variable]
		beego.Debug("build variable:" + variable)
		beego.Debug("build criteria" + criteria)
		switch paramValue := paramValue.(type) {
		case string:
			replaceCriteria = column + " " + operator + " '" + paramValue + "'"
			beego.Debug("replaceCriteria:", replaceCriteria, "!")
			buildSQL = strings.Replace(buildSQL, criteria, replaceCriteria, -1)
		case nil:
			buildSQL = strings.Replace(buildSQL, criteria, "1=1", -1)
		//case []string
		//slice current not support yet
		default:
			replaceCriteria = column + " " + operator + " " + fmt.Sprint(paramValue)
			beego.Debug("replaceCriteria:", replaceCriteria, "!")
			buildSQL = strings.Replace(buildSQL, criteria, replaceCriteria, -1)
		}
	}
	return strings.TrimSpace(buildSQL)
}

func (query *basicCRUD) BuildAndInsert(table string, value interface{}) string {
	return buildInsert(table, value)
}
func (b *basicCRUD) BuildAndUpdateOne(table string, value interface{}) {
	fields, values := ExtractToStringFromObject(value)
	idFieldIndex := pos("id", fields)
	updateSql := "update " + table + " set "
	setArray := make([]string, 0)
	for index, value := range fields {
		setArray = append(setArray, value+" = "+values[index])
	}
	updateSql += strings.Join(setArray, ",")
	updateSql += " where id = " + values[idFieldIndex]
	_, err = db.Exec(updateSql)
	checkErr(err)
}

func pos(value string, slice []string) int {
	for p, v := range slice {
		if strings.ToLower(v) == value {
			return p
		}
	}
	return -1
}

func buildInsert(table string, value interface{}) string {
	fields, values := ExtractToStringFromObject(value)
	//ignore id fields
	idFieldIndex := pos("id", fields)
	fields = append(fields[:idFieldIndex], fields[idFieldIndex+1:]...)
	values = append(values[:idFieldIndex], values[idFieldIndex+1:]...)
	insertSql := "INSERT INTO " + table + " (" + strings.Join(fields, ",") + ") values (" + strings.Join(values, ",") + ");"
	beego.Debug("insertSql:" + insertSql)
	_, err := db.Exec(insertSql)
	checkErr(err)
	return insertSql
}

func ExtractToStringFromObject(value interface{}) ([]string, []string) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	fields := make([]string, 0)
	values := make([]string, 0)
	for i := 0; i < valueType.NumField(); i++ {
		fieldValue := valueValue.Field(i).Interface()
		switch fieldValue := fieldValue.(type) {
		case string:
			if len(fieldValue) == 0 {
				continue
			}
			values = append(values, "'"+fieldValue+"'")
		case nil:
			continue
		default:
			values = append(values, fmt.Sprint(fieldValue))
		}

		field := valueType.Field(i)
		fields = append(fields, underscore(field.Name))
	}
	return fields, values
}

func CloseRowsAndCheckError(rows *sql.Rows) {
	defer rows.Close()
	err := rows.Err()
	if err != nil {
		beego.Error(err)
	}
}

func Destory() {
	db.Close()
}

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}
