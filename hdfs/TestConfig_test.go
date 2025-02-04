package hdfs

import (
	"fmt"
	"testing"
)

func getNodes() *Folder {
	return &Folder{
		"root",
		[]*Folder{},
		[]*File{{
			"data.txt",
			1024,
			[]FileChunk{},
			0,
			"",
			"",
		}},
	}
}

func TestMkdir(t *testing.T) {
	Nodes := getNodes()
	fmt.Println(Nodes.CreateFolder("/root", "data"))
	fmt.Println(Nodes.Folder[0].Name)
}

// func TestGetFileList(t *testing.T) {
// 	Nodes := getNodes()
// 	FileList := Nodes.GetFileList("/root")
// 	t.Log(FileList)
// 	if FileList != nil {
// 		t.Log(len(FileList))
// 	}
// }

func TestGetFileNode(t *testing.T) {
	Nodes := getNodes()
	File, err := Nodes.GetFileNode("/root/data.txt")
	t.Log(File)
	t.Log(err)
}

//func TestFunc(t *testing.T) {
//	x := "/root/teset"
//	fmt.Println(strings.Split(x, "/"))
//	fmt.Println(len(strings.Split(x, "/")))
//}
