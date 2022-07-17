package test

import (
	"context"
	"database/sql"
	"fmt"

	"log"
	"strconv"
	"strings"
	"time"

	"example.com/m/graph/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func TableInsert(ctx context.Context, input *model.Studententer) (*model.StudentResponse, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	resultdata := &model.StudentResponse{}
	_, e := IsValidDateWithDateObject(input.Dob)
	validPhn := checkPhno(*input.Phno)
	input.BloodGroup = strings.ToUpper(input.BloodGroup)
	validBloodgroup := checkBlood(input.BloodGroup)
	*input.Gender = strings.ToUpper(*input.Gender)
	validGender := checkGender(*input.Gender)
	if e == nil && validPhn && validBloodgroup && validGender && input.ID == nil {
		insertStd := `insert into mrinal_db.Student("firstname","lastname","dob","phno","bloodgroup","address","gender")values( $1,$2,$3,$4,$5,$6,$7)RETURNING ID`
		log.Println("write query")
		err := db.QueryRowContext(ctx, insertStd, input.Firstname, input.Lastname, input.Dob, input.Phno, input.BloodGroup, input.Address, input.Gender).Scan(&input.ID)
		if err != nil {
			resultdata.ID = nil
			var output1 string = "failed to insert data"
			resultdata.Status = append(resultdata.Status, &output1)
			return resultdata, err
		}
		resultdata.ID = input.ID
		var output2 string = "success-inserted into table"
		resultdata.Status = append(resultdata.Status, &output2)
		return resultdata, err
	} else if input.ID != nil {
		updatestd := `update mrinal_db.student set firstname=$1,lastname=$2,dob=$3,phno=$4,bloodgroup=$5,address=$6,gender=$7 where id=$8`
		_, err := db.Exec(updatestd, input.Firstname, input.Lastname, input.Dob, input.Phno, input.BloodGroup, input.Address, input.Gender, input.ID)
		if err != nil {
			resultdata.ID = nil
			var output1 string = "failed to update data"
			resultdata.Status = append(resultdata.Status, &output1)
			return resultdata, err
		}
		resultdata.ID = input.ID
		var output2 string = "success-updated into table"
		resultdata.Status = append(resultdata.Status, &output2)
		return resultdata, err
	}
	resultdata.ID = nil
	if e != nil {
		var err1 string = "failed-date format"
		resultdata.Status = append(resultdata.Status, &err1)
	}
	if !validPhn {
		var err2 string = "failed-Phno format"
		resultdata.Status = append(resultdata.Status, &err2)
	}
	if !validBloodgroup {
		var err3 string = "failed-bloodgroup format"
		resultdata.Status = append(resultdata.Status, &err3)
	}
	if !validGender {
		var err4 string = "failed-gender format"
		resultdata.Status = append(resultdata.Status, &err4)
	}
	return resultdata, nil
}
func IsValidDateWithDateObject(date string) (time.Time, error) {
	const (
		layoutISO = "2006/01/02"
	)
	timeobject, err := time.Parse(layoutISO, date)
	if err != nil {
		return timeobject, err
	}
	return timeobject, nil
}
func checkPhno(phno int) bool {
	m := strconv.Itoa(phno)
	if len(m) == 10 {
		return true
	}
	return false
}
func checkBlood(blood string) bool {
	bloodmap := make(map[string]string)

	bloodmap["group1"] = "A+"
	bloodmap["group2"] = "A-"
	bloodmap["group3"] = "B+"
	bloodmap["group4"] = "B-"
	bloodmap["group5"] = "O+"
	bloodmap["group6"] = "O-"
	bloodmap["group7"] = "AB+"
	bloodmap["group8"] = "AB-"
	for _, bgroup := range bloodmap {
		if blood == bgroup {
			return true
		}

	}
	return false
}
func checkGender(gender string) bool {
	gendermap := make(map[string]string)
	gendermap["type1"] = "MALE"
	gendermap["type2"] = "FEMALE"
	for _, gendtype := range gendermap {
		if gender == gendtype {
			return true
		}
	}
	return false
}
