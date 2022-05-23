package konstruktor

import (
	"math/rand"
	"time"

	pb "github.com/mhpixxio/pb"
)

type RandomData struct {
	Aaabcdef int32  `json:"aaabcdef,omitempty"`
	Ababcdef bool   `json:"ababcdef,omitempty"`
	Acabcdef bool   `json:"acabcdef,omitempty"`
	Adabcdef bool   `json:"adabcdef,omitempty"`
	Aeabcdef bool   `json:"aeabcdef,omitempty"`
	Afabcdef bool   `json:"afabcdef,omitempty"`
	Agabcdef bool   `json:"agabcdef,omitempty"`
	Ahabcdef bool   `json:"ahabcdef,omitempty"`
	Aiabcdef bool   `json:"aiabcdef,omitempty"`
	Ajabcdef bool   `json:"ajabcdef,omitempty"`
	Akabcdef bool   `json:"akabcdef,omitempty"`
	Alabcdef string `json:"alabcdef,omitempty"`
	Amabcdef string `json:"amabcdef,omitempty"`
	Anabcdef string `json:"anabcdef,omitempty"`
	Aoabcdef string `json:"aoabcdef,omitempty"`
	Apabcdef string `json:"apabcdef,omitempty"`
	Aqabcdef string `json:"aqabcdef,omitempty"`
	Arabcdef string `json:"arabcdef,omitempty"`
	Asabcdef string `json:"asabcdef,omitempty"`
	Atabcdef string `json:"atabcdef,omitempty"`
	Auabcdef string `json:"auabcdef,omitempty"`
	Avabcdef string `json:"avabcdef,omitempty"`
	Awabcdef string `json:"awabcdef,omitempty"`
	Axabcdef string `json:"axabcdef,omitempty"`
	Ayabcdef string `json:"ayabcdef,omitempty"`
	Azabcdef string `json:"azabcdef,omitempty"`
	Baabcdef string `json:"baabcdef,omitempty"`
	Bbabcdef string `json:"bbabcdef,omitempty"`
	Bcabcdef string `json:"bcabcdef,omitempty"`
	Bdabcdef string `json:"bdabcdef,omitempty"`
	Beabcdef string `json:"beabcdef,omitempty"`
	Bfabcdef string `json:"bfabcdef,omitempty"`
	Bgabcdef string `json:"bgabcdef,omitempty"`
	Bhabcdef string `json:"bhabcdef,omitempty"`
	Biabcdef string `json:"biabcdef,omitempty"`
	Bjabcdef string `json:"bjabcdef,omitempty"`
	Bkabcdef string `json:"bkabcdef,omitempty"`
	Blabcdef string `json:"blabcdef,omitempty"`
	Bmabcdef string `json:"bmabcdef,omitempty"`
	Bnabcdef string `json:"bnabcdef,omitempty"`
	Boabcdef string `json:"boabcdef,omitempty"`
	Bpabcdef string `json:"bpabcdef,omitempty"`
	Bqabcdef string `json:"bqabcdef,omitempty"`
	Brabcdef string `json:"brabcdef,omitempty"`
	Bsabcdef string `json:"bsabcdef,omitempty"`
	Btabcdef string `json:"btabcdef,omitempty"`
	Buabcdef string `json:"buabcdef,omitempty"`
	Bvabcdef string `json:"bvabcdef,omitempty"`
	Bwabcdef string `json:"bwabcdef,omitempty"`
	Bxabcdef string `json:"bxabcdef,omitempty"`
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
		random := RandomData{Aaabcdef: int32(rand.Intn(8999999) + 1000000),
			Ababcdef: true,
			Acabcdef: true,
			Adabcdef: false,
			Aeabcdef: false,
			Afabcdef: false,
			Agabcdef: false,
			Ahabcdef: false,
			Aiabcdef: false,
			Ajabcdef: false,
			Akabcdef: false,
			Alabcdef: "",
			Amabcdef: "",
			Anabcdef: "",
			Aoabcdef: "",
			Apabcdef: "",
			Aqabcdef: "",
			Arabcdef: "",
			Asabcdef: "",
			Atabcdef: "",
			Auabcdef: "",
			Avabcdef: "",
			Awabcdef: RandStringRunes(lengthString),
			Axabcdef: RandStringRunes(lengthString),
			Ayabcdef: RandStringRunes(lengthString),
			Azabcdef: RandStringRunes(lengthString),
			Baabcdef: RandStringRunes(lengthString),
			Bbabcdef: RandStringRunes(lengthString),
			Bcabcdef: RandStringRunes(lengthString),
			Bdabcdef: RandStringRunes(lengthString),
			Beabcdef: RandStringRunes(lengthString),
			Bfabcdef: RandStringRunes(lengthString),
			Bgabcdef: RandStringRunes(lengthString),
			Bhabcdef: RandStringRunes(lengthString),
			Biabcdef: RandStringRunes(lengthString),
			Bjabcdef: RandStringRunes(lengthString),
			Bkabcdef: RandStringRunes(lengthString),
			Blabcdef: RandStringRunes(lengthString),
			Bmabcdef: RandStringRunes(lengthString),
			Bnabcdef: RandStringRunes(lengthString),
			Boabcdef: RandStringRunes(lengthString),
			Bpabcdef: RandStringRunes(lengthString),
			Bqabcdef: RandStringRunes(lengthString),
			Brabcdef: RandStringRunes(lengthString),
			Bsabcdef: RandStringRunes(lengthString),
			Btabcdef: RandStringRunes(lengthString),
			Buabcdef: RandStringRunes(lengthString),
			Bvabcdef: RandStringRunes(lengthString),
			Bwabcdef: RandStringRunes(lengthString),
			Bxabcdef: RandStringRunes(lengthString)}
		bigdata = append(bigdata, random)

	}

	return bigdata

}

