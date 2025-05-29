package main

import "golang.org/x/text/language"

func main() {
	ctr, err := language.ParseRegion("")
	if err != nil {
		println("Error parsing region: ", err)
		return
	}

	println(ctr.IsCountry())
	println(ctr.IsGroup())
	println(ctr.ISO3())
	println(ctr.String())
}
