# Golang MAP

map adalah tipe data asosiatif yang ada di GO, berbentuk _key-value pair_, untuk setiap data (value) yang disimpan, harus juga disiapkan _key_-nya, _key_ harus unik karena digunakan sebagai penanda untuk mengakses value.

jika dilihat `map` seperti tipe data object yang ada di JavaScript, dimana setiap element di `map` berisi `key` dan `value`.

Cara penggunaannya cuku dengan menuliskan keyword `map` diikuti tipe data `key` dan tipe data `value`-nya.

```go
var moon map[byte]string

moon = map[byte]string{}
```

maksud dari kode `map[byte]string` tersebut adalah, tipe data _map_ dengan `byte` sebagai tipe data dari _key_ nya dan `string` sebagai tipe data _value_-nya.

default nilai dari tipe data _map_ adalah _nill_, oleh karena itu perlu di inisialisasi nilai default di awal, caranya dengan menambahkan tanda kurung kurawal di akhir tipe nya, contoh-nya seperti `map[byte]string{}`.

Untuk mendeklarasikan variabel bertipe _map_ selain dengan cara di atas bisa juga dengan menggunakan keyword `make` dan `new`

```go
var moons = make(map[byte]string)
var days = new(map[byte]string)
```

khusus untuk inisialisasi data dengan keyword `new`, yang dihasilkan berupa _pointer_, untuk mengambil nilai aslinya bisa dengan menggunakan tanda asterisk (`*`).

Untuk meng-set nilai pada sebuah nilai dengan cara menuliskan variabel-nya, kemudian disisipkan _key_ pada kurung siku, lalu di isi nilainya, contoh-nya seperti berikut.

```go
var moon = map[byte]string{}

moon[1] = "Januari"
moon[2] = "februari"
```

pengisian data pada sebuah map bersifat **overwrite**, ketika sebuah variabel sudah memiliki item dengan key yang sama, maka value lama akan ditimpa dengan value baru.

variabel dengan tipe data _map_ isinya juga bisa dideklarasikan atau didefinisikan di awal, caranya dengan menambahkan kurung kurawal setelah tipe datanya, didalam kurung kurawal lalu dimasukkan _key_ dan _value_.

_key_ dan _value_ dipisahkan dengan menggunakan (`:`) sedangkan tiap item-nya dipisahkan dengan menggunakan `,`. 
Pengisian-nya bisa dilakukan dengan cara _horizontal_ maupun _vertikal_, khusus untuk _vertikal_ tanda koma (`,`) wajib dituliskan setelah _value_ nya meskipun itu _value_ terakhir.

```go
var moon = map[byte]string{
  1:"januari",
  2:"februari",
  3:"maret",
  4:"april",
  5:"mei",
  6:"juni",
  7:"juli",
  8:"agustus",
  9:"september",
  10:"oktober",
  11:"november",
  12:"desember",
}
```

untuk mengakses atau mengambil data dari item-nya bisa dilakukan seperti pada `slice`, 

```go
var data = map[string]int{
  "a":1,
  "b":2,
  "c":3,
  "d":4,
}

fmt.Println(data["a"]) // 1
fmt.Println(data["b"]) // 2
```

atau bisa juga dilakukan dengan menggunakan keyword `for` - `range`, data yang dikembalikan di tiap perulangan berupa `key` dan `value`, bukan `indeks` dan `element`.

```go
var data = map[string]int{
  "a":1,
  "b":2,
  "c":3,
  "d":4,
}

for key, value := range data {

  fmt.Printf("key-%s : %d\n",key,value)
}
```

lalu untuk menghapus element didalam `map` dengan cara menggunakan fungsi `delete()`, untuk cara penggunaannya dengan memasukkan object `map` dan `key` untuk parameter dari fungsi `delete()`.

```go
var data = map[int]string{}

data[1] = "a"
data[2] = "b"

delete(data,2)
```

item yang memiliki key _2_ dalam variabel _data_ akan dihapus, sebaliknya jika key yang akan dihapus tidak ada di dalam variabel maka akan diabaikan.

Kadang saat ingin menghapus sebuah item dalam sebuah map, kita ingin mengecek apakah item tersebut ada atau tidak

untuk mengetahui apakah dalam sebuah map terdapat item dengan key tertentu atau tidak, yaitu dengan memanfaatkan 2 variabel sebagai penampung nilai pengembalian pengaksesan item, return value ke 2 adalah opsional berupa `bool` yang menunjukkan ada atau tidak-nya sebuah item berada dalam sebuah map tertentu.

```go
var workDays = map[byte]string{
  1:"senin",
  2:"selasa",
  3:"rabu",
  4:"kamis",
  5:"jumat",
}

const value, isExist = workDays[5]

if isExist {
  fmt.Println(value) // "jumat"
}
```

dari contoh di atas, jika variabel `workDays` tersebut memiliki key `5` maka variabel `value` berisi value dengan key `5` dan `isExist` berisi `bool` dengan isi `true`, sebaliknya jika item yang dicari tidak ada maka variabel pertama berisi default value dari tipe data value di map, dan variabel pertama berisi `bool` dengan value `false`.

Selain itu `map` juga bisa dikombinasikan dengan `slice`, dan cara penggunaannya cukup mudah, contohnya seperti ini `[]map[byte]string` artinya adalah sebuah slice yang tiap-tiap item-nya berupa `map[byte]string`.

```go
var days = []map[byte]string{
  map[byte]string{1:"senin",0:"hari kerja"},
  map[byte]string{2:"selasa",0:"hari kerja"},
  map[byte]string{3:"rabu",0:"hari kerja"},
  map[byte]string{4:"kamis",0:"hari kerja"},
  map[byte]string{5:"jumat",0:"hari kerja"},
  map[byte]string{6:"sabtu",0:"hari kerja"},
  map[byte]string{7:"minggu",0:"hari libur"},
}

for i,element := range days{
  fmt.Println()
  fmt.Printf("index ke-%d berisi : %v\n",i,element)

  for key, value := range element {
    fmt.Printf("key map-%d : %s\n",key,value)
  }
}
```

variabel `days` diatas berisikan informasi bertipe `map[byte]string` dimana tiap element mempunyai 2 key yang sama, dan jika menggunakan GO versi terbaru, cara pendeklarasian `slice` `map` bisa dipersingkat dengan cara tiap tiap element tidak perlu dituliskan tipe data-nya.

```go
var days = []map[byte]string{
  {1:"senin",0:"hari kerja"},
  {2:"selasa",0:"hari kerja"},
  {3:"rabu",0:"hari kerja"},
  {4:"kamis",0:"hari kerja"},
  {5:"jumat",0:"hari kerja"},
  {6:"sabtu",0:"hari kerja"},
  {7:"minggu",0:"hari libur"},
}
```

dan didalam sebuah `[]map[byte]string` tiap element bisa saja memiliki `key` yang berbeda-beda.