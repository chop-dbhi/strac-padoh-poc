package converter

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
)

var STRACColumns = []string{
	"id",
	"street",
	"city",
	"state",
	"zip",
	"patient_first_name",
	"patient_last_name",
	"patient_dob",
	"patient_address",
	"patient_city",
	"patient_state",
	"patient_county",
	"patient_zip",
	"patient_callback_number",
	"patient_sex",
	"patient_race",
	"patient_ethnicity",
	"patient_results",
	"patient_positive",
	"reason",
}

// STRACRecord represents a STRAC record of data.
type STRACRecord struct {
	ID                    string `json:"id"`
	Street                string `json:"street"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Zip                   string `json:"zip"`
	PatientFirstName      string `json:"patient_first_name"`
	PatientLastName       string `json:"patient_last_name"`
	PatientDOB            string `json:"patient_dob"`
	PatientAddress        string `json:"patient_address"`
	PatientCity           string `json:"patient_city"`
	PatientCounty         string `json:"patient_county"`
	PatientState          string `json:"patient_state"`
	PatientZip            string `json:"patient_zip"`
	PatientCallbackNumber string `json:"patient_callback_number"`
	PatientSex            string `json:"patient_sex"`
	PatientRace           string `json:"patient_race"`
	PatientEthnicity      string `json:"patient_ethnicity"`
	PatientResults        string `json:"patient_results"`
	PatientPositive       string `json:"patient_positive"`
	Reason                string `json:"reason"`
}

type Column struct {
	Name     string
	Required bool
	Values   map[string]struct{}
	Mapper   func(record *STRACRecord) (string, error)
}

type ValidationResult struct {
	Errors   []string
	Warnings []string
}

// indexSTRACHeader validates and builds an index from column to index.
func indexSTRACHeader(header []string) (map[string]int, *ValidationResult) {
	index := make(map[string]int)

	// Keep track of the columns that are found in the provided header.
	foundColumns := make(map[string]bool)
	for _, col := range STRACColumns {
		foundColumns[col] = false
	}

	var (
		errors   []string
		warnings []string
	)

	for i, col := range header {
		// Skip empty columns.
		if col == "" {
			continue
		}

		found, ok := foundColumns[col]
		// Found previously, note the column.
		if found {
			errors = append(errors, fmt.Sprintf("duplicate column: %s", col))
			continue
		}

		// Not an expected column. Treat as warning since it will be ignored.
		if !ok {
			warnings = append(warnings, fmt.Sprintf("unexpected column: %s", col))
			continue
		}

		// Expected and not found, so mark good.
		foundColumns[col] = true
		index[col] = i
	}

	// Check for expected columns that were not in the header.
	for col, found := range foundColumns {
		if !found {
			warnings = append(warnings, fmt.Sprintf("column not found: %s", col))
		}
	}

	return index, &ValidationResult{
		Errors:   errors,
		Warnings: warnings,
	}
}

func readBom(r io.Reader) (io.Reader, error) {
	br := bufio.NewReader(r)
	rn, _, err := br.ReadRune()
	if err != nil {
		return nil, err
	}
	// Check if BOM
	if rn != '\uFEFF' {
		br.UnreadRune()
	}
	return br, nil
}

func rowValue(row []string, index map[string]int, col string) string {
	idx, ok := index[col]
	if !ok {
		return ""
	}
	return row[idx]
}

func Convert(r io.Reader, w io.Writer, columns []*Column) error {
	// Warn on columns.
	for _, c := range columns {
		if c.Mapper != nil {
			continue
		}

		if len(c.Values) > 0 {
			log.Printf("column %v: no mapper defined, but a value set defined", c.Name)
		}
		if c.Required {
			log.Printf("column %v: no mapper defined, but marked required", c.Name)
		}
	}

	r, err := readBom(r)
	if err != nil {
		return err
	}

	// Parse STRAC CSV data.
	cr := csv.NewReader(r)
	head, err := cr.Read()
	if err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	// Validate the head columns match the expected ones.

	// Build an index of the STRAC columns and validate the column names.
	stracIndex, validationResult := indexSTRACHeader(head)
	// Log warnings
	for _, warn := range validationResult.Warnings {
		log.Printf("validation warning: %s", warn)
	}

	// Fail on validation error.
	if len(validationResult.Errors) > 0 {
		for _, err := range validationResult.Warnings {
			log.Printf("validation error: %s", err)
		}
		return errors.New("validation errors")
	}

	cw := csv.NewWriter(w)
	defer cw.Flush()

	// Write the target header.
	var targetHeader []string
	for _, col := range columns {
		targetHeader = append(targetHeader, col.Name)
	}
	cw.Write(targetHeader)

	// Read the rows in the STRAC file and produce rows in the target format.
	var rowNum int

	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("read row: %w", err)
		}

		rowNum++

		// Produce a record that will be used within the mapper functions.
		record := &STRACRecord{
			ID:                    rowValue(row, stracIndex, "id"),
			Street:                rowValue(row, stracIndex, "street"),
			City:                  rowValue(row, stracIndex, "city"),
			State:                 rowValue(row, stracIndex, "state"),
			Zip:                   rowValue(row, stracIndex, "zip"),
			PatientFirstName:      rowValue(row, stracIndex, "patient_first_name"),
			PatientLastName:       rowValue(row, stracIndex, "patient_last_name"),
			PatientDOB:            rowValue(row, stracIndex, "patient_dob"),
			PatientAddress:        rowValue(row, stracIndex, "patient_address"),
			PatientCity:           rowValue(row, stracIndex, "patient_city"),
			PatientState:          rowValue(row, stracIndex, "patient_state"),
			PatientCounty:         rowValue(row, stracIndex, "patient_county"),
			PatientZip:            rowValue(row, stracIndex, "patient_zip"),
			PatientCallbackNumber: rowValue(row, stracIndex, "patient_callback_number"),
			PatientSex:            rowValue(row, stracIndex, "patient_sex"),
			PatientRace:           rowValue(row, stracIndex, "patient_race"),
			PatientEthnicity:      rowValue(row, stracIndex, "patient_ethnicity"),
			PatientResults:        rowValue(row, stracIndex, "patient_resuls"),
			PatientPositive:       rowValue(row, stracIndex, "patient_positive"),
			Reason:                rowValue(row, stracIndex, "reason"),
		}

		data := make([]string, len(columns))

		for i, col := range columns {
			// Nothing to map, skip.
			if col.Mapper == nil {
				continue
			}

			value, err := col.Mapper(record)
			if err != nil {
				log.Println(err)
			}

			if col.Required && value == "" {
				log.Printf("row %d: missing value for %v", rowNum, col.Name)
				continue
			}

			if len(col.Values) > 0 {
				if _, ok := col.Values[value]; !ok {
					log.Printf("row %d: invalid value for %v: %v", rowNum, col.Name, value)
				}
			}

			data[i] = value
		}

		if err := cw.Write(data); err != nil {
			return fmt.Errorf("write row: %w", err)
		}
	}

	return nil
}
