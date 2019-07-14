package fileserver

import (
	"net/http"
	"log"
	"io/ioutil"
	"FileServer/server/entity"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//type Callback func(x, y int) int

var (
	server *http.Server
)

type StartCallBack interface {
	OnError(err error)
	OnSuccess(files []*entity.File)
}

func StopFileServer() {
	if server != nil {
		err := server.Shutdown(nil)
		if err != nil {
			log.Print("StopFileServer error:", err.Error())
		}
	}
}

func StartFileServer(conf *entity.FileServerConf, startCallBack StartCallBack) {
	StopFileServer()
	// 获取文件的MD5
	files, err := queryFileList(conf.FilePath)
	if err != nil {
		log.Println("queryFileList err:", err.Error())
		startCallBack.OnError(err)
		return
	}
	log.Println("startCallBack:", startCallBack)
	// 回调通知文件查询成功，后面服务开启
	startCallBack.OnSuccess(files)
	// 开启文件服务
	err = openFileServer(conf)
	if err != nil {
		log.Println("ListenAndServe err:", err.Error())
		startCallBack.OnError(err)
		return
	}
}

func openFileServer(conf *entity.FileServerConf) (err error) {
	if server == nil {
		http.Handle(conf.Route, http.StripPrefix(conf.Route, http.FileServer(http.Dir(conf.FilePath))))
	}
	server = &http.Server{
		Addr: ":" + conf.Port,
		Handler: nil,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Println("ListenAndServe err:", err.Error())
		return err
	}
	return nil
}

func queryFileList(filePath string) (fileList []*entity.File, err error) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	fileList = make([]*entity.File, 0)
	for _, file := range files {
		md5 := md5.New()
		realFile, err := os.Open(filePath + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		io.Copy(md5, realFile)
		MD5Str := hex.EncodeToString(md5.Sum(nil))
		fileEntity := &entity.File{
			FileName: file.Name(),
			FilePath: filePath + "/" + file.Name(),
			FileMD5: MD5Str,
			FileSize: file.Size(),
		}
		fileList = append(fileList, fileEntity)
	}
	return fileList, nil
}
