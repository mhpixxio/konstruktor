package konstruktor

import (
	"math/rand"
	"time"

	pb "github.com/mhpixxio/pb"
)

type RandomData struct {
	a string
	b string
	c string
	d string
	e string
	f string
	g string
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
		random := RandomData{a: RandStringRunes(lengthString), b: RandStringRunes(lengthString), c: RandStringRunes(lengthString), d: RandStringRunes(lengthString), e: RandStringRunes(lengthString), f: RandStringRunes(lengthString), g: RandStringRunes(lengthString)}
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
