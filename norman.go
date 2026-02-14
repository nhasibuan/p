package main

import f "fmt"

var (
	k        int
	p        *int
	student1 siswa
)

type address struct {
	jalan     string
	kelurahan string
}

type siswa struct {
	nama   string
	umur   int
	status bool
	alamat address
}

func main() {
	var a int
	// a = 1
	a = a * 1
	i := 1
	i = i + 1
	f.Println(a, i)
	k = 6
	p = &k
	f.Println(k, p, *p)
	m(&k)
	//
	student1.nama = "Retno"
	student1.umur = 18
	student1.status = true
	student1.alamat.jalan = "Jalan Raya"
	student1.alamat.kelurahan = "Kelurahan A"
	f.Println(s(&student1))
	//
	student2 := siswa{
		nama:   "Norman",
		umur:   40,
		status: true,
		alamat: address{
			jalan:     "Jalan Raya Jagakarsa",
			kelurahan: "Jagakarsa",
		},
	}
	f.Println(student2)
}

func m(p *int) {
	*p = *p * 10
	f.Println(p, *p)
}

func s(murid *siswa) bool {
	f.Println(*&murid.status)
	return !murid.status
} //

// go mod init p
// git init
// git add .\norman.go 1
// git commit -m "Initial commit" 1
// git branch -M p
// git remote add origin https://github.com/nhasibuan/p.git
// git push -u origin p 1
