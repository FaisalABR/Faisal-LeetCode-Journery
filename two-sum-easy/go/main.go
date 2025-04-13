package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if result != nil {
		println("Index ke-", result[0], "dan", result[1])
	} else {
		println("Tidak ada pasangan yang ditemukan")
	}

	resultHashMap := twoSumHashMap(nums, target)
	if resultHashMap != nil {
		println("Index ke-", resultHashMap[0], "dan", resultHashMap[1])
	} else {
		println("Tidak ada pasangan yang ditemukan")
	}

}

// Brute force
func twoSum(nums []int, target int) []int {
	// buat nested loop
	// outer loop untuk brute force indeks pertama
	for i := 0; i < len(nums); i++ {
		// inner loop untuk brute force indeks kedua
		for j := i + 1; j < len(nums); j++ {
			// total kan dan cek apakah nilai indeks j dan i == target
			if nums[i]+nums[j] == target {
				// assign nilai i dan j ke variable result
				return []int{i, j}
			}
		}
	}
	return nil
}

// HashMap
func twoSumHashMap(nums []int, target int) []int {
	// inisiasi map
	seen := make(map[int]int)
	// looping elemen pada array
	for i, _ := range nums {
		// hitung sisa target yang dikurangi dengan elemen ke i
		sisa := target - nums[i]
		// cek apakah sisa ada dimap
		_, ok := seen[sisa]
		if ok {
			// jika ada maka kita sudah pernah melihat pasangannya
			// return int{seen[sisa], i}
			return []int{seen[sisa], i}
		} else {
			// jika belum, masukan nums ke-i ke map dengan nums key-i sebagai key, indeks sebagai value
			seen[nums[i]] = i
		}
	}
	// return niil
	return nil

}
