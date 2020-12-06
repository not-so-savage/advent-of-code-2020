package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var requiredFields = []field{
	field{k: "byr", regex: `^(19[2-8][0-9]|199[0-9]|200[0-2])$`},
	field{k: "iyr", regex: `^(201[0-9]|2020)$`},
	field{k: "eyr", regex: `^(202[0-9]|2030)$`},
	field{k: "hgt", regex: `^(((15[0-9]|1[6-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in))$`},
	field{k: "hcl", regex: `^#[0-9a-f]{6}$`},
	field{k: "ecl", regex: `^((amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth))$`},
	field{k: "pid", regex: `^\d{9}$`},
}

type passport struct {
	fields map[string]string
}

type field struct {
	k     string
	regex string
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(inputBytes)

	passports := extractPassports(input)

	invalidPassportsPart1 := countInvalidPassports(passports)
	invalidPassportsPart2 := countInvalidPassportsUsingValidation(passports)

	fmt.Println("Part 1: ", len(passports)-invalidPassportsPart1)
	fmt.Println("Part 2: ", len(passports)-invalidPassportsPart2)
}

func countInvalidPassports(passports []passport) int {
	invalidPassports := 0

	for i := range passports {
		for _, field := range requiredFields {
			if _, ok := passports[i].fields[field.k]; !ok {
				invalidPassports++
				break
			}
		}
	}

	return invalidPassports
}

func countInvalidPassportsUsingValidation(passports []passport) int {
	invalidPassports := 0

	for _, passport := range passports {
		for _, requiredField := range requiredFields {
			if _, ok := passport.fields[requiredField.k]; !ok {
				invalidPassports++
				break
			}
			matches, _ := regexp.MatchString(requiredField.regex, passport.fields[requiredField.k])
			fmt.Println(passport.fields[requiredField.k], " matches? ", matches)
			if !matches {
				invalidPassports++
				break
			}
		}
	}

	return invalidPassports
}

func extractPassports(passportBatch string) []passport {
	passportStrings := strings.Split(passportBatch, "\n\n")

	passports := make([]passport, len(passportStrings))
	for i := range passports {
		passports[i] = passport{fields: extractFields(passportStrings[i])}
	}

	return passports
}

func extractFields(fieldsString string) map[string]string {
	fieldStrings := strings.FieldsFunc(fieldsString, func(r rune) bool {
		return r == ' ' || r == '\n'
	})

	fields := make(map[string]string)

	for i := range fieldStrings {
		fieldSlice := strings.Split(fieldStrings[i], ":")
		k := fieldSlice[0]
		v := fieldSlice[1]
		fields[k] = v
	}

	return fields
}
