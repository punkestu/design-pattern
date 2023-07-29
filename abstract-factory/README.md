# ABSTRACT FACTORY
## Review
Abstract factory adalah pola dimana kita akan membuat beberapa factory yang bisa membuat sebuah objek berbeda dengan factory method yang hanya menggunakan 1 method factorynya.

## Alur Pembuatan
Pada dasarnya kita akan banyak bermain dengan interface di sini. Jadi kita harus membuat interface untuk factory dan untuk objek produksinya.
maka alurnya akan menjadi seperti ini:
1. Buat beberapa interface untuk objek yang akan diproduksi pada tiap factory
2. Buat struct (atau class) yang mengimplementasikan interface tersebut
3. Buat interface untuk factory
4. Buat method produksi untuk tiap interface objek yang kita buat pada interface factory
5. Buat struct (atau class) yang mengimplementasikan interface factory
6. Pada method produksi di struct ini, kita bebas menggunakan struct objek produksi yang mana saja (yang penting sudah mengimplementasikan interface objek produksi yang digunakan pada factory ini)
7. Untuk membuat objek baru, kita tinggal membuat variabel dari struct factory mana saja yang kita inginkan dan tinggal memanggil fungsi produksi yang ada
