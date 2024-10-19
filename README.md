# Website Strategi Algoritma

Program ini bertujuan untuk membantu para pendaki dalam menghitung jumlah makanan dan minuman yang perlu dibawa selama ekspedisi, dengan memperhitungkan kapasitas maksimal beban tas, kebutuhan kalori harian, serta kebutuhan minimum air. Program ini akan diimplementasikan menggunakan dua pendekatan algoritma, yaitu greedy dan dynamic programming, dalam bahasa pemrograman JavaScript. Selanjutnya, program ini juga akan menganalisis efisiensi dari masing-masing algoritma tersebut.

# Fitur Utama
•	Input Data: Pengguna dapat memasukkan jenis dan jumlah barang (makanan/minuman) yang akan dibawa, serta berat masing-masing barang.
•	Data Kalori: Program menyediakan data kalori dari berbagai makanan dan minuman yang umum dibawa oleh pendaki.
•	Algoritma Greedy: Terdapat tiga strategi pemilihan menggunakan algoritma greedy:
    o	Greedy by Weight: Memprioritaskan barang dengan berat paling ringan terlebih dahulu.
    o	Greedy by Calories: Memprioritaskan barang yang memberikan kalori paling tinggi.
    o	Greedy by Density: Mengutamakan barang dengan perbandingan kalori terhadap berat yang paling tinggi.
•	Algoritma Dynamic Programming (Knapsack): Menyediakan solusi optimal untuk memilih barang berdasarkan kapasitas tas dengan mempertimbangkan kebutuhan kalori minimal dan syarat berat air minum.
•	Analisis Efisiensi: Membandingkan waktu eksekusi antara strategi greedy dan dynamic programming untuk memahami performa masing-masing metode.
# Struktur Program
•	Algoritma Dynamic Programming: Menggunakan pendekatan knapsack problem untuk menentukan kombinasi barang terbaik berdasarkan berat dan kalori yang memenuhi kapasitas tas.
•	Algoritma Greedy: Mengimplementasikan tiga strategi berbeda dalam memilih barang, yaitu berdasarkan berat, kalori, dan kepadatan kalori. Pemilihan dilakukan secara iteratif sampai mencapai kapasitas maksimal tas atau terpenuhinya kebutuhan minimal kalori dan air minum.
# Hasil dan Analisis
•	Hasil akhir mencakup barang-barang yang terpilih untuk dibawa, total berat yang dibawa, dan total kalori yang diperoleh.
•	Program juga mengevaluasi apakah persyaratan minimum berat air dan kalori harian terpenuhi.
•	Terdapat perbandingan waktu eksekusi antara pendekatan greedy dan dynamic programming, yang ditampilkan setelah pemilihan selesai, memberikan wawasan tentang efisiensi masing-masing metode dalam konteks masalah ini.


