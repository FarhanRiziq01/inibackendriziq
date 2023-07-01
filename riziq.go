package inibackendriziq

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPelanggan(db string, pelanggan Pelanggan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("pelanggan").InsertOne(context.TODO(), pelanggan)
	if err != nil {
		fmt.Printf("InsertPelanggan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertTagihan(db string, tagihan Tagihan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("tagihan").InsertOne(context.TODO(), tagihan)
	if err != nil {
		fmt.Printf("InsertTagihan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPembayaran(db string, pembayaran Pembayaran) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("pembayaran").InsertOne(context.TODO(), pembayaran)
	if err != nil {
		fmt.Printf("InsertPembayaran: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProduk(db string, produk Produk) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("produk").InsertOne(context.TODO(), produk)
	if err != nil {
		fmt.Printf("InsertProduk: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAbout(db string, about About) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDataPelanggan(stats string) (data []Pelanggan) {
	user := MongoConnect("tagihan").Collection("pelanggan")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPelanggan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataTagihan(stats string) (data []Tagihan) {
	user := MongoConnect("tagihan").Collection("tagihan")
	filter := bson.M{"tanggaltagihan": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataTagihan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPembayaran(stats string) (data []Pembayaran) {
	user := MongoConnect("tagihan").Collection("pembayaran")
	filter := bson.M{"tanggalpembayaran": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPembayaran :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataProduk(stats string) (data []Produk) {
	user := MongoConnect("tagihan").Collection("produk")
	filter := bson.M{"namaproduk": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataProduk :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataAbout(stats string) (data []About) {
	user := MongoConnect("tagihan").Collection("about")
	filter := bson.M{"isisatu": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAbout :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
