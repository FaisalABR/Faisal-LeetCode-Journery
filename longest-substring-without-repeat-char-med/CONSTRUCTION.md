### 🔢 Judul Soal: [Two Sum](https://leetcode.com/problems/longest-substring-without-repeating-characters)

#### 📌 Deskripsi Ulang

Diberikan sebuah string `s`, cari panjang dari substring terpanjang tanpa ada pengulangan karakter

#### 🧠 Contoh Kasus

Input: string = "abcabcaa"
Output: 3

Penjelasan Karna substring "abc" merupakan substring terpanjang tanpa pengulangan karakter, dengan panjang karakter 3

#### 🔎 Ide Solusi / Strategi

- Brute Force
- Sliding window & Hash Map

#### 📝 Pseudocode

- Brute Force \* **timecomplexity O(n^3)**

1. Buat Nested Loop
2. Outer loop untuk mengecek semua angka untuk pointer kiri
3. Inner loop untuk mengecek semua angka untuk pointer kanan
4. Totalkan dan cek apakah total dari nilai indeks pertama dan kedua sama dengan nilai target
5. Jika iya kembalikan hasilnya dalam bentuk array integers

- HashMap \* **timecomplexity O(n)**

1. Inisiasi Map
2. Inisiasi pointer window left (untuk menunjukkan indeks yang ada dikiri) dan maxLength
3. Looping string
   a. cek apakah character ada di map?, jika ada maka geser pointer left ke setelah karakter yang sama.
   b. jika tidak maka masukan karakter ke map
   c. hitung panjang windowLength dengan menghitung selisih pointer window left dan right + 1
   d. jika windowLength lebih besar dari maxLength, maka perbaharui nilai maxLength dengan windowLength
4. mengembalikan nilai maxLength

#### 💻 Kode Final

Cek Directory tiap bahasa
