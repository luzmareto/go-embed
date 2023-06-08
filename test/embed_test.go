package test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

/*
ALUR :
1. bikin file version txt yang isinya 1.0.0-SNAPSHOT
2. buat file baru dan Import _ "embed"
3. deklarasi var version di luar func
4. tepat diatas deklarasi, masukan //go:embed version.txt
*/

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)  //1.0.0-SNAPSHOT
	fmt.Println(version2) //1.0.0-SNAPSHOT
}

/*
ALUR LOAD GAMBAR / BINARY TEXT:
1. ctrl + shift + e
2. drag gambar dari file explorer ke vs code
3. maka akan muncul file baru yang isinya gambar
4. deklarasi var slice of byet di luar func
5. bikin func untuk load gambar tersebut
6. maka akan muncul file baru
*/
//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {

	//load gambar
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
ALUR MULTIPLES FILES
1. bikin folder di vsc,misalkan namanya files
2. bikin beberapa file pada folder tersebut
3. deklarasikan nama folders di luar func. contoh : var files embed.FS
4. masukan embed sebanyak file di folders tersebut
5. bikin func untuk load embed
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS //syntax masuk ke folders

func TestMultipleFiles(t *testing.T) {
	//ignore error
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

}

/*
ALUR PATHMATCHER :
1. gunakan syntax *.txt pada deklarasi path
2. artinya membaca seluruh folder files yang isi filenya txt
*/

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	//melakukan pembacaan dari folder files
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
