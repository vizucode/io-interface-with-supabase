package main

import (
	"encoding/json"
	"io"
	supastorage "iowithsupabase/repositories/supa_storage"
	"log"
	"math/rand"
	"os"
)

type SuperPower struct{}

/*
Cast will create new random file then will write into io.WriteCloser
*/
func (w *SuperPower) Cast(f io.WriteCloser) (resp error) {

	words := []string{"lorem ipsum", "dolor sit amet", "consectetur adipiscing elit"}
	randomize := words[rand.Intn(len(words))]

	data, err := json.Marshal(randomize)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		log.Println(err)
		return err
	}

	// To start uploading data to the write close.
	// This case is supabase needs close() for start uploading.
	defer f.Close()

	return
}

/*
Will Download file from io.ReadCloser data and send to temp folder
*/
func (w *SuperPower) Observe(f io.ReadCloser) (resp error) {

	path, err := os.MkdirTemp(".", "temp")
	if err != nil {
		log.Println(err)
		return err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return
}

func main() {

	// initiate supabase object
	supaCli := supastorage.NewSupaClient("https://qytwpijdbksagnkuqgay.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InF5dHdwaWpkYmtzYWdua3VxZ2F5Iiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTcxNTM2MTg0MiwiZXhwIjoyMDMwOTM3ODQyfQ.idV0yiC5VvRSBmwmVNDCpXcZti-XAlnAUTHvqRxHb98", "testing")

	Wizard := SuperPower{}

	// attach Writer() to wizard.Cast() with io.WriteCloser interface it's can be implemented.
	err := Wizard.Cast(supaCli.Writer("wizard_spell.txt"))
	if err != nil {
		panic(err)
	}

	// attach Reader() to wizard.Observe() with io.ReadCloser interface it's can be implemented.
	err = Wizard.Observe(supaCli.Reader("wizard_spell.txt"))
	if err != nil {
		panic(err)
	}

}
