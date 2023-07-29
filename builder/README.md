# BUILDER
## Review
Builder adalah pattern yang memungkinkan kita untuk membuat sebuah objek tanpa harus memikirkan komponennya diawal inisialisasi. Jadi pada pola ini nantinya akan dibuatkan beberapa method yang digunakan untuk mengubah komponen-komponennya.

## Alur Pembuatan
Komponen utama pada pola ini adalah Builder dan Objek yang akan dibuild. Sehingga alurnya akan menjadi seperti berikut:
1. Buat struct yang ingin dilakukan proses building
2. Buat struct Builder dimana struct ini akan menyimpan objek dari struct yang kita buat sebelumnya
3. Buat method untuk melakukan modifikasi pada atribut dari objek Building tadi
4. Buat method reset untuk menghapus semua perubahan pada objek Building jika diperlukan
5. Buat method generate untuk mendapatkan objek yang sudah kita build sebelumnya
6. Untuk menggunakan builder, kita hanya perlu membuat objek builder dan memanggil setiap method untuk menambahkan komponen. Setelah semua komponen yang dibutuhkan sudah ditambahkan, maka panggil method generate yang akan mereturn objek hasil dari proses building