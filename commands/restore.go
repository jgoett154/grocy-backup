package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Restore(ctx *cli.Context) error {
	// Open backup file given
	backupFileName := ctx.Args().First()
	backupFile, err := os.Open(backupFileName)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	// Decode as JSON file
	dec := json.NewDecoder(backupFile)

	var backupData = make(map[string][]map[string]interface{})

	err = dec.Decode(&backupData)
	if err != nil {
		return err
	}

	for entity, dataList := range backupData {
		var client = &http.Client{}
		var url = fmt.Sprintf("%s/objects/%s", ctx.String("server"), entity)

		// As each entity must be added one at a time, loop over array
		for _, data := range dataList {
			// Encode data as JSON for request
			b := new(bytes.Buffer)
			err = json.NewEncoder(b).Encode(data)
			if err != nil {
				return err
			}

			// Create the request and set API key for access
			req, err := http.NewRequest("POST", url, b)
			if err != nil {
				return err
			}

			req.Header.Add("GROCY-API-KEY", ctx.String("api-key"))
			req.Header.Add("Content-Type", "application/json")

			// Send request
			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			entityName := data["name"]
			if entityName == nil {
				entityName = data["barcode"]
			}
			if entityName == nil {
				entityName = "N/A"
			}

			// Log status info
			if resp.StatusCode == 200 {
				log.Printf("Successfully restored `%s` entity `%s`", entity, entityName)
			} else {
				log.Printf("Failed to restore `%s` entity `%s`. Likely already exists in Grocy", entity, entityName)
			}
		}
	}
	
	return nil
}
