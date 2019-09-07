package main 

import "os"

func main() {
	buffer := make([]byte, 10)
	f, _ := os.Open("E:\\gopath\\src\\youaresherlock.net\\test.json")
	defer f.Close()

	for {
		n, _ := f.Read(buffer)
		if n == 0 {
			break 
		}
	}

	os.Stdout.Write(buffer)
}