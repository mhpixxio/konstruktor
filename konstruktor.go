package konstruktor

import (
	"math/rand"
	"time"

	pb "github.com/mhpixxio/pb"
)

type RandomData struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C string `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E string `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	F string `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G string `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
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
