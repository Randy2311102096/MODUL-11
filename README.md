1.Program membaca satu baris input menggunakan bufio.Scanner.
Input dipisahkan menjadi array string menggunakan strings.Fields().
Program memproses setiap vote, mengkonversi ke integer, dan memvalidasi (1-20).
Vote yang valid disimpan dalam slice validVotes dan dihitung dalam map voteCount.
Jika ditemukan angka 0, program berhenti memproses input lebih lanjut.
Program mencetak jumlah suara masuk (total input - 1) dan jumlah suara sah.
Terakhir, program mencetak jumlah suara untuk setiap kandidat yang mendapat suara.

2.Program ini sederhana namun efektif untuk menghitung suara dan menentukan pemenang dalam pemilihan. Dengan memvalidasi input dan menggunakan struktur data yang tepat, program ini dapat memberikan hasil yang akurat.

3.Program ini efektif untuk mencari posisi sebuah bilangan dalam array, dengan kemampuan menangani input hingga 1 juta elemen.
