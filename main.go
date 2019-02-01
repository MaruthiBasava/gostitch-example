package main

var db = InitializeDb()

func main() {

	PopulateFromCSV("datasets/akc_breed.csv")
	presetup(db, true)

}
