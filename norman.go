package main

import f "fmt"

var (
	k        int
	p        *int
	student1 siswa
	shape    bangund
	rect     persegip
	square   bujurs
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

type bangund interface {
	luas() int
	keliling() int
}

type persegip struct {
	panjang, lebar int
}

func (p persegip) luas() int {
	return p.panjang * p.lebar
}

func (p persegip) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type bujurs struct {
	sisi int
}

func (b bujurs) luas() int {
	return b.sisi * b.sisi
}

func (b bujurs) keliling() int {
	return 4 * b.sisi
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
	ss(siswa{
		nama:   "Rulita",
		umur:   16,
		status: true,
		alamat: address{
			jalan:     "Jalan Raya Jagakarsa",
			kelurahan: "Jagakarsa",
		},
	})
	//
	student3 := struct {
		nama   string
		umur   int
		status bool
		alamat address
	}{
		nama:   "Delima",
		umur:   10,
		status: true,
		alamat: address{
			jalan:     "Jalan Raya Jagakarsa",
			kelurahan: "Jagakarsa",
		},
	}
	ss(student3)
	//
	rect.panjang = 15
	rect.lebar = 10
	f.Println(rect.luas(), rect.keliling())
	rect := persegip{
		panjang: 10,
		lebar:   6,
	}
	f.Println(rect.luas(), rect.keliling())
}

func m(p *int) {
	*p = *p * 10
	f.Println(p, *p)
}

func s(murid *siswa) bool {
	f.Println(*&murid.status)
	return !murid.status
} //

func ss(pelajar siswa) {
	f.Println(pelajar.nama)
}

// go mod init p
// git init
// git add .\norman.go 1
// git commit -m "Initial commit" 1
// git branch -M p
// git remote add origin https://github.com/nhasibuan/p.git
// git push -u origin p 1
