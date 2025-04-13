### 🔢 Judul Soal: [Two Sum](https://leetcode.com/problems/two-sum/)

#### 📌 Deskripsi Ulang

Diberikan sebuah array integers dengan nama variabels `nums` dan sebuah integer dengan variabel `target`.
Mengembalikan 2 indeks yang dimana saat ditambah nilai dari 2 indeks tersebut, hasilnya sama dengan nilai `target`.

#### 🧠 Contoh Kasus

Input: nums = [3,2,4], target = 6  
Output: [1,2]

#### 🔎 Ide Solusi / Strategi

- Brute Force
- Hash Map

#### 📝 Pseudocode

- Brute Force \* **timecomplexity O(n^2)**

1. Buat Nested Loop
2. Outer loop untuk mengecek semua angka untuk indeks pertama
3. Inner loop untuk mengecek semua angka untuk indeks kedua
4. Totalkan dan cek apakah total dari nilai indeks pertama dan kedua sama dengan nilai target
5. Jika iya kembalikan hasilnya dalam bentuk array integers

- HashMap \* **timecomplexity O(n)**

1. Inisiasi map
2. Iterasi array
3. hitung sisa = target - nums[i]
4. cek apakah sisa ada di dalam map
5. Jika iya, maka return [seen[sisa], i]
6. Jika tidak, maka masukan nums[i] ke map, dengan nums[i] sebagai key, indeks sebagai valuenya

#### 💻 Kode Final

Cek Directory tiap bahasa
