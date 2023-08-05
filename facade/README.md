# FACADE
## Review
Facade sesuai namanya akan menjadi eksterior dari sekumpulan program. Jadi daripada kita mengakses semua library dan class secara langsung untuk melakukan sebuah operasi yang sebenarnya dapat dikelompokan menjadi 1 class, lebih baik kita abstraksikan operasi yang berhubungan tersebut menjadi method-method dalam sebuah class sehingga kita juga bisa menyimpan dependensi pada class tersebut.

## Alur Pembuatan
Pada dasarnya Facade hanya sebuah abstraksi operasi namun disusun menjadi sebuah struktur. Dengan begitu maka alurnya akan menjadi:
1. Buat sebuah class atau struct untuk menyimpan dependensi yang dibutuhkan
2. Pada constructor inisialisasi tiap dependensi atau bisa dengan inject dependensi
3. Buat method-method abstraksi untuk mengelola operasi-operasi yang akan dikerjakan
4. Buat parameter sesimpel mungkin
5. Dengan begitu maka untuk menggunakan facade kita hanya perlu membuat objek facade dan tinggal memanggil method-method abstraknya