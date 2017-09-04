package grades

import (
	"testing"
	"os"
	"log"
	"bufio"
)

func SetupFile(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err)
	}

	return file
}

func TestCan_not_read_date_from_empty_file(t *testing.T) {
	file := SetupFile("files/grades_empty.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	date, _ := ReadString(line, 40, 7)
	if date != "" {
		t.Error("Expected empty date, got", date)
	}
}

func TestCan_not_read_grade_from_empty_file(t *testing.T) {
	file := SetupFile("files/grades_empty.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	date, _ := ReadString(line, 87, 1)
	if date != "" {
		t.Error("Expected empty grade, got", date)
	}
}

func TestCan_read_date_from_file(t *testing.T) {
	file := SetupFile("files/grades_ymd.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	date, _ := ReadString(line, 40, 7)
	if date != "YYYYMMDD" {
		t.Error("Expected YYYYMMDD, got", date)
	}
}

func TestCan_parse_and_validate_invalid_date(t *testing.T) {
	file := SetupFile("files/grades_invalid_date.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	date, _ := ReadString(line, 40, 7)
	valid := ValidateDate(date)
	if valid {
		t.Error("Expected false, got", valid)
	}
}

func TestCan_parse_and_validate_valid_date(t *testing.T) {
	file := SetupFile("files/grades_valid_date.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	date, _ := ReadString(line, 40, 7)
	valid := ValidateDate(date)
	if !valid {
		t.Error("Expected true, got", valid)
	}
}

func TestCan_read_grade_from_file(t *testing.T) {
	file := SetupFile("files/grades_ymd.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	grade, _ := ReadString(line, 87, 1)
	if grade == "" {
		t.Error("Expected a number, got", grade)
	}
}

func TestCan_read_grade_from_file_and_validate_invalid_danish_grade(t *testing.T) {
	file := SetupFile("files/grades_invalid_grade.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	grade, _ := ReadString(line, 87, 1)
	validDanishGrade := ValidateGrade(grade)
	if validDanishGrade {
		t.Error("Expected false got ", validDanishGrade)
	}
}

func TestCan_read_grade_from_file_and_validate_valid_danish_grade(t *testing.T) {
	file := SetupFile("files/grades_valid_grade.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	grade, _ := ReadString(line, 87, 1)
	validDanishGrade := ValidateGrade(grade)
	if !validDanishGrade {
		t.Error("Expected true got ", validDanishGrade)
	}
}

func TestCan_read_grade_from_file_and_validate_valid_danish_grade_and_convert_to_ects(t *testing.T) {
	file := SetupFile("files/grades_valid_grade_12.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	grade, _ := ReadString(line, 87, 1)
	validDanishGrade := ValidateGrade(grade)
	if !validDanishGrade {
		t.Error("Expected true got ", validDanishGrade)
	}
	ectsGrade := ConvertToECTS(grade)
	expectedECTS := "A"
	if ectsGrade != expectedECTS {
		t.Error("Expected grade "+expectedECTS+" got", ectsGrade)
	}
}