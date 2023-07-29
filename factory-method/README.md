# FACTORY METHOD
## Review
Factory method adalah pola dimana kita akan membuat sebuah factory yang bisa membuat sebuah objek. Namun daripada kita membuat sebuah objek
yang spesifik dengan 1 factory, kita bisa membuat banyak factory turunan dari factory utama yang bisa menghasilkan objek yang berbeda-beda.

## Alur Pembuatan
Pada dasarnya kita akan banyak bermain dengan interface di sini. Jadi kita harus membuat interface untuk factory dan untuk objek produksinya.
maka alurnya akan menjadi seperti ini:
1. Buat interface untuk objek yang akan diproduksi
2. Buat struct (atau class) yang mengimplementasikan interface tersebut
3. Buat interface untuk factory
4. Buat method produksi di interface factory yang akan mereturn interface objek yang kita buat tadi
5. Buat struct (atau class) yang mengimplementasikan interface factory
6. Pada method produksi di struct ini, kita bebas menggunakan struct objek produksi yang mana saja (yang penting sudah mengimplementasikan interface objek produksi)
7. Untuk membuat objek baru, kita tinggal membuat variabel dari struct factory mana saja yang kita inginkan dan tinggal memanggil fungsi produksi yang ada
