-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 29, 2024 at 05:23 PM
-- Server version: 10.4.8-MariaDB
-- PHP Version: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_puri`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_comments`
--

CREATE TABLE `tb_comments` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `comment` longtext DEFAULT NULL,
  `reply_comment` longtext DEFAULT NULL,
  `rating` double DEFAULT NULL,
  `id_galery` char(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_comments`
--

INSERT INTO `tb_comments` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `comment`, `reply_comment`, `rating`, `id_galery`) VALUES
('258f808c-9334-4565-8335-a182b18b7850', '2024-06-26 17:37:27.458', '2024-06-26 17:37:27.458', NULL, 'komen ke 2', 'ke 2', '', 1, 'f0debca4-2eda-48b0-b426-a4b52bc60b92'),
('2b4292fb-72b6-4369-a544-b824d0478fc5', '2024-06-17 14:34:53.975', '2024-06-17 14:35:25.005', NULL, 'komen ke 2', 'ke 2', 'Maaf saya tidak bisa melakukan itu', 5, 'f0debca4-2eda-48b0-b426-a4b52bc60b92'),
('52ee2d10-4b02-4137-96fd-f417b72b873e', '2024-06-26 18:05:44.705', '2024-06-26 18:05:44.705', NULL, 'komen ke 2', 'ke 2', '', 5, 'b980378a-5272-488f-9373-e12e505cd91a'),
('58bd7f25-db0b-4fb6-b232-29914404aade', '2024-06-26 17:39:12.094', '2024-06-26 17:39:12.094', NULL, 'komen ke 2', 'ke 2', '', 5, '9f8d3fe3-8c01-4a93-bd0f-92a242af06e8'),
('5d12618f-5a70-45f5-891d-dded0e2f6b7d', '2024-06-26 17:39:06.346', '2024-06-26 17:39:06.346', NULL, 'komen ke 2', 'ke 2', '', 1, '9f8d3fe3-8c01-4a93-bd0f-92a242af06e8'),
('72cad0d3-7dc8-43cf-b823-d50055ae3ff9', '2024-06-26 18:05:42.244', '2024-06-26 18:05:42.244', NULL, 'komen ke 2', 'ke 2', '', 5, 'b980378a-5272-488f-9373-e12e505cd91a'),
('7ea2afff-220b-4fc2-b221-e837862f571a', '2024-06-26 17:37:15.083', '2024-06-26 17:37:15.083', NULL, 'komen ke 2', 'ke 2', '', 5, 'f0debca4-2eda-48b0-b426-a4b52bc60b92'),
('875f9ca1-9570-4793-9191-65ed8382339b', '2024-06-26 17:37:20.275', '2024-06-26 17:37:20.275', NULL, 'komen ke 2', 'ke 2', '', 2, 'f0debca4-2eda-48b0-b426-a4b52bc60b92'),
('b86de0b7-e5d3-47a1-b5fa-6229d52d3aea', '2024-06-26 18:06:06.796', '2024-06-26 18:06:06.796', NULL, 'komen ke 2', 'ke 2', '', 5, '31a25f4f-1d39-4526-96ae-f849207ba374'),
('c4d42175-4ea7-4be4-b2b8-a4bed1bdce1d', '2024-06-11 23:47:16.632', '2024-06-17 14:24:17.562', NULL, 'Surya', 'test1', 'Maaf saya tidak bisa melakukan itu', 4, 'f0debca4-2eda-48b0-b426-a4b52bc60b92'),
('d65f6b5d-cc7d-4ca3-b706-2672b8df976e', '2024-06-26 17:39:10.474', '2024-06-26 17:39:10.474', NULL, 'komen ke 2', 'ke 2', '', 5, '9f8d3fe3-8c01-4a93-bd0f-92a242af06e8'),
('db6dfa7f-1e94-40c2-83e4-4bb420d692ee', '2024-06-26 17:39:11.306', '2024-06-26 17:39:11.306', NULL, 'komen ke 2', 'ke 2', '', 5, '9f8d3fe3-8c01-4a93-bd0f-92a242af06e8'),
('f1707619-dbcc-47c2-ba70-da45eb97555f', '2024-06-26 18:06:11.813', '2024-06-26 18:06:11.813', NULL, 'komen ke 2', 'ke 2', '', 4, '31a25f4f-1d39-4526-96ae-f849207ba374');

-- --------------------------------------------------------

--
-- Table structure for table `tb_galeries`
--

CREATE TABLE `tb_galeries` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_information` char(36) NOT NULL,
  `image` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_galeries`
--

INSERT INTO `tb_galeries` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_information`, `image`) VALUES
('0c8f715d-5df9-4efb-8d71-410978f6837e', '2024-06-26 18:05:09.816', '2024-06-26 18:05:09.816', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\Contract Staff Sinta Wisnu_page-0001.jpg'),
('31a25f4f-1d39-4526-96ae-f849207ba374', '2024-06-26 18:04:48.913', '2024-06-26 18:04:48.913', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\4abfe89b-f028-4980-974b-8130c625f75b-removebg-preview.png'),
('9f8d3fe3-8c01-4a93-bd0f-92a242af06e8', '2024-06-12 14:21:40.389', '2024-06-12 14:21:40.389', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\412968-horror-Mariusz-Lewandowski-digital-art-fan-art-Photoshop.jpg'),
('b980378a-5272-488f-9373-e12e505cd91a', '2024-06-26 18:04:58.791', '2024-06-26 18:04:58.791', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\usecase update.drawio (5).png'),
('d2e99fb4-d8ad-47e9-9a7b-62c772d454a7', '2024-06-19 18:17:54.391', '2024-06-19 18:17:54.391', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\412968-horror-Mariusz-Lewandowski-digital-art-fan-art-Photoshop.jpg'),
('f0debca4-2eda-48b0-b426-a4b52bc60b92', '2024-06-19 18:20:58.302', '2024-06-19 18:20:58.302', NULL, '046d975e-c9a9-415b-a413-3b6a44b2ab27', 'uploads\\lost_haven_by_chriscold_d96cl0d.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `tb_informations`
--

CREATE TABLE `tb_informations` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_login` bigint(20) UNSIGNED DEFAULT NULL,
  `judul_foto` longtext DEFAULT NULL,
  `nama_lokasi` longtext DEFAULT NULL,
  `deskripsi` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_informations`
--

INSERT INTO `tb_informations` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_login`, `judul_foto`, `nama_lokasi`, `deskripsi`) VALUES
('046d975e-c9a9-415b-a413-3b6a44b2ab27', '2024-06-11 23:45:23.451', '2024-06-11 23:45:23.451', NULL, 3, 'judul foto baru', 'Nama lokasi baru', 'Desc'),
('7baa6f63-0dcc-4c53-8c8a-b5bf5bf9ced1', '2024-06-19 18:10:48.599', '2024-06-19 18:10:48.599', NULL, 3, 'foto monument 1', 'monument 1', 'Desc monument 1');

-- --------------------------------------------------------

--
-- Table structure for table `tb_logins`
--

CREATE TABLE `tb_logins` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `username` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_logins`
--

INSERT INTO `tb_logins` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `password`, `username`) VALUES
(1, NULL, NULL, NULL, 'test1', 'surya9029', NULL),
(2, '2024-02-02 20:52:05.837', '2024-02-02 20:52:05.837', NULL, 'test1', '$2a$10$XUntPo4K5f4x9Trzjcx9Ne.28zlO8t03owRjB6NxgNnSptT5959zy', NULL),
(3, '2024-03-06 15:16:32.690', '2024-03-06 15:16:32.690', NULL, 'admin', '$2a$10$7vX91Y5EYcyEqjXgjBGMaugzK.nfteWVruphEG20M23/JrPcMLGhe', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tb_monuments`
--

CREATE TABLE `tb_monuments` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `id_information` char(36) NOT NULL,
  `image` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_monuments`
--

INSERT INTO `tb_monuments` (`id`, `created_at`, `updated_at`, `deleted_at`, `id_information`, `image`) VALUES
('0a4fa7ff-6d6b-4f72-9dc2-a39d125c265b', '2024-06-19 18:22:47.458', '2024-06-19 18:22:47.458', NULL, '7baa6f63-0dcc-4c53-8c8a-b5bf5bf9ced1', 'uploads\\lost_haven_by_chriscold_d96cl0d.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_comments`
--
ALTER TABLE `tb_comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_comments_tb_galery` (`id_galery`);

--
-- Indexes for table `tb_galeries`
--
ALTER TABLE `tb_galeries`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_galeries_tb_information` (`id_information`);

--
-- Indexes for table `tb_informations`
--
ALTER TABLE `tb_informations`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tb_logins`
--
ALTER TABLE `tb_logins`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_tb_logins_deleted_at` (`deleted_at`);

--
-- Indexes for table `tb_monuments`
--
ALTER TABLE `tb_monuments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_tb_monuments_tb_information` (`id_information`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_logins`
--
ALTER TABLE `tb_logins`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tb_comments`
--
ALTER TABLE `tb_comments`
  ADD CONSTRAINT `fk_tb_comments_tb_galery` FOREIGN KEY (`id_galery`) REFERENCES `tb_galeries` (`id`);

--
-- Constraints for table `tb_galeries`
--
ALTER TABLE `tb_galeries`
  ADD CONSTRAINT `fk_tb_galeries_tb_information` FOREIGN KEY (`id_information`) REFERENCES `tb_informations` (`id`);

--
-- Constraints for table `tb_monuments`
--
ALTER TABLE `tb_monuments`
  ADD CONSTRAINT `fk_tb_monuments_tb_information` FOREIGN KEY (`id_information`) REFERENCES `tb_informations` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
