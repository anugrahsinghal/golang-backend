package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	path string
}

func NewClient(path string) Client {
	cl := Client{
		path,
	}
	return cl
}

func (c Client) createDB() error {
	dbSchema := databaseSchema{
		make(map[string]User),
		make(map[string]Post),
	}
	jsonData, err := json.Marshal(dbSchema)
	if err != nil {
		fmt.Println("createDB", err)
		return err
	}
	err = os.WriteFile(c.path, jsonData, 0666)
	if err != nil {
		fmt.Println("write", err)
		return err
	}
	return nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.path)
	if err == nil {
		// file found
		return nil
	}
	fmt.Println("EnsureDB", err)
	// else create file
	fmt.Println("DB CREATE")
	return c.createDB()
}

// should save data in given schema to db path and overwrite
func (c Client) updateDB(db databaseSchema) error {
	jsonData, err := json.Marshal(db)
	// fmt.Println(string(jsonData))
	if err != nil {
		fmt.Println("updateDB", err)
		return err
	}
	if err := os.WriteFile(c.path, jsonData, 0666); err != nil {
		fmt.Println("updateDB2", err)
		return err
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	dbSchema := databaseSchema{}
	data, err := os.ReadFile(c.path)
	if err != nil {
		fmt.Println("readDB", err)
		return dbSchema, err
	}
	err = json.Unmarshal(data, &dbSchema)
	if err != nil {
		fmt.Println("Unmarshal", err)
		return dbSchema, err
	}
	// fmt.Println(dbSchema)
	return dbSchema, nil
}
