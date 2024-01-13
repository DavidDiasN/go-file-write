package main

const (
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Oi, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name

}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case "French":
		{
			prefix = frenchHelloPrefix
		}
	case "Spanish":
		{
			prefix = spanishHelloPrefix
		}
	case "Portuguese":
		{
			prefix = portugueseHelloPrefix
		}
	default:
		{
			prefix = englishHelloPrefix
		}
	}
	return prefix
}

func main() {

	// fmt.println(hello("mom"))
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
