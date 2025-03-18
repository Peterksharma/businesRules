# Rules Engine Data Transformations

A powerful data transformation and rules processing engine that helps you clean, transform, and validate your data with configurable business rules.

## Features

### Data Cleaning and Transformation
- **Data Masking**: Intelligent masking of sensitive data with support for:
  - Emails (e.g., `****@example.com`)
  - Phone Numbers (last 4 digits preserved)
  - Dates (XXXX-XX-XX format)
  - Credit Cards (last 4 digits preserved)
  - IP Addresses (last octet preserved)
  - Numbers and Text (smart masking)

- **Deduplication**: Remove duplicate records based on specified columns
  - Configurable column selection
  - Preserves header rows
  - Generates detailed statistics

### Business Rules
- Configurable rule mappings
- Fuzzy matching capabilities
- Extensible transformation framework

## Project Structure

```
businessRules/
├── front/
│   └── ui/           # Next.js frontend application
├── back/
│   ├── mappings/     # Rule mapping configurations
│   └── transformations/
│       └── cleaning/
│           ├── dataMasking/    # Data masking functionality
│           └── deduplication/  # Deduplication tools
```

## Technology Stack

### Frontend
- Next.js 15.2
- React 19
- TypeScript
- Tailwind CSS
- Radix UI Components

### Backend
- Go (Golang)
- Redis (for caching and data storage)

## Getting Started

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd businessRules/front/ui
```

2. Install dependencies:
```bash
npm install
# or
yarn install
```

3. Start the development server:
```bash
npm run dev
# or
yarn dev
```

The frontend will be available at [http://localhost:3000](http://localhost:3000).

### Backend Setup

1. Ensure Go is installed on your system
2. Navigate to the backend directory:
```bash
cd businessRules/back
```

3. Install Go dependencies (if any)
4. Run the specific transformation tools as needed:

For data masking:
```bash
cd transformations/cleaning/dataMasking
go run masking.go <csv_file_path> <column_index>
```

For deduplication:
```bash
cd transformations/cleaning/deduplication
go run dedup.go <csv_file_path> <column_index>
```

## Usage Examples

### Data Masking

Mask a specific column:
```bash
go run masking.go data.csv 1
```

Mask all columns:
```bash
go run masking.go data.csv all
```

### Deduplication

Remove duplicates based on the first column:
```bash
go run dedup.go data.csv 0
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is proprietary and confidential. All rights reserved. 