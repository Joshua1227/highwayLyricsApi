package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetDbCreds(fileName string) string {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var creds DbCreds

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &creds)

	return creds.DbPassword
}

type DbCreds struct {
	DbPassword string `json:"mongodb"`
}

const Uri = `mongodb+srv://sandeepjoshuadaniel:%s@lyricsdb0.1rri3.mongodb.net/?retryWrites=true&w=majority&appName=LyricsDB0`
