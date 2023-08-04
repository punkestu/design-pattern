# BRIDGE
## Review
Bridge dapat digunakan ketika kita mempunyai sebuah struktur yang implementasinya dapat menghasilkan banyak kombinasi. Jadi daripada membuat implementasinya satu per satu, lebih baik struktur tersebut dipecah menjadi komponen utama dan komponen pendukungnya. Dengan begitu, ketika kita ingin membuat kombinasi kita tidak perlu membuat struktur baru melainkan hanya dengan menggunakan 1 struktur yang sudah ada lalu komponen pendukungnya dikomposisikan ke dalam struktur yang sudah dibuat tadi.

## Alur Pembuatan
Inti dari pola ini sebenarnya hanya membuat komponen-komponen yang dapat kombinasikan satu sama lain sehingga menghindari pembuatan struktur yang berulang yang hanya berbeda pada kombinasinya saja. Sehingga alurnya akan menjadi seperti berikut:
1. Buat interface untuk struktur utama dan pendukung (agar kita bisa membuat berbagai varian struktur utama dan pendukung)
2. Implementasikan interface struktur utama dengan mengkomposisikan interface pendukung didalamnya
3. Implementasikan interface struktur pendukung
4. Untuk menggunakan pola ini kita tinggal membuat objek dari struktur pendukung lalu ketika kita membuat objek dari struktur utama, maka kita tinggal memasukkan objek pendukung yang sudah kita buat tadi sesuai dengan kebutuhan.
5. Dengan begitu maka kita tidak perlu membuat struct baru untuk setiap kombinasi melainkan hanya tinggal merubah komponen pendukung di dalam objek sesuai kombinasi yang diinginkan