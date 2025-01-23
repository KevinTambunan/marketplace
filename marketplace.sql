-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 12, 2025 at 01:04 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `marketplace`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Electronics', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(2, 'Fashion', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(3, 'Books', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(4, 'Home Appliances', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(5, 'Toys', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(6, 'Sports', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(7, 'Furniture', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(8, 'Beauty Products', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(9, 'Automotive', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(10, 'Groceries', '2025-01-12 06:21:30', '2025-01-12 06:21:30'),
(11, 'test update', '2025-01-12 01:25:31', '2025-01-12 01:33:48'),
(13, 'test create product', '2025-01-12 02:57:10', '2025-01-12 02:57:10');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `status_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `user_id`, `product_id`, `quantity`, `total_price`, `status_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 2, 1400, 1, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(2, 3, 4, 1, 15, 2, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(3, 5, 2, 1, 1200, 3, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(4, 7, 3, 3, 150, 4, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(5, 2, 5, 2, 200, 5, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(6, 6, 7, 1, 30, 6, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(7, 8, 9, 4, 80, 7, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(8, 9, 8, 2, 300, 8, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(9, 10, 10, 1, 200, 9, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(10, 4, 6, 5, 125, 10, '2025-01-12 06:22:46', '2025-01-12 06:22:46'),
(11, 1, 1, 2, 1400, 1, '2025-01-12 03:05:53', '2025-01-12 03:06:34');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `stock` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `price`, `category_id`, `stock`, `created_at`, `updated_at`) VALUES
(1, 'Smartphone', 'Latest 5G smartphone', 700, 1, 50, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(2, 'Laptop', 'Powerful laptop for professionals', 1200, 1, 30, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(3, 'Jeans', 'Comfortable denim jeans', 50, 2, 100, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(4, 'Novel', 'Bestselling fiction book', 15, 3, 200, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(5, 'Blender', 'High-speed kitchen blender', 100, 4, 40, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(6, 'Teddy Bear', 'Soft and cuddly toy', 25, 5, 80, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(7, 'Football', 'Standard size football', 30, 6, 60, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(8, 'Office Chair', 'Ergonomic office chair', 150, 7, 20, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(9, 'Lipstick', 'Matte finish lipstick', 20, 8, 150, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(10, 'Car Tire', 'Durable all-season tire', 200, 9, 15, '2025-01-12 06:22:12', '2025-01-12 06:22:12'),
(11, 'Car Tire', 'Durable all-season tire', 200, 9, 15, '2025-01-12 02:57:24', '2025-01-12 02:58:04');

-- --------------------------------------------------------

--
-- Table structure for table `status`
--

CREATE TABLE `status` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `status`
--

INSERT INTO `status` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Pending', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(2, 'Shipped', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(3, 'Delivered', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(4, 'Cancelled', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(5, 'Returned', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(6, 'Processing', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(7, 'Failed', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(8, 'Completed', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(9, 'On Hold', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(10, 'Awaiting Payment', '2025-01-12 06:21:54', '2025-01-12 06:21:54'),
(11, 'test update', '2025-01-12 01:46:24', '2025-01-12 01:47:04');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `created_at`, `updated_at`) VALUES
(1, 'Alice Johnson', 'alice@mail.com', 'hashed_password1', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(2, 'Bob Smith', 'bob@mail.com', 'hashed_password2', 'admin', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(3, 'Charlie Brown', 'charlie@mail.com', 'hashed_password3', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(4, 'Daisy White', 'daisy@mail.com', 'hashed_password4', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(5, 'Evan Green', 'evan@mail.com', 'hashed_password5', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(6, 'Fiona Red', 'fiona@mail.com', 'hashed_password6', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(7, 'George Black', 'george@mail.com', 'hashed_password7', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(8, 'Helen Blue', 'helen@mail.com', 'hashed_password8', 'admin', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(9, 'Ian Silver', 'ian@mail.com', 'hashed_password9', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(10, 'Julia Gold', 'julia@mail.com', 'hashed_password10', 'customer', '2025-01-12 06:22:29', '2025-01-12 06:22:29'),
(11, '', 'user@example.com', '$2a$10$cPKJAti1YQK9MW0s8CLCvuODU3so6E8RfyurbJSG2w3lcgY6oT7wi', 'user', '2025-01-12 10:29:51', '2025-01-12 10:29:51'),
(12, 'John Doe', 'john.doe@example.com', '$2a$10$CkrihClPF6gQXdjfarn.buCl1if5mtuoifmbBAUCKKxkQT31.A7bW', 'user', '2025-01-12 10:29:59', '2025-01-12 10:29:59');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `status`
--
ALTER TABLE `status`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `status`
--
ALTER TABLE `status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
