# Number Converter

Program sederhana untuk mengkonversi bilangan ke berbagai sistem bilangan menggunakan bahasa Go.

Program ini dibuat untuk memenuhi tugas UTS pengganti.

## Identitas

- **Nama:** Bagus Bimawan Sembodo
- **NPM:** 10125182

## Fitur

- Konversi bilangan desimal.
- Konversi bilangan biner.
- Konversi bilangan oktal.
- Konversi bilangan heksadesimal.
- Mendukung bilangan negatif menggunakan 8-bit Two's Complement.
- Mendukung bilangan pecahan (Floating Point) untuk konversi ke Binary.
- Validasi input.
- Program dapat dijalankan berulang kali.

## Format Input

| Prefix | Sistem Bilangan |
|---|---|
| `0b` | Biner |
| `0o` | Oktal |
| `0x` | Heksadesimal |
| Tanpa prefix | Desimal |

Contoh input:

```text
10
-5
10.625
0b1010
0o12
0xA
```

## Contoh Output

### Bilangan negatif

Input:

```text
-5
```

Output:

```text
Desimal      : -5
Biner        : 11111011
Oktal        : 373
Heksadesimal : FB
```

Representasi bilangan negatif menggunakan 8-bit Two's Complement.

### Floating Point

Input:

```text
10.625
```

Output:

```text
Desimal      : 10.625
Biner        : 1010.101
Oktal        : Integer saja
Heksadesimal : Integer saja
```

Konversi Floating Point dilakukan dengan memisahkan bagian integer dan bagian pecahan.

## Batasan Program

- Integer menggunakan representasi 8-bit signed dengan range -128 sampai 127.
- Two's Complement digunakan untuk representasi bilangan integer negatif.
- Floating Point didukung untuk konversi ke Binary.
- Konversi Oktal dan Heksadesimal hanya untuk integer.
- Konversi bagian pecahan Binary dibatasi hingga 10 digit.

## Cara Menjalankan

Pastikan Go sudah terinstall.

Clone repository:

```bash
git clone https://github.com/Bagusbimawan/number-converter.git
```

Masuk ke folder project:

```bash
cd number-converter
```

Jalankan program:

```bash
go run main.go
```

## Teknologi

- Go (Golang)

## Penutup

Program ini dibuat untuk memahami sistem bilangan, proses konversi bilangan, representasi bilangan negatif menggunakan Two's Complement, serta konversi bilangan pecahan ke Binary.
