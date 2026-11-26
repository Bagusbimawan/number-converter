# Konversi Bilangan

**Tugas Pengganti UTS**

Nama: Bagus Bimawan Sembodo  
NPM: 10125182

## Tentang Program

Program ini merupakan program sederhana yang dibuat menggunakan bahasa Go untuk mengonversi bilangan ke beberapa sistem bilangan, yaitu desimal, biner, oktal, dan hexadecimal.

Pengguna dapat memasukkan angka dengan prefix tertentu untuk menentukan jenis bilangan yang digunakan. Jika tidak menggunakan prefix, angka akan dianggap sebagai desimal.

## Cara Penggunaan

Prefix yang digunakan:

- `0b` untuk Biner
- `0o` untuk Oktal
- `0x` untuk Hexadecimal
- Tanpa prefix untuk Desimal

Contoh input:

```text
15
0b1111
0o17
0xF
```

Program akan menampilkan hasil konversi dari angka yang dimasukkan ke dalam empat sistem bilangan.

## Menjalankan Program

Pastikan Go sudah terinstal, kemudian jalankan program melalui terminal:

```bash
go run main.go
```

Untuk keluar dari program, masukkan `q`.

## Contoh Hasil

Input:

```text
15
```

Output:

```text
Jenis: Desimal
Desimal      : 15
Biner        : 1111
Oktal        : 17
Heksadesimal : F
```

## Penutup

Program ini dibuat sebagai bagian dari tugas pengganti UTS untuk mempelajari konversi sistem bilangan menggunakan bahasa Go.
