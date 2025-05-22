package main
import "fmt"
const NMAX int = 100
type atlet struct{
	nama string
	gol int
	assist int
}
type arrPemain [NMAX] atlet
func main(){
	var A, temp arrPemain
	var n, i, cetakGol, cetakAssist, idxGol, j, idx, idxAssist  int
	var namaAtlet string
	fmt.Scan(&n)
	for i =0; i < n; i++{
		fmt.Scan(&namaAtlet, &cetakGol, &cetakAssist)
		A[i].nama = namaAtlet
		A[i].gol = cetakGol
		A[i].assist = cetakAssist
	}
	for i =0; i < n; i++{
		idx = i
		idxGol = A[i].gol
		idxAssist =A[i].assist
		for j = i+1; j < n; j++{
			if A[j].gol > idxGol || (A[j].gol == idxGol && A[j].assist > idxAssist) {
				idx = j
				idxGol = A[j].gol
				idxAssist = A[j].assist
			}
		}
		temp[0] =A[i]
		A[i] = A[idx]
		A[idx] = temp[0]
	}
	for i = 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].gol, A[i].assist)
	}
	fmt.Println()
}

