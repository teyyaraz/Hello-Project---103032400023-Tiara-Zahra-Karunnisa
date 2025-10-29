package main
import "fmt"

type Tenant struct {
	Nama            string
	JumlahTransaksi int
	TotalPendapatan float64
}

const NMAX = 100

var daftarTenant [NMAX]Tenant
var jumlahTenant int = 0

func menu() {
	fmt.Println("--------- MENU ---------")
	fmt.Println("1. Lihat Data Tenant")
	fmt.Println("2. Tambah Data Tenant")
	fmt.Println("3. Hapus Data Tenant")
	fmt.Println("4. Catat Transaksi")
	fmt.Println("5. Ubah Nama Tenant")
	fmt.Println("6. Keluar")
	fmt.Println("------------------------")
}

func cariTenant(nama string) int {
	for i := 0; i < jumlahTenant; i++ {
		if daftarTenant[i].Nama == nama {
			return i
		}
	}
	return -1
}

func tambahTenant(nama string) {
	if jumlahTenant >= NMAX {
		fmt.Println("Kapasitas tenant penuh.")
		return
	}
	if cariTenant(nama) != -1 {
		fmt.Println("Nama tenant sudah ada.")
	} else {
		daftarTenant[jumlahTenant] = Tenant{Nama: nama}
		jumlahTenant++
		fmt.Println("Tenant berhasil ditambahkan.")
	}
}

func tambahTransaksi(nama string, nominal float64) {
	idx := cariTenant(nama)
	if idx == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}
	adminKantin := 0.20 * nominal
	pendapatanTenant := nominal - adminKantin

	daftarTenant[idx].JumlahTransaksi++
	daftarTenant[idx].TotalPendapatan += pendapatanTenant
	fmt.Println("Transaksi berhasil dicatat.")
}

func hapusDataTenant(nama string) {
	idx := cariTenant(nama)
	if idx == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahTenant-1; i++ {
		daftarTenant[i] = daftarTenant[i+1]
	}
	jumlahTenant--
	fmt.Println("Tenant berhasil dihapus.")
}

func ubahNamaTenant(namaLama string, namaBaru string) {
	idx := cariTenant(namaLama)
	if idx == -1 {
		fmt.Println("Tenant tidak ditemukan.")
		return
	}
	if cariTenant(namaBaru) != -1 {
		fmt.Println("Nama tenant baru sudah digunakan.")
		return
	}
	daftarTenant[idx].Nama = namaBaru
	fmt.Println("Nama tenant berhasil diubah.")
}

func tampilkanDaftarTenant() {
	if jumlahTenant == 0 {
		fmt.Println("Belum ada data tenant.")
		return
	}

	for i := 1; i < jumlahTenant; i++ {
		key := daftarTenant[i]
		j := i - 1
		for j >= 0 && daftarTenant[j].TotalPendapatan < key.TotalPendapatan {
			daftarTenant[j+1] = daftarTenant[j]
			j--
		}
		daftarTenant[j+1] = key
	}


	fmt.Println("Daftar Tenant:")
	for i := 0; i < jumlahTenant; i++ {
		fmt.Printf("%d. %s | Transaksi: %d | Pendapatan: %.2f\n",
			i+1, daftarTenant[i].Nama, daftarTenant[i].JumlahTransaksi, daftarTenant[i].TotalPendapatan)
	}
}

func main() {
	var pilihan int
	var nama, namaBaru string
	var nominal float64

	for {
		menu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tampilkanDaftarTenant()
			fmt.Println()
		} else if pilihan == 2 {
			fmt.Print("Masukkan nama tenant: ")
			fmt.Scanln(&nama)
			tambahTenant(nama)
			fmt.Println()
		} else if pilihan == 3 {
			fmt.Print("Masukkan nama tenant yang ingin dihapus: ")
			fmt.Scanln(&nama)
			hapusDataTenant(nama)
			fmt.Println()
		} else if pilihan == 4 {
			fmt.Print("Masukkan nama tenant: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan nominal transaksi: ")
			fmt.Scanln(&nominal)
			fmt.Println()
			tambahTransaksi(nama, nominal)
		} else if pilihan == 5 {
			fmt.Print("Masukkan nama tenant lama: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan nama tenant baru: ")
			fmt.Scanln(&namaBaru)
			fmt.Println()
			ubahNamaTenant(nama, namaBaru)
		} else if pilihan == 6 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}