package utils

import (
	"fmt"
	"log"
	"os"
)

func getPreferenceFile() string {

	home, err := os.UserHomeDir()

	if err != nil {
		return ""
	}

	return home + "/.smoovka"
}

func GetPreference(key string) string {

	file := getPreferenceFile()
	if file == "" {
		log.Fatal("Could not get preference file")
		return ""
	}

	// create file if it does not exist
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		_, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
			return ""
		}
	}

	// read file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get all preferences as a map
	preferences := make(map[string]string)
	for {
		var k string
		var v string
		_, err := fmt.Fscanf(f, "%s %s\n", &k, &v)
		if err != nil {
			log.Println(err)
			break
		}
		preferences[k] = v
	}

	log.Println("preferences", preferences)

	// return the value of the key
	return preferences[key]
}

func SetPreference(key string, value string) {

	file := getPreferenceFile()
	if file == "" {
		return
	}

	// create file if it does not exist
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		_, err := os.Create(file)
		if err != nil {
			return
		}
	}

	// read file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get all preferences as a map
	preferences := make(map[string]string)
	for {
		var k string
		var v string
		log.Println(f)
		_, err := fmt.Fscanf(f, "%s %s\n", &k, &v)
		if err != nil {
			break
		}
		preferences[k] = v
	}

	// set the value of the key
	preferences[key] = value

	// write all preferences to the file
	f, err = os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for k, v := range preferences {
		fmt.Fprintf(f, "%s %s\n", k, v)
	}

}
