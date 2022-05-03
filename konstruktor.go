package konstruktor

import (
	"math/rand"
	"time"

	pb "github.com/mhpixxio/pb"
)

type RandomData struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
	F string `json:"f"`
	G string `json:"g"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateBigData(lengthString int, lengthSlice int) []RandomData {

	bigdata := []RandomData{}

	for i := 0; i < lengthSlice; i++ {
		random := RandomData{A: RandStringRunes(lengthString), B: RandStringRunes(lengthString), C: RandStringRunes(lengthString), D: RandStringRunes(lengthString), E: RandStringRunes(lengthString), F: RandStringRunes(lengthString), G: RandStringRunes(lengthString)}
		bigdata = append(bigdata, random)
	}

	return bigdata

}

func CreateBigData_proto(lengthString int, lengthSlice int) []*pb.RandomData {

	bigdata := []*pb.RandomData{}

	for i := 0; i < lengthSlice; i++ {
		random := &pb.RandomData{A: RandStringRunes(lengthString), B: RandStringRunes(lengthString), C: RandStringRunes(lengthString), D: RandStringRunes(lengthString), E: RandStringRunes(lengthString), F: RandStringRunes(lengthString), G: RandStringRunes(lengthString)}
		bigdata = append(bigdata, random)
	}

	return bigdata

}
