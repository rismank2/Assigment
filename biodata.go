package main

import (
	"fmt"
	"os"
	"strconv"
)

type data struct {
	no_absen  int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Datas = []data{
	{no_absen: 0, nama: "Tidak Ada"},
	{no_absen: 1, nama: "Yuniar Wahyu Pratama"},
	{no_absen: 2, nama: "Fahri Novaldi"},
	{no_absen: 3, nama: "Addumairi Azhar Maksum"},
	{no_absen: 4, nama: "Arji Surya Maulana"},
	{no_absen: 5, nama: "Muhammad Fikri Sandi Pratama"},
	{no_absen: 6, nama: "Aditya Dhian Pratama"},
	{no_absen: 7, nama: "Yohannes Julius"},
	{no_absen: 8, nama: "Fariz Dandy Lazuardi"},
	{no_absen: 9, nama: "Saritalia"},
	{no_absen: 10, nama: "Bima Karo Karo"},
}

func view(data) {
	var mu data
	mu.alamat = "Jakarta Pusat"
	mu.pekerjaan = "Karyawan Swasta"
	mu.alasan = "Ingin menambah wawasan bahasa Pemrograman"

	ab, err := strconv.Atoi(os.Args[1])
	fmt.Println("Data Peserta \t:", Datas[ab])
	fmt.Println("Alamat \t\t:", mu.alamat)
	fmt.Println("Pekerjaan \t:", mu.pekerjaan)
	fmt.Println("Alasan \t\t:", mu.alasan)
	if err != nil {
		panic(err)
	}
}

func main() {
	view(data{})
	// fmt.Println(mu.nama[1])

}
