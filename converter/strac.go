package converter

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Define STRAC field names as constants to have compile-time safety of lookups.
var STRACColumns = []string{
	"Reporting_Facility_Name",
	"CLIA_Number",
	"Performing_Organization_Name",
	"Performing_Organization_Address",
	"Performing_Organization_City",
	"Performing_Organization_Zip",
	"Performing_Organization_State",
	"Device_Identifier",
	"Ordered_Test_Name",
	"LOINC_Code",
	"LOINC_Text",
	"Result",
	"Result_Units",
	"Reference_Range",
	"Date_Test_Performed",
	"Test_Result_Date",
	"Pt_Fname",
	"Pt_Middle_Initial",
	"Pt_Lname",
	"Date_of_Birth",
	"Patient Age",
	"Sex",
	"Pt_Race",
	"Pt_Ethnicity",
	"Pt_Phone",
	"Pt_Str",
	"Pt_City",
	"Pt_ST",
	"Pt_Zip",
	"Pt_County",
	"Accession_Number",
	"Ordering_Facility",
	"Ordering_Facility_Address",
	"Ordering_Facility_City",
	"Ordering_Facility_State",
	"Ordering_Facility_Zip",
	"Ordering_Provider_Last_Name",
	"Ordering_Provider_First_Name",
	"Ordering_Provider_NPI",
	"Ordering_Provider_Street_Address",
	"Ordering_Provider_City",
	"Ordering_Provider_State",
	"Ordering_Provider_Zip",
	"Ordering_Provider_Phone",
	"Specimen_ID",
	"Specimen_Type",
	"Date_Test_Ordered",
	"Date_Specimen_Collected",
}

// STRACRecord represents a STRAC record of data.
type STRACRecord struct {
	Reporting_Facility_Name          string
	CLIA_Number                      string
	Performing_Organization_Name     string
	Performing_Organization_Address  string
	Performing_Organization_City     string
	Performing_Organization_Zip      string
	Performing_Organization_State    string
	Device_Identifier                string
	Ordered_Test_Name                string
	LOINC_Code                       string
	LOINC_Text                       string
	Result                           string
	Result_Units                     string
	Reference_Range                  string
	Date_Test_Performed              string
	Test_Result_Date                 string
	Pt_Fname                         string
	Pt_Middle_Initial                string
	Pt_Lname                         string
	Date_of_Birth                    string
	Patient_Age                      string
	Sex                              string
	Pt_Race                          string
	Pt_Ethnicity                     string
	Pt_Phone                         string
	Pt_Str                           string
	Pt_City                          string
	Pt_ST                            string
	Pt_Zip                           string
	Pt_County                        string
	Accession_Number                 string
	Ordering_Facility                string
	Ordering_Facility_Address        string
	Ordering_Facility_City           string
	Ordering_Facility_State          string
	Ordering_Facility_Zip            string
	Ordering_Provider_Last_Name      string
	Ordering_Provider_First_Name     string
	Ordering_Provider_NPI            string
	Ordering_Provider_Street_Address string
	Ordering_Provider_City           string
	Ordering_Provider_State          string
	Ordering_Provider_Zip            string
	Ordering_Provider_Phone          string
	Specimen_ID                      string
	Specimen_Type                    string
	Date_Test_Ordered                string
	Date_Specimen_Collected          string
}

type Column struct {
	Name     string
	Required bool
	Values   map[string]struct{}
	Mapper   func(record *STRACRecord) (string, error)
}

func Main(columns []*Column) error {
	log.SetFlags(0)

	// Supports both input and paths paths, input path (and STDOUT),
	// or no paths (STDIN and STDOUT).
	var (
		input      io.Reader
		output     io.Writer
		inputPath  string
		outputPath string
	)

	args := os.Args[1:]

	switch len(args) {
	case 0:
	case 1:
		inputPath = args[0]
	case 2:
		inputPath = args[0]
		outputPath = args[1]
	}

	if inputPath != "" {
		f, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("open input file: %w", err)
		}
		defer f.Close()
		input = f
	} else {
		input = os.Stdin
	}

	if outputPath != "" {
		f, err := os.Create(args[0])
		if err != nil {
			return fmt.Errorf("create output file: %w", err)
		}
		defer f.Close()
		output = f
	} else {
		output = os.Stdout
	}

	return Convert(input, output, columns)
}

type ValidationResult struct {
	Errors   []string
	Warnings []string
}

