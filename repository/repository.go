package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"zinkworks/grad/server/models"

	log "github.com/inconshreveable/log15"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DbInit() {
	dsn := ""
	dsn = fmt.Sprintf("host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("repository didnt connect")
		panic("failed to connect database")

	}
	insertTablesAndDefaultData()

}
func insertTablesAndDefaultData() {
	checkReportsTable := db.Migrator().HasTable("reports")
	pwd, _ := os.Getwd()
	editpwd := strings.TrimSuffix(pwd, "/controller")
	log.Debug(fmt.Sprintf("pwd %v", pwd))
	if checkReportsTable {
		log.Debug("checkReportTable is there")
	} else {

		reportJsonFile, _ := ioutil.ReadFile(editpwd + "/repository/reports.json")
		var report []models.Report
		JSONStringToStructure(string(reportJsonFile), &report)
		db.AutoMigrate(&models.Report{})
		{
			for index := range report {

				db.Create(&report[index])
				fmt.Print(index)
			}
		}
		log.Debug("crated the report table ")
	}

}
func JSONStringToStructure(jsonString string, structure interface{}) error {
	jsonBytes := []byte(jsonString)
	return json.Unmarshal(jsonBytes, structure)
}
func GetALLReport() []models.Report {
	var results []models.Report
	if results := db.Find(&results).Error; results != nil {
		log.Error("unable to connect to database ")
	}
	return results
}
