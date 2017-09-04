package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"io"

	"github.com/steffen25/convert-danish-grades-to-ects/grades"
)

func main()  {
	file, err := os.Open("grades/files/grades.txt")
	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		fmt.Printf(" > Read %d characters\n", len(line))
		date, err := grades.ReadString(line, 40, 7)
		if err != nil {
			log.Println(err)
			break
		}
		validDate := grades.ValidateDate(date)
		if err != nil {
			fmt.Printf("could not parse date err %v", err)
		}
		fmt.Printf(" > Read date %s \n", date)
		fmt.Printf(" > Is valid date? %v \n", validDate)
		grade, err := grades.ReadString(line, 87, 1)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf(" > Read grade %s \n", grade)
		if err != nil {
			log.Println(err)
			break
		}
	}

	if err == io.EOF {
		fmt.Println(" > End of file")
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
}