// indexSTRACHeader validates and builds an index from column to index.
// TODO: normalize column names, account for lowercase, etc?
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
			errors = append(errors, fmt.Sprintf("column not found: %s", col))
		}
	}

	return index, &ValidationResult{
		Errors:   errors,
		Warnings: warnings,
	}
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

	r, err := ReadBom(r)
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
			Reporting_Facility_Name:          row[stracIndex["Reporting_Facility_Name"]],
			CLIA_Number:                      row[stracIndex["CLIA_Number"]],
			Performing_Organization_Name:     row[stracIndex["Performing_Organization_Name"]],
			Performing_Organization_Address:  row[stracIndex["Performing_Organization_Address"]],
			Performing_Organization_City:     row[stracIndex["Performing_Organization_City"]],
			Performing_Organization_Zip:      row[stracIndex["Performing_Organization_Zip"]],
			Performing_Organization_State:    row[stracIndex["Performing_Organization_State"]],
			Device_Identifier:                row[stracIndex["Device_Identifier"]],
			Ordered_Test_Name:                row[stracIndex["Ordered_Test_Name"]],
			LOINC_Code:                       row[stracIndex["LOINC_Code"]],
			LOINC_Text:                       row[stracIndex["LOINC_Text"]],
			Result:                           row[stracIndex["Result"]],
			Result_Units:                     row[stracIndex["Result_Units"]],
			Reference_Range:                  row[stracIndex["Reference_Range"]],
			Date_Test_Performed:              row[stracIndex["Date_Test_Performed"]],
			Test_Result_Date:                 row[stracIndex["Test_Result_Date"]],
			Pt_Fname:                         row[stracIndex["Pt_Fname"]],
			Pt_Middle_Initial:                row[stracIndex["Pt_Middle_Initial"]],
			Pt_Lname:                         row[stracIndex["Pt_Lname"]],
			Date_of_Birth:                    row[stracIndex["Date_of_Birth"]],
			Patient_Age:                      row[stracIndex["Patient Age"]],
			Sex:                              row[stracIndex["Sex"]],
			Pt_Race:                          row[stracIndex["Pt_Race"]],
			Pt_Ethnicity:                     row[stracIndex["Pt_Ethnicity"]],
			Pt_Phone:                         row[stracIndex["Pt_Phone"]],
			Pt_Str:                           row[stracIndex["Pt_Str"]],
			Pt_City:                          row[stracIndex["Pt_City"]],
			Pt_ST:                            row[stracIndex["Pt_ST"]],
			Pt_Zip:                           row[stracIndex["Pt_Zip"]],
			Pt_County:                        row[stracIndex["Pt_County"]],
			Accession_Number:                 row[stracIndex["Accession_Number"]],
			Ordering_Facility:                row[stracIndex["Ordering_Facility"]],
			Ordering_Facility_Address:        row[stracIndex["Ordering_Facility_Address"]],
			Ordering_Facility_City:           row[stracIndex["Ordering_Facility_City"]],
			Ordering_Facility_State:          row[stracIndex["Ordering_Facility_State"]],
			Ordering_Facility_Zip:            row[stracIndex["Ordering_Facility_Zip"]],
			Ordering_Provider_Last_Name:      row[stracIndex["Ordering_Provider_Last_Name"]],
			Ordering_Provider_First_Name:     row[stracIndex["Ordering_Provider_First_Name"]],
			Ordering_Provider_NPI:            row[stracIndex["Ordering_Provider_NPI"]],
			Ordering_Provider_Street_Address: row[stracIndex["Ordering_Provider_Street_Address"]],
			Ordering_Provider_City:           row[stracIndex["Ordering_Provider_City"]],
			Ordering_Provider_State:          row[stracIndex["Ordering_Provider_State"]],
			Ordering_Provider_Zip:            row[stracIndex["Ordering_Provider_Zip"]],
			Ordering_Provider_Phone:          row[stracIndex["Ordering_Provider_Phone"]],
			Specimen_ID:                      row[stracIndex["Specimen_ID"]],
			Specimen_Type:                    row[stracIndex["Specimen_Type"]],
			Date_Test_Ordered:                row[stracIndex["Date_Test_Ordered"]],
			Date_Specimen_Collected:          row[stracIndex["Date_Specimen_Collected"]],
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

func ReadBom(r io.Reader) (io.Reader, error) {
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
