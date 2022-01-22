package commands

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
)

var defaultEntities = []string{"products", "chores", "product_barcodes", "batteries", "locations", "quantity_units", "quantity_unit_conversions", "shopping_list", "shopping_lists", "shopping_locations", "recipes", "recipes_pos", "recipes_nestings", "tasks", "task_categories", "product_groups", "equipment", "userfields", "userentities", "userobjects", "meal_plan", "stock_log", "stock", "stock_current_locations", "chores_log", "meal_plan_sections"}

func Backup(ctx *cli.Context) error {
	var backupData = make(map[string]interface{})

	// Run API requests (/api/objects/%s) to dump entire DB to JSON
	for _, entity := range defaultEntities {
		var client = &http.Client{}
		var url = fmt.Sprintf("%s/objects/%s", ctx.String("server"), entity)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		// Add the API key header
		req.Header.Add("GROCY-API-KEY", ctx.String("api-key"))

		// Make the request and check for errors
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Decode the JSON response
		dec := json.NewDecoder(resp.Body)
		var data interface{}

		err = dec.Decode(&data)
		if err != nil {
			return err
		}

		// Store it in the map for backup
		backupData[entity] = data
	}

	// Open file for writing
	fileName := fmt.Sprintf("backup-%s.json", time.Now().Format(time.RFC3339))
	outFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write backup data to flat JSON file
	enc := json.NewEncoder(outFile)
	err = enc.Encode(backupData)
	if err != nil {
		return err
	}

	return nil
}
