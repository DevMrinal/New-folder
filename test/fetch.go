package test

import (
	"context"
	"database/sql"
	"fmt"

	//"strings"

	"example.com/m/graph/model"
	_ "github.com/lib/pq"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

const (
	host1     = "localhost"
	port1     = 5432
	user1     = "postgres"
	password1 = ""
	dbname1   = "postgres"
)

func Fetch(ctx context.Context, input *model.DataFetch) ([]*model.Studentoutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host1, port1, user1, password1, dbname1)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	if input.Inst != nil {
		fetchrow := `SELECT * FROM mrinal_db.student WHERE firstname ILIKE  $1 `
		row, err := db.Query(fetchrow, "%"+*input.Inst+"%")
		rowentry := []*model.Studentoutput{}
		if err != nil {
			return nil, fmt.Errorf("RowFromQuery %v: %v", input.Inst, err)
		}
		for row.Next() {
			resultRow := &model.Studentoutput{}
			err := row.Scan(&resultRow.ID, &resultRow.Firstname, &resultRow.Lastname, &resultRow.Dob, &resultRow.Phno, &resultRow.BloodGroup, &resultRow.Address, &resultRow.Gender)
			if err != nil {
				return nil, fmt.Errorf("RowFromQuery %v:%v", input.Inst, err)
			}
			if resultRow.Lastname == nil {
				a := ""
				resultRow.Lastname = &a
			}
			if resultRow.Gender == nil {
				b := ""
				resultRow.Gender = &b
			}
			rowentry = append(rowentry, resultRow)
		}
		return rowentry, nil
	}

	fetchrow := `SELECT *FROM mrinal_db.student`
	row, err := db.Query(fetchrow)

	rowentry := []*model.Studentoutput{}
	if err != nil {
		return nil, fmt.Errorf("RowFromquery %v:%v", input.Inst, err)
	}

	for row.Next() {
		resultRow := &model.Studentoutput{}
		err := row.Scan(&resultRow.Firstname, &resultRow.Lastname, &resultRow.Dob, &resultRow.Phno, &resultRow.BloodGroup, &resultRow.Address, &resultRow.Gender)
		if err != nil {
			return nil, fmt.Errorf("RowFromQuery %v:%v", input.Inst, err)
		}
		if resultRow.Lastname == nil {
			a := ""
			resultRow.Lastname = &a
		}
		if resultRow.Gender == nil {
			b := ""
			resultRow.Gender = &b
		}
		rowentry = append(rowentry, resultRow)
	}
	return rowentry, nil
}
