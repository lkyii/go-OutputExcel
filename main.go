package main

import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	//
	//var rows int = 10000
	var dbusername string = "root"
	var dbpassword string = "root"
	var dbhostip string = "127.0.0.1:3306"
	var dbname string = "beego"

	db,err:=sql.Open("mysql",dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")

	checkErr(err)
	Get(db)
}
//
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}


func Get(db *sql.DB) {
	rows, err := db.Query("select  * from users limit ?", 10)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		line := Users{}
		err = rows.Scan(&line)
		log.Println(line)
	}

	var arr []map[string]interface{}
	for rows.Next() {
		var m map[string]interface{}
		var userName string
		var userPassword string
		if err := rows.Scan(&userName,&userPassword); err != nil { 
			log.Fatal(err)
		}

		m["userName"] = userName
		m["userPassword"] = userPassword

		arr = append(arr,m)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	CreatExcel		(arr)
}

func CreatExcel(arr []map[string]interface{}) {

	t1 := time.Now()

	f := excelize.NewFile()

	index := f.NewSheet("Sheet1")

	for key, value := range arr {
		A := fmt.Sprintf("A%d",key)
		B := fmt.Sprintf("B%d",key)


		f.SetCellValue("Sheet1", A, value["ciss_id"])
		f.SetCellValue("Sheet1", B, value["part_name"])
		f.SetCellValue("Sheet1", B, value["part_description"])
	}

	f.SetActiveSheet(index)

	err := f.SaveAs("./Book1.xlsx")

	if err != nil {
		fmt.Printf("%s",err)
	}

	elapsed := time.Since(t1)
	result := fmt.Sprintf("%s", elapsed)
	fmt.Println(result)

}