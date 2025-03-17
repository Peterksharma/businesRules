# CSV Deduplication Tool

This tool removes duplicate rows from CSV files based on a specified column index.

## Usage

The basic syntax is:
```bash
go run dedup.go <csv_file_path> <column_index>
```

Where:
- `<csv_file_path>`: Path to your input CSV file
- `<column_index>`: The zero-based index of the column to check for duplicates (0 for first column, 1 for second, etc.)

## Examples

Remove duplicates based on the first column:
```bash
go run dedup.go data.csv 0
```

Remove duplicates based on the second column:
```bash
go run dedup.go data.csv 1
```

Remove duplicates from a file in a different directory:
```bash
go run dedup.go /path/to/your/data.csv 2
```

## Output

The tool will:
1. Create a new file with "_deduped" suffix in the same directory as the input file
2. Preserve the header row
3. Remove any rows where the specified column has duplicate values
4. Print statistics about how many duplicates were removed

For example, if your input file is `data.csv`, the output will be `data_deduped.csv`

## Error Handling

The tool will show helpful error messages if:
- Wrong number of arguments is provided
- File doesn't exist
- File isn't a CSV
- Column index is invalid
- Column index is out of range
- Any row has insufficient columns

## Notes

- The tool preserves the first occurrence of each unique value in the specified column
- Column indexing starts at 0 (zero-based)
- The header row is always preserved
- The original file is not modified; a new file is created instead

## this will run on the first file
go run dedup.go

## this will 