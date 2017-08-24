package db

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type conf struct {
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
}

func getConf() conf {

	yamlFile, err := ioutil.ReadFile("config/db.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var c conf
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func initDb() error {
	c := getConf()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=5432 host=vasilesk.ru",
		c.User, c.Password, c.Database)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	// checkErr(err)

	return err
}

func randKey() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func createKey(userId int) string {
	key := randKey()

	db.QueryRow("INSERT INTO access_keys(user_id,key) VALUES($1,$2);", userId, key)

	return key
}

func addUser(login string, password string) int {
	hash, _ := HashPassword(password)

	var lastInserted int
	err := db.QueryRow("INSERT INTO users(login,password) VALUES($1,$2) returning id;", login, hash).Scan(&lastInserted)
	checkErr(err)

	return lastInserted
}

func CheckUser(login string, password string) int {
	rows, err := db.Query("SELECT id, password FROM users WHERE login=$1 LIMIT 1;", login)
	checkErr(err)

	for rows.Next() {
		var userId int
		var dbHash string
		err = rows.Scan(&userId, &dbHash)
		checkErr(err)
		if CheckPasswordHash(password, dbHash) {
			return userId
		}
	}

	return 0
}

func init() {
	err := initDb()
	checkErr(err)
}

// func main() {
// 	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=5432 host=vasilesk.ru",
// 	// 	DB_USER, DB_PASSWORD, DB_NAME)
// 	// db, err := sql.Open("postgres", dbinfo)
// 	// db := initDb()
// 	// defer db.Close()
// 	//
//
// 	// inserted := addUser("vasilesk", "12345678")
// 	inserted := checkUser("vasilesk", "12345678")
// 	// inserted := createKey(1)
// 	fmt.Println("# Inserting values ", inserted)
//
// 	// fmt.Println("# Inserting values")
// 	//
// 	// var lastInsertId int
// 	// err := db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
// 	// checkErr(err)
// 	// fmt.Println("last inserted id =", lastInsertId)
// 	//
// 	// fmt.Println("# Updating")
// 	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
// 	// checkErr(err)
// 	//
// 	// res, err := stmt.Exec("astaxieupdate", lastInsertId)
// 	// checkErr(err)
// 	//
// 	// affect, err := res.RowsAffected()
// 	// checkErr(err)
// 	//
// 	// fmt.Println(affect, "rows changed")
// 	//
// 	// fmt.Println("# Querying")
// 	// rows, err := db.Query("SELECT * FROM userinfo")
// 	// checkErr(err)
// 	//
// 	// for rows.Next() {
// 	// 	var uid int
// 	// 	var username string
// 	// 	var department string
// 	// 	var created time.Time
// 	// 	err = rows.Scan(&uid, &username, &department, &created)
// 	// 	checkErr(err)
// 	// 	fmt.Println("uid | username | department | created ")
// 	// 	fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
// 	// }
// 	//
// 	// fmt.Println("# Deleting")
// 	// stmt, err = db.Prepare("delete from userinfo where uid=$1")
// 	// checkErr(err)
// 	//
// 	// res, err = stmt.Exec(lastInsertId)
// 	// checkErr(err)
// 	//
// 	// affect, err = res.RowsAffected()
// 	// checkErr(err)
// 	//
// 	// fmt.Println(affect, "rows changed")
// }
