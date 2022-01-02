// package database will provide a way to connect to SQL databases with GORM
package database

import (
	"database/sql"
	"fmt"

	"github.com/bseto/arcade/backend/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GenerateMySQLConnectionString will generate the string for the normal connection
// and the root connection
// root:cubskwadskwad@tcp(127.0.0.1:3306)
// user:cubcubskwad@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local"
func GenerateMySQLConnectionString(
	rootUsername,
	rootPassword,
	username,
	password,
	address,
	port,
	databaseName string,
) (string, string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		address,
		port,
		databaseName,
	)

	rootDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		rootUsername,
		rootPassword,
		address,
		port,
	)
	return dsn, rootDsn
}

// SetupDB will create the database if it does not exist
// the function will then connect as user and setup the database if it is newly
// created
func SetupDB(
	rootUsername,
	rootPassword,
	username,
	password,
	address,
	port,
	databaseName string,
) (*gorm.DB, bool, error) {
	rootDSN, DSN := GenerateMySQLConnectionString(
		rootUsername,
		rootPassword,
		username,
		password,
		address,
		port,
		databaseName,
	)
	newlyCreated, err := createDBIfNotExists(rootDSN, databaseName)
	if err != nil {
		log.Errorf("unable to create DB: %v\n", err)
		return nil, newlyCreated, err
	}
	db, err := connectWithGorm(DSN)
	if err != nil {
		log.Errorf("unable to open db: %v\n", err)
		return db, newlyCreated, err
	}

	return db, newlyCreated, nil
}

func connectWithGorm(dsn string) (*gorm.DB, error) {
	fmt.Println("connecting gorm")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("unable to open db: %v", err)
	}
	return db, err
}

func createDBIfNotExists(rootDSN, databaseName string) (createdDB bool, err error) {
	fmt.Println("connecting to db")
	db, err := sql.Open("mysql", rootDSN)
	if err != nil {
		log.Errorf("unable to open database: %v\n", err)
		return
	}
	defer db.Close()

	res := db.QueryRow("SHOW DATABASES like '%" + databaseName + "%'")
	var existingDatabase string
	err = res.Scan(&existingDatabase)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("unable to call rows affected: %v\n", err)
		return
	}
	if existingDatabase != databaseName {
		createdDB = true
		log.Infof("creating db since it doesn't exist")
		_, err = db.Exec("CREATE DATABASE " + databaseName)
		if err != nil {
			log.Errorf("unable to create database: %v\n", err)
			return
		}
		fmt.Println("granting access")
		_, err = db.Exec(`grant all on *.* TO 'user'@'%'`)
		if err != nil {
			log.Errorf("unable to get rows affected: %v", err)
			return
		}
	}

	return
}
