package libs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	cnf "learninggolang/config"
	"log"
	"reflect"
	"runtime"
)

func Views(i interface{}) string{
	name :=runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	log.Println("packagename:"+name)
	return name
}


func GetDb() *sql.DB{
	db, err := sql.Open(cnf.DB_TYPE, cnf.DB_USERNAME+":"+cnf.DB_PASSWORD+"@tcp("+cnf.DB_HOST+":"+cnf.DB_PORT+")/"+cnf.DB_DATABASE)
	if err != nil {
		log.Println("connect database faild")
		log.Println(err.Error())
	}
	defer db.Close()
	log.Println("connect database success")
	return db
}