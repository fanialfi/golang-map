package main

import (
	"fmt"
)

func main() {
	// deklarasi map
	var chicken1 map[byte]string
	chicken1 = map[byte]string{} // digunakan untuk inisialisasi nilai map dengan nilainya default dari type nya

	chicken2 := map[byte]string{}
	chicken3 := make(map[byte]string)
	chicken4 := new(map[byte]string)

	// inisialisasi nilai map
	chicken1[1] = "a"
	chicken1[2] = "b"

	chicken2[1] = "januari"
	chicken2[2] = "februari"

	chicken3[1] = "a"
	chicken3[2] = "b"
	chicken3[1] = "c"
	chicken3[3] = "d"

	fmt.Println(chicken1, chicken2, chicken3, *chicken4) // mengambil nilai-nya secara langsung

	// mengiterasi nilai-nya dengan menggunakan keyword for - range
	for key, value := range chicken1 {
		fmt.Println(key, ":", value)
	}

	for key, value := range chicken2 {
		fmt.Println(key, ":", value)
	}

	for key, value := range chicken3 {
		fmt.Println(key, ":", value)
	}

	// menghapus element map
	fmt.Println()
	fmt.Println(len(chicken3))
	delete(chicken3, 4)
	fmt.Println(len(chicken3))

	for key, value := range chicken3 {
		fmt.Println(key, ":", value)
	}

	// deteksi keberadaan item dengan menggunakan key

	var val, key = chicken1[1]

	if key {
		fmt.Println(val, key)
	} else {
		fmt.Println("Item Not Exists")
	}

	// kombinasi slice dan map
	var data = make([]map[string]string, 5)

	data[0] = map[string]string{"nama": "fani alfirdaus", "alamat": "indonesia"}
	data[1] = map[string]string{"nama": "fani", "alamat": "indonesia"}
	data[2] = map[string]string{"nama": "alfi", "alamat": "indonesia"}
	data[3] = map[string]string{"nama": "dea", "alamat": "indonesia"}
	data[4] = map[string]string{"nama": "amel", "alamat": "indonesia"}

	fmt.Println(len(data), cap(data))
	fmt.Println()

	for i, isi := range data {
		fmt.Printf("index ke-%d isinya %v\n", i, isi)
	}

	{
		// mengecek keberadaan item dengan key tertentu
		var workDays = map[string]byte{
			"senin":  1,
			"selasa": 2,
			"rabu":   3,
			"kamis":  4,
			"jumat":  5,
		}

		var day, isExist = workDays["minggu"]
		if isExist {
			fmt.Println("hari bekerja :", day)
		} else {
			fmt.Println(day)
		}
	}

	fmt.Println()

	{
		var days = []map[byte]string{
			map[byte]string{1: "senin", 0: "hari kerja"},
			map[byte]string{2: "selasa", 0: "hari kerja"},
			map[byte]string{3: "rabu", 0: "hari kerja"},
			map[byte]string{4: "kamis", 0: "hari kerja"},
			map[byte]string{5: "jumat", 0: "hari kerja"},
			map[byte]string{6: "sabtu", 0: "hari kerja"},
			map[byte]string{7: "minggu", 0: "hari libur"},
		}

		for i, element := range days {
			fmt.Println()
			fmt.Printf("index ke-%d berisi : %v\n", i, element)

			for key, value := range element {
				fmt.Printf("key map-%d : %s\n", key, value)
			}
		}
	}
}
