-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 14, 2021 at 09:14 PM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 8.0.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `astara`
--

-- --------------------------------------------------------

--
-- Table structure for table `Areas`
--

CREATE TABLE `Areas` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Name` varchar(50) NOT NULL,
  `Id_user` int(10) UNSIGNED NOT NULL,
  `Deleteable` tinyint(1) NOT NULL DEFAULT 1,
  `Slug` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Areas`
--

INSERT INTO `Areas` (`Id`, `Name`, `Id_user`, `Deleteable`, `Slug`) VALUES
(1, 'pulvinar lobortis', 7, 1, 'pulvinar-lobortis'),
(2, 'hac', 7, 0, 'hac'),
(3, 'nascetur', 7, 1, 'nascetur'),
(4, 'venenatis non', 7, 1, 'venenatis-non'),
(5, 'vulputate', 7, 1, 'vulputate'),
(50, 'universidad', 14, 1, 'universidad'),
(51, 'no se', 14, 1, 'no-se'),
(52, 'Index', 16, 0, 'index'),
(53, 'Index', 17, 0, 'index'),
(55, 'Nueva', 17, 1, 'Nueva'),
(56, 'novedoso', 18, 0, 'index'),
(233, 'tfg', 18, 1, 'tfg'),
(234, 'nueva', 18, 1, 'nueva');

-- --------------------------------------------------------

--
-- Table structure for table `Goals`
--

CREATE TABLE `Goals` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Id_target` int(10) UNSIGNED NOT NULL,
  `Description` text DEFAULT NULL,
  `Children` tinyint(3) UNSIGNED DEFAULT 0,
  `ChildrenDone` tinyint(3) UNSIGNED DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Goals`
--

