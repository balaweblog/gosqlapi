package model

import (
	"database/sql"
	"fmt"
	//gosqldriver
	_ "github.com/go-sql-driver/mysql"
)

/*ExecuteNonQuery Execute Non Query from sql database.*/
func ExecuteNonQuery(db *sql.DB, query string) (interface{}, error) {
	var result interface{}
	stmtIns, err := db.Prepare(query)
	if err != nil {
		return result, err
	}

	defer stmtIns.Close()
	result, err = stmtIns.Exec()
	if err != nil {
		return result, err
	}

	return result, err
}

/*ExecuteQuery  Execute Query from sql database */
func ExecuteQuery(db *sql.DB, query string) (map[int]map[string]string, error) {

	result := map[int]map[string]string{}

	stmtIns, err := db.Prepare(query)
	if err != nil {
		return result, err
	}
	defer stmtIns.Close()

	rows, err := db.Query(query)
	columns, err := rows.Columns()
	colcount := len(columns)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	values := make([]interface{}, colcount)
	valuePtrs := make([]interface{}, colcount)

	resultid := 0
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		tmpstruct := map[string]string{}
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmpstruct[col] = fmt.Sprintf("%s", v)
		}

		result[resultid] = tmpstruct
		resultid++
	}
	return result, err
}

/*ParseQuery parse sql query for the sql input */
func ParseQuery(input map[string]interface{}, parsetype string) string {
	dbname, tablename, columnames, wherenames, properties := readinputvariables(input)
	query := parsewritequery(dbname, tablename, columnames, wherenames, properties, parsetype)
	fmt.Println(query)
	return query
}

/*NewConnection open up new connection from sql database */
func NewConnection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:dhiva@tcp(192.168.99.100:3306)/devdb")
	return
}

/* readinputvariables from the given string */
func readinputvariables(input map[string]interface{}) (string, string, map[string]string, map[int]map[string]string, map[string]string) {
	var tablename, dbname string
	columnames := make(map[string]string)
	wherenames := map[int]map[string]string{}
	properties := make(map[string]string)

	for key, value := range input {
		switch key {
		case "database":
			{
				dbname = value.(string)
			}
		case "table":
			{
				tablename = value.(string)
			}
		case "column":
			{
				for _, innerval := range value.([]interface{}) {
					switch innerval.(type) {
					case map[string]interface{}:
						for key, val := range innerval.(map[string]interface{}) {
							columnames[key] = val.(string)
						}
					}
				}
			}
		case "where":
			{
				i := 0
				for _, innerval := range value.([]interface{}) {
					switch innerval.(type) {
					case map[string]interface{}:
						tmpstruct := map[string]string{}
						for key, val := range innerval.(map[string]interface{}) {
							tmpstruct[key] = val.(string)
						}
						i += 1
						wherenames[i] = tmpstruct
					}
				}
			}
		case "properties":
			{
				for _, innerval := range value.([]interface{}) {
					switch innerval.(type) {
					case map[string]interface{}:
						for key, val := range innerval.(map[string]interface{}) {
							properties[key] = val.(string)
						}
					}
				}
			}
		}
	}
	return dbname, tablename, columnames, wherenames, properties
}

/*parsewritequery */
func parsewritequery(dbname string, tablename string, columnames map[string]string, wherenames map[int]map[string]string, properties map[string]string, parsetype string) string {
	var query string
	switch parsetype {

	case "CREATEDB":
		{
			query = "CREATE DATABASE "
			query = query + dbname + ";"
		}
	case "TRUNCATE":
		{
			query = "TRUNCATE TABLE "
			query = query + tablename + ";"
		}
	case "DROPTABLE":
		{
			query = "DROP TABLE "
			query = query + tablename + ";"
		}
	case "DROPDB":
		{
			query = "DROP DATABASE "
			query = query + dbname + ";"
		}
	case "SHOWALLTABLEINDEX":
		{
			query = "SHOW INDEX FROM "
			query = query + tablename + ";"
		}
	case "CREATE":
		{
			query = "CREATE TABLE "
			query = query + tablename + " ("
			var subquery string
			var primary string
			var i int
			for key, val := range columnames {
				subquery = subquery + key + " " + val
				i = i + 1
				if len(columnames) != i {
					subquery += ","
				}
			}
			if len(properties) > 0 {
				primary = ", PRIMARY KEY ("
				for key, val := range properties {
					if key == "primary" {
						primary += val
						primary += ")"
					}
				}
			}
			query = query + subquery + primary + ");"
		}
	case "SELECT":
		{
			query = "SELECT "
			var subquery string
			var i int
			for _, val := range columnames {
				subquery = subquery + " " + val
				i = i + 1
				if len(columnames) != i {
					subquery += " ,"
				} else {
					subquery += " "
				}
			}
			query = query + subquery
			query = query + " FROM " + tablename

			if len(wherenames) > 0 {
				wherequery := " WHERE "
				for _, val := range wherenames {
					wherequery += val["StartTag"] + val["LHS"] + val["Operator"] + val["RHS"] + " " + val["NextOp"] + " " + val["EndTag"]
				}
				query = query + wherequery
			}
			query += ";"
		}
	case "LISTALLUSERS":
		{
			query = "SELECT USER,HOST FROM mysql.user"
		}
	case "ALTER":
		{
			query = "ALTER TABLE "
			var subquery string
			var i int
			query = query + tablename + " "
			for key, val := range properties {
				subquery = subquery + "ADD COLUMN " + key + " " + val
				i = i + 1
				if len(properties) != i {
					subquery += ","
				} else {
					subquery += ";"
				}
			}
			query = query + subquery
		}
	case "DESCTABLE":
		{
			query = "describe " + tablename + ";"
		}
	case "SHOWDB":
		{
			query = "SHOW DATABASES;"
		}
	case "CURRENTDB":
		{
			query = "select database();"
		}
	case "SHOWALLTABLES":
		{
			query = "show tables;"
		}
	case "EXPLAINTABLE":
		{
			query = "EXPLAIN SELECT * FROM " + tablename + ";"
		}
	case "INSERT":
		{
			query = "INSERT INTO " + tablename
			var colquery string = "("
			var valquery string = " VALUES ("
			var i int
			for key, val := range columnames {
				colquery += key
				valquery += val
				i = i + 1
				if len(columnames) != i {
					colquery += " ,"
					valquery += " ,"
				} else {
					colquery += ") "
					valquery += "); "
				}
			}
			query = query + colquery + valquery
		}
	case "DELETE":
		{
			query = "DELETE FROM " + tablename
			var wherequery string
			if len(wherenames) > 0 {
				wherequery := " WHERE "
				for _, val := range wherenames {
					wherequery += val["StartTag"] + val["LHS"] + val["Operator"] + val["RHS"] + " " + val["NextOp"] + " " + val["EndTag"]
				}
				query = query + wherequery
			}
			query = query + wherequery + ";"

		}
	case "UPDATE":
		{
			query = "UPDATE " + tablename + " SET "
			var colquery string
			var wherequery string
			var i int
			for key, val := range columnames {
				colquery += key + "=" + val
				i = i + 1
				if len(columnames) != i {
					colquery += ", "
				}
			}
			if len(wherenames) > 0 {
				wherequery = " WHERE "
				for _, val := range wherenames {
					wherequery = wherequery + val["StartTag"] + val["LHS"] + val["Operator"] + val["RHS"] + " " + val["NextOp"] + " " + val["EndTag"]
				}
			}
			query = query + colquery + wherequery + ";"
		}
	}
	return query
}
