mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 8.0.37, for Linux (aarch64)
--
-- Host: localhost    Database: shopdevgo
-- ------------------------------------------------------
-- Server version	8.0.37

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `shopdevgo`
--

/*!40000 DROP DATABASE IF EXISTS `shopdevgo`*/;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `shopdevgo` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `shopdevgo`;

--
-- Table structure for table `go_crm_user`
--

DROP TABLE IF EXISTS `go_crm_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `go_crm_user` (
  `usr_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
  `usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'User name',
  `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
  `usr_created_at` int NOT NULL DEFAULT '0' COMMENT 'PCreation time',
  `usr_update_at` int NOT NULL DEFAULT '0' COMMENT 'Update time',
  `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Createion IP',
  `usr_last_login_at` int NOT NULL DEFAULT '0' COMMENT 'last login time',
  `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
  `usr_login_times` int NOT NULL DEFAULT '0' COMMENT 'Login times',
  `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disnable, -1:deleted',
  PRIMARY KEY (`usr_id`),
  KEY `idx_email` (`usr_email`),
  KEY `idx_phone` (`usr_phone`),
  KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Account';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-23  4:25:08
