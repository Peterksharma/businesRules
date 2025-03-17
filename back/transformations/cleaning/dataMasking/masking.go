package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"regexp"
	"unicode"
)

// MaskingStrategy defines how different types of data should be masked
type MaskingStrategy struct {
	Value       string // Custom masking value (if provided)
	MaskAll     bool   // Whether to mask all columns
	ColumnIndex int    // Specific column to mask (-1 for all columns)
}

// DataTypePatterns contains regex patterns for different data types
var DataTypePatterns = map[string]*regexp.Regexp{
	"email":    regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	"phone":    regexp.MustCompile(`^\+?[\d\s-]{10,}$`),
	"date":     regexp.MustCompile(`^\d{4}[-/]\d{1,2}[-/]\d{1,2}$`),
	"number":   regexp.MustCompile(`^-?\d*\.?\d+$`),
	"creditCard": regexp.MustCompile(`^\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}$`),
	"ipAddress": regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`),
}

// maskValue applies appropriate masking based on the data type
func maskValue(value, customMask string) string {
	if customMask != "" {
		return customMask
	}

	// If empty value, return as is
	if value == "" {
		return value
	}

	// Detect data type and apply appropriate masking
	switch {
	case DataTypePatterns["email"].MatchString(value):
		parts := strings.Split(value, "@")
		return "****@" + parts[1]
	
	case DataTypePatterns["phone"].MatchString(value):
		return strings.Repeat("*", len(value)-4) + value[len(value)-4:]
	
	case DataTypePatterns["date"].MatchString(value):
		return "XXXX-XX-XX"
	
	case DataTypePatterns["creditCard"].MatchString(value):
		return "****-****-****-" + strings.Replace(value[len(value)-4:], " ", "", -1)
	
	case DataTypePatterns["ipAddress"].MatchString(value):
		parts := strings.Split(value, ".")
		return fmt.Sprintf("***.***.***.%s", parts[3])
	
	case DataTypePatterns["number"].MatchString(value):
		return strings.Repeat("#", len(value))
	
	default:
		// For text, preserve first and last character
		if len(value) <= 2 {
			return strings.Repeat("*", len(value))
		}
		return value[0:1] + strings.Repeat("*", len(value)-2) + value[len(value)-1:]
	}
}

func processFile(filePath string, strategy MaskingStrategy) error {
	// Verify file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Verify file is CSV
	if ext := filepath.Ext(filePath); ext != ".csv" {
		return fmt.Errorf("file must be a CSV file, got: %s", ext)
	}

	// Open and read file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %v", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("empty CSV file")
	}

	// Validate column index if specified
	if !strategy.MaskAll && strategy.ColumnIndex >= len(data[0]) {
		return fmt.Errorf("invalid column index: %d. Must be between 0 and %d", 
			strategy.ColumnIndex, len(data[0])-1)
	}

	// Create masked data
	maskedData := make([][]string, len(data))
	maskedData[0] = data[0] // Preserve header row

	// Process each row
	for i := 1; i < len(data); i++ {
		row := make([]string, len(data[i]))
		for j := range data[i] {
			if strategy.MaskAll || j == strategy.ColumnIndex {
				row[j] = maskValue(data[i][j], strategy.Value)
			} else {
				row[j] = data[i][j]
			}
		}
		maskedData[i] = row
	}

	// Create output filename
	dir := filepath.Dir(filePath)
	filename := filepath.Base(filePath)
	ext := filepath.Ext(filename)
	nameWithoutExt := filename[:len(filename)-len(ext)]
	outputPath := filepath.Join(dir, nameWithoutExt + "_masked" + ext)

	// Write output
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	err = writer.WriteAll(maskedData)
	if err != nil {
		return fmt.Errorf("error writing output: %v", err)
	}

	fmt.Printf("Successfully masked data. Output written to: %s\n", outputPath)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("Mask specific column: go run masking.go <csv_file_path> <column_index> [mask_value]")
		fmt.Println("Mask all columns: go run masking.go <csv_file_path> all [mask_value]")
		fmt.Println("\nExamples:")
		fmt.Println("go run masking.go data.csv 1")
		fmt.Println("go run masking.go data.csv 2 XXXXX")
		fmt.Println("go run masking.go data.csv all")
		fmt.Println("go run masking.go data.csv all ****")
		os.Exit(1)
	}

	filePath := os.Args[1]
	strategy := MaskingStrategy{}

	// Parse arguments
	if len(os.Args) > 2 {
		if strings.ToLower(os.Args[2]) == "all" {
			strategy.MaskAll = true
			strategy.ColumnIndex = -1
		} else {
			columnIndex, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("Invalid column index. Must be a number or 'all': %v", err)
			}
			strategy.ColumnIndex = columnIndex
		}
	}

	// Check for custom mask value
	if len(os.Args) > 3 {
		strategy.Value = os.Args[3]
	}

	if err := processFile(filePath, strategy); err != nil {
		log.Fatal(err)
	}
}
