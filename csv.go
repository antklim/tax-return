package taxreturn

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// TODO: add headers constants
// TODO: add record map function
// TODO: add test
// TODO: add reports

// ReadCsv reads CSV file and maps it to Bills.
func ReadCsv(file *os.File) (Bills, error) {
	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Header: %v\n", header)

	records, err := reader.ReadAll()
	fmt.Println(records)

	bills := make(Bills, len(records))
	for _, record := range records {
		period, err := NewBillPeriod(record[0], record[1])
		if err != nil {
			return nil, err
		}

		paid, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			return nil, err
		}

		bill := Bill{
			Period: period,
			Paid:   float32(paid),
		}
		bills = append(bills, bill)
	}

	return bills, nil
}
