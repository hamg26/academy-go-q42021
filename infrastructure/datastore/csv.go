package datastore

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

type myCSV interface {
	FindAll(string, int, int) (error, [][]string)
	Save([]string) error
}

// Holds the CSV file information (path)
type MyCSV struct {
	Filepath string
}

func closeFile(f *os.File) error {
	log.Println("Closing file")
	err := f.Close()
	if err != nil {
		log.Fatalln("Unable to close file", err)
		return err
	}
	return nil
}

/*
Return all the rows from the csv file specified in the config file (C.CSV.Path)
The first time it's called it reads the file content
- If items=-1 there is no limit for the number of items returned
- If itemsPerWorker=-1 there is no limit for the number of items returned
- Uses 1 worker by default
*/
func (mycsv *MyCSV) FindAll(filter string, items, itemsPerWorker int) (error, [][]string) {
	log.Println("Reading records", mycsv.Filepath, filter, items, itemsPerWorker)

	f, err := os.Open(mycsv.Filepath)
	if err != nil {
		return err, nil
	}

	jobs := make(chan []string)
	results := make(chan []string)

	wg := new(sync.WaitGroup)

	// Start up some workers
	numberOfWorkers := int(math.Ceil(float64(items) / float64(itemsPerWorker)))
	if numberOfWorkers <= 0 {
		numberOfWorkers = 1
		itemsPerWorker = items
	}
	for w := 1; w <= numberOfWorkers; w++ {
		wg.Add(1)
		go filterByIdType(items, itemsPerWorker, jobs, results, wg, filter)
	}

	// Enqueue all the rows of the file to be processed
	go func() {
		csvReader := csv.NewReader(f)
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			jobs <- record
		}

		close(jobs)

		if err := closeFile(f); err != nil {
			log.Fatal(err)
		}
	}()

	// Aggregate the results and close the result channel
	go func() {
		wg.Wait()
		close(results)
	}()

	allrows := make([][]string, 0)
	for r := range results {
		if len(r) != 0 && (items == -1 || len(allrows) < items) {
			allrows = append(allrows, r)
		}
	}

	return nil, allrows
}

func filterByIdType(items, itemsPerWorker int, rows <-chan []string, results chan<- []string, wg *sync.WaitGroup, filter string) {
	// Remove the worker from the wait-group once it finishes
	defer wg.Done()

	//Count the number of results found by this worker
	counter := 0

	for row := range rows {
		// Already processed all items for this worker
		if items != -1 && itemsPerWorker != -1 && counter >= itemsPerWorker {
			break
		}

		// the row is empty
		if len(row) == 0 {
			continue
		}

		id, _ := strconv.ParseUint(row[0], 10, 64)
		switch {
		case filter == "":
			fallthrough
		case filter == "even" && id%2 == 0:
			fallthrough
		case filter == "odd" && id%2 != 0:
			results <- row
			counter = counter + 1
		}
	}
}

/*
Saves a new record to the CSV file specified in the config file (C.CSV.Path)
Reads all the records first and just adds the new record, overwriting the entire file content
*/
func (mycsv *MyCSV) Save(record []string) error {
	log.Println("Saving record", record, mycsv.Filepath)

	err, records := mycsv.FindAll("", -1, -1)
	if err != nil {
		return err
	}
	records = append(records, record)

	f, err := os.Create(mycsv.Filepath)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	if w.WriteAll(records) == nil {
		w.Flush()
	} else {
		log.Fatalln("Error saving record", err)
	}

	if err := closeFile(f); err != nil {
		return err
	}

	return nil
}

/*
Returns a new instance of the CSV datastore
*/
func NewCSV(filePath string) *MyCSV {
	return &MyCSV{Filepath: filePath}
}
