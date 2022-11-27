package main
import "fmt"

func main(){
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)
	fmt.Println(codes["en"])

	//value, found := codes["en"]
	//fmt.Println(value, found)

	//value, found := codes["hh"]
	//fmt.Println(value, found)

	code := make(map[string]int)
	fmt.Println(code)

	codes["de"] = "German"
	codes["en"] = "English Language"
	delete(codes, "fr")
	fmt.Println(codes)
	fmt.Println(len(codes))

	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}
	//* truncate a map - 1
	// for key := range codes {
	// 	delete(codes, key)
	// }
	// fmt.Println(codes)

	//* truncate a map -2
	codes = make(map[string]string)
	fmt.Println(codes)
}