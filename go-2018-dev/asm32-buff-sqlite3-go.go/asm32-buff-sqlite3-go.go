// asm32-buff-sqlite3-go.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	strDatabase = "file::memory:?mode=memory&cache=shared&loc=auto"
	strTableName = "table_buff"
)

var db *sql.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}

	var strHtmlBody, strQuery string
	db, err := sql.Open("sqlite3", strDatabase)
	if err != nil {
		strHtmlBody = "SQLite: " + err.Error()
	}
	defer db.Close()

	strQuery = fmt.Sprintf("insert into `%s`(`id`,`strContent`,`strDateCreated`,`strDateModified`) values(1, null, null, null)", strTableName)
	_, err = db.Exec(strQuery)
	if err != nil {
		strHtmlBody += "Error: " + err.Error() + "\n"
	}

	strQuery = fmt.Sprintf("select count(*) from `%s`", strTableName)
	rows, err := db.Query(strQuery)
	if err != nil {
		strHtmlBody += "Error: " + err.Error() + "\n"
	} else {
		nCount := 0
		if rows.Next() {
			rows.Scan(&nCount)
			strHtmlBody += fmt.Sprintf("nCount: %d\n", nCount)
		}
		rows.Close()
	}

	strHtmlBody += "indexHandler"

	io.WriteString(w, strHtmlBody)
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	var strHtmlBody string
	strQuery := fmt.Sprintf("create table `%s`(\n" +
		"\t`id` int,\n" +
		"\t`strContent` text,\n" +
		"\t`strDateCreated` datetime,\n" +
		"\t`strDateModified` datetime,\n" +
		"\tprimary key(`id`))",
		strTableName);
	db, err := sql.Open("sqlite3", strDatabase)
	if err != nil {
		strHtmlBody = "SQLite connect error: " + err.Error()
	} else {
		defer db.Close()
		_, err = db.Exec(strQuery)
		if err != nil {
			strHtmlBody = "Error: " + err.Error()
		} else {
			strHtmlBody = "Install OK !"
		}
	}
	io.WriteString(w, strHtmlBody)
}

func favriconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/install", installHandler)
	http.HandleFunc("/favricon.ico", favriconHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}