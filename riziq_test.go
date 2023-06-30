package riziq

import (
	"fmt"
	"testing"
)

// func TestInsertPelanggan(t *testing.T) {
// 	dbname := "tagihan"
// 	pelanggan := Pelanggan{
// 		ID:        primitive.NewObjectID(),
// 		Nama:      "Farhan Riziq",
// 		Alamat:    "bandung",
// 		NoTelepon: "081238765694",
// 		Email:     "farhan@gmail.com",
// 	}
// 	insertedID := InsertPelanggan(dbname, pelanggan)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertTagihan(t *testing.T) {
// 	dbname := "tagihan"
// 	tagihan := Tagihan{
// 		ID:               primitive.NewObjectID(),
// 		TanggalTagihan:   "08.01.2021",
// 		TotalTagihan:     "$5000.00",
// 		StatusPembayaran: "complete",
// 	}
// 	insertedID := InsertTagihan(dbname, tagihan)
// 	if insertedID == nil {
// 		t.Error("Failed to insert surat")
// 	}
// }

// func TestInsertPembayaran(t *testing.T) {
// 	dbname := "tagihan"
// 	pembayaran := Pembayaran{
// 		ID:                primitive.NewObjectID(),
// 		TanggalPembayaran: "05.01.2021",
// 		JumlahPembayaran:  "$2500.00",
// 		MetodePembayaran:  "transfer",
// 	}
// 	insertedID := InsertPembayaran(dbname, pembayaran)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

// func TestInsertProduk(t *testing.T) {
// 	dbname := "tagihan"
// 	produk := Produk{
// 		ID:          primitive.NewObjectID(),
// 		NamaProduk:  "Spotify",
// 		HargaProduk: "$5000.00",
// 	}
// 	insertedID := InsertProduk(dbname, produk)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

// func TestInsertAbout(t *testing.T) {
// 	dbname := "tagihan"
// 	about := About{
// 		ID:      primitive.NewObjectID(),
// 		IsiSatu: "isi satu",
// 		IsiDua:  "isi dua",
// 		Image:   "tes",
// 	}
// 	insertedID := InsertAbout(dbname, about)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

func TestGetDataAbout(t *testing.T) {
	stats := "tes"
	data := GetDataAbout(stats)
	fmt.Println(data)
}

func TestGetDataPelanggan(t *testing.T) {
	stats := "Farhan Riziq"
	data := GetDataPelanggan(stats)
	fmt.Println(data)
}

func TestGetDataPembayaran(t *testing.T) {
	stats := "05.01.2021"
	data := GetDataPembayaran(stats)
	fmt.Println(data)
}

func TestGetDataProduk(t *testing.T) {
	stats := "Spotify"
	data := GetDataProduk(stats)
	fmt.Println(data)
}
func TestGetTagihan(t *testing.T) {
	stats := "08.01.2021"
	data := GetDataTagihan(stats)
	fmt.Println(data)
}
