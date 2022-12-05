package utils

import (
	"log"
	"toolforthi/model"

	"github.com/xuri/excelize/v2"
)

func ReadExcel() (map[string]model.Algorithms, error) {
	resp := map[string]model.Algorithms{}
	file, error := excelize.OpenFile(model.EXCEL_PATH)
	if error != nil {
		log.Fatal(error)
	}
	rowIndex, err := file.GetRows(model.SHEET_NAME)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// caching Issue name
	StageIssue := ""
	useCache := false
	// loop all row of sheet
	rowCount := 0
	for _, row := range rowIndex {
		rowCount++
		if len(row) == 0 {
			// skip empty row
			continue
		}
		if len(row) < 6 {
			break
		}
		// row is an array of string
		// do: paste string to response
		issueName := row[5]
		description := row[6]
		quickfix := row[7]
		if issueName == "" {
			issueName = StageIssue
			useCache = true
		} else {
			StageIssue = issueName
			useCache = false
		}
		if useCache { // algorithm is the same as previous row, append fixes
			t := resp[issueName]
			t.Inspector.Fixes = append(t.Inspector.Fixes, model.Fix{
				Name:     quickfix,
				Patterns: []string{},
				Strings:  []string{},
			})
			resp[issueName] = t
		} else { // create new algorithm
			algori := model.Algorithms{
				Id:               issueName,
				BriefDescription: description,
				Inspector: model.Inspector{
					Pattern:      "",
					GroupPattern: []string{},
					Fixes: []model.Fix{
						{
							Name:     quickfix,
							Patterns: []string{},
							Strings:  []string{},
						},
					},
					HightLightType: "WARNING",
				},
			}
			resp[issueName] = algori
		}
	}
	return resp, nil
}
