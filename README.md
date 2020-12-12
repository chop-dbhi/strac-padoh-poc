# STRAC to State Mapping

This repo contains a standard process for defining a mapping between the Southwest Texas Regional Advisory Council (STRAC) point-of-care record structure to state-specific records to be processed locally.

## States

Don't see your state? Please contribute!

- [PA](./PA)

## Target implementations

The goal is to support an increasing number of workflows at the state level, including:

- Manual translation
  - Used by data entry workers who can map the STRAC format to the state-specific record structure.
- Command-line tool
  - Provides machine translation of the STRAC format to the state-specific record structure.
- Self-hosted Web interface
  - Operators can host this locally for users to enable uploading of the STRAC records and translate it to the state-specific record structure.

## STRAC Format

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
