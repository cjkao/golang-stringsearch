package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	// "runtime/pprof"

	_ "net/http/pprof"
	"time"

	data "github.com/kaopeter/search/data"
	imp "github.com/kaopeter/search/imp"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var keyWord = flag.String("key", "高...", "key word")

func main() {
	flag.Parse()

	// http.ListenAndServe("localhost:6060", nil)

	start := time.Now()
	employees := data.GenerateEmployee(80000)
	elapsed := time.Since(start)
	log.Printf("gen took %s", elapsed)
	for _, employee := range employees {
		fmt.Println(employee)
		break
	}
	pattern := *keyWord
	newString := data.BigHungString(employees)
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 1000; i++ {
		newStart := time.Now()
		indexs := imp.SearchBigHunk(newString, pattern, 200)
		logTime(newStart, "BigHunk took ", indexs)

		newStart = time.Now()
		indexs = imp.SearchByStruct(employees, pattern)
		logTime(newStart, "Struct String took ", indexs)

		newStart = time.Now()
		indexs = imp.SearchByStructRegex(employees, pattern)
		logTime(newStart, "Struct Regex took ", indexs)
	}

}

var isPrint = false

func logTime(start time.Time, name string, indexs []int) {
	if isPrint {
		elapsed := time.Since(start)
		log.Printf("%s took %s %v", name, elapsed, indexs)
		isPrint = true
	}

}
