# CSV Data Masking Tool

This tool provides intelligent data masking for CSV files, with support for various data types and masking strategies.

## Features

### Smart Data Type Detection and Masking

The tool automatically detects and appropriately masks different data types:

- **Emails**: `john.doe@example.com` → `****@example.com`
- **Phone Numbers**: `+1-234-567-8901` → `**********8901`
- **Dates**: `2024-03-15` → `XXXX-XX-XX`
- **Credit Cards**: `4111-1111-1111-1111` → `****-****-****-1111`
- **IP Addresses**: `192.168.1.100` → `***.***.***.100`
- **Numbers**: `12345` → `#####`
- **Text**: `Hello` → `H***o`

### Flexible Masking Options

- Mask specific columns by index
- Mask all columns at once
- Use custom masking values
- Preserve header row

## Usage

Basic syntax:
```bash
# Mask specific column
go run masking.go <csv_file_path> <column_index> [mask_value]

# Mask all columns
go run masking.go <csv_file_path> all [mask_value]
```

### Examples

1. Mask a specific column (using smart masking):
```bash
go run masking.go data.csv 1
```

2. Mask a column with custom value:
```bash
go run masking.go data.csv 2 XXXXX
```

3. Mask all columns:
```bash
go run masking.go data.csv all
```

4. Mask all columns with custom value:
```bash
go run masking.go data.csv all ****
```

## Output

- Creates a new file with "_masked" suffix in the same directory
- Example: `data.csv` → `data_masked.csv`
- Original file remains unchanged

## Error Handling

The tool provides clear error messages for:
- Missing or invalid files
- Invalid column indices
- Non-CSV files
- Empty files
- Insufficient columns

## Notes

- Column indexing starts at 0 (zero-based)
- Header row is always preserved
- Empty values remain empty
- Original file is not modified 