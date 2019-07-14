# FileServer
GoLang实现的文件服务
## 使用方式
参考main.go即可,开启之后就可以访问http://127.0.0.1/staticfile下面的文件了
```
type StartCallBackTest struct {

}

func (scbt StartCallBackTest) OnError(err error) {
	log.Println("error:", err.Error())
}

func (scbt StartCallBackTest) OnSuccess(files []*entity.File) {
	log.Println("on success:", len(files))
}
func main()  {
	conf := &entity.FileServerConf{
		Port: "8081",
		Route: "/staticfile/",
		FilePath: "./staticfile",
	}
	startCallBacktest := StartCallBackTest{}
	fileserver.StartFileServer(conf, startCallBacktest)
}
```
### 作者联系方式：QQ：975804495
### 疯狂的程序员群：186305789，没准你能遇到绝影大神
