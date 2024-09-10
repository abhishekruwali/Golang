package main

func main() {
	//  file create work is of os

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error ehile creating file", err)
	// 	return
	// }
	// defer file.Close() // essential because the resource which are used to create file should be free

	// content := "hello world by prince"
	// _, errors := io.WriteString(file, content+"\n")
	// byte, error := io.WriteString(file,"content")
	// if err != nil {
	// 	fmt.Println("Error while writing a file", errors)
	// }

	// fmt.Println("file is created")
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }
	// defer file.Close()

	// // create a buffer to read the file content
	// buffer := make([]byte, 1024)

	// // Read the file content into the buffer
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error while reading a file", err)
	// 		return
	// 	}
	// 	// Process the read content
	// 	fmt.Println(string(buffer[:n]))

	// }
	// Read the file content using function
	// content, err = ioutil.ReadFile("example.txt")
	// if err !=nil{
	// 	fmt.Println("Error while reading")
	// }

}
