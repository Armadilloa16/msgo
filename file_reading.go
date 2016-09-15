package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	
	dat, err := ioutil.ReadFile("examples/hpc/fid/0_A20/1/1SRef/acqu")
	check(err)
	// fmt.Print(string(dat))
	
	// Gets all args except multi-line args into an awkward [][]string format
	re := regexp.MustCompile("\n##\\$?([^\\$\n\r]*) *= *<?([^>\n\r]*)>?")
	args := re.FindAllStringSubmatch(string(dat), -1)
	
	// fmt.Println(args)

	fmt.Println("\nn args detected:")
	fmt.Println(len(args))
	
	i := 235
	fmt.Println("\nargs[i], i:", i)
	fmt.Println(args[i])
	fmt.Println("\nargs[i][0]")
	fmt.Println(args[i][0])
	fmt.Println("\nargs[i][1]")
	fmt.Println(args[i][1])
	fmt.Println("\nargs[i][2]")
	fmt.Println(args[i][2])
	
}