-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 29, 2022 at 05:21 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `booking_hotel`
--

-- --------------------------------------------------------

--
-- Table structure for table `booking`
--

CREATE TABLE `booking` (
  `id` int(11) NOT NULL,
  `status` varchar(25) NOT NULL,
  `room_id` int(11) NOT NULL,
  `hotel_id` int(11) NOT NULL,
  `pic_name` varchar(100) NOT NULL,
  `pic_contact` varchar(100) NOT NULL,
  `event_start` datetime NOT NULL,
  `event_end` datetime NOT NULL,
  `invoice_number` varchar(100) NOT NULL,
  `invoice_date` date NOT NULL,
  `invoice_grandtotal` double NOT NULL,
  `discount_request` varchar(25) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `facility`
--

CREATE TABLE `facility` (
  `id` int(11) NOT NULL,
  `type` varchar(100) NOT NULL,
  `description` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `floor`
--

CREATE TABLE `floor` (
  `id` int(11) NOT NULL,
  `number` int(11) NOT NULL,
  `hotel_id` int(11) NOT NULL,
  `room_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `floor`
--

INSERT INTO `floor` (`id`, `number`, `hotel_id`, `room_id`, `created_at`, `updated_at`) VALUES
(38, 1, 1, 1, '2022-07-22 12:55:46', '2022-07-22 12:55:46'),
(40, 1, 1, 0, '2022-07-22 13:43:49', '2022-07-22 13:43:49'),
(41, 1, 1, 0, '2022-07-22 13:44:43', '2022-07-22 13:44:43'),
(42, 1, 1, 4, '2022-07-22 13:45:18', '2022-07-22 13:45:18');

-- --------------------------------------------------------

--
-- Table structure for table `hotel`
--

CREATE TABLE `hotel` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `province` varchar(100) NOT NULL,
  `city` varchar(100) NOT NULL,
  `zip_code` varchar(100) NOT NULL,
  `star` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `hotel`
--

INSERT INTO `hotel` (`id`, `name`, `address`, `province`, `city`, `zip_code`, `star`, `created_at`, `updated_at`) VALUES
(1, 'futura', 'jl bukit pajajaran no 1', 'jawa barat', 'bandung', '40145', 5, '2022-07-09 02:21:46', '2022-07-09 02:21:46'),
(2, 'alexa', 'jl bukit tinggi 2', 'jawa timur', 'yogyakarta', '20332', 4, '2022-07-09 02:23:10', '2022-07-09 02:23:10'),
(3, 'safir', 'jl tanpa bayangan no. 1', 'jawa tengah', 'semarang', '76427', 3, '2022-07-09 06:28:32', '2022-07-09 06:28:32'),
(4, 'intan', 'jl bukit hijau no.1', 'sumatra selatan', 'palembang', '64232', 5, '2022-07-09 06:28:32', '2022-07-09 06:28:32'),
(6, 'adibas', 'jl bukit pajajaran no 999', 'jawa barat', 'bandung', '40666', 3, '2022-07-21 08:07:41', '2022-07-21 08:07:41');

-- --------------------------------------------------------

--
-- Table structure for table `role`
--

CREATE TABLE `role` (
  `id` int(11) NOT NULL,
  `type` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `role`
--

INSERT INTO `role` (`id`, `type`, `created_at`, `updated_at`) VALUES
(1, 'owner', '2022-07-22 10:09:09', '2022-07-22 10:09:09'),
(2, 'employee', '2022-07-22 10:10:32', '2022-07-22 10:10:32');

-- --------------------------------------------------------

--
-- Table structure for table `room`
--

CREATE TABLE `room` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `capacity` int(11) NOT NULL,
  `price_per_hour` double NOT NULL,
  `price_per_day` double NOT NULL,
  `facility_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `room`
--

INSERT INTO `room` (`id`, `name`, `capacity`, `price_per_hour`, `price_per_day`, `facility_id`, `created_at`, `updated_at`) VALUES
(1, 'Maisyah', 20, 200000, 1000000, 0, '2022-07-09 04:35:07', '2022-07-09 04:35:07'),
(2, 'Marwah', 30, 300000, 1200000, 0, '2022-07-09 04:35:07', '2022-07-09 04:35:07'),
(3, 'Barakah', 50, 500000, 1500000, 0, '2022-07-09 04:37:28', '2022-07-09 04:37:28'),
(4, 'Burdah', 60, 600000, 1600000, 0, '2022-07-09 04:41:50', '2022-07-09 04:41:50');

-- --------------------------------------------------------

--
-- Table structure for table `user_hotel`
--

CREATE TABLE `user_hotel` (
  `id` int(11) NOT NULL,
  `user_profile_id` int(11) NOT NULL,
  `hotel_id` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_hotel`
--

INSERT INTO `user_hotel` (`id`, `user_profile_id`, `hotel_id`) VALUES
(20, 39, 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_profile`
--

CREATE TABLE `user_profile` (
  `id` int(11) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8 NOT NULL,
  `gender` varchar(100) CHARACTER SET utf8 NOT NULL,
  `email` varchar(100) CHARACTER SET utf8 NOT NULL,
  `password` varchar(100) CHARACTER SET utf8 NOT NULL,
  `role_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_profile`
--

INSERT INTO `user_profile` (`id`, `name`, `gender`, `email`, `password`, `role_id`, `created_at`, `updated_at`) VALUES
(36, 'Budi', 'Male', 'budiowners123@gmail.com', 'budiownerhotel123213', 1, '2022-07-22 10:12:08', '2022-07-22 10:12:08'),
(37, 'Aria', 'female', 'aria@gmail.com', 'ariel123', 2, '2022-07-22 10:13:31', '2022-07-22 10:13:31'),
(39, 'Arias', 'female', 'arias@gmail.com', 'ariel123', 2, '2022-07-22 10:14:04', '2022-07-22 10:14:04'),
(40, 'Ariast', 'female', 'ariast@gmail.com', 'ariel123', 2, '2022-07-22 10:16:12', '2022-07-22 10:16:12');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `booking`
--
ALTER TABLE `booking`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `facility`
--
ALTER TABLE `facility`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `type` (`type`);

--
-- Indexes for table `floor`
--
ALTER TABLE `floor`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `hotel`
--
ALTER TABLE `hotel`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `address` (`address`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `type` (`type`);

--
-- Indexes for table `room`
--
ALTER TABLE `room`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `user_hotel`
--
ALTER TABLE `user_hotel`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_profile_id` (`user_profile_id`);

--
-- Indexes for table `user_profile`
--
ALTER TABLE `user_profile`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `booking`
--
ALTER TABLE `booking`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `facility`
--
ALTER TABLE `facility`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `floor`
--
ALTER TABLE `floor`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=43;

--
-- AUTO_INCREMENT for table `hotel`
--
ALTER TABLE `hotel`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `role`
--
ALTER TABLE `role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `room`
--
ALTER TABLE `room`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user_hotel`
--
ALTER TABLE `user_hotel`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `user_profile`
--
ALTER TABLE `user_profile`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
