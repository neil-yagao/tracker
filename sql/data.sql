 -- MySQL dump 10.13  Distrib 5.7.9, for Win64 (x86_64)
--
-- Host: localhost    Database: powerlift
-- ------------------------------------------------------
-- Server version	5.6.34

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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movement`
--

LOCK TABLES `movement` WRITE;
/*!40000 ALTER TABLE `movement` DISABLE KEYS */;
INSERT INTO `movement` VALUES (1,'Barbell Clean','FullBody;UpperBody;LowerBody','','Barbell Clean Movement for strength training or CrossFit'),(2,'Squart','Leg','','Squart'),(3,'Deadlift','Leg','','Deadlift'),(4,'Front Squart','Leg','','Front Squart'),(5,'Dumbbell Walking Lunge','Leg','','Dumbbell Walking Lunge'),(6,'Bench Press','Upper Body','','Bench Press'),(7,'Barbell Row','Upper Body','','Barbell Row'),(8,'Overhead Barbell Press','Upper Body','','Overhead Barbell Press'),(9,'Barbell Biceps Curl','Upper Body','','Barbell Biceps Curl'),(10,'Incline Dumbbell Bench Press','Upper Body','','Incline Dumbbell Bench Press'),(11,'Dip - BodyWeight','Upper Body','','Dip'),(12,'Incline Barbell Bench Press','Chest','','Barbell version of incline bench press'),(13,'Dumbbell Incline Bench Press','Chest','','Dumbbell version of Incline bench Press');
/*!40000 ALTER TABLE `movement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `useridentity` varchar(128) NOT NULL DEFAULT '',
  `username` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `useridentity_UNIQUE` (`useridentity`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'d68e939882371200637d5024b360fc20','neil',''),(2,'1a833da63a6b7e20098dae06d06602e1','helloWorld',''),(4,'5d41402abc4b2a76b9719d911017c592','hello',''),(6,'d41d8cd98f00b204e9800998ecf8427e','','');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_workout`
--

DROP TABLE IF EXISTS `user_workout`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_workout` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `workout` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_workout`
--

LOCK TABLES `user_workout` WRITE;
/*!40000 ALTER TABLE `user_workout` DISABLE KEYS */;
INSERT INTO `user_workout` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(6,1,6),(7,1,7),(8,1,8);
/*!40000 ALTER TABLE `user_workout` ENABLE KEYS */;
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
  `acheive_weight` double DEFAULT '0',
  `sequence` smallint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `working_set`
--

LOCK TABLES `working_set` WRITE;
/*!40000 ALTER TABLE `working_set` DISABLE KEYS */;
INSERT INTO `working_set` VALUES (1,1,2,87.5,6,6,0,1),(2,1,2,87.5,6,6,0,2),(3,1,2,87.5,6,6,0,3),(4,1,3,115,6,6,0,1),(5,1,3,115,6,6,0,2),(6,1,3,115,6,6,0,3),(7,1,4,60,8,8,0,1),(8,1,4,60,8,8,0,2),(9,1,4,60,8,8,0,3),(10,1,5,21,8,8,0,1),(11,1,5,21,8,8,0,2),(12,1,5,21,8,8,0,3),(13,2,2,90,6,0,0,1),(14,2,2,90,6,0,0,2),(15,2,2,90,6,0,0,3),(16,2,3,117.5,6,0,0,1),(17,2,3,117.5,6,0,0,2),(18,2,3,117.5,6,0,0,3),(19,2,4,62.5,8,0,0,1),(20,2,4,62.5,8,0,0,2),(21,2,4,62.5,8,0,0,3),(22,2,5,23.5,8,0,0,1),(23,2,5,23.5,8,0,0,2),(24,2,5,23.5,8,0,0,3),(25,3,2,92.5,6,0,0,1),(26,3,2,92.5,6,0,0,2),(27,3,2,92.5,6,0,0,3),(28,3,3,120,6,0,0,1),(29,3,3,120,6,0,0,2),(30,3,3,120,6,0,0,3),(31,3,4,65,8,0,0,1),(32,3,4,65,8,0,0,2),(33,3,4,65,8,0,0,3),(34,3,5,26,8,0,0,1),(35,3,5,26,8,0,0,2),(36,3,5,26,8,0,0,3),(37,4,2,95,6,0,0,1),(38,4,2,95,6,0,0,2),(39,4,2,95,6,0,0,3),(40,4,3,122.5,6,0,0,1),(41,4,3,122.5,6,0,0,2),(42,4,3,122.5,6,0,0,3),(43,4,4,67.5,8,0,0,1),(44,4,4,67.5,8,0,0,2),(45,4,4,67.5,8,0,0,3),(46,4,5,28.5,8,0,0,1),(47,4,5,28.5,8,0,0,2),(48,4,5,28.5,8,0,0,3),(49,5,6,85,6,0,0,1),(50,5,6,85,6,0,0,2),(51,5,6,85,6,0,0,3),(52,5,6,85,6,0,0,4),(53,5,7,65,6,0,0,1),(54,5,7,65,6,0,0,2),(55,5,7,65,6,0,0,3),(56,5,7,65,6,0,0,4),(57,5,8,40,8,0,0,1),(58,5,8,40,8,0,0,2),(59,5,9,37.5,8,0,0,1),(60,5,9,37.5,8,0,0,2),(61,5,9,37.5,8,0,0,3),(62,5,10,23.5,8,0,0,1),(63,5,10,23.5,8,0,0,2),(64,5,10,23.5,8,0,0,3),(65,5,11,5,10,0,0,1),(66,5,11,5,10,0,0,2),(67,5,11,5,10,0,0,3),(68,6,6,87.5,6,0,0,1),(69,6,6,87.5,6,0,0,2),(70,6,6,87.5,6,0,0,3),(71,6,6,87.5,6,0,0,4),(72,6,7,67.5,6,0,0,1),(73,6,7,67.5,6,0,0,2),(74,6,7,67.5,6,0,0,3),(75,6,7,67.5,6,0,0,4),(76,6,8,42.5,8,0,0,1),(77,6,8,42.5,8,0,0,2),(78,6,9,40,8,0,0,1),(79,6,9,40,8,0,0,2),(80,6,9,40,8,0,0,3),(81,6,10,26,8,0,0,1),(82,6,10,26,8,0,0,2),(83,6,10,26,8,0,0,3),(84,6,11,7.5,10,0,0,1),(85,6,11,7.5,10,0,0,2),(86,6,11,7.5,10,0,0,3),(87,7,6,90,6,0,0,1),(88,7,6,90,6,0,0,2),(89,7,6,90,6,0,0,3),(90,7,6,90,6,0,0,4),(91,7,7,70,6,0,0,1),(92,7,7,70,6,0,0,2),(93,7,7,70,6,0,0,3),(94,7,7,70,6,0,0,4),(95,7,8,45,8,0,0,1),(96,7,8,45,8,0,0,2),(97,7,9,42.5,8,0,0,1),(98,7,9,42.5,8,0,0,2),(99,7,9,42.5,8,0,0,3),(100,7,10,28.5,8,0,0,1),(101,7,10,28.5,8,0,0,2),(102,7,10,28.5,8,0,0,3),(103,7,11,10,10,0,0,1),(104,7,11,10,10,0,0,2),(105,7,11,10,10,0,0,3),(106,8,6,92.5,6,0,0,1),(107,8,6,92.5,6,0,0,2),(108,8,6,92.5,6,0,0,3),(109,8,6,92.5,6,0,0,4),(110,8,7,72.5,6,0,0,1),(111,8,7,72.5,6,0,0,2),(112,8,7,72.5,6,0,0,3),(113,8,7,72.5,6,0,0,4),(114,8,8,47.5,8,0,0,1),(115,8,8,47.5,8,0,0,2),(116,8,9,45,8,0,0,1),(117,8,9,45,8,0,0,2),(118,8,9,45,8,0,0,3),(119,8,10,31,8,0,0,1),(120,8,10,31,8,0,0,2),(121,8,10,31,8,0,0,3),(122,8,11,12.5,10,0,0,1),(123,8,11,12.5,10,0,0,2),(124,8,11,12.5,10,0,0,3);
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workout`
--

LOCK TABLES `workout` WRITE;
/*!40000 ALTER TABLE `workout` DISABLE KEYS */;
INSERT INTO `workout` VALUES (1,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-22',1),(2,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-29',0),(3,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-06',0),(4,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-13',0),(5,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-11-23',0),(6,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-11-30',0),(7,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-12-07',0),(8,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-12-14',0);
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

-- Dump completed on 2016-12-02 11:41:35
