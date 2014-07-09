package main

import "fmt"
import "strings"
import "bufio"
// import "log"
import "os"
import "encoding/csv"

func main() {
	// fmt.Println(os.Args)
	sourcefile_add := os.Args[1]
	file, error := os.Create("output.csv")

    if error != nil {panic(error)}
    defer file.Close()

    // New Csv writer
    writer := csv.NewWriter(file)

    // Headers
    var new_headers = []string { "SSN", "FIRSTNAME", "MIDDLENAME", "LASTNAME", "SUFFIX","BDAY","BMONTH","BYEAR", "DDAY", "DMONTH", "DYEAR", "RAW" }        
    returnError := writer.Write(new_headers)
    if returnError != nil {
        fmt.Println(returnError)
    }
    // sourcefile, _ := os.Open("/home/ec2-user/go/src/github.com/roberthan/va_process_ssdm/test_data.txt")
    sourcefile, _ := os.Open(sourcefile_add)
    // sourcefile, _ := os.Open("/Users/roberthan/Documents/go/src/github.com/roberthan/va_process_ssdm/test_data.txt")
	scanner := bufio.NewScanner(sourcefile)
	count := 15
	var slice [15]string
    j := 0
	for scanner.Scan() {

		slice[j] = scanner.Text()
		j++
	    if count <= j{
			for i := range slice {
				values := parse_ssdm(slice[i])
				returnError := writer.Write(values)
	    		if returnError != nil {
	        		fmt.Println(returnError)
	    		}
        		i++
    		}
    		writer.Flush() 
			j = 0
	    }
	}
	
}

func parse_ssdm(str string) []string{
	lastName := strings.TrimSpace(string(str[10:30]))
	ssn := strings.TrimSpace(string(str[1:10]))
	nameSuffix := strings.TrimSpace(string(str[30:34]))
	firstName := strings.TrimSpace(string(str[34:49]))
	middleName := strings.TrimSpace(string(str[49:64]))
	dod := strings.TrimSpace(string(str[65:73]))
	dob := strings.TrimSpace(string(str[73:81]))
	// fmt.Println(middleName + ":"+dod)
	values := []string {ssn, firstName, middleName, lastName, nameSuffix, dob[0:2], dob[2:4], dob[4:8], dod[0:2], dod[2:4], dod[4:8]}
	return values
}