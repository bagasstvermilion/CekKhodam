package main
import (
	"fmt"
	"math/rand"
	"time"
)

type khodam [11]string

func bacaData() string {
	var nama string

	fmt.Print("Masukkan nama kamu: ")
	fmt.Scan(&nama)
	return nama
}

func variasiKhodam(A *khodam) {
	A[0] = "Elang Hitam Dari Negeri Jawa"
	A[1] = "Macan Putih Ki Darsono"
	A[2] = "Landak Berduri Emas"
	A[3] = "Belut Putih Yang Ganas"
	A[4] = "Kelabang Madura si Raja Iblis"
	A[5] = "Singa Laut Yang Bermutasi"
	A[6] = "Pembangkit Listrik Tenaga Angin"
	A[7] = "Legenda Hantu Kota Mataram"
	A[8] = "Monster Hutan Kalimantan"
	A[9] = "Bandara Tua Yang Terbengkalai"
	A[10] = "Ular Berkepala Dua Yang Mengerikan"
}

func main() {
	var khodamS khodam
	fmt.Println("")
	nama := bacaData()
	fmt.Println("")

	variasiKhodam(&khodamS)

	rand.Seed(time.Now().UnixNano())
	acakIndex := rand.Intn(len(khodamS))
	fmt.Printf("Hai %s!, Khodam kamu adalah: %s\n", nama, khodamS[acakIndex])

	acakAtk := rand.Intn(300)
	acakDef := rand.Intn(300)
	fmt.Printf("ATK: %d\n", acakAtk)
	fmt.Printf("DEF: %d\n", acakDef)
	total := (acakAtk + acakDef) / 2
	if (total < 50) {
		fmt.Println("Tingkat Rarity: D")
	} else if (total < 100) {
		fmt.Println("Tingkat Rarity: C")
	} else if (total < 150) {
		fmt.Println("Tingkat Rarity: B")
	} else if (total < 200) {
		fmt.Println("Tingkat Rarity: A")
	} else if (total < 250) {
		fmt.Println("Tingkat Rarity: S")
	} else if (total < 300) {
		fmt.Println("Tingkat Rarity: SS")
	} else {
		fmt.Println("Tingkat Rarity: SSR")
	}
}