INSERT INTO `Goals` (`Id`, `Id_target`, `Description`, `Children`, `ChildrenDone`) VALUES
(1, 15, 'quam a odio in', 0, 0),
(2, 28, 'venenatis turpis', 0, 0),
(3, 22, 'tincidunt eget tempus vel', 0, 0),
(4, 12, 'lectus aliquam sit amet', 0, 0),
(5, 12, 'mus etiam vel augue', 0, 0),
(6, 28, 'enim in tempor', 0, 0),
(7, 13, 'et ultrices posuere cubilia curae', 0, 0),
(8, 18, 'in quis justo maecenas', 0, 0),
(9, 17, 'nunc proin', 0, 0),
(10, 27, 'luctus cum sociis natoque penatibus', 0, 0),
(11, 14, 'non velit donec', 0, 0),
(12, 15, 'erat id mauris', 0, 0),
(13, 18, 'augue luctus tincidunt nulla', 0, 0),
(14, 18, 'tempor turpis nec euismod', 0, 0),
(15, 28, 'nulla neque', 0, 0),
(16, 23, 'id nulla ultrices aliquet', 0, 0),
(17, 12, 'orci luctus et', 0, 0),
(18, 25, 'vestibulum ante ipsum primis in', 0, 0),
(19, 23, 'erat id mauris vulputate elementum', 0, 0),
(20, 13, 'eros suspendisse accumsan', 0, 0),
(42, 109, 'Guardar los pañuelos en otro sitio', 0, 0),
(43, 110, 'nose algo random por aquí no sé', 0, 0),
(48, 148, 'nose', 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `Rols`
--

CREATE TABLE `Rols` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Rols`
--

INSERT INTO `Rols` (`Id`, `Name`) VALUES
(1, 'user'),
(2, 'admin');

-- --------------------------------------------------------

--
-- Table structure for table `Status`
--

CREATE TABLE `Status` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Status`
--

INSERT INTO `Status` (`Id`, `Name`) VALUES
(1, 'Ok'),
(2, 'Confirmed'),
(10, 'Pending'),
(20, 'Banned'),
(30, 'Gone'),
(50, 'Undone'),
(51, 'Done'),
(70, 'Coming');

-- --------------------------------------------------------

--
-- Table structure for table `Targets`
--

CREATE TABLE `Targets` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Id_parent` int(10) UNSIGNED DEFAULT NULL,
  `Id_usu` int(10) UNSIGNED NOT NULL,
  `Id_area` int(10) UNSIGNED NOT NULL,
  `Id_status` int(10) UNSIGNED NOT NULL DEFAULT 50,
  `Name` varchar(155) NOT NULL,
  `Deadline` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Targets`
--

INSERT INTO `Targets` (`Id`, `Id_parent`, `Id_usu`, `Id_area`, `Id_status`, `Name`, `Deadline`) VALUES
(11, 22, 7, 3, 51, 'cubilia curae donec pharetra magna', '2020-12-07'),
(12, 15, 7, 3, 51, 'pellentesque at nulla', '2021-02-14'),
(13, NULL, 7, 4, 50, 'ipsum primis in faucibus orci', '2020-11-25'),
(14, 15, 7, 2, 50, 'turpis donec posuere metus', '2020-07-16'),
(15, NULL, 7, 4, 51, 'cras in purus eu magna', '2020-10-28'),
(16, 20, 7, 1, 50, 'ridiculus mus etiam vel', '2020-12-15'),
(17, NULL, 7, 2, 50, 'velit vivamus vel nulla eget', '2020-09-02'),
(18, NULL, 7, 1, 51, 'ultrices phasellus id sapien in', '2021-01-26'),
(19, NULL, 7, 1, 50, 'ut at dolor quis', '2020-12-03'),
(20, 13, 7, 5, 50, 'ut erat id', '2021-03-31'),
(21, NULL, 7, 1, 50, 'amet consectetuer adipiscing', '2021-01-26'),
(22, 23, 7, 2, 51, 'dui vel nisl duis', '2020-08-25'),
(23, 30, 7, 3, 51, 'at nibh in hac habitasse', '2021-04-21'),
(24, 25, 7, 4, 51, 'penatibus et magnis', '2020-07-05'),
(25, NULL, 7, 4, 51, 'id ligula suspendisse', '2020-06-18'),
(26, NULL, 7, 5, 50, 'malesuada in imperdiet', '2020-09-14'),
(27, 12, 7, 2, 50, 'mattis pulvinar nulla', '2021-04-26'),
(28, 30, 7, 5, 50, 'phasellus sit amet erat nulla', '2020-09-30'),
(29, NULL, 7, 5, 51, 'augue vel accumsan tellus nisi', '2020-11-12'),
(30, 26, 7, 3, 50, 'ipsum primis in faucibus', '2021-04-20'),
(31, NULL, 14, 51, 50, 'algo', '2021-06-29'),
(32, NULL, 14, 51, 50, 'prueba', '2021-06-23'),
(33, NULL, 14, 51, 50, 'prueba2', '2021-06-28'),
(34, NULL, 14, 51, 50, 'pruebaTres', '2021-06-16'),
(35, NULL, 14, 51, 50, 'nuevo', '2021-06-16'),
(36, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(37, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(38, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(39, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(40, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(41, NULL, 14, 51, 50, 'nose', '2021-06-16'),
(42, NULL, 14, 51, 50, 'new', '2021-06-16'),
(43, NULL, 14, 51, 50, 'new', '2021-06-16'),
(44, NULL, 14, 51, 50, 'new', '2021-06-16'),
(45, NULL, 14, 51, 50, 'new', '2021-06-16'),
(46, NULL, 14, 51, 50, 'new', '2021-06-16'),
(47, NULL, 14, 51, 50, 'new', '2021-06-16'),
(109, NULL, 18, 233, 50, 'pañuelos', '2021-06-07'),
(110, NULL, 18, 233, 50, 'un objetivo de prueba', '2021-06-24'),
(129, NULL, 18, 233, 50, 'muy nuevo', '2021-06-21'),
(130, NULL, 18, 233, 50, 'uno nuevo', '2021-06-22'),
(132, NULL, 18, 233, 50, 'botella', '2021-07-05'),
(133, NULL, 18, 233, 50, 'botella', '2021-07-05'),
(134, NULL, 18, 233, 50, 'vaso de cristal', '2021-06-16'),
(135, NULL, 18, 233, 50, 'galletas', '2021-06-16'),
(139, 109, 18, 233, 50, 'nueva tarea', '2021-06-08'),
(148, NULL, 18, 233, 50, 'nuevo objetivo', '2021-06-17'),
(154, 109, 18, 233, 50, 'newone', '2021-06-10'),
(161, 109, 18, 233, 50, 'novedoso', '2021-06-10');

-- --------------------------------------------------------

--
-- Table structure for table `Task`
--

CREATE TABLE `Task` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Id_target` int(10) UNSIGNED NOT NULL,
  `Dated` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Task`
--

INSERT INTO `Task` (`Id`, `Id_target`, `Dated`) VALUES
(1, 31, '2021-06-23'),
(2, 32, '2021-06-29'),
(3, 33, '2021-06-23'),
(4, 34, '2021-06-28'),
(5, 35, '2021-06-22'),
(6, 36, '2021-06-22'),
(7, 37, '2021-06-22'),
(8, 38, '2021-06-22'),
(9, 39, '2021-06-22'),
(10, 40, '2021-06-22'),
(11, 41, '2021-06-22'),
(12, 42, '2021-06-21'),
(13, 43, '2021-06-21'),
(14, 44, '2021-06-21'),
(15, 45, '2021-06-21'),
(16, 46, '2021-06-21'),
(17, 47, '2021-06-21'),
(73, 129, '2021-06-21'),
(74, 130, '2021-06-23'),
(76, 133, '2021-06-09'),
(77, 134, '2021-06-07'),
(78, 135, '2021-06-14'),
(80, 139, '2021-06-09'),
(100, 161, '2021-06-18');

-- --------------------------------------------------------

--
-- Table structure for table `Users`
--

CREATE TABLE `Users` (
  `Id` int(10) UNSIGNED NOT NULL,
  `Name` varchar(100) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Email` varchar(100) DEFAULT NULL,
  `Theme` tinyint(1) NOT NULL DEFAULT 0,
  `Id_rol` int(10) UNSIGNED NOT NULL DEFAULT 1,
  `Id_status` int(10) UNSIGNED NOT NULL DEFAULT 10
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `Users`
--

INSERT INTO `Users` (`Id`, `Name`, `Password`, `Email`, `Theme`, `Id_rol`, `Id_status`) VALUES
(7, 'probando', 'David', 'JuanDavid@hotmail.com', 0, 1, 1),
(8, 'probando', 'algo', 'pedro@gmail.com', 0, 1, 10),
(9, 'probando', 'algo', 'alguien@gmail.com', 0, 1, 10),
(10, 'probando', '$2a$10$Hs3nUL1NQcwIVFBi0bNNKegpx0Bf98WCRuoDgNi4DnudZPzg.uuFq', 'ana@gmail.com', 0, 1, 10),
(11, 'probando', '$2a$10$BLlOyUG6wyxBVC8D6jvNM.RmUG/4vB8fVxBVu.TPpnGeIiXy4S6QK', 'antonio32@gmail.com', 0, 1, 10),
(12, 'probando', '$2a$10$tn.c6dlqMYrymq/orVwwPOiFizRm2mI.HBDorNFugwxpPgU4ni20q', 'rafit@gmail.com', 0, 1, 10),
(13, 'probando', '$2a$10$yUeJwXtxtZuB5SBX6mEYX.Jq51uZbS5kQY3HP0AgGmN0nzFTRCftu', 'laurta@gmail.com', 0, 1, 10),
(14, 'probando', '$2a$10$JQiyigosl2uc246.U9aK8u6SW8HjxlTxTiI/NQoF3zoz/b3obQTky', 'pedrito@gmail.com', 0, 1, 10),
(15, 'probando', '$2a$10$49wy/GVwPktRDwg9EImZROEAQpqbST/dS4MbPYAWJAQezIfok3LH.', 'antonio@gmail.com', 0, 1, 10),
(16, 'probando', '$2a$10$vQcjZUkfZBUb7dOI0TEGQ.GhutOvbvGFv65o2Hb0EMyPWBy5UIpZ6', 'pedritoRechulon@yahoo.com', 0, 1, 10),
(17, 'probandolololo', '$2a$10$3T1CBEvhIY0AiDIMgHJHa.MBr1a2lT7556xkw/8EJM3YAUHiqegYe', 'prueba@gmail.com', 0, 1, 10),
(18, 'pruebados', '$2a$10$DXClnbFpCSAO8jaAnhKaHOQDhHIwlBrkDRmzT.q/LKIrrgwK8ZX3G', '', 0, 1, 10);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Areas`
--
ALTER TABLE `Areas`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Id_user` (`Id_user`);

--
-- Indexes for table `Goals`
--
ALTER TABLE `Goals`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Id_target` (`Id_target`);

--
-- Indexes for table `Rols`
--
ALTER TABLE `Rols`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `Status`
--
ALTER TABLE `Status`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `Targets`
--
ALTER TABLE `Targets`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Id_parent` (`Id_parent`),
  ADD KEY `Id_usu` (`Id_usu`),
  ADD KEY `Id_area` (`Id_area`),
  ADD KEY `Id_status` (`Id_status`);

--
-- Indexes for table `Task`
--
ALTER TABLE `Task`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Id_target` (`Id_target`);

--
-- Indexes for table `Users`
--
ALTER TABLE `Users`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `id_rol_users` (`Id_rol`),
  ADD KEY `id_status_users` (`Id_status`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Areas`
--
ALTER TABLE `Areas`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=235;

--
-- AUTO_INCREMENT for table `Goals`
--
ALTER TABLE `Goals`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- AUTO_INCREMENT for table `Rols`
--
ALTER TABLE `Rols`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `Status`
--
ALTER TABLE `Status`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=71;

--
-- AUTO_INCREMENT for table `Targets`
--
ALTER TABLE `Targets`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=162;

--
-- AUTO_INCREMENT for table `Task`
--
ALTER TABLE `Task`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=101;

--
-- AUTO_INCREMENT for table `Users`
--
ALTER TABLE `Users`
  MODIFY `Id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Areas`
--
ALTER TABLE `Areas`
  ADD CONSTRAINT `Areas_ibfk_1` FOREIGN KEY (`Id_user`) REFERENCES `Users` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Goals`
--
ALTER TABLE `Goals`
  ADD CONSTRAINT `Goals_ibfk_1` FOREIGN KEY (`Id_target`) REFERENCES `Targets` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Targets`
--
ALTER TABLE `Targets`
  ADD CONSTRAINT `Targets_ibfk_1` FOREIGN KEY (`Id_parent`) REFERENCES `Targets` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Targets_ibfk_2` FOREIGN KEY (`Id_usu`) REFERENCES `Users` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Targets_ibfk_3` FOREIGN KEY (`Id_area`) REFERENCES `Areas` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Targets_ibfk_4` FOREIGN KEY (`Id_status`) REFERENCES `Status` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Task`
--
ALTER TABLE `Task`
  ADD CONSTRAINT `Task_ibfk_1` FOREIGN KEY (`Id_target`) REFERENCES `Targets` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Users`
--
ALTER TABLE `Users`
  ADD CONSTRAINT `id_rol_users` FOREIGN KEY (`Id_rol`) REFERENCES `Rols` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `id_status_users` FOREIGN KEY (`Id_status`) REFERENCES `Status` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
