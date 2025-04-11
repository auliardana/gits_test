## Kompleksitas waktu dan ruang

### Tabel Perhitungan Kompleksitas Waktu

| Langkah | Fungsi / Operasi                | Deskripsi Operasi                                         | Waktu per Operasi | Frekuensi Eksekusi | Total Waktu |
|---------|----------------------------------|------------------------------------------------------------|--------------------|---------------------|--------------|
| 1       | `for _, char := range s`        | Loop setiap karakter                                       | O(1)               | n                   | O(n)         |
| 2a      | `append(stack, char)`           | Push ke stack (bracket buka)                              | O(1)               | ≤ n                 | ≤ O(n)       |
| 2b      | `stack[len(stack)-1]`           | Akses elemen paling atas (peek)                           | O(1)               | ≤ n                 | ≤ O(n)       |
| 2c      | `stack = stack[:len(stack)-1]`  | Pop dari stack (bracket tutup valid)                      | O(1)               | ≤ n                 | ≤ O(n)       |
| 3       | `if len(stack) == 0`            | Cek apakah semua bracket sudah ditutup                    | O(1)               | 1                   | O(1)         |


Total  Time Complexity: **O(n)**


### Tabel Perhitungan Kompleksitas Ruang (Space Complexity)

| Struktur Data     | Deskripsi Penggunaan                                                            | Ukuran Maksimum             | Kompleksitas Ruang |
|-------------------|----------------------------------------------------------------------------------|------------------------------|---------------------|
| `stack []rune`    | Menyimpan semua bracket buka yang ditemukan selama parsing                      | Maksimum sebanyak `n` elemen | **O(n)**            |
| `pairs map[rune]rune` | Map untuk mencocokkan bracket tutup dengan pasangannya (')' → '(', dst.)    | Selalu 3 pasang, tetap       | O(1)                |
| Loop variables    | Variabel sementara seperti `char` dan indeks loop                               | Tidak tergantung ukuran `n`  | O(1)                |

#### Maka total penggunaan ruang adalah:
- Dalam worst case (semua karakter adalah bracket buka), kita simpan semuanya dalam stack.
- Map tetap konstan.
- Maka:

Total Space Complexity:  **O(n)** 