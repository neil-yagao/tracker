package models

import (
	"database/sql"
	"fmt"
	"log"
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

var Querier *basicCRUD

var QueryBuilder *queryBuilder

//sequence is important
var OPERATION = "<>|=|!=|like|( not +in )|( in )|<=|>=|<|>"

var OPERATION_PATTERN = regexp.MustCompile(OPERATION)
var WHERE_CLAUSE_MATCHING_PATTERN = regexp.MustCompile("\\S+\\s*(" + OPERATION + ")\\s*:\\S+")

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "powerlift:password@/powerlift")
	checkErr(err)
}

//todo return a slice of map[string]interface
//based on the struct passing in
func (b *basicCRUD) Query(sql string) *sql.Rows {
	rows, err := db.Query(sql)
	checkErr(err)
	return rows
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (query *queryBuilder) Build(sql string, param map[string]interface{}) string {
	var buildSQL, replaceCriteria string
	buildSQL = sql
	for _, criteria := range WHERE_CLAUSE_MATCHING_PATTERN.FindAllString(sql, -1) {
		variable := OPERATION_PATTERN.Split(criteria, 2)[0]
		variable = strings.TrimSpace(variable)
		operator := OPERATION_PATTERN.FindString(criteria)
		operator = strings.TrimSpace(operator)
		paramValue := param[variable]
		switch paramValue := paramValue.(type) {
		case string:
			replaceCriteria = variable + " " + operator + " '" + paramValue + "'"
			beego.Debug("replaceCriteria:", replaceCriteria, "!")
			buildSQL = strings.Replace(buildSQL, criteria, replaceCriteria, -1)
		case nil:
			buildSQL = strings.Replace(buildSQL, criteria, "1=1", -1)
		//case []string
		//slice current not support yet
		default:
			replaceCriteria = variable + " " + operator + " " + fmt.Sprint(paramValue)
			beego.Debug("replaceCriteria:", replaceCriteria, "!")
			buildSQL = strings.Replace(buildSQL, criteria, replaceCriteria, -1)
		}
	}
	return strings.TrimSpace(buildSQL)
}

func (query *queryBuilder) BuildInsert(table string, value interface{}) string {
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
	insertSql := "INSERT INTO " + table + " (" + strings.Join(fields, ",") + ") values (" + strings.Join(values, ",") + ")"
	return insertSql
}

func CloseRowsAndCheckError(rows *sql.Rows) {
	defer rows.Close()
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
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
