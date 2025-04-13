### 🔢 Judul Soal: [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

#### 📌 Deskripsi Ulang

Diberikan 2 linked list tidak kosong merepresentasikan 2 bilangan bulat tak negatif. Digit disimpan pada urutan terbalik.
dan setiap nodes mengandung satu digit. Tambahkan 2 angka dan kembalikan total seperti linked list.

#### 🧠 Contoh Kasus

Input: l1 = [2,4,3] dan l2 = [5,6,4]
Output: [7,0,8]

Penjelasan: 243 + 564 = 807 -> dibalik jadi 708

#### 🔎 Ide Solusi / Strategi

- Singly Linked List Traversal

#### 📝 Pseudocode

- Brute Force \* **timecomplexity O(n^2)**

1. Buat dummy head node sebagai awal list hasil
2. Buat pointer curr dan variabel carry = 0
3. Loop selama l1 != nil || l2 != nil || carry > 0:
   a. Ambil val1 dari l1, val2 dari l2 (jika nil, pakai 0)
   b. total = val1 + val2 + carry
   c. carry = total / 10
   d. curr.Next = node dengan nilai (total % 10)
   e. curr = curr.Next
   f. l1 = l1.Next, l2 = l2.Next
4. return dummy.Next

#### 💻 Kode Final

Cek Directory tiap bahasa
