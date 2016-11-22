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

func (query *queryBuilder) BuildInsert(table string, value interface{}) string {
	return buildInsert(table, value)
}

func buildInsert(table string, value interface{}) string {
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
	insertSql := "INSERT INTO " + table + " (" + strings.Join(fields, ",") + ") values (" + strings.Join(values, ",") + ");"
	return insertSql
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
