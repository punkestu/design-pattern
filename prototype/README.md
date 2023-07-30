# PROTOTYPE
## Review
Prototype dapat digunakan ketika kita ingin men-copy suatu objek tapi hasil copy tersebut sudah tidak bisa digunakan untuk mengubah nilai di dalam objek yang dicopy (hanya men-copy nilainya saja). Hal ini dilakukan karena jika kita men-copy sebuah objek secara langsung, maka objek tersebut akan ikut berubah jika kita mengubah hasil dari copy-annya.

## Alur Pembuatan
Untuk pembuatan fungsi copy ini, kita harus menanamkannya langsung ke dalam class atau struct yang ingin kita copy (atau bisa dengan meletakkannya dalam 1 package). Ini harus dilakukan karena jika ada kondisi dimana kita juga harus men-copy nilai yang private, maka jika kita membuat fungsinya di scope lain kita tidak dapat men-copy nilai private tersebut. Maka alur pembuatannya akan menjadi seperti ini:
1. Buat sebuah interface (jika diperlukan) dengan fungsi Copy di dalamnya yang akan mereturn interface itu sendiri
2. Buat struct dengan mengimplementasikan interface tersebut
3. Pada method Copy di struct, return object baru dimana atribut yang digunakan dicopy sepenuhnya dari struct itu sendiri.
4. Dengan mekanisme ini, untuk memproduksi objek baru dengan Copy maka kita hanya perlu membuat sebuah objek baru. 
5. Ketika ingin mulai memproduksi, kita hanya perlu memanggil method Copy dan menampung nilainya dalam sebuah variabel