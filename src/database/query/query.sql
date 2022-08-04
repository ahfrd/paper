-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 04, 2022 at 05:21 PM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.4.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `paper`
--
CREATE DATABASE paper;
-- --------------------------------------------------------

--
-- Table structure for table `financial_account`
--

CREATE TABLE `financial_account` (
  `id` int(11) NOT NULL,
  `accountNumber` varchar(11) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `phoneNumber` varchar(10) NOT NULL,
  `created_date` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `financial_account`
--

INSERT INTO `financial_account` (`id`, `accountNumber`, `firstname`, `lastname`, `address`, `username`, `phoneNumber`, `created_date`, `updated_at`, `deleted_at`) VALUES
(1, '32412390', 'Farid', 'a', '2', 'aa', '0877863721', '2022-08-03 23:09:24', NULL, '2022-08-03 23:31:05'),
(8, '2147483647', 'Test', 'Update Acc', 'Cibanteng', 'errq', '0877863721', '2022-08-04 10:58:26', '2022-08-04 22:14:48', '2022-08-04 22:14:54'),
(10, '72290666', 'Ahmad', 'Farid', 'Cibanteng', 'errq', '0877863721', '2022-08-04 11:00:00', NULL, NULL),
(11, '1251585643', 'Ahmad', 'Farid', 'Cibanteng', 'errq', '0877863721', '2022-08-04 11:00:09', NULL, NULL),
(14, '2147483647', 'Ahmad', 'Farid', 'Cibanteng', 'errq', '0877863721', '2022-08-04 11:01:41', NULL, NULL),
(15, '79679912848', 'Ahmad', 'Farid', 'Cibanteng', 'errq', '0877863721', '2022-08-04 11:04:28', NULL, NULL),
(16, '31358475509', 'Ahmad', 'Farid', 'Cibanteng', 'errq', '0877863721', '2022-08-04 22:05:24', NULL, NULL),
(17, '40974962159', 'Ahmad', 'Farid', 'Cibanteng', 'errqs', '0877863721', '2022-08-04 22:09:38', NULL, NULL),
(18, '96307322703', 'Ahmad', 'Farid', 'Cibanteng', 'errqs', '0877863721', '2022-08-04 22:12:37', NULL, NULL),
(19, '61898912274', 'Ahmad', 'Farid', 'Cibanteng', 'errqs', '0877863721', '2022-08-04 22:14:42', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `financial_transaction`
--

CREATE TABLE `financial_transaction` (
  `id` int(11) NOT NULL,
  `accountNumber` int(11) NOT NULL,
  `productName` varchar(255) NOT NULL,
  `productCategory` varchar(255) NOT NULL,
  `nominal` varchar(60) NOT NULL,
  `transaction_date` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `financial_transaction`
--

INSERT INTO `financial_transaction` (`id`, `accountNumber`, `productName`, `productCategory`, `nominal`, `transaction_date`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 941248, 'tes', 'test', '500', '2022-08-02 20:00:59', '2022-08-04 01:01:20', '2022-08-03 20:00:59', NULL),
(2, 941248, 'tes', 'test', '500', '2022-08-02 20:00:59', '2022-08-04 01:01:35', '2022-08-03 20:00:59', NULL),
(3, 941248, 'tes', 'test', '500', '2022-08-03 20:00:59', '2022-08-04 01:01:45', '2022-08-03 20:00:59', NULL),
(4, 941248, 'tes', 'test', '500', '2022-08-04 20:00:59', '2022-08-04 01:01:50', '2022-08-03 20:00:59', NULL),
(9, 2147483647, 'DAIA', 'Sabun', '50000', '2022-08-04 11:39:55', '2022-08-04 11:39:55', NULL, NULL),
(10, 2147483647, 'DAIA', 'Sabun', '50000', '2022-08-04 11:39:56', '2022-08-04 11:39:56', NULL, '2022-08-04 15:05:20'),
(11, 2147483647, 'DAIA', 'Sabun', '50000', '2022-08-04 11:39:58', '2022-08-04 11:39:58', NULL, NULL),
(12, 2147483647, 'DAIA', 'Sabun', '50000', '2022-08-04 11:39:59', '2022-08-04 11:39:59', NULL, NULL),
(13, 2147483647, 'DAIA', 'Sabun', '50000', '2022-08-04 11:40:00', '2022-08-04 11:40:00', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tbl_user`
--

CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `login_retry` int(11) DEFAULT NULL,
  `next_login_date` datetime DEFAULT NULL,
  `last_login` datetime DEFAULT NULL,
  `session_id` varchar(255) DEFAULT NULL,
  `status` int(2) NOT NULL DEFAULT 0,
  `created_date` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tbl_user`
--

INSERT INTO `tbl_user` (`id`, `username`, `password`, `email`, `login_retry`, `next_login_date`, `last_login`, `session_id`, `status`, `created_date`) VALUES
(1, 'errq', '8920413b7f29dbe6f0e5fff94ece754c6e9af2815a661b280971682090bf9a1f', 'farid2@gmail.com', 0, NULL, '2022-08-04 10:30:56', 'f30bff84dcaea0dfc070da5a9ae6f21e77cf145497308e5cefe617cd4c530743', 0, '2022-08-04 10:26:13'),
(2, 'errqs', '8cea1b549e21bde240df5cff359217d71325afeeb64da13a988be9925794a054', 'farid2@gmail.com', 0, NULL, '2022-08-04 22:14:37', '8ac52e67b82fd3e57f925dd52c47ddbb8c7f1f2c4210ed41237892b039edbe70', 1, '2022-08-04 21:57:52');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `financial_account`
--
ALTER TABLE `financial_account`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `financial_transaction`
--
ALTER TABLE `financial_transaction`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_user`
--
ALTER TABLE `tbl_user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `financial_account`
--
ALTER TABLE `financial_account`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- AUTO_INCREMENT for table `financial_transaction`
--
ALTER TABLE `financial_transaction`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `tbl_user`
--
ALTER TABLE `tbl_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
