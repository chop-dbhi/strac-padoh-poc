# STRAC-PADOH POC Translator

Southwest Texas Regional Advisory Council (STRAC) COVID-19 point-of-care (POC) test template to the PA department of health (DOH) point-of-care (POC) test translator.

## Motivation

Given POC of test records in the STRAC format, translate this file structure into the PA DOH record structure.

## Target implementations

- [ ] Command-line tool
  - Supported on Windows, macOS, and Linux
  - Users can use anywhere in their workflow
- [ ] Web interface
  - Use cURL or other HTTP client to upload and translate
  - Simple form interface to graphically upload and download
  - Operators can host locally for users

## Source File

### Headers

- Reporting_Facility_Name
- CLIA_Number
- Performing_Organization_Name
- Performing_Organization_Address
- Performing_Organization_City
- Performing_Organization_Zip
- Performing_Organization_State
- Device_Identifier
- Ordered_Test_Name
- LOINC_Code
- LOINC_Text
- Result
- Result_Units
- Reference_Range
- Date_Test_Performed
- Test_Result_Date
- Pt_Fname
- Pt_Middle_Initial
- Pt_Lname
- Date_of_Birth
- "Patient Age"
- Sex
- Pt_Race
- Pt_Ethnicity
- Pt_Phone
- Pt_Str
- Pt_City
- Pt_ST
- Pt_Zip
- Pt_County
- Accession_Number
- Ordering_Facility
- Ordering_Facility_Address
- Ordering_Facility_City
- Ordering_Facility_State
- Ordering_Facility_Zip
- Ordering_Provider_Last_Name
- Ordering_Provider_First_Name
- Ordering_Provider_NPI
- Ordering_Provider_Street_Address
- Ordering_Provider_City
- Ordering_Provider_State
- Ordering_Provider_Zip
- Ordering_Provider_Phone
- Specimen_ID
- Specimen_Type
- Date_Test_Ordered
- Date_Specimen_Collected

## Target File

### Headers

- PatientFirstName
- PatientMiddleInitial
- PatientLastName
- PatientSuffix
- PatientDOB
- PatientAddress1
- PatientCity
- PatientState
- PatientZipCode
- PatientPhoneNumber
- PatientGender
- PatientRace
- PatientEthnicity
- TestID
- SpecimenCollectedDate
- SpecimenSource
- TestName
- TestQualitativeResult
- Notes
- PerformingFacilityName
