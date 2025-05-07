func isiJurnal(A *tabTidur, n *int){
/*I.S terdapat array kosong sebagai tempat untuk menyimpan durasi tidur
F.S array durasi tidur terisi oleh jam tidur, menit tidur, jam bangun, menit bangun selama n hari */
	
	var i int
	var tSiang bool
	//pengondisian apabila n melebihi NMAX
	if *n > NMAX{
		*n = NMAX
	}

	for i = 0; i < *n; i++{
		fmt.Println("Hari ke-", i++)
		fmt.Print("Apakah anda tidur siang?(true/false) ");fmt.Scan(&tSiang)

		//Pengondisian apakah pengguna melakukan tidur siang atau tidak
		
		if tSiang == true{
		//Tidur siang
		fmt.Println("Silahkan isi jurnal dengan format (jam tidur siang, menit tidur siang, jam bangun siang, menit bangun siang)")
		fmt.Scan(&A[i].jTidurSiang, &A[i].mTidurSiang, &A[i].jBangunSiang, &A[i].mBangunSiang)

		//Tidur malam
		fmt.Println("Silahkan isi jurnal dengan format (jam tidur malam, menit tidur malam, jam bangun malam, menit bangun malam)")
		fmt.Scan(&A[i].jTidurMalam, &A[i].mTidurMalam, &A[i].jBangunMalam, &A[i].mBangunMalam)
		}else if tSiang == false{
			//Tidur malam
		fmt.Println("Silahkan isi jurnal dengan format (jam tidur malam, menit tidur malam, jam bangun malam, menit bangun malam)")
		fmt.Scan(&A[i].jTidurMalam, &A[i].mTidurMalam, &A[i].jBangunMalam, &A[i].mBangunMalam)	
		}
	}
}
