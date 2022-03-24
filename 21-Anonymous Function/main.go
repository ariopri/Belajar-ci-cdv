package main

import "fmt"

/*
	Anonymous Function -> Sebelumnya setiap membuat function,
	kita harus menentukan nama function terlebih dahulu
	namun kadang ada kalanya lebih mudah membuat function secara langsung
	di variable atau parameter tanpa harus membuat function terlebih dahulu
	hal tersebut dinamakan Anonymous function, atau function tanpa nama
	bahasa awamnya function dalam functuiin
	atau kita biikin varial ntr valuenya itu Function
	bisa juga dimasukan langsung ke dalam parameter
	contoh :

*/

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) { // -> jika benar maka akan di return true ini tuh boolean
		fmt.Println("User is blacklisted")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		// bisa dilihat tidak ada functionya makanya di panggil anonymous function
		return name == "Admin"
		/*ini kan kita masukan namaya Adimin
		jika benar maka dia akan masukan ke dalam kondisi true
		maka dia akan ngeprint "User is blacklisted"
		jika salah maka dia akan masukan ke dalam kondisi false
		maka dia akan ngeprint "Welcome name"
		*/
	}
	RegisterUser("Admin", blacklist)
	RegisterUser("Yuji", blacklist)
	/* 	maka hasilnya adalah User is blacklisted
	dan Welcome Yuji
	*/
	//Pengen langsung
	RegisterUser("admin", func(name string) bool {
		return name == "Admin"
	})
	/*
		Jadi disini di dalam RegisterUser()
		di dalam tanda kurung itu kita bikin function

		"name", func(name string) bool {
			return name == "Admin"
		}
	*/

}
