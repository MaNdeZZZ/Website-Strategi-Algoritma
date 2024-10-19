package main

import (
	"fmt"
	"time"
)

const nmax int = 10

type pendaki struct {
	barang, jenis              string
	kalori, berat, totalkalori float64
}

type tabPendaki [nmax]pendaki

func main() {
	var X tabPendaki
	var barang string
	var jumlahB int
	var berat, kapasitas float64

	dataKalori(&X)
	fmt.Print("Berapa barang yang ingin dimasukkan? ")
	fmt.Scan(&jumlahB)
	fmt.Print("Masukkan kapasitas maksimum tas: ")
	fmt.Scan(&kapasitas)
	for i := 0; i < jumlahB; i++ {
		fmt.Print("Masukan nama makanan/minuman dan beratnya (dalam gram): ")
		fmt.Scan(&barang, &berat)
		jenisBarang(&X, barang)
		// fmt.Println("=====================================")
		totalKalori(&X, barang, berat)
		fmt.Println("")
	}

	fmt.Println("== Dynamic Programming ==")
	dpSolution := knapsackDynamic(X, kapasitas)

	fmt.Println("== Solusi ==")
	fmt.Println("Total Berat yang Dipilih:", dpSolution.totalBerat)
	fmt.Println("Total Kalori yang Dipilih:", dpSolution.totalKalori)
	if !dpSolution.minumanSyarat {
		fmt.Println("Syarat minuman atau kalori tidak terpenuhi setelah pemilihan barang.")
	}
	fmt.Println("Barang yang Dipilih:")
	for _, item := range dpSolution.selectedItems {
		fmt.Printf("- %s: Berat=%d, Kalori=%d\n", item.barang, int(item.berat), int(item.totalkalori))
	}
	fmt.Print("")
	start := time.Now()
	defer trackTime(start, "program")
}

func jenisBarang(T *tabPendaki, barang string) {
	var index int = -1
	var jenis string

	for i := 0; i < nmax; i++ {
		if T[i].barang == barang {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Print("Jenis barang yang dimasukkan: ")
		fmt.Scan(&jenis)
		T[index].jenis = jenis
	}
}

func totalKalori(T *tabPendaki, barang string, berat float64) {
	var index int = -1

	for i := 0; i < nmax; i++ {
		if T[i].barang == barang {
			index = i
			break
		}
	}

	if index != -1 {
		T[index].totalkalori = (berat / 100) * T[index].kalori
		T[index].berat = berat
		// fmt.Printf("Total kalori untuk %s: %d\nJenis Barang: %s\n", barang, int(T[index].totalkalori), T[index].jenis)
	} else {
		fmt.Printf("Barang %s tidak ditemukan dalam daftar.\n", barang)
	}
}

func dataKalori(T *tabPendaki) {
	T[0].barang = "Nasi"
	T[0].kalori = 129
	T[1].barang = "Indomie"
	T[1].kalori = 430
	T[2].barang = "Sarden"
	T[2].kalori = 100
	T[3].barang = "Sosis"
	T[3].kalori = 220
	T[4].barang = "Wafer"
	T[4].kalori = 500
	T[5].barang = "Roti"
	T[5].kalori = 310
	T[6].barang = "Biskuit"
	T[6].kalori = 500
	T[7].barang = "AirPutih"
	T[7].kalori = 0
	T[8].barang = "Susu"
	T[8].kalori = 90
	T[9].barang = "Kopi"
	T[9].kalori = 450
}

type dpSolution struct {
	selectedItems []pendaki
	totalBerat    float64
	totalKalori   float64
	minumanSyarat bool
}

func knapsackDynamic(items tabPendaki, capacity float64) dpSolution {
	n := len(items)

	dp := make([][]dpSolution, n+1)
	for i := range dp {
		dp[i] = make([]dpSolution, int(capacity)+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= int(capacity); w++ {
			if items[i-1].berat <= float64(w) {
				withItem := dp[i-1][w-int(items[i-1].berat)]
				withItem.totalBerat += items[i-1].berat
				withItem.totalKalori += items[i-1].totalkalori

				withoutItem := dp[i-1][w]

				if items[i-1].jenis == "Minuman" {
					if withItem.totalBerat < 3000 {
						if withItem.totalBerat+items[i-1].berat >= 3000 {
							if withItem.totalKalori+items[i-1].totalkalori >= 6000 {
								withItem.selectedItems = append(withItem.selectedItems, items[i-1])
								withItem.totalBerat += items[i-1].berat
								withItem.totalKalori += items[i-1].totalkalori
								dp[i][w] = withItem
							}
						}
					} else {
						if withItem.totalKalori+items[i-1].totalkalori >= 6000 && withItem.totalBerat <= capacity {
							withItem.selectedItems = append(withItem.selectedItems, items[i-1])
							dp[i][w] = withItem
						}
					}
				} else {
					if withItem.totalKalori > withoutItem.totalKalori && withItem.totalBerat <= capacity {
						withItem.selectedItems = append(withItem.selectedItems, items[i-1])
						dp[i][w] = withItem
					} else {
						dp[i][w] = withoutItem
					}
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	solution := dp[n][int(capacity)]

	solution.minumanSyarat = solution.totalBerat >= 3000 || solution.totalKalori >= 6000

	return solution
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Waktu eksekusi %s: %.7f detik\n", name, elapsed)
}
