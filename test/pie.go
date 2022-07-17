package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	//"log"
	//	"strconv"
	_ "strings"

	//"example.com/m/graph/model"
	//"example.com/m/graph/model"
	"example.com/m/graph/model"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

const (
	host5     = "localhost"
	port5     = 5432
	user5     = "postgres"
	password5 = ""
	dbname5   = "postgres"
)

func Enterpie(ctx context.Context) (*model.PieResponse, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host5, port5, user5, password5, dbname5)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	type Marksheet struct {
		name  string
		marks int
		// id    string
	}
	var s  Marksheet
	//log.Println("fetching")
	var stat string
	student := make([]opts.PieData, 0)
	fetchRow := `SELECT name,marks FROM mrinal_db.results`
	row, err := db.Query(fetchRow)
	if err != nil {
		log.Println("the error from query is:", err)
		return nil, err
	}
	//log.Println("before loop")
	// for row.Next() {
		// log.Println(i)
		for row.Next() {
			if err := row.Scan(&s.name, &s.marks); err != nil {
				return nil, err
			}
			student = append(student, opts.PieData{
				Name:  s.name,
				Value: s.marks})
		}
	// }
	pie := charts.NewPie()

	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{Title: "Annual result contribution"}),
	)
	pie.AddSeries("Marks", student)

	f, e := os.Create("Pie.html")
	if e != nil {
		stat = "Failed due to-" + e.Error()
	} else {
		stat = "Success"
	}
	pie.Render(f)

	return &model.PieResponse{Status: &stat}, nil
}
