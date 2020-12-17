package pa

import (
	"github.com/chop-dbhi/strac/converter"
)

var Columns = []*converter.Column{
	{
		Name:     "PatientFirstName",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientFirstName, nil
		},
	},
	{
		Name: "PatientMiddleInitial",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return "", nil
		},
	},
	{
		Name:     "PatientLastName",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientLastName, nil
		},
	},
	{
		Name: "PatientSuffix",
	},
	{
		Name:     "PatientDOB",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientDOB, nil
		},
	},
	{
		Name: "PatientAddress1",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientAddress, nil
		},
	},
	{
		Name: "PatientCity",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientCity, nil
		},
	},
	{
		Name: "PatientState",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientState, nil
		},
	},
	{
		Name:     "PatientZipCode",
		Required: true,
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientZip, nil
		},
	},
	{
		Name: "PatientPhoneNumber",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientCallbackNumber, nil
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
			return r.PatientSex, nil
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
			return r.PatientRace, nil
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
			return r.PatientEthnicity, nil
		},
	},
	{
		Name: "TestID",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.ID, nil
		},
	},
	{
		Name: "SpecimenCollectedDate",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientResults, nil
		},
	},
	{
		Name: "SpecimenSource",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return "nasal swab", nil
		},
	},
	{
		Name: "TestName",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return "SARS-CoV-2 (COVID-19) Ag [Presence] in Respiratory specimen by Rapid immunoassay", nil
		},
	},
	{
		Name: "TestQualitativeResult",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.PatientPositive, nil
		},
	},
	{
		Name: "Notes",
	},
	{
		Name: "PerformingFacilityName",
		Mapper: func(r *converter.STRACRecord) (string, error) {
			return r.Reason, nil
		},
	},
}