func CreateBigData_proto(lengthString int, lengthSlice int) []*pb.RandomData {

	bigdata := []*pb.RandomData{}

	for i := 0; i < lengthSlice; i++ {
		random := &pb.RandomData{Aaabcdef: int32(rand.Intn(8999999) + 1000000),
			Ababcdef: true,
			Acabcdef: true,
			Adabcdef: false,
			Aeabcdef: false,
			Afabcdef: false,
			Agabcdef: false,
			Ahabcdef: false,
			Aiabcdef: false,
			Ajabcdef: false,
			Akabcdef: false,
			Awabcdef: RandStringRunes(lengthString),
			Axabcdef: RandStringRunes(lengthString),
			Ayabcdef: RandStringRunes(lengthString),
			Azabcdef: RandStringRunes(lengthString),
			Baabcdef: RandStringRunes(lengthString),
			Bbabcdef: RandStringRunes(lengthString),
			Bcabcdef: RandStringRunes(lengthString),
			Bdabcdef: RandStringRunes(lengthString),
			Beabcdef: RandStringRunes(lengthString),
			Bfabcdef: RandStringRunes(lengthString),
			Bgabcdef: RandStringRunes(lengthString),
			Bhabcdef: RandStringRunes(lengthString),
			Biabcdef: RandStringRunes(lengthString),
			Bjabcdef: RandStringRunes(lengthString),
			Bkabcdef: RandStringRunes(lengthString),
			Blabcdef: RandStringRunes(lengthString),
			Bmabcdef: RandStringRunes(lengthString),
			Bnabcdef: RandStringRunes(lengthString),
			Boabcdef: RandStringRunes(lengthString),
			Bpabcdef: RandStringRunes(lengthString),
			Bqabcdef: RandStringRunes(lengthString),
			Brabcdef: RandStringRunes(lengthString),
			Bsabcdef: RandStringRunes(lengthString),
			Btabcdef: RandStringRunes(lengthString),
			Buabcdef: RandStringRunes(lengthString),
			Bvabcdef: RandStringRunes(lengthString),
			Bwabcdef: RandStringRunes(lengthString),
			Bxabcdef: RandStringRunes(lengthString)}
		bigdata = append(bigdata, random)
	}

	return bigdata

}
