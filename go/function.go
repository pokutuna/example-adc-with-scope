package function

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	spreadsheet "google.golang.org/api/sheets/v4"
)

var SPREADSHEET_ID = os.Getenv("SPREADSHEET_ID")
var SHEET_RANGE = os.Getenv("SHEET_RANGE")

type RowChange struct {
	From []interface{} `json:"from"`
	To   []interface{} `json:"to"`
}

func App(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	// もし ADC に quota project を設定しない派であれば
	// option.WithQuotaProject(projectID) を渡す
	client, err := spreadsheet.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	resp, err := client.Spreadsheets.Values.Get(SPREADSHEET_ID, SHEET_RANGE).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	count, _ := strconv.Atoi(resp.Values[0][0].(string))

	valueRange := &spreadsheet.ValueRange{
		Values: [][]interface{}{
			[]interface{}{count + 1, time.Now().Format(time.RFC3339), "go"},
		},
	}
	updated, err := client.Spreadsheets.Values.Update(SPREADSHEET_ID, SHEET_RANGE, valueRange).ValueInputOption("USER_ENTERED").IncludeValuesInResponse(true).Do()
	if err != nil {
		log.Fatalf("Unable to update data to sheet: %v", err)
	}

	change := &RowChange{
		From: resp.Values[0],
		To:   updated.UpdatedData.Values[0],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.Encode(change)
}
