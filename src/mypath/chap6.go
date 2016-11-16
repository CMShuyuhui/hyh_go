package mypath

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func Chap6() {
	TestString := "Hi, I'm HuYuhui\n"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("\n\n%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	fmt.Println("-------------------------------------\n\n")
	hash2()
}

func hash2() {
	TestFile := "123.txt"
	infile, inerr := os.Open(TestFile)
	defer infile.Close()
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("\n\n%x  %s\n\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("\n\n%x  %s\n\n", sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}
