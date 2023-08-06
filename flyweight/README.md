# FLYWEIGHT
## Review
Flyweight adalah pola penyusunan struktur dimana nantinya jika beberapa objek memiliki beberapa komponen yang bernilai 100% sama (dan bukan merupakan data yang mudah berubah), maka lebih baik komponen tersebut diekstraksi menjadi struktur baru dan disimpan sebagai komposisi saja.

## Alur Pembuatan
Pada dasarnya cara pembuatannya sama seperti struktur yang mempunyai komposisi pada umumnya. Tapi disini flyweight hanya mengajarkan kita tentang komponen apa saja yang perlu diekstrak lalu dijadikan komposisi. Ciri-cirinya antara lain:
1. Nilainya tidak pernah berubah pada semua objek yang memiliki
2. Nilainya punya variasi tapi tidak banyak
3. Nilainya mudah berubah tapi masih dapat digunakan oleh banyak objek

Jika muncul komponen dengan kondisi seperti di atas ketika membuat sebuah struktur, maka lebih baik komponen tersebut diekstrak dan dijadikan struktur tersendiri sehingga dapat menghemat penggunaan memori.