package datafile

import (
	"bufio"
	"os"
)

func GetStrings(filename string)([]string, error) {

	var lines []string

	file, err := os.Open(filename)
    if err != nil{
        return nil, err
    }
    
	

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		if err != nil {
			return nil, err			
		}
		lines = append(lines, line)
    }

    err = file.Close()

    if err != nil{
        return nil, err
    }

    if scanner.Err() != nil{
        return nil, scanner.Err() 
    }

	return lines, nil
}