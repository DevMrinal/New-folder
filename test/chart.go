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
	host4     = "localhost"
	port4     = 5432
	user4     = "postgres"
	password4 = ""
	dbname4   = "postgres"
)

func Enterexcel(ctx context.Context, input *model.Chartenter) (*model.ChartResponse, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host4, port4, user4, password4, dbname4)
	mn, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer mn.Close()
	type chartdata struct {
		country_name string
		population   int
	}
	var stat string
	var chartArr chartdata
	chart := excelize.NewFile()
	chart.SetSheetName("sheet1", "density")

	log.Println("New sheet created")

	chart.SetCellValue("density", "A1", "id")
	chart.SetCellValue("density", "B1", "country_name")
	//chart.SetCellValue("density", "C1", "continent")
//	chart.SetCellValue("density", "D1", "population")
//	if *input.Cht != "" {
		log.Println("Inside if for input")
		fetch := `SELECT * FROM mrinal_db.densities`
		log.Println("Inside if for input1")
		row, err := mn.Query(fetch)
		log.Println("Inside if for input2")
		if err != nil {
			return nil, err
		}
		log.Println("before loop")
		var rowNumber int = 2
	//	for _, col := range chartArr {
			for row.Next() {
				err := row.Scan( &chartArr.country_name, &chartArr.population)
				if err != nil {
					return nil, err
				}
			//	chart.SetCellValue("density", "A"+strconv.Itoa(rowNumber), chartArr.id)
				chart.SetCellValue("density", "B"+strconv.Itoa(rowNumber), chartArr.country_name)
			// 	chart.SetCellValue("density", "C"+strconv.Itoa(rowNumber), chartArr.continent)
				chart.SetCellValue("density", "D"+strconv.Itoa(rowNumber), chartArr.population)
				rowNumber++
			}
	log.Println("Inside  for chart")
	er := chart.AddChart("density", "E1", `{
		"type":"col","series":[
			{"name":"density!$B$2","country_name":"density!$B$1:$B$5","population":"density!$D$1:$D$5"}
			]
			"title":{"name":"population record list of 2015"}}`)
	crr := chart.SaveAs("D://Records.xlsx")
	if crr != nil {

		fmt.Println(er)
		stat = "failed"

	} else {
		stat = "Success"
	}
	return &model.ChartResponse{Status: &stat}, nil
}
