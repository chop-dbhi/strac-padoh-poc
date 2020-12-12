# PA

Translation from STRAC to the PA department of health (DOH) point-of-care (POC) template.

## Implementations

### Manual entry

Open the [XLSX](./padoh_covid19_poc_template.xlsx) and follow the instructions to translate the data in the STRAC format to the PA-DOH format.

### Command-line tool

This provides an automated method to convert STRAC records into the PA-DOH record format.

```
strac-pa-converter < strac_data.csv > padoh_data.csv
```

Note, this does not preclude the ability to edit or correct the translate the converted data. Any validation errors such as missing fields, invalid controlled values will print warning messages in order for a user to review.

Download a release. Windows, macOS, and Linux are supported.

### Web API/interface

This builds utilizes the same logic as the command-line tool, but is presented as a simple Web API or interface to upload a CSV file.

## Data Dictionary

*Corresponds to version 5 of the template as of 2020/12/10.*

PA DOH | Required | Example | STRAC Field
-------|----------|---------|-------
PatientFirstName | yes | Jane Pt_Fname
PatientMiddleInitial | no | M | Pt_Middle_Initial
PatientLastName | yes | Doe | Pt_Lname
PatientSuffix | no | |
PatientDOB | yes | 11/20/1997 | Date_of_Birth
PatientAddress1 | no | 123 Main St | Pt_Str
PatientCity |  no | Atown | Pt_City
PatientState | no | PA | Pt_ST
PatientZipCode | yes | 19104 | Pt_Zip
PatientPhoneNumber | no | 215-555-5555 | Pt_Phone
PatientGender |  yes | Female | Pt_Sex
PatientRace | no | Asian | Pt_Race
PatientEthnicity | no | | Pt_Ethnicity
TestID | yes | 20RH-196-01554 |
SpecimenCollectedDate | yes | 12/10/2020 |
SpecimenSource | no | Saliva |
TestName | yes | COVID-19 ANTIGEN test - Point-of-care | Ordered_Test_Name
TestQualitativeResult | yes | [ Result, Result_Units ]
Notes | no | [ Reference_Range
PerformingFacilityName | no | Performing_Organization_Name

### Controlled Values

- TestName
  - COVID-19 ANTIGEN test - Point-of-care
  - COVID-19 PCR test - Point-of-care
  - Influenza A ANTIGEN (positives only)
  - Influenza A PCR (positives only)
  - Influenza B ANTIGEN (positives only)
  - Influenza B PCR (positives only)
  - RSV ANTIGEN (positives only)
  - RSV PCR (positives only)
- TestQualitativeResult
  - Detected
  - Not Detected
  - Inconclusive
- SpecimenSource
  - None
  - NP swab
  - Saliva
  - Throat
  - Unknown
- PatienGender
  - Female
  - Male
  - Unknown
- PatientRace
  - Asian
  - Black
  - Native America
  - Other
  - Pacific Islander
  - Unknown
  - White
- PatientEthnicity
  - Hispanic
  - Non-Hispanic
  - Unkown
