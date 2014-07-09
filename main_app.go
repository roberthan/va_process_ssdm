package main

import "fmt"
import "bufio"
import "log"
import "os"
import (
	_ "github.com/lib/pq"
	"database/sql"
)


func main() {
	// db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full") ?sslmode=verify-full
	db, err := sql.Open("postgres", "postgres://dbuser:vitaepassword@ec2-54-183-9-10.us-west-1.compute.amazonaws.com/data_warehouse")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from ssdm_raw;")
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}

    file, _ := os.Open("/Users/roberthan/Documents/go/src/github.com/roberthan/va_process_ssdm/test_data.txt")
	scanner := bufio.NewScanner(file)
	count := 15
	var slice [15]string
    var j int
    j = 0
	for scanner.Scan() {
		// fmt.Println(j)
		slice[j] = scanner.Text()
		j++
	    if count <= j{
			for i := range slice {
				// byteArray := []byte(slice[i])
    			parse_ssdm(slice[i])
    	// 		if old_ssn > ssn{
					// fmt.Println(j)
    	// 		}
        		i++
    		}
			j = 0
	    }
	    
	}
	if err := scanner.Err(); err != nil {
    	log.Fatal(err)
	}
}

func parse_ssdm(str string) string{
	// lastName := string(str[10:30])
	ssn := string(str[1:10])
	// nameSuffix := string(str[30:33])
	// firstName := string(str[34:49])
	middleName := string(str[49:64])
	dod := string(str[65:73])
	// dob := string(str[73:81])
	fmt.Println(middleName + ":"+dod)
	return ssn
}