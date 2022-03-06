# CalendarCSVConverter
Converts an xls file into a googleCalendar CSV

## Usage

`csv2csv --range 8:52 --title A import.xls`

## Development

Run the development script

`go run csv2csv.go --range 8:52 --title A import.xls`

`go run csv2csv.go --row "A8:A52|title" --row "B8:A52|description" import.xls`

Build the binary 

`go build`
