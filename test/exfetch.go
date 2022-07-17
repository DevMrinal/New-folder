package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	//"strings"

	"example.com/m/graph/model"
	_ "github.com/lib/pq"

	//"golang.org/x/tools/go/analysis/passes/nilfunc"

	//"golang.org/x/tools/go/analysis/passes/nilfunc"

	//	"github.com/360EnSecGroup-Skylar/excelize/v2"
	"github.com/xuri/excelize/v2"
)

const (
	host2     = "localhost"
	port2     = 5432
	user2     = "postgres"
	password2 = ""
	dbname2   = "postgres"
)

func Exfetch(ctx context.Context) (*model.Exeloutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host2, port2, user2, password2, dbname2)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var stat string
	//err = db.Ping()
	if err != nil {
		panic(err)
	}
	f, err := excelize.OpenFile("D://biodata.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rows, err := f.GetRows("practice1")
	if err != nil {
		panic(err)
	}
	for row, col := range rows {
		if row == 0 {
			continue
		}
		if len(col) > 1 {
			if col[0] == "" {
				//log.Println("for insert")
				insertstmt := `Insert into mrinal_db.Student("firstname","lastname","dob","phno","bloodgroup","address","gender")values($1,$2,$3,$4,$5,$6,$7)`
				_, e := db.Query(insertstmt, col[1], col[2], col[3], col[4], col[5], col[6], col[7])
				if e != nil {
					log.Println(e)
					stat = "failed to insert"
				}
				stat = "success"
			} else if col[0] != "" {
				log.Println("for update")
				updatestmt := `Update mrinal_db.student set firstname=$1,lastname=$2,dob=$3,phno=$4,bloodgroup=$5,address=$6,gender=$7 where id=$8`
				_, er := db.Exec(updatestmt, col[1], col[2], col[3], col[4], col[5], col[6], col[7], col[0])
				log.Println(er)
				if er != nil {
					stat = "failed to update"
				}
				stat = "success"
			}
		} else {
			log.Println("Inside for delete")
			deletestmt := `delete from mrinal_db.student where id = $1`
			_, er := db.Exec(deletestmt, col[0])
			if er != nil {
				stat = "failed to update"
			}
			stat = "success"
		}
	}
	return &model.Exeloutput{Status: &stat}, nil
}
