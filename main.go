package main

func main() {
	folder := make(map[string]string)
	//正向隔离
	folder["E:\\root\\oneArea\\send\\request"] = "E:\\root\\safeArea\\receive\\request"
	folder["E:\\root\\oneArea\\send\\response"] = "E:\\root\\safeArea\\receive\\response"
	//反向隔离
	folder["E:\\root\\safeArea\\send\\request"] = "E:\\root\\oneArea\\receive\\request"
	folder["E:\\root\\safeArea\\send\\response"] = "E:\\root\\oneArea\\receive\\response"
	forever := make(chan bool)
	go syncDirs(folder)
	<-forever
}
