package extractor

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

// TableData
// table of data extracted from resource
type TableData struct {
	RowHeaders []string
	ColHeaders []string
	Data       [][]string
}

// ExtractTableFromFile
// extract local file contents as TableData
func ExtractTableFromFile(fileName string) (TableData, error) {

	t := TableData{}

	// step 1 - open file to read
	fl, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		return TableData{}, err
	}
	defer fl.Close()

	// step 2 - read contents
	raw, err := readCsvContents(fl)
	if err != nil {
		return TableData{}, err
	}

	// step 3 - extract data
	t.Data = extractData(raw)
	t.RowHeaders, t.ColHeaders = extractHeaders(raw)

	log.WithFields(log.Fields{"cols:":len(t.ColHeaders),  "data:": strconv.Itoa(len(t.Data)) + "x" + strconv.Itoa(len(t.Data[0])),"rows:":len(t.RowHeaders) }).Info("extracted")

	return t, nil
}

// ExtractTableFromUrl
// downloads the url and extracts data and headers
func ExtractTableFromUrl(url string) (TableData, error){

	t := TableData{}

	// step 1 - download
	raw, err := DownloadLink(url)
	if err!= nil {
		log.WithFields(log.Fields{"err":err}).Error("cant download link")
		return TableData{}, nil
	}

	// step 2 - extract data and headers
	if raw != nil {
		t.Data = extractData(raw)
		t.RowHeaders, t.ColHeaders = extractHeaders(raw)
	}

	return t, nil
}

// GetDataHeadersFromUrl
// downloads the url and gets data headers
func GetDataHeadersFromUrl(url string) (rowHeaders, colHeaders []string){

	linkContents, err := DownloadLink(url)
	if err!= nil {
		log.WithFields(log.Fields{"err":err}).Error("cant download link")
		return nil, nil
	}

	rowHeaders, colHeaders = extractHeaders(linkContents)

	return rowHeaders, colHeaders
}

// readCsvContents
// read file contents
func readCsvContents(r io.ReadCloser) ([][]string, error ){

	reader := csv.NewReader(r)

	contents, err := reader.ReadAll()
	if err!=nil {
		log.WithFields(log.Fields{"err":err}).Error("cant read reader contents")
		return nil, nil
	}

	return contents, nil
}

// extractHeaders
// return headers ( first column and first row )
func extractHeaders(contents [][]string) (rowHeaders, colHeaders []string){

	// step 1 - extract col headers
	colHeaders = contents[0][1:]

	// step 2 - extract row headers
	numRowHeaders := len(contents) - 1
	rowHeaders = make([]string, numRowHeaders)
	for i:=0;i<numRowHeaders;i++ {
		rowHeaders[i] = contents[i+1][0]
	}

	return rowHeaders, colHeaders
}

// extractData
// extract only the data section, omitting headers
func extractData(contents [][]string) [][]string {


	numRows := len(contents) - 1
	numCols := len(contents[0])-1
	log.WithFields(log.Fields{"numRows:":numRows, "numCols:":numCols}).Info("")

	data := make([][]string, numRows)

	for i:=0;i<numRows;i++ {
		data[i] = make([]string, numCols)
		data[i] = contents[i+1][1:]
	}

	return data
}

