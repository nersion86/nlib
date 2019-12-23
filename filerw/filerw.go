package filerw

import (
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

/* 	O_RDONLY = Read Only Open
O_WRONLY = Write Only Open
O_RDWR = Read/Write Open
O_APPEND = File Append Mode
O_CREATE = File Does Not Exist then Create. Or Exist then Open Mode
O_EXCL = File Does Not Exist, Create. Exist then Error
O_SYNC = Open Sync Mode
O_TRUNC = File Exist Then Clear File
*/

//FileRW is Simple file read/wirter using deserialize/serialize
type FileRW struct {
	file          *os.File    //OpenFile
	filePath      string      //File Path
	isAlreadyLoad bool        //load flag
	openFlag      int         //file open flags
	myMutex       *sync.Mutex //thread lock
}

//SerializeData is SomeData Convert To []byte
type SerializeData interface {
	Serialize() ([]byte, error) //SomeData to byte array
}

//DeserializeData is []byte Convert To SomeData
type DeserializeData interface {
	DeSerialize([]byte) error //Byte array to SomeData
}

//NewFileRW is Create New File Read/Write I/O
func NewFileRW(filePath string, flag int) (*FileRW, error) {

	file, err := os.OpenFile(filePath, flag, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	myMutex := &sync.Mutex{}

	nwFile := &FileRW{
		file:          file,
		filePath:      filePath,
		isAlreadyLoad: true,
		openFlag:      flag,
		myMutex:       myMutex,
	}

	return nwFile, nil
}

//NewFileRWAppend is Append & Create Mode
func NewFileRWAppend(filePath string) (*FileRW, error) {
	return NewFileRW(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND)
}

//NewFileRWClear is Open & Clear File Mode
func NewFileRWClear(filePath string) (*FileRW, error) {
	return NewFileRW(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC)
}

//WriteFile is SomeData Serialize And Write File
func (fw *FileRW) WriteFile(data SerializeData) (int, error) {

	if fw.isAlreadyLoad == false {
		return 0, errors.New("Not Opened File")
	}

	buf, err := data.Serialize()

	if err != nil {
		return 0, err
	}

	fw.myMutex.Lock()
	defer fw.myMutex.Unlock()

	//m... defer is useful?? defer not fast.
	return fw.file.Write(buf)
}

//ReadFile is Read And Buffer Convert To SomeData (Deserialize)
func (fw *FileRW) ReadFile(data DeserializeData) (interface{}, error) {

	if fw.isAlreadyLoad == false {
		return nil, errors.New("Not Opend File")
	}

	fw.myMutex.Lock()
	defer fw.myMutex.Unlock()

	buf, err := ioutil.ReadAll(fw.file)

	if err != nil {
		return nil, err
	}

	if err = data.DeSerialize(buf); err != nil {
		return nil, err
	}

	//data is not ref? hm....
	return data, nil
}

//Close is File I/O close
func (fw *FileRW) Close() {
	if fw.isAlreadyLoad == true {

		fw.myMutex.Lock()

		fw.file.Close()
		fw.isAlreadyLoad = false

		fw.myMutex.Unlock()
	}
}

//TODO : ReOpenFile & Basic String & MapData Serializer
