# STRAC to State Mapping

This repo contains a standard process for defining a mapping between the Southwest Texas Regional Advisory Council (STRAC) point-of-care record structure to state-specific records to be processed locally.

## States

Don't see your state? Please contribute!

- [PA](./states/pa)

## Target implementations

The goal is to support an increasing number of workflows at the state level, including:

### Manual entry

Used by data entry workers who can map the STRAC format to the state-specific record structure. Each state will have instructions in the directory for the translation process.

### Command-line interface

Download the latest [release](https://github.com/chop-dbhi/strac/releases).

Provides machine translation of the STRAC format to the state-specific record structure.

For example, the following will take a STRAC dataset and convert it into the corresponding PA DOH record structure.

```
strac convert --state=PA strac_example.csv
```

During the conversion, it will also validate the data being produced to ensure values are correct and missing data is not present. These will be reported as warnings with the line and column names where they need to be corrected. This enables for human follow-up if required.

### Self-hosted Web interface

Operators can host this locally for users to enable uploading of the STRAC records and translate it to the state-specific record structure.

*Work in progress.*
