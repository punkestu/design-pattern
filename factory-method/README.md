# FACTORY METHOD
## Review
Factory method adalah pola dimana kita akan membuat sebuah factory yang bisa membuat sebuah objek dimana factory tersebut bisa memproduksi sebuah objek yang sama namun dengan berbagai varian.

## Alur Pembuatan
Pada dasarnya kita akan banyak bermain dengan interface di sini. Jadi kita harus membuat interface untuk objek yang ingin diproduksi sehingga kita bisa membuat varian dari produk tersebut.
maka alurnya akan menjadi seperti ini:
1. Buat interface untuk objek yang akan diproduksi
2. Buat implementasi interface tersebut (bisa lebih dari 1)
3. Buat struct untuk factory dengan method untuk produksi yang akan mereturn interface objek yang ingin diproduksi tadi.
6. Pada method produksi ini, kita bebas menggunakan struct objek produksi yang mana saja yang penting sudah mengimplementasikan interface objek produksi
7. Untuk membuat objek baru, kita tinggal membuat variabel dari struct factory dan tinggal memanggil fungsi produksi yang ada dengan menambahkan parameter untuk memilih varian
