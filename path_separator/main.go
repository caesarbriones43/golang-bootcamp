package main
import ("fmt")
import ("path")


func main(){
	//_,file := path.Split("css/main.css")
	dir,file := path.Split("css/main.css")
	fmt.Println("dir:",dir)
	fmt.Println("file:",file)

}