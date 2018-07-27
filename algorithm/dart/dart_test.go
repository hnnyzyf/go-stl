package algorithm

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

//func Test_trie(t *testing.T) {
//	ac := NewDart()
//
//	keywords := []string{
//		"space",
//		"keyword",
//		"ch",
//	}
//
//	ac.Build(keywords)
//
//	data := []struct {
//		text string
//		res  bool
//	}{
//		{"space", true},
//		{"keyword", true},
//		{"ch", true},
//		{"  ch", true},
//		{"chkeyword", true},
//		{"oooospace2", true},
//		{"c", false},
//		{"", false},
//		{"spac", false},
//		{"nothing", false},
//		{"dsadasdasspacdsadasdasdas", false},
//	}
//
//	for i := range data {
//		if ac.Matches(data[i].text) != data[i].res {
//			t.Error("Fail(", data[i].text, "),Expect:", data[i].res)
//		}
//	}
//
//}

//loadDictionary read key from a text file
func loadDictionary(path string) ([]string, error) {
	var (
		file   *os.File
		err    error
		part   []byte
		prefix bool
	)
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)
	text := bufio.NewReader(file)
	buffer := bytes.NewBuffer([]byte{})

	for {
		part, prefix, err = text.ReadLine()
		if err != nil {
			break
		}

		if _, err = buffer.Write(part); err != nil {
			return nil, err
		}

		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}

	if err == io.EOF {
		return lines, nil
	} else {
		return nil, err
	}
}

func loadText(path string) (io.RuneReader, error) {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil

}

func findhit(begin int, end int, pattern string) {
	fmt.Println("Find pattern:", pattern, "(", begin, ":", end, ")")
}

func Test_parseText(t *testing.T) {
	dictionary, err := loadDictionary("./resource/cn/dictionary.txt")
	if err != nil {
		t.Error(err)
	} //
	//_, err = loadText("./resource/cn/text.txt")
	//if err != nil {
	//t.Error(err)
	//}

	ac := NewDart()
	err = ac.Build(dictionary)
	fmt.Println(err)

	//ac.ParseText(text, findhit)

}
