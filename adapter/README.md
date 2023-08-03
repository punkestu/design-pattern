# ADAPTER
## Review
Adapter berguna untuk membuat sebuah struct atau class berubah bentuk menjadi mirip dengan class yang lain. Dengan begitu walaupun struct yang ingin kita gunakan tidak sesuai dengan yang dibutuhkan, kita tetap bisa menggunakannya asalkan struct tersebut memiliki komponen yang bisa digunakan sebagai komponen class yang diinginkan.

## Alur Pembuatan
Pada dasarnya pola ini dapat digunakan ketika kita memiliki 2 struct atau lebih yang memiliki sifat yang berbeda namun kita ingin semua struct tersebut dapat diperlakukan dengan method yang sama. Sehingga kita hanya perlu menyiapkan 1 atau lebih struct adapter menyesuaikan struct apa saja yang ingin diubah sifatnya. Maka alurnya akan menjadi seperti berikut:
1. Buat sebuah struct yang ingin diadapt
2. Buat sebuah struct adapter yang akan menyimpan struct yang ingin diadapt tadi
3. Tambahkan method menyesuaikan dengan sifat interface yang diinginkan pada struct adapter dengan mengoperasikan komponen yang ada pada struct yang diadapt tadi
4. Dengan ini maka jika kita ingin menggunakan struct yang kita buat tadi sebagai parameter dengan interface yang ditentukan, kita tinggal membuat objek adapter dan menyimpan objek yang ingin diadapt sebagai parameter