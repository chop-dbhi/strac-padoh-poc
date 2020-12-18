package paphilly

import (
	"github.com/chop-dbhi/strac/converter"
)

var Columns = []*converter.Column{
	{
		Name:     "First Name",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientFirstName, nil
		},
	},
	{
		Name:     "Last Name",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientLastName, nil
		},
	},
	{
		Name:     "DOB",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: format date
			return r.PatientDOB, nil
		},
	},
	{
		Name: "Street",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientAddress, nil
		},
	},
	{
		Name: "City",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientCity, nil
		},
	},
	{
		Name: "State",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientState, nil
		},
	},
	{
		Name:     "Zip",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientZip, nil
		},
	},
	{
		Name: "Phone",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientCallbackNumber, nil
		},
	},
	{
		Name: "Email",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: STRAC field?
			return "", nil
		},
	},
	{
		Name: "Gender",
		Values: map[string]struct{}{
			"FEMALE":  {},
			"MALE":    {},
			"UNKNOWN": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientSex, nil
		},
	},
	{
		Name: "Race",
		Values: map[string]struct{}{
			"AFRICAN AMERICAN":                  {},
			"AMERICAN INDIAN OR ALASKAN NATIVE": {},
			"ASIAN":                             {},
			"NATIVE HAWAIIAN OR OTHER PACIFIC ISLANDER": {},
			"WHITE": {},
			"OTHER": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: map from STRAC options
			return r.PatientRace, nil
		},
	},
	{
		Name: "Ethnicity",
		Values: map[string]struct{}{
			"HISPANIC":     {},
			"NON-HISPANIC": {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientEthnicity, nil
		},
	},
	{
		Name: "Symptoms",
		Values: map[string]struct{}{
			"Yes":                                 {},
			"No":                                  {},
			"Priority acute respiratory symptoms": {},
			"Cough":                               {},
			"Difficulty Breathing":                {},
			"Shortness of Breath":                 {},
			"Fever":                               {},
		},
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: STRAC field?
			return "", nil
		},
	},
	{
		Name: "Ordering Facility",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return "PDPH Ambulatory Health", nil
		},
	},
	{
		Name: "Lab",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return "LABCORP", nil
		},
	},

	{
		Name: "Collection Date",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: correct column? parse date
			return r.PatientResults, nil
		},
	},
	{
		Name: "Result Date",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: correct column? parse date
			return r.PatientResults, nil
		},
	},

	{
		Name: "TestCode",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			// TODO: STRAC field?
			return "", nil
		},
	},
	{
		Name: "Result",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientPositive, nil
		},
	},
}
