package main

import (
	"encoding/json"
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"errors"
	"strconv"
	"path/filepath"
	"github.com/go-redis/redis"
)

// DedupConfig holds the configuration for deduplication
type DedupConfig struct {
	ColumnIndex int    `json:"columnIndex"`
}

func removeDuplicates(data [][]string, columnIndex int) ([][]string, error) {
    if len(data) == 0 {
        return data, nil
    }
    
    // Validate column index
    if columnIndex < 0 || (len(data) > 0 && columnIndex >= len(data[0])) {
        return nil, fmt.Errorf("invalid column index: %d. Column index must be between 0 and %d", columnIndex, len(data[0])-1)
    }

    seen := make(map[string]bool)
    var unique [][]string

    // Add header row if exists
    if len(data) > 0 {
        unique = append(unique, data[0])
    }

    // Start from 1 to skip header row
    for i := 1; i < len(data); i++ {
        row := data[i]
        if len(row) <= columnIndex {
            return nil, fmt.Errorf("row %d has insufficient columns (needs at least %d columns)", i, columnIndex+1)
        }
        
        key := row[columnIndex]
        if !seen[key] {
            seen[key] = true
            unique = append(unique, row)
        }
    }
    return unique, nil
}

func processFile(filePath string, columnIndex int) error {
    // Verify file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return fmt.Errorf("file does not exist: %s", filePath)
    }

    // Verify file is CSV
    if ext := filepath.Ext(filePath); ext != ".csv" {
        return fmt.Errorf("file must be a CSV file, got: %s", ext)
    }

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

    uniqueData, err := removeDuplicates(data, columnIndex)
    if err != nil {
        return fmt.Errorf("error removing duplicates: %v", err)
    }

    // Create output filename
    dir := filepath.Dir(filePath)
    filename := filepath.Base(filePath)
    ext := filepath.Ext(filename)
    nameWithoutExt := filename[:len(filename)-len(ext)]
    outputPath := filepath.Join(dir, nameWithoutExt + "_deduped" + ext)

    // Write output to new CSV file
    outputFile, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    err = writer.WriteAll(uniqueData)
    if err != nil {
        return fmt.Errorf("error writing output: %v", err)
    }

    fmt.Printf("Successfully processed file. Output written to: %s\n", outputPath)
    fmt.Printf("Removed %d duplicate rows\n", len(data)-len(uniqueData))
    return nil
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run dedup.go <csv_file_path> <column_index>")
        fmt.Println("Example: go run dedup.go ./data.csv 0")
        fmt.Println("Note: column_index starts at 0")
        os.Exit(1)
    }

    filePath := os.Args[1]
    columnIndex, err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Fatalf("Invalid column index. Must be a number: %v", err)
    }

    if err := processFile(filePath, columnIndex); err != nil {
        log.Fatal(err)
    }
}

