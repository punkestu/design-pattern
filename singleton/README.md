# SINGLETON
## Review
Singleton adalah pola yang memastikan bahwa sebuah objek hanya dibuat 1 kali saja. Pola ini biasa digunakan saat kita membuat objek yang selalu dibuat ulang saat kita melakukan sebuah operasi. Jadi ketimbang membuat ulang objek pada setiap operasi, lebih baik kita buat satu kali dan dapat digunakan berkali-kali.

## Alur Pembuatan
Kunci dari singleton adalah bagaimana cara kita menghalangi pengguna untuk membuat ulang objek dari class atau struct tertentu. Cara yang umumnya digunakan adalah membuat constructor menjadi private dan hanya menggunakan method static untuk menginisialisasi objek diawal. Ketika method static tersebut dipanggil setelah objek sudah dinisialisasi, maka objek juga tidak akan dibuat ulang. Dengan begitu maka alurnya akan menjadi:
1. Buat sebuah class atau struct dan pastikan constructornya adalah private
2. Buat sebuah atribut pointer objek static berdasarkan class tersebut di dalam class itu sendiri (jika menggunakan Go maka buat saja sebuah objek dengan pointer struct)
3. Buat sebuah method static untuk menginisialisasi objek
4. Buat method getter (jika memang diperlukan)
5. Buat method static yang diperlukan untuk mengoperasikan objek static yang kita buat sebelumnya.
6. Dengan seperti ini untuk menggunakannya kita hanya perlu memanggil method static untuk inisialisasi objek, maka objek singleton akan otomatis terbentuk
7. Jika ingin melakukan operasi dengan objek tersebut, maka kita hanya perlu memanggil method-method static yang kita buat sebelumnya