package main

import (
	"fmt"
	"os"

	"github.com/chop-dbhi/strac-state-mapping/converter"
)

var columns = []*converter.Column{
	{
		Name:     "PatientFirstName",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Fname, nil
		},
	},
	{
		Name: "PatientMiddleInitial",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Middle_Initial, nil
		},
	},
	{
		Name:     "PatientLastName",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Lname, nil
		},
	},
	{
		Name: "PatientSuffix",
	},
	{
		Name:     "PatientDOB",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Date_of_Birth, nil
		},
	},
	{
		Name: "PatientAddress1",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Str, nil
		},
	},
	{
		Name: "PatientCity",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_City, nil
		},
	},
	{
		Name: "PatientState",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_ST, nil
		},
	},
	{
		Name:     "PatientZipCode",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Zip, nil
		},
	},
	{
		Name: "PatientPhoneNumber",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Phone, nil
		},
	},
	{
		Name: "PatientGender",
		Values: map[string]struct{}{
			"Female":  {},
			"Male":    {},
			"Unknown": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Sex, nil
		},
	},
	{
		Name: "PatientRace",
		Values: map[string]struct{}{
			"Asian":            {},
			"Black":            {},
			"Native America":   {},
			"Other":            {},
			"Pacific Islander": {},
			"Unknown":          {},
			"White":            {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Race, nil
		},
	},
	{
		Name: "PatientEthnicity",
		Values: map[string]struct{}{
			"Hispanic":     {},
			"Non-Hispanic": {},
			"Unkown":       {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Ethnicity, nil
		},
	},
	{
		Name: "TestID",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Specimen_ID, nil
		},
	},
	{
		Name: "SpecimenCollectedDate",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Date_Specimen_Collected, nil
		},
	},
	{
		Name: "SpecimenSource",
		Values: map[string]struct{}{
			"None":    {},
			"NP swab": {},
			"Saliva":  {},
			"Throat":  {},
			"Unknown": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Specimen_Type, nil
		},
	},
	{
		Name: "TestName",
		Values: map[string]struct{}{
			"COVID-19 ANTIGEN test - Point-of-care": {},
			"COVID-19 PCR test - Point-of-care":     {},
			"Influenza A ANTIGEN (positives only)":  {},
			"Influenza A PCR (positives only)":      {},
			"Influenza B ANTIGEN (positives only)":  {},
			"Influenza B PCR (positives only)":      {},
			"RSV ANTIGEN (positives only)":          {},
			"RSV PCR (positives only)":              {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Pt_Middle_Initial, nil
		},
	},
	{
		Name: "TestQualitativeResult",
		Values: map[string]struct{}{
			"Detected":     {},
			"Not Detected": {},
			"Inconclusive": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Result, nil
		},
	},
	{
		Name: "Notes",
	},
	{
		Name: "PerformingFacilityName",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Performing_Organization_Name, nil
		},
	},
}

func main() {
	if err := converter.Main(columns); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
