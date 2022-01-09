package imp

import (
	"regexp"
	"strings"

	"github.com/kaopeter/search/data"
)

//search string on big string that return index mapping to employees
func SearchBigHunk(bigHunk, pattern string, recLens int) []int {
	// structLens := 200
	var result []int = []int{}
	idx := 0
	for {

		if len(result) > 0 {
			idx = result[len(result)-1]*recLens + recLens
		}

		tidx := strings.Index((bigHunk)[idx:], pattern)
		if tidx > 0 {
			base := 0
			if len(result) > 0 {
				base = result[len(result)-1] + 1
			}
			tidx = tidx/recLens + base
			result = append(result, tidx)
			if len(result) > 5 {
				break
			}
		} else {
			break
		}
	}
	return result
}

//string matching on employees
func SearchByStruct(employees []data.Employee, searchString string) []int {
	var result []int = []int{}
	for i, employee := range employees {
		contain := strings.Contains(employee.Account, searchString) ||
			strings.Contains(employee.Address, searchString) ||
			strings.Contains(employee.Chinesename, searchString) ||
			strings.Contains(employee.DepartmentEng, searchString) ||
			strings.Contains(employee.DeptID, searchString) ||
			strings.Contains(employee.Englishname, searchString) ||
			strings.Contains(employee.Phone, searchString)
		if contain {
			result = append(result, i)
		}
		if len(result) > 5 {
			// fmt.Print(employees[result[0]].Chinesename)
			break
		}
	}
	return result
}

//string matching on employees
func SearchByStructRegex(employees []data.Employee, pat string) []int {
	re := regexp.MustCompile(pat)
	var result []int = []int{}
	for i, employee := range employees {
		match := len(re.FindStringIndex(employee.Account)) > 0 ||
			len(re.FindStringIndex(employee.Address)) > 0 ||
			len(re.FindStringIndex(employee.Chinesename)) > 0 ||
			re.MatchString(employee.DepartmentEng) ||
			re.MatchString(employee.DeptID) ||
			re.MatchString(employee.Englishname) ||
			re.MatchString(employee.Phone)
		if match {
			result = append(result, i)
		}
		if len(result) > 5 {
			break
		}
	}
	return result
}
