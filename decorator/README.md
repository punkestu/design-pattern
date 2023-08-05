# DECORATOR
## Review
Decorator dapat digunakan untuk membentuk sebuah struktur berantai dari sebuah interface. Jadi daripada kita menjalankan setiap operasi satu per satu, lebih baik kita siapkan sebuah struktur berlapis dimana nantinya akan menjalankan operasi pada setiap lapisan sehingga operasinya dapat terjadi secara berantai hingga akhir dari lapisan struktur.

## Alur Pembuatan
Untuk membuat decorator, kita hanya tinggal membuat struktur yang mengimplementasikan sebuah interface dan mengkomposisi interface itu juga di dalamnya. Untuk membuat base decorator juga sama, bedanya kita tidak perlu mengkomposisi interface di dalam struct tersebut. Dengan begitu maka alurnya akan menjadi:
1. Buat interface decorator dan buat base decorator yang mengimplementasikan interface tersebut
2. Buat struct decorator yang akan mengimplementasikan interface decorator
3. Simpan komposisi interface di dalam struct tersebut
4. Pada method yang mengimplementasikan method interface, jalankan method yang sama dengan objek komposisi yang disimpan sehingga nanti method pada objek tersebut akan terpanggil ketika method pada decoratornya dipanggil
5. Dengan begitu maka untuk menggunakan decorator kita tinggal membuat objek base decoratornya, Lalu kita bungkus dengan objek dari struct decorator yang kita buat dengan menginject objek base ke dalamnya
6. Dengan begini jika kita memanggil method yang diimplementasikan dari interface, maka method dari objek yang diwrap juga akan dipanggil secara berantai