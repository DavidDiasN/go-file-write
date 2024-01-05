package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(s string) string {
	if s == "" {
		s = "World"
	}
	return englishHelloPrefix + s
}

func main() {

	fmt.Println(Hello("mom"))
	// TestHello(&testing.T{})

	// var mine string = "Hello FileSystem!"
	// fmt.Println(mine)
	// file, err := os.Open("file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Read %d bytes: %q\n", count, data[:count])

	// fmt.Println("Type something")
	// scanner := bufio.NewScanner(os.Stdin)
	// var input string
	// if scanner.Scan() {

	// 	input = scanner.Text()
	// }

	// err = os.WriteFile("test.txt", []byte(input), 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }

}
