package gohelper

import (
	"testing"
)

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if KFile.GetExt(filename) !="go" {
		t.Error("file extension error")
		return
	}
}

func BenchmarkGetExt(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.GetExt(filename)
	}
}

func TestFileSize(t *testing.T)  {
	filename := "./file.go"
	if KFile.FileSize(filename) <=0 {
		t.Error("file size error")
		return
	}
}

func BenchmarkFileSize(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.FileSize(filename)
	}
}

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkIsExist(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsExist(filename)
	}
}

func TestIsWritable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsWritable(filename) {
		t.Error("file can not write")
		return
	}
}

func BenchmarkIsWritable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsWritable(filename)
	}
}

func TestIsReadable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsReadable(filename) {
		t.Error("file can not read")
		return
	}
}

func BenchmarkIsReadable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsReadable(filename)
	}
}

func TestIsFile(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsFile(filename) {
		t.Error("isn`t a file")
		return
	}
}

func BenchmarkIsFile(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsFile(filename)
	}
}

func TestIsDir(t *testing.T) {
	dirname := "./"
	if !KFile.IsDir(dirname) {
		t.Error("isn`t a dir")
		return
	}
}

func BenchmarkIsDir(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsDir(filename)
	}
}

func TestIsBinary(t *testing.T) {
	filename := "./file.go"
	if KFile.IsBinary(filename) {
		t.Error("file isn`t binary")
		return
	}
}

func BenchmarkIsBinary(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsBinary(filename)
	}
}

func TestIsImg(t *testing.T) {
	filename := "./testdata/diglett.png"
	if !KFile.IsImg(filename) {
		t.Error("file isn`t img")
		return
	}
}

func BenchmarkIsImg(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i:=0;i<b.N;i++{
		KFile.IsImg(filename)
	}
}

func TestAbsPath(t *testing.T) {
	filename := "./testdata/diglett.png"
	abspath := KFile.AbsPath(filename)
	if !KFile.IsExist(abspath) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkAbsPath(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i:=0;i<b.N;i++{
		KFile.AbsPath(filename)
	}
}

func TestCopyFile(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"

	num, err := KFile.CopyFile(src, des, FCOVER_ALLOW)
	if err != nil || num ==0 {
		t.Error("copy file fail")
		return
	}
}

func BenchmarkCopyFile(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"
	for i:=0;i<b.N;i++{
		_,_ = KFile.CopyFile(src, des, FCOVER_ALLOW)
	}
}

func TestFastCopy(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"

	num, err := KFile.FastCopy(src, des, FCOVER_ALLOW)
	if err != nil || num ==0 {
		t.Error("copy file fail")
		return
	}
}

func BenchmarkFastCopy(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"
	for i:=0;i<b.N;i++{
		_,_ = KFile.FastCopy(src, des, FCOVER_ALLOW)
	}
}
