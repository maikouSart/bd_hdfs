package hdfs

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func splitToFileAndStore(fileName string, storeFile string) (chunkLen int,offsetLast int){
	data := readFileByBytes(fileName)
	var i int = 0;
	for i<len(data)/SPLIT_UNIT {
		FastWrite(storeFile+strconv.Itoa(i), data[i*SPLIT_UNIT:(i+1)*SPLIT_UNIT])
		i++
	}
	FastWrite(storeFile+strconv.Itoa(i), data[i*SPLIT_UNIT:])
	chunkLen = i
	offsetLast = len(data) - i*SPLIT_UNIT

	return chunkLen+1, offsetLast
}

func FastWrite(fileName string, data []byte){
	err := ioutil.WriteFile(fileName, data, 0666)
    if err != nil {
			sugarLogger.Error(err)
    }
}

func CreateFile(fileName string) (newFile *os.File) {
	defer newFile.Close()
	fmt.Println(fileName)
	newFile, err := os.Create(fileName)
  if err != nil {
		sugarLogger.Error(err)
	}
	// TDFSLogger.Println(newFile)
	return newFile
}

func showFileInfo(fileName string){
	fileInfo, err := os.Stat(fileName)
    if err != nil {
			sugarLogger.Error(err)
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
    fmt.Println("Is Directory: ", fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n", fileInfo.Sys())
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func DeleteFile(fileName string){
	err := os.Remove(fileName)
    if err != nil {
			sugarLogger.Error(err)
    }
}

func CleanFile(fileName string){
	// DeleteFile(fileName)
	// CreateFile(fileName)
	var data []byte
	FastWrite(fileName, data)
}

func OpenFile(fileName string) (file *os.File) {
	defer file.Close()
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
			sugarLogger.Error(err)
    }
	return file
}

func copyFile(oriFilename string, newFilename string){
	oriFile, err := os.Open(oriFilename)
    if err != nil {
			sugarLogger.Error(err)
    }
	defer oriFile.Close()

	newFile, err := os.Create(newFilename)
    if err != nil {
			sugarLogger.Error(err)
    }
    defer newFile.Close()
	
	bytesWritten, err := io.Copy(newFile, oriFile)
    if err != nil {
			sugarLogger.Error(err)
    }
		sugarLogger.Infof("Copied %d bytes.", bytesWritten)

    err = newFile.Sync()
    if err != nil {
			sugarLogger.Error(err)
    }
}

func readFileLimitedBytes(fileName string, limit int64){
	file, err := os.Open(fileName)
    if err != nil {
			sugarLogger.Error(err)
	}
	byteSlice := make([]byte, limit)
    numBytesRead, err := io.ReadFull(file, byteSlice)
    if err != nil {
			sugarLogger.Error(err)
    }
    fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: \n%s", byteSlice)
	fmt.Println()
}

func readFileByBytes(fileName string)([]byte){
	file, err := os.Open(fileName)
    if err != nil {
			sugarLogger.Error(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil { 
		sugarLogger.Error(err)
	}
	// fmt.Printf("Data as hex: %x\n", data)
	// fmt.Printf("Data as bytes: %b\n", data)
    // fmt.Printf("Data as string: %s\n", data)
	// fmt.Println("Number of bytes read:", len(data))
	return data
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SplitToChunksByName(bigFileName string) (chunklist []ChunkUnit, dataLen int){
	data := readFileByBytes(bigFileName)
	var i int = 0;
	dataLen = len(data)
	// fmt.Printf("dataLen:%d\n",dataLen)
	for i < len(data)/SPLIT_UNIT {
		chunklist = append(chunklist, data[i*SPLIT_UNIT:(i+1)*SPLIT_UNIT])
		i++
	}
	chunklist = append(chunklist, data[i*SPLIT_UNIT:])
	return chunklist, dataLen
}

func SplitToChunksByFobj(bigFile *os.File) (chunklist []ChunkUnit, dataLen int){
	data, err := ioutil.ReadAll(bigFile)
    if err != nil {
			sugarLogger.Error(err)
	}
	var i int = 0;
	dataLen = len(data)
	for i < dataLen/SPLIT_UNIT {
		chunklist[i] = data[i*SPLIT_UNIT:(i+1)*SPLIT_UNIT]
		i++
	}
	chunklist[i] = data[i*SPLIT_UNIT:]
	return chunklist, dataLen
}