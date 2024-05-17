# RAPIKAN

## About The Project
Rapikan adalah aplikasi untuk memesan layanan kebersihan secara *online* dimana para pengguna dapat memilih layanan dan durasi serta jadwal kedatangan *housekeeper* sesuai keinginan pengguna.

## Features
Pada aplikasi terdapat 2 role yaitu *User* dan *Housekeeper*

### User
- Register dan login
- Melihat layanan beserta deskripsi layanan yang tersedia.
- Memilih housekeeper sesuai dengan layanan yang dipilih.
- Memilih durasi laynan.
- Melakukan pemesanan dan pembayaran secara online.

### Admin
- Melihat daftar pemesanan yang diterima.
- Mengelola jadwal layanan.


## Tech Stacks
- Programming Language: Go
- Framework: Echo
- Database: MySQL
- Deployment: Railway
- Authentication: JWT
- Version Control System: Github & Github Desktop

## Design API


## Entity Relationship Diagram 

[ERD](https://dbdiagram.io/d/ERD-Mini-Project-Back-End-Golang-ALTA-66313bc35b24a634d033af5c)
![ERD](https://github.com/Ingeylas/MiniProject_HousekeepingService/assets/114483889/ced995a5-b024-43ae-801a-bb325aef6793)


## Setup

**Clone the repository:**

```
git clone https://github.com/Ingeylas/MiniProject_HousekeepingService.git
```

**Navigate to the project directory:**

```
git clone https://github.com/Ingeylas/MiniProject_HousekeepingService.git
```

**Copy the .env.example file to .env and configure the environment variables:**

```
git clone https://github.com/Ingeylas/MiniProject_HousekeepingService.git
```

**Install the dependencies:**

```
go mod tidy
```

Untuk konfigurasi DB lakukan pengaturan terhadap nama DB, Password DB dan Host DB pada file `configs/config.go`

**Run the application:**

```
go run main.go
```

