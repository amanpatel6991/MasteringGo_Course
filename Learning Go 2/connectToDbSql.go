package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type TestConn struct {
	Id   int             `json:"id"`
	Name  string         `json:"name"`
}

func (TestConn) TableName() string {
	return "TestConn"
}

type CustomTable struct {
	Id   int             `json:"id"`
	Name  string         `json:"name"`
}

func (CustomTable) TableName() string {
	return "t_custom"
}

func main() {
	//dbUser := "root"
	//dbPassword := "password"
	//dbName := "godb1"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	//	dbUser,
	//	dbPassword,
	//	"127.0.0.1:3306",
	//	dbName)
	//db, err := sql.Open("mysql", dsn)

	//db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	db, err := gorm.Open("mysql", "root:password@tcp(35.244.22.130)/godb1")
	if err != nil {
		fmt.Println(err)
	} else  {
		fmt.Println("Connected !")
	}

	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

	defer db.Close()

	var testData []TestConn

	fmt.Println("Testing Queries..........")
	db.Debug().Model(&TestConn{}).Find(&testData)

	fmt.Println(db.Debug().HasTable("TestConn"))
	db.Debug().AutoMigrate(&CustomTable{})
	//db.Raw("Select name From TestConn").Scan(&testData)

	fmt.Println(testData)


}
