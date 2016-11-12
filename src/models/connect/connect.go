package connect

import (
	"database/sql"
	"fmt"
	"log"
	"models/connect"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type QueryBuilder struct {
}

type JSONObject map[string]interface{}

type CRUD interface {
	Query(sql string) *sql.Rows
}

type BasicCRUD struct {
}

var Querier *connect.BasicCRUD
var QueryBuilder *connect.QueryBuilder

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
func (b *BasicCRUD) Query(sql string) *sql.Rows {
	rows, err := db.Query(sql)
	checkErr(err)
	return rows
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (query *QueryBuilder) Build(sql string, param map[string]interface{}) string {
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
