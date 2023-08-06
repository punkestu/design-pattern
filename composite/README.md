# COMPOSITE
## Review
Composite sedikit mirip dengan decorator. Persamaannya adalah decorator dan composite sama-sama dapat memanggil sebuah method secara berantai. 

Perbedaannya adalah decorator ada sebuah objek yang menjadi basis dan dibungkus oleh objek decorator yang hanya akan menambah fungsi pada method dari objek basisnya. Sedangkan pada composite setiap objek punya independensi sendiri jadi walaupun akhirnya akan membungkus objek lain, objek tersebut akan tetap berlaku seperti objek biasa.

## Alur Pembuatan
Karena composite mirip dengan decorator, maka pembuatannya juga akan mirip dengan decorator. Hanya saja jika pada decorator kita hanya menyimpan 1 objek yang akan dibungkus, pada composite kita bisa menyimpan banyak objek sebagai child. Sehingga alur pembuatannya akan menjadi seperti berikut:
1. Buat interface untuk node yang akan dibuat
2. Buat implementasi dari interface tersebut dan tambahkan array interface node sebagai children di dalam struct atau classnya
3. Untuk method yang bisa berjalan secara berantai, lakukan iterasi pada setiap child dan panggil method yang sama pada setiap child
4. Dengan begini kita hanya perlu membuat objek-objek node dan simpan children ke dalam objek yang menjadi parent
5. Dan untuk memanggil method berantai, kita hanya perlu memanggil method yang diinginkan