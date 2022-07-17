package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	_ "strings"

	"example.com/m/graph/model"
	_ "github.com/lib/pq"

	//"golang.org/x/tools/go/analysis/passes/nilfunc"

	//"golang.org/x/tools/go/analysis/passes/nilfunc"

	// "github.com/360EnSecGroup-Skylar/excelize/v2"
	"github.com/xuri/excelize/v2"
)

const (
	host3     = "localhost"
	port3     = 5432
	user3     = "postgres"
	password3 = ""
	dbname3   = "postgres"
)

func Insertexcel(ctx context.Context, input *model.DbInput) (*model.Response, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host3, port3, user3, password3, dbname3)
	db, err := sql.Open("postgres", psqlinfo)
	log.Println(db, err)
	log.Println("Starting..")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	type exceldata struct {
		id         string
		firstname  string
		lastname   sql.NullString
		dob        string
		Phno       int
		bloodgroup string
		address    string
		Gender     string
	}
	var stat string
	var excelArr = [20]exceldata{}
	f := excelize.NewFile()
	//s := f.NewSheet("survey")
	f.SetSheetName("Sheet1", "survey")
	log.Println("New sheet created")
	//	fmt.Println(s)
	f.SetCellValue("survey", "A1", "id")
	f.SetCellValue("survey", "B1", "firstname")
	f.SetCellValue("survey", "C1", "lastname")
	f.SetCellValue("survey", "D1", "dob")
	f.SetCellValue("survey", "E1", "phno")
	f.SetCellValue("survey", "F1", "bloodGroup")
	f.SetCellValue("survey", "G1", "address")
	f.SetCellValue("survey", "H1", "Gender")
	if *input.Inp != "" {
		log.Println("Inside if for input")
		fetchRow := `SELECT * FROM mrinal_db.student WHERE firstname ILIKE  '%` + *input.Inp + `%' `
		row, err := db.Query(fetchRow)
		if err != nil {
			return nil, fmt.Errorf("RowFromQuery %v: %v", input.Inp, err)
		}
		log.Println("before loop")

		var rowNumber int = 2
		for _, colEx := range excelArr {
			for row.Next() {
				err := row.Scan(&colEx.id, &colEx.firstname, &colEx.lastname, &colEx.dob, &colEx.Phno, &colEx.bloodgroup, &colEx.address, &colEx.Gender)
				if err != nil {
					return nil, fmt.Errorf("RowFromQuery %v: %v", input.Inp, err)
				}
				//	log.Println("Id found is", colEx.id)

				f.SetCellValue("survey", "A"+strconv.Itoa(rowNumber), colEx.id)
				f.SetCellValue("survey", "B"+strconv.Itoa(rowNumber), colEx.firstname)
				if !colEx.lastname.Valid {
					f.SetCellValue("survey", "C"+strconv.Itoa(rowNumber), "")
				} else {
					f.SetCellValue("survey", "C"+strconv.Itoa(rowNumber), colEx.lastname.String)
				}
				f.SetCellValue("survey", "D"+strconv.Itoa(rowNumber), colEx.dob)
				f.SetCellValue("survey", "E"+strconv.Itoa(rowNumber), colEx.Phno)
				f.SetCellValue("survey", "F"+strconv.Itoa(rowNumber), colEx.bloodgroup)
				f.SetCellValue("survey", "G"+strconv.Itoa(rowNumber), colEx.address)
				f.SetCellValue("survey", "H"+strconv.Itoa(rowNumber), colEx.Gender)
				rowNumber++
			}
		}
	} else {
		log.Println("into else")
		fetchRow := `SELECT * FROM mrinal_db.student`
		row, err := db.Query(fetchRow)
		log.Println("query result-", row)
		if err != nil {

			return nil, fmt.Errorf("RowFromQuery %v: %v", input.Inp, err)
			var rowNumber int = 2
			for _, colEx := range excelArr {
				for row.Next() {
					err := row.Scan(&colEx.id, &colEx.firstname, &colEx.lastname, &colEx.dob, &colEx.Phno, &colEx.bloodgroup, &colEx.address, &colEx.Gender)
					if err != nil {
						return nil, fmt.Errorf("RowFromQuery %v: %v", input.Inp, err)
					}
					log.Println("query result-", colEx)
					f.SetCellValue("survey", "A"+strconv.Itoa(rowNumber), colEx.id)
					f.SetCellValue("survey", "B"+strconv.Itoa(rowNumber), colEx.firstname)
					if !colEx.lastname.Valid {
						f.SetCellValue("survey", "C"+strconv.Itoa(rowNumber), "")
					} else {
						f.SetCellValue("survey", "C"+strconv.Itoa(rowNumber), colEx.lastname.String)
					}
					f.SetCellValue("survey", "D"+strconv.Itoa(rowNumber), colEx.dob)
					f.SetCellValue("survey", "E"+strconv.Itoa(rowNumber), colEx.Phno)
					f.SetCellValue("survey", "F"+strconv.Itoa(rowNumber), colEx.bloodgroup)
					f.SetCellValue("survey", "G"+strconv.Itoa(rowNumber), colEx.address)
					f.SetCellValue("survey", "H"+strconv.Itoa(rowNumber), colEx.Gender)
					rowNumber++
				}
			}
		}

	}
	er := f.SaveAs("D://villagers.xlsx")
	if er != nil {

		fmt.Println(er)
		stat = "failed"

	} else {
		stat = "Success"
	}
	return &model.Response{Status: &stat}, nil
}
