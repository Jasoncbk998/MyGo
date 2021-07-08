package mysqlUtils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)
type Student struct {
	Id   string `db:"SID"`
	Name string `db:"Sname"`
	Age  string `db:"Sage"`
	Sex  string `db:"Ssex"`
}
func Insert2table(s Student,db *sqlx.DB) {
	sql := fmt.Sprintf("insert into Student(SID,Sname,Sage,Ssex)values (%s,'%s','%s','%s')", s.Id, s.Name, s.Age, s.Sex)
	//fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("insert into table err:%s", err)
	}
}
