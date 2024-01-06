-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 06 Jan 2024 pada 17.09
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_mahasiswa`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `nilaimahasiswa`
--

CREATE TABLE `nilaimahasiswa` (
  `id` int(11) NOT NULL,
  `nama` varchar(60) NOT NULL,
  `matakuliah` varchar(60) NOT NULL,
  `nilai` int(3) NOT NULL,
  `indeksnilai` varchar(3) NOT NULL,
  `created_at` varchar(30) NOT NULL,
  `updated_at` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `nilaimahasiswa`
--

INSERT INTO `nilaimahasiswa` (`id`, `nama`, `matakuliah`, `nilai`, `indeksnilai`, `created_at`, `updated_at`) VALUES
(3, 'Mirza', 'Golang Web', 80, 'A', '1704557298', '1704557298'),
(4, 'Azrim', 'Golang Web', 100, 'A', '1704557310', '1704557310');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `nilaimahasiswa`
--
ALTER TABLE `nilaimahasiswa`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `nilaimahasiswa`
--
ALTER TABLE `nilaimahasiswa`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
