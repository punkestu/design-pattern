# ABSTRACT FACTORY
## Review
Abstract factory mirip dengan factory method. Tapi bedanya jika pada factory method kita hanya membuat factory yang menghasilkan sebuah produk dengan berbagai varian, pada abstract factory kita bisa membuat beberapa varian factory yang bisa menghasilkan banyak jenis dan varian produk. Misal kita akan membuat sekelompok factory yang bisa menghasilkan produk A dan produk B, maka pada setiap factory juga bisa menghasilkan varian produk A dan produk B. Dan karena setiap factory bisa berdiri sendiri, maka produk yang dihasilkan dari setiap factory bisa saja berbeda 1 dengan yang lain.

## Alur Pembuatan
Pada dasarnya kita akan banyak bermain dengan interface di sini. Jadi kita harus membuat interface untuk factory dan untuk objek produksinya.
maka alurnya akan menjadi seperti ini:
1. Buat beberapa beberapa interface untuk objek yang akan diproduksi pada tiap factory
2. Buat struct (atau class) yang mengimplementasikan interface-interface tersebut
3. Buat interface untuk factory
4. Buat beberapa method produksi pada interface factory
5. Buat struct (atau class) yang mengimplementasikan interface factory
6. Pada beberapa method produksi di struct ini sesuai dengan interface objek yang sudah kita buat yang akan mereturn objek sesuai dengan interface objek yang digunakan pada method tersebut.
7. Untuk membuat objek baru, kita tinggal membuat variabel dari struct factory mana saja yang kita inginkan dan tinggal memanggil fungsi produksi yang kita inginkan
