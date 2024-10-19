package main

import (
	"fmt"
	"sort"
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
		totalKalori(&X, barang, berat)
		fmt.Println("")
	}

	fmt.Println("== Greedy by Berat ==")
	bestBeratSolution := greedyByBerat(X, kapasitas)
	fmt.Println("== Greedy by Kalori ==")
	bestKaloriSolution := greedyByKalori(X, kapasitas)
	fmt.Println("== Greedy by Density ==")
	bestDensitySolution := greedyByDensity(X, kapasitas)

	bestSolution := bestBeratSolution
	if bestKaloriSolution.totalKalori > bestSolution.totalKalori {
		bestSolution = bestKaloriSolution
	}
	if bestDensitySolution.totalKalori > bestSolution.totalKalori {
		bestSolution = bestDensitySolution
	}

	fmt.Println("== Solusi Terbaik ==")
	fmt.Println("Total Berat yang Dipilih:", bestSolution.totalBerat)
	fmt.Println("Total Kalori yang Dipilih:", bestSolution.totalKalori)
	fmt.Println("Barang yang Dipilih:")
	for _, item := range bestSolution.selectedItems {
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

func greedyByBerat(T tabPendaki, capacity float64) greedySolution {
	sort.Slice(T[:], func(i, j int) bool {
		if T[i].jenis == "Minuman" && T[j].jenis != "Minuman" {
			return true
		} else if T[i].jenis != "Minuman" && T[j].jenis == "Minuman" {
			return false
		}
		return T[i].berat < T[j].berat
	})

	var selected []pendaki
	var totalKalori, totalBerat float64
	fmt.Println("Greedy by Berat:")
	for _, item := range T {
		if item.berat == 0 {
			continue
		}
		if totalBerat+item.berat <= capacity {
			selected = append(selected, item)
			totalBerat += item.berat
			totalKalori += item.totalkalori
			fmt.Printf("Mengambil %s: Berat=%d, Kalori=%d\n", item.barang, int(item.berat), int(item.totalkalori))
		}
	}

	if !checkMinumanCondition(totalBerat, totalKalori, selected) {
		fmt.Println("Syarat minuman atau kalori tidak terpenuhi setelah pemilihan barang.")
		totalKalori = 0
	} else {
		fmt.Println("Syarat minuman dan kalori terpenuhi setelah pemilihan barang.")
	}

	fmt.Printf("Total kalori yang diambil: %d\n\n", int(totalKalori))

	return greedySolution{selectedItems: selected, totalBerat: totalBerat, totalKalori: totalKalori}
}

func greedyByKalori(T tabPendaki, capacity float64) greedySolution {
	sort.Slice(T[:], func(i, j int) bool {
		if T[i].jenis == "Minuman" && T[j].jenis != "Minuman" {
			return true
		} else if T[i].jenis != "Minuman" && T[j].jenis == "Minuman" {
			return false
		}
		return T[i].totalkalori > T[j].totalkalori
	})

	var selected []pendaki
	var totalKalori, totalBerat float64
	fmt.Println("Greedy by Kalori:")
	for _, item := range T {
		if item.berat == 0 {
			continue
		}
		if totalBerat+item.berat <= capacity {
			selected = append(selected, item)
			totalBerat += item.berat
			totalKalori += item.totalkalori
			fmt.Printf("Mengambil %s: Berat=%d, Kalori=%d\n", item.barang, int(item.berat), int(item.totalkalori))
		}
	}

	if !checkMinumanCondition(totalBerat, totalKalori, selected) {
		fmt.Println("Syarat minuman atau kalori tidak terpenuhi setelah pemilihan barang.")
		totalKalori = 0
	} else {
		fmt.Println("Syarat minuman dan kalori terpenuhi setelah pemilihan barang.")
	}

	fmt.Printf("Total kalori yang diambil: %d\n\n", int(totalKalori))

	return greedySolution{selectedItems: selected, totalBerat: totalBerat, totalKalori: totalKalori}
}

func greedyByDensity(T tabPendaki, capacity float64) greedySolution {
	sort.Slice(T[:], func(i, j int) bool {
		if T[i].jenis == "Minuman" && T[j].jenis != "Minuman" {
			return true
		} else if T[i].jenis != "Minuman" && T[j].jenis == "Minuman" {
			return false
		}
		return float64(T[i].totalkalori)/float64(T[i].berat) > float64(T[j].totalkalori)/float64(T[j].berat)
	})

	var selected []pendaki
	var totalKalori, totalBerat float64
	fmt.Println("Greedy by Density:")
	for _, item := range T {
		if item.berat == 0 {
			continue
		}
		if totalBerat+item.berat <= capacity {
			selected = append(selected, item)
			totalBerat += item.berat
			totalKalori += item.totalkalori
			fmt.Printf("Mengambil %s: Berat=%d, Kalori=%d\n", item.barang, int(item.berat), int(item.totalkalori))
		}
	}

	if !checkMinumanCondition(totalBerat, totalKalori, selected) {
		fmt.Println("Syarat minuman atau kalori tidak terpenuhi setelah pemilihan barang.")
		totalKalori = 0
	} else {
		fmt.Println("Syarat minuman dan kalori terpenuhi setelah pemilihan barang.")
	}

	fmt.Printf("Total kalori yang diambil: %d\n\n", int(totalKalori))

	return greedySolution{selectedItems: selected, totalBerat: totalBerat, totalKalori: totalKalori}
}

func checkMinumanCondition(totalBerat, totalKalori float64, selected []pendaki) bool {
	var minumanBerat float64 = 0
	var totalBerat2 float64 = 0

	for _, item := range selected {
		totalBerat2 += item.berat
		if item.jenis == "Minuman" {
			minumanBerat += item.berat
		}
	}

	fmt.Println("Total Berat Minuman:", minumanBerat)
	fmt.Println("Total Berat Barang Terpilih:", totalBerat2)
	fmt.Println("Total Kalori Keseluruhan:", totalKalori)

	if minumanBerat >= 3000 && totalKalori >= 6000 {
		return true
	}

	return false
}

type greedySolution struct {
	selectedItems []pendaki
	totalBerat    float64
	totalKalori   float64
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Waktu eksekusi %s: %.7f detik\n", name, elapsed)
}
