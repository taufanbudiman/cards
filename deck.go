package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// membuat slice baru dengan nama deck
//
// dimana deck tersebut berisi kumpulan string
type deck []string

func newDeck() deck {
	// membuat deck string baru sesuai string yang sudah ditentukan di
	// variabel cardSuites dan cardValues
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	// menampilkan semua kartu berdasarkan reviever-nya yaitu, deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// return card of hand, remaining card in deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// mengubah slice string menjadi string, dengan menggabungkan semua
	// yang ada pada slice dengan pemisah ","
	// contoh : Ace of Spades,Two of Diomonds,Two of Spades
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// Menyimpan hasil dari `toString()` ke hard drive dalam bentuk file
	// dengan permission 0666 (semua orang bisa write dan read)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFormFile(filename string) deck {
	// mengambil isi deck, dengan cara membaca file yang ada pada hard drive
	// jika file tidak ada menampilkan error, dan keluar dari program
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		// `os.Exit(code int)`, jika 0 maka sukses, jika selain 0 maka ada error
		os.Exit(1)
	}
	// mengubah bs (byte dari readfile) menjadi string, karena deck berupa slice string
	// memotong string dengan "," agar menghasilkan slice string
	split_strings := strings.Split(string(bs), ",")
	// memasukkan hasil dari split, ke object deck [] string
	return deck(split_strings)
}

func (d deck) shuffle() {
	// swap index dari deck (mengacak urutan pada deck slice string)
	// UnixNano digunakan untuk generate unix int64 dengan menggunakan tanggal
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
