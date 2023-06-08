package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

/*
1. compile dengan cara  go build di terminal saat sudah selesai ngoding
2. muncul file baru
3. jalankan file tersebut di terminal dengan cara ./namafile
4. compiler sifatnya permanen.
5. meskipun isi codinganya diganti, namun di compiler tidak merubah output yang dari terakhir
*/

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")

	//perulangan untuk membaca seluruh file
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			//mencetak seluruh nama file pada folder files
			fmt.Println(entry.Name())
			//membaca + mencetak seluruh value pada semua file pada folder files
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
