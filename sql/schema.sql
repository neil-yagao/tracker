CREATE SCHEMA `powerlift` ;


/*user name powerlift password p@ssw0rd*/

-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: powerlift
-- ------------------------------------------------------
-- Server version 5.6.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `movement`
--

DROP TABLE IF EXISTS `movement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `target_muscle` varchar(128) NOT NULL DEFAULT '',
  `secondary_muscle` varchar(256) DEFAULT '',
  `description` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movement`
--

LOCK TABLES `movement` WRITE;
/*!40000 ALTER TABLE `movement` DISABLE KEYS */;
INSERT INTO `movement` VALUES (1,'movement1','UpperBody','','movement1'),(2,'Squart','Leg','','Squart'),(3,'Deadlift','Leg','','Deadlift'),(4,'Front Squart','Leg','','Front Squart'),(5,'Dumbbell Walking Lunge','Leg','','Dumbbell Walking Lunge');
/*!40000 ALTER TABLE `movement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `working_set`
--

DROP TABLE IF EXISTS `working_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `working_set` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workout` int(11) NOT NULL,
  `movement` int(11) NOT NULL,
  `target_weight` double NOT NULL DEFAULT '0',
  `target_number` int(11) NOT NULL DEFAULT '10',
  `acheive_number` int(11) NOT NULL,
  `sequence` smallint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `working_set`
--

LOCK TABLES `working_set` WRITE;
/*!40000 ALTER TABLE `working_set` DISABLE KEYS */;
INSERT INTO `working_set` VALUES (1,1,2,87.5,6,0,1),(2,1,2,87.5,6,0,2),(3,1,2,87.5,6,0,3),(4,1,3,115,6,0,1),(5,1,3,115,6,0,2),(6,1,3,115,6,0,3),(7,1,4,60,8,0,1),(8,1,4,60,8,0,2),(9,1,4,60,8,0,3),(10,1,5,21,8,0,1),(11,1,5,21,8,0,2),(12,1,5,21,8,0,3),(13,2,2,90,6,0,1),(14,2,2,90,6,0,2),(15,2,2,90,6,0,3),(16,2,3,117.5,6,0,1),(17,2,3,117.5,6,0,2),(18,2,3,117.5,6,0,3),(19,2,4,62.5,8,0,1),(20,2,4,62.5,8,0,2),(21,2,4,62.5,8,0,3),(22,2,5,23.5,8,0,1),(23,2,5,23.5,8,0,2),(24,2,5,23.5,8,0,3),(25,3,2,92.5,6,0,1),(26,3,2,92.5,6,0,2),(27,3,2,92.5,6,0,3),(28,3,3,120,6,0,1),(29,3,3,120,6,0,2),(30,3,3,120,6,0,3),(31,3,4,65,8,0,1),(32,3,4,65,8,0,2),(33,3,4,65,8,0,3),(34,3,5,26,8,0,1),(35,3,5,26,8,0,2),(36,3,5,26,8,0,3),(37,4,2,95,6,0,1),(38,4,2,95,6,0,2),(39,4,2,95,6,0,3),(40,4,3,122.5,6,0,1),(41,4,3,122.5,6,0,2),(42,4,3,122.5,6,0,3),(43,4,4,67.5,8,0,1),(44,4,4,67.5,8,0,2),(45,4,4,67.5,8,0,3),(46,4,5,28.5,8,0,1),(47,4,5,28.5,8,0,2),(48,4,5,28.5,8,0,3);
/*!40000 ALTER TABLE `working_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workout`
--

DROP TABLE IF EXISTS `workout`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `workout` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `target` varchar(128) DEFAULT '',
  `description` varchar(512) DEFAULT NULL,
  `perform_date` varchar(32) NOT NULL DEFAULT '2016-11-11',
  `is_finalized` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workout`
--

LOCK TABLES `workout` WRITE;
/*!40000 ALTER TABLE `workout` DISABLE KEYS */;
INSERT INTO `workout` VALUES (1,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-22',0),(2,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-29',0),(3,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-06',0),(4,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-13',0);
/*!40000 ALTER TABLE `workout` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-22 16:25:22
