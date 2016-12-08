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
  `dividable` tinyint(2) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movement`
--

LOCK TABLES `movement` WRITE;
/*!40000 ALTER TABLE `movement` DISABLE KEYS */;
INSERT INTO `movement` VALUES (1,'Barbell Clean','Traps','','Barbell Clean Movement for strength training or CrossFit',1),(2,'Squart','Leg','','Squart',1),(3,'Deadlift','Leg','','Deadlift',1),(4,'Front Squart','Leg','','Front Squart',1),(5,'Dumbbell Walking Lunge','Leg','','Dumbbell Walking Lunge',0),(6,'Bench Press','Upper Body','','Bench Press',1),(7,'Barbell Row','Upper Body','','Barbell Row',1),(8,'Overhead Barbell Press','Upper Body','','Overhead Barbell Press',1),(9,'Barbell Biceps Curl','Upper Body','','Barbell Biceps Curl',1),(10,'Incline Dumbbell Bench Press','Upper Body','','Incline Dumbbell Bench Press',0),(11,'Dip - BodyWeight','Upper Body','','Dip',0),(12,'Incline Barbell Bench Press','Chest','','Barbell version of incline bench press',1),(13,'Dumbbell Incline Bench Press','Chest','','Dumbbell version of Incline bench Press',0),(15,'Barbell Shrug','Traps','','Stand up straight with your feet at shoulder width as you hold a barbell with both hands in front of you using a pronated grip (palms facing the thighs). Tip: Your hands should be a little wider than shoulder width apart. You can use wrist wraps for this exercise for a better grip. This will be your starting position.\nRaise your shoulders up as far as you can go as you breathe out and hold the contraction for a second. Tip: Refrain from trying to lift the barbell by using your biceps.\nSlowly return to the starting position as you breathe in.\nRepeat for the recommended amount of repetitions',1),(17,'Dumbbell Shrug','Traps','','Stand erect with a dumbbell on each hand (palms facing your torso), arms extended on the sides.\nLift the dumbbells by elevating the shoulders as high as possible while you exhale. Hold the contraction at the top for a second. Tip: The arms should remain extended at all times. Refrain from using the biceps to help lift the dumbbells. Only the shoulders should be moving up and down.\nLower the dumbbells back to the original position.\nRepeat for the recommended amount of repetitions.',0),(18,'Upright Row','Traps','','To begin, stand on an exercise band so that tension begins at arm length. Grasp the handles using a pronated (palms facing your thighs) grip that is slightly less than shoulder width. The handles should be resting on top of your thighs. Your arms should be extended with a slight bend at the elbows and your back should be straight. This will be your starting position.\nUse your side shoulders to lift the handles as you exhale. The handles should be close to the body as you move them up. Continue to lift the handles until they nearly touches your chin. Tip: Your elbows should drive the motion. As you lift the handles, your elbows should always be higher than your forearms. Also, keep your torso stationary and pause for a second at the top of the movement.\nLower the handles back down slowly to the starting position. Inhale as you perform this portion of the movement.\nRepeat for the recommended amount of repetitions.',1),(20,'ALTERNATE INCLINE DUMBBELL CURL','Biceps','','Sit down on an incline bench with a dumbbell in each hand being held at arms length. Tip: Keep the elbows close to the torso.This will be your starting position.\nWhile holding the upper arm stationary, curl the right weight forward while contracting the biceps as you breathe out. As you do so, rotate the hand so that the palm is facing up. Continue the movement until your biceps is fully contracted and the dumbbells are at shoulder level. Hold the contracted position for a second as you squeeze the biceps. Tip: Only the forearms should move.\nSlowly begin to bring the dumbbell back to starting position as your breathe in.\nRepeat the movement with the left hand. This equals one repetition.\nContinue alternating in this manner for the recommended amount of repetitions.',0),(23,'Barbell Curl','Biceps','','Stand up with your torso upright while holding a barbell at a shoulder-width grip. The palm of your hands should be facing forward and the elbows should be close to the torso. This will be your starting position.\nWhile holding the upper arms stationary, curl the weights forward while contracting the biceps as you breathe out. Tip: Only the forearms should move.\nContinue the movement until your biceps are fully contracted and the bar is at shoulder level. Hold the contracted position for a second and squeeze the biceps hard.\nSlowly begin to bring the bar back to starting position as your breathe in.\nRepeat for the recommended amount of repetitions.',1),(24,'Barbell PREACHER CURL','Biceps','','Place a preacher bench about 2 feet in front of a pulley machine.\nAttach a straight bar to the low pulley.\nSit at the preacher bench with your elbow and upper arms firmly on top of the bench pad and have someone hand you the bar from the low pulley.\nGrab the bar and fully extend your arms on top of the preacher bench pad. This will be your starting position.\nNow start pilling the weight up towards your shoulders and squeeze the biceps hard at the top of the movement. Exhale as you perform this motion. Also, hold for a second at the top.\nNow slowly lower the weight to the starting position.\nRepeat for the recommended amount of repetitions.',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_workout`
--

LOCK TABLES `user_workout` WRITE;
/*!40000 ALTER TABLE `user_workout` DISABLE KEYS */;
INSERT INTO `user_workout` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(6,1,6),(7,1,7),(8,1,8),(9,1,41),(10,1,42),(11,1,43),(12,1,44);
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
) ENGINE=InnoDB AUTO_INCREMENT=3194 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `working_set`
--

LOCK TABLES `working_set` WRITE;
/*!40000 ALTER TABLE `working_set` DISABLE KEYS */;
INSERT INTO `working_set` VALUES (1,1,2,87.5,6,6,0,1),(2,1,2,87.5,6,6,0,2),(3,1,2,87.5,6,6,0,3),(4,1,3,115,6,6,0,1),(5,1,3,115,6,6,0,2),(6,1,3,115,6,6,0,3),(7,1,4,60,8,8,0,1),(8,1,4,60,8,8,0,2),(9,1,4,60,8,8,0,3),(10,1,5,21,8,8,0,1),(11,1,5,21,8,8,0,2),(12,1,5,21,8,8,0,3),(13,2,2,90,6,0,0,1),(14,2,2,90,6,0,0,2),(15,2,2,90,6,0,0,3),(16,2,3,117.5,6,0,0,1),(17,2,3,117.5,6,0,0,2),(18,2,3,117.5,6,0,0,3),(19,2,4,62.5,8,0,0,1),(20,2,4,62.5,8,0,0,2),(21,2,4,62.5,8,0,0,3),(22,2,5,23.5,8,0,0,1),(23,2,5,23.5,8,0,0,2),(24,2,5,23.5,8,0,0,3),(25,3,2,92.5,6,0,0,1),(26,3,2,92.5,6,0,0,2),(27,3,2,92.5,6,0,0,3),(28,3,3,120,6,0,0,1),(29,3,3,120,6,0,0,2),(30,3,3,120,6,0,0,3),(31,3,4,65,8,0,0,1),(32,3,4,65,8,0,0,2),(33,3,4,65,8,0,0,3),(34,3,5,26,8,0,0,1),(35,3,5,26,8,0,0,2),(36,3,5,26,8,0,0,3),(37,4,2,95,6,0,0,1),(38,4,2,95,6,0,0,2),(39,4,2,95,6,0,0,3),(40,4,3,122.5,6,0,0,1),(41,4,3,122.5,6,0,0,2),(42,4,3,122.5,6,0,0,3),(43,4,4,67.5,8,0,0,1),(44,4,4,67.5,8,0,0,2),(45,4,4,67.5,8,0,0,3),(46,4,5,28.5,8,0,0,1),(47,4,5,28.5,8,0,0,2),(48,4,5,28.5,8,0,0,3),(49,5,6,85,6,6,85,1),(50,5,6,85,6,6,85,2),(51,5,6,85,6,6,85,3),(52,5,6,85,6,6,85,4),(53,5,7,65,6,6,65,1),(54,5,7,65,6,6,65,2),(55,5,7,65,6,6,65,3),(56,5,7,65,6,6,65,4),(57,5,8,40,8,7,40,1),(58,5,8,40,8,8,40,2),(59,5,9,37.5,8,8,37.5,1),(60,5,9,37.5,8,8,37.5,2),(61,5,9,37.5,8,8,37.5,3),(62,5,10,23.5,8,8,23.5,1),(63,5,10,23.5,8,8,23.5,2),(64,5,10,23.5,8,8,23.5,3),(65,5,11,5,10,10,5,1),(66,5,11,5,10,10,5,2),(67,5,11,5,10,10,5,3),(68,6,6,87.5,6,0,0,1),(69,6,6,87.5,6,0,0,2),(70,6,6,87.5,6,0,0,3),(71,6,6,87.5,6,0,0,4),(72,6,7,67.5,6,0,0,1),(73,6,7,67.5,6,0,0,2),(74,6,7,67.5,6,0,0,3),(75,6,7,67.5,6,0,0,4),(76,6,8,42.5,8,0,0,1),(77,6,8,42.5,8,0,0,2),(78,6,9,40,8,0,0,1),(79,6,9,40,8,0,0,2),(80,6,9,40,8,0,0,3),(81,6,10,26,8,0,0,1),(82,6,10,26,8,0,0,2),(83,6,10,26,8,0,0,3),(84,6,11,7.5,10,0,0,1),(85,6,11,7.5,10,0,0,2),(86,6,11,7.5,10,0,0,3),(87,7,6,90,6,0,0,1),(88,7,6,90,6,0,0,2),(89,7,6,90,6,0,0,3),(90,7,6,90,6,0,0,4),(91,7,7,70,6,0,0,1),(92,7,7,70,6,0,0,2),(93,7,7,70,6,0,0,3),(94,7,7,70,6,0,0,4),(95,7,8,45,8,0,0,1),(96,7,8,45,8,0,0,2),(97,7,9,42.5,8,0,0,1),(98,7,9,42.5,8,0,0,2),(99,7,9,42.5,8,0,0,3),(100,7,10,28.5,8,0,0,1),(101,7,10,28.5,8,0,0,2),(102,7,10,28.5,8,0,0,3),(103,7,11,10,10,0,0,1),(104,7,11,10,10,0,0,2),(105,7,11,10,10,0,0,3),(106,8,6,92.5,6,0,0,1),(107,8,6,92.5,6,0,0,2),(108,8,6,92.5,6,0,0,3),(109,8,6,92.5,6,0,0,4),(110,8,7,72.5,6,0,0,1),(111,8,7,72.5,6,0,0,2),(112,8,7,72.5,6,0,0,3),(113,8,7,72.5,6,0,0,4),(114,8,8,47.5,8,0,0,1),(115,8,8,47.5,8,0,0,2),(116,8,9,45,8,0,0,1),(117,8,9,45,8,0,0,2),(118,8,9,45,8,0,0,3),(119,8,10,31,8,0,0,1),(120,8,10,31,8,0,0,2),(121,8,10,31,8,0,0,3),(122,8,11,12.5,10,0,0,1),(123,8,11,12.5,10,0,0,2),(124,8,11,12.5,10,0,0,3),(2564,17,3,20,6,0,0,0),(2565,17,3,30,6,0,0,0),(2566,17,3,40,6,0,0,0),(2567,17,3,50,6,0,0,0),(2568,17,3,60,6,0,0,0),(2569,17,3,70,6,0,0,0),(2570,17,3,80,6,0,0,0),(2571,17,3,90,6,0,0,0),(2572,17,3,120,6,0,0,1),(2573,17,3,120,6,0,0,2),(2574,17,3,120,6,0,0,3),(2575,17,3,120,6,0,0,4),(2576,17,6,20,6,0,0,0),(2577,17,6,25,6,0,0,0),(2578,17,6,30,6,0,0,0),(2579,17,6,35,6,0,0,0),(2580,17,6,40,6,0,0,0),(2581,17,6,45,6,0,0,0),(2582,17,6,60,6,0,0,1),(2583,17,6,60,6,0,0,2),(2584,17,6,60,6,0,0,3),(2585,17,6,60,6,0,0,4),(2586,17,8,20,6,0,0,0),(2587,17,8,30,6,0,0,1),(2588,17,8,30,6,0,0,2),(2589,17,8,30,6,0,0,3),(2590,17,8,30,6,0,0,4),(2591,17,5,5.2,6,0,0,0),(2592,17,5,15.600000000000001,6,0,0,0),(2593,17,5,26,6,0,0,1),(2594,17,5,26,6,0,0,2),(2595,17,5,26,6,0,0,3),(2596,17,5,26,6,0,0,4),(2597,18,3,20,6,0,0,0),(2598,18,3,30,6,0,0,0),(2599,18,3,40,6,0,0,0),(2600,18,3,50,6,0,0,0),(2601,18,3,60,6,0,0,0),(2602,18,3,70,6,0,0,0),(2603,18,3,80,6,0,0,0),(2604,18,3,90,6,0,0,0),(2605,18,3,122.5,6,0,0,1),(2606,18,3,122.5,6,0,0,2),(2607,18,3,122.5,6,0,0,3),(2608,18,3,122.5,6,0,0,4),(2609,18,6,20,6,0,0,0),(2610,18,6,25,6,0,0,0),(2611,18,6,30,6,0,0,0),(2612,18,6,35,6,0,0,0),(2613,18,6,40,6,0,0,0),(2614,18,6,45,6,0,0,0),(2615,18,6,50,6,0,0,0),(2616,18,6,62.5,6,0,0,1),(2617,18,6,62.5,6,0,0,2),(2618,18,6,62.5,6,0,0,3),(2619,18,6,62.5,6,0,0,4),(2620,18,8,20,6,0,0,0),(2621,18,8,32.5,6,0,0,1),(2622,18,8,32.5,6,0,0,2),(2623,18,8,32.5,6,0,0,3),(2624,18,8,32.5,6,0,0,4),(2625,18,5,5.7,6,0,0,0),(2626,18,5,17.1,6,0,0,0),(2627,18,5,28.5,6,0,0,1),(2628,18,5,28.5,6,0,0,2),(2629,18,5,28.5,6,0,0,3),(2630,18,5,28.5,6,0,0,4),(2631,19,3,20,6,0,0,0),(2632,19,3,30,6,0,0,0),(2633,19,3,40,6,0,0,0),(2634,19,3,50,6,0,0,0),(2635,19,3,60,6,0,0,0),(2636,19,3,70,6,0,0,0),(2637,19,3,80,6,0,0,0),(2638,19,3,90,6,0,0,0),(2639,19,3,100,6,0,0,0),(2640,19,3,125,6,0,0,1),(2641,19,3,125,6,0,0,2),(2642,19,3,125,6,0,0,3),(2643,19,3,125,6,0,0,4),(2644,19,6,20,6,0,0,0),(2645,19,6,25,6,0,0,0),(2646,19,6,30,6,0,0,0),(2647,19,6,35,6,0,0,0),(2648,19,6,40,6,0,0,0),(2649,19,6,45,6,0,0,0),(2650,19,6,50,6,0,0,0),(2651,19,6,65,6,0,0,1),(2652,19,6,65,6,0,0,2),(2653,19,6,65,6,0,0,3),(2654,19,6,65,6,0,0,4),(2655,19,8,20,6,0,0,0),(2656,19,8,25,6,0,0,0),(2657,19,8,35,6,0,0,1),(2658,19,8,35,6,0,0,2),(2659,19,8,35,6,0,0,3),(2660,19,8,35,6,0,0,4),(2661,19,5,6.2,6,0,0,0),(2662,19,5,18.6,6,0,0,0),(2663,19,5,31,6,0,0,1),(2664,19,5,31,6,0,0,2),(2665,19,5,31,6,0,0,3),(2666,19,5,31,6,0,0,4),(2667,20,3,20,6,0,0,0),(2668,20,3,30,6,0,0,0),(2669,20,3,40,6,0,0,0),(2670,20,3,50,6,0,0,0),(2671,20,3,60,6,0,0,0),(2672,20,3,70,6,0,0,0),(2673,20,3,80,6,0,0,0),(2674,20,3,90,6,0,0,0),(2675,20,3,100,6,0,0,0),(2676,20,3,127.5,6,0,0,1),(2677,20,3,127.5,6,0,0,2),(2678,20,3,127.5,6,0,0,3),(2679,20,3,127.5,6,0,0,4),(2680,20,6,20,6,0,0,0),(2681,20,6,30,6,0,0,0),(2682,20,6,40,6,0,0,0),(2683,20,6,50,6,0,0,0),(2684,20,6,67.5,6,0,0,1),(2685,20,6,67.5,6,0,0,2),(2686,20,6,67.5,6,0,0,3),(2687,20,6,67.5,6,0,0,4),(2688,20,8,20,6,0,0,0),(2689,20,8,25,6,0,0,0),(2690,20,8,30,6,0,0,0),(2691,20,8,37.5,6,0,0,1),(2692,20,8,37.5,6,0,0,2),(2693,20,8,37.5,6,0,0,3),(2694,20,8,37.5,6,0,0,4),(2695,20,5,6.7,6,0,0,0),(2696,20,5,11.7,6,0,0,0),(2697,20,5,16.7,6,0,0,0),(2698,20,5,21.7,6,0,0,0),(2699,20,5,26.7,6,0,0,0),(2700,20,5,33.5,6,0,0,1),(2701,20,5,33.5,6,0,0,2),(2702,20,5,33.5,6,0,0,3),(2703,20,5,33.5,6,0,0,4),(2704,21,3,20,6,0,0,0),(2705,21,3,30,6,0,0,0),(2706,21,3,40,6,0,0,0),(2707,21,3,50,6,0,0,0),(2708,21,3,60,6,0,0,0),(2709,21,3,70,6,0,0,0),(2710,21,3,80,6,0,0,0),(2711,21,3,90,6,0,0,0),(2712,21,3,100,6,0,0,0),(2713,21,3,120,6,0,0,1),(2714,21,3,120,6,0,0,2),(2715,21,3,120,6,0,0,3),(2716,21,3,120,6,0,0,4),(2717,21,6,20,6,0,0,0),(2718,21,6,25,6,0,0,0),(2719,21,6,30,6,0,0,0),(2720,21,6,35,6,0,0,0),(2721,21,6,40,6,0,0,0),(2722,21,6,45,6,0,0,0),(2723,21,6,50,6,0,0,0),(2724,21,6,60,6,0,0,1),(2725,21,6,60,6,0,0,2),(2726,21,6,60,6,0,0,3),(2727,21,6,60,6,0,0,4),(2728,21,8,20,6,0,0,0),(2729,21,8,25,6,0,0,0),(2730,21,8,30,6,0,0,1),(2731,21,8,30,6,0,0,2),(2732,21,8,30,6,0,0,3),(2733,21,8,30,6,0,0,4),(2734,21,5,5,6,0,0,0),(2735,21,5,15,6,0,0,0),(2736,21,5,26,6,0,0,1),(2737,21,5,26,6,0,0,2),(2738,21,5,26,6,0,0,3),(2739,21,5,26,6,0,0,4),(2740,22,3,20,6,0,0,0),(2741,22,3,30,6,0,0,0),(2742,22,3,40,6,0,0,0),(2743,22,3,50,6,0,0,0),(2744,22,3,60,6,0,0,0),(2745,22,3,70,6,0,0,0),(2746,22,3,80,6,0,0,0),(2747,22,3,90,6,0,0,0),(2748,22,3,100,6,0,0,0),(2749,22,3,110,6,0,0,0),(2750,22,3,122.5,6,0,0,1),(2751,22,3,122.5,6,0,0,2),(2752,22,3,122.5,6,0,0,3),(2753,22,3,122.5,6,0,0,4),(2754,22,6,20,6,0,0,0),(2755,22,6,25,6,0,0,0),(2756,22,6,30,6,0,0,0),(2757,22,6,35,6,0,0,0),(2758,22,6,40,6,0,0,0),(2759,22,6,45,6,0,0,0),(2760,22,6,50,6,0,0,0),(2761,22,6,55,6,0,0,0),(2762,22,6,62.5,6,0,0,1),(2763,22,6,62.5,6,0,0,2),(2764,22,6,62.5,6,0,0,3),(2765,22,6,62.5,6,0,0,4),(2766,22,8,20,6,0,0,0),(2767,22,8,25,6,0,0,0),(2768,22,8,32.5,6,0,0,1),(2769,22,8,32.5,6,0,0,2),(2770,22,8,32.5,6,0,0,3),(2771,22,8,32.5,6,0,0,4),(2772,22,5,5,6,0,0,0),(2773,22,5,15,6,0,0,0),(2774,22,5,28.5,6,0,0,1),(2775,22,5,28.5,6,0,0,2),(2776,22,5,28.5,6,0,0,3),(2777,22,5,28.5,6,0,0,4),(2778,23,3,20,6,0,0,0),(2779,23,3,30,6,0,0,0),(2780,23,3,40,6,0,0,0),(2781,23,3,50,6,0,0,0),(2782,23,3,60,6,0,0,0),(2783,23,3,70,6,0,0,0),(2784,23,3,80,6,0,0,0),(2785,23,3,90,6,0,0,0),(2786,23,3,100,6,0,0,0),(2787,23,3,110,6,0,0,0),(2788,23,3,125,6,0,0,1),(2789,23,3,125,6,0,0,2),(2790,23,3,125,6,0,0,3),(2791,23,3,125,6,0,0,4),(2792,23,6,20,6,0,0,0),(2793,23,6,25,6,0,0,0),(2794,23,6,30,6,0,0,0),(2795,23,6,35,6,0,0,0),(2796,23,6,40,6,0,0,0),(2797,23,6,45,6,0,0,0),(2798,23,6,50,6,0,0,0),(2799,23,6,55,6,0,0,0),(2800,23,6,65,6,0,0,1),(2801,23,6,65,6,0,0,2),(2802,23,6,65,6,0,0,3),(2803,23,6,65,6,0,0,4),(2804,23,8,20,6,0,0,0),(2805,23,8,25,6,0,0,0),(2806,23,8,30,6,0,0,0),(2807,23,8,35,6,0,0,1),(2808,23,8,35,6,0,0,2),(2809,23,8,35,6,0,0,3),(2810,23,8,35,6,0,0,4),(2811,23,5,5,6,0,0,0),(2812,23,5,17.5,6,0,0,0),(2813,23,5,31,6,0,0,1),(2814,23,5,31,6,0,0,2),(2815,23,5,31,6,0,0,3),(2816,23,5,31,6,0,0,4),(2817,24,3,20,6,0,0,0),(2818,24,3,30,6,0,0,0),(2819,24,3,40,6,0,0,0),(2820,24,3,50,6,0,0,0),(2821,24,3,60,6,0,0,0),(2822,24,3,70,6,0,0,0),(2823,24,3,80,6,0,0,0),(2824,24,3,90,6,0,0,0),(2825,24,3,100,6,0,0,0),(2826,24,3,110,6,0,0,0),(2827,24,3,127.5,6,0,0,1),(2828,24,3,127.5,6,0,0,2),(2829,24,3,127.5,6,0,0,3),(2830,24,3,127.5,6,0,0,4),(2831,24,6,20,6,0,0,0),(2832,24,6,30,6,0,0,0),(2833,24,6,40,6,0,0,0),(2834,24,6,50,6,0,0,0),(2835,24,6,60,6,0,0,0),(2836,24,6,67.5,6,0,0,1),(2837,24,6,67.5,6,0,0,2),(2838,24,6,67.5,6,0,0,3),(2839,24,6,67.5,6,0,0,4),(2840,24,8,20,6,0,0,0),(2841,24,8,25,6,0,0,0),(2842,24,8,30,6,0,0,0),(2843,24,8,37.5,6,0,0,1),(2844,24,8,37.5,6,0,0,2),(2845,24,8,37.5,6,0,0,3),(2846,24,8,37.5,6,0,0,4),(2847,24,5,5,6,0,0,0),(2848,24,5,10,6,0,0,0),(2849,24,5,15,6,0,0,0),(2850,24,5,20,6,0,0,0),(2851,24,5,25,6,0,0,0),(2852,24,5,33.5,6,0,0,1),(2853,24,5,33.5,6,0,0,2),(2854,24,5,33.5,6,0,0,3),(2855,24,5,33.5,6,0,0,4),(2856,25,3,20,6,0,0,0),(2857,25,3,30,6,0,0,0),(2858,25,3,40,6,0,0,0),(2859,25,3,50,6,0,0,0),(2860,25,3,60,6,0,0,0),(2861,25,3,70,6,0,0,0),(2862,25,3,80,6,0,0,0),(2863,25,3,90,6,0,0,0),(2864,25,3,100,6,0,0,0),(2865,25,3,120,6,0,0,1),(2866,25,3,120,6,0,0,2),(2867,25,3,120,6,0,0,3),(2868,25,3,120,6,0,0,4),(2869,25,6,20,6,0,0,0),(2870,25,6,25,6,0,0,0),(2871,25,6,30,6,0,0,0),(2872,25,6,35,6,0,0,0),(2873,25,6,40,6,0,0,0),(2874,25,6,45,6,0,0,0),(2875,25,6,50,6,0,0,0),(2876,25,6,60,6,0,0,1),(2877,25,6,60,6,0,0,2),(2878,25,6,60,6,0,0,3),(2879,25,6,60,6,0,0,4),(2880,25,8,20,6,0,0,0),(2881,25,8,25,6,0,0,0),(2882,25,8,30,6,0,0,1),(2883,25,8,30,6,0,0,2),(2884,25,8,30,6,0,0,3),(2885,25,8,30,6,0,0,4),(2886,25,5,5,6,0,0,0),(2887,25,5,15,6,0,0,0),(2888,25,5,26,6,0,0,1),(2889,25,5,26,6,0,0,2),(2890,25,5,26,6,0,0,3),(2891,25,5,26,6,0,0,4),(2892,26,3,20,6,0,0,0),(2893,26,3,30,6,0,0,0),(2894,26,3,40,6,0,0,0),(2895,26,3,50,6,0,0,0),(2896,26,3,60,6,0,0,0),(2897,26,3,70,6,0,0,0),(2898,26,3,80,6,0,0,0),(2899,26,3,90,6,0,0,0),(2900,26,3,100,6,0,0,0),(2901,26,3,110,6,0,0,0),(2902,26,3,122.5,6,0,0,1),(2903,26,3,122.5,6,0,0,2),(2904,26,3,122.5,6,0,0,3),(2905,26,3,122.5,6,0,0,4),(2906,26,6,20,6,0,0,0),(2907,26,6,25,6,0,0,0),(2908,26,6,30,6,0,0,0),(2909,26,6,35,6,0,0,0),(2910,26,6,40,6,0,0,0),(2911,26,6,45,6,0,0,0),(2912,26,6,50,6,0,0,0),(2913,26,6,55,6,0,0,0),(2914,26,6,62.5,6,0,0,1),(2915,26,6,62.5,6,0,0,2),(2916,26,6,62.5,6,0,0,3),(2917,26,6,62.5,6,0,0,4),(2918,26,8,20,6,0,0,0),(2919,26,8,25,6,0,0,0),(2920,26,8,32.5,6,0,0,1),(2921,26,8,32.5,6,0,0,2),(2922,26,8,32.5,6,0,0,3),(2923,26,8,32.5,6,0,0,4),(2924,26,5,5,6,0,0,0),(2925,26,5,15,6,0,0,0),(2926,26,5,28.5,6,0,0,1),(2927,26,5,28.5,6,0,0,2),(2928,26,5,28.5,6,0,0,3),(2929,26,5,28.5,6,0,0,4),(2930,27,3,20,6,0,0,0),(2931,27,3,30,6,0,0,0),(2932,27,3,40,6,0,0,0),(2933,27,3,50,6,0,0,0),(2934,27,3,60,6,0,0,0),(2935,27,3,70,6,0,0,0),(2936,27,3,80,6,0,0,0),(2937,27,3,90,6,0,0,0),(2938,27,3,100,6,0,0,0),(2939,27,3,110,6,0,0,0),(2940,27,3,125,6,0,0,1),(2941,27,3,125,6,0,0,2),(2942,27,3,125,6,0,0,3),(2943,27,3,125,6,0,0,4),(2944,27,6,20,6,0,0,0),(2945,27,6,25,6,0,0,0),(2946,27,6,30,6,0,0,0),(2947,27,6,35,6,0,0,0),(2948,27,6,40,6,0,0,0),(2949,27,6,45,6,0,0,0),(2950,27,6,50,6,0,0,0),(2951,27,6,55,6,0,0,0),(2952,27,6,65,6,0,0,1),(2953,27,6,65,6,0,0,2),(2954,27,6,65,6,0,0,3),(2955,27,6,65,6,0,0,4),(2956,27,8,20,6,0,0,0),(2957,27,8,25,6,0,0,0),(2958,27,8,30,6,0,0,0),(2959,27,8,35,6,0,0,1),(2960,27,8,35,6,0,0,2),(2961,27,8,35,6,0,0,3),(2962,27,8,35,6,0,0,4),(2963,27,5,5,6,0,0,0),(2964,27,5,17.5,6,0,0,0),(2965,27,5,31,6,0,0,1),(2966,27,5,31,6,0,0,2),(2967,27,5,31,6,0,0,3),(2968,27,5,31,6,0,0,4),(2969,28,3,20,6,0,0,0),(2970,28,3,30,6,0,0,0),(2971,28,3,40,6,0,0,0),(2972,28,3,50,6,0,0,0),(2973,28,3,60,6,0,0,0),(2974,28,3,70,6,0,0,0),(2975,28,3,80,6,0,0,0),(2976,28,3,90,6,0,0,0),(2977,28,3,100,6,0,0,0),(2978,28,3,110,6,0,0,0),(2979,28,3,127.5,6,0,0,1),(2980,28,3,127.5,6,0,0,2),(2981,28,3,127.5,6,0,0,3),(2982,28,3,127.5,6,0,0,4),(2983,28,6,20,6,0,0,0),(2984,28,6,30,6,0,0,0),(2985,28,6,40,6,0,0,0),(2986,28,6,50,6,0,0,0),(2987,28,6,60,6,0,0,0),(2988,28,6,67.5,6,0,0,1),(2989,28,6,67.5,6,0,0,2),(2990,28,6,67.5,6,0,0,3),(2991,28,6,67.5,6,0,0,4),(2992,28,8,20,6,0,0,0),(2993,28,8,25,6,0,0,0),(2994,28,8,30,6,0,0,0),(2995,28,8,37.5,6,0,0,1),(2996,28,8,37.5,6,0,0,2),(2997,28,8,37.5,6,0,0,3),(2998,28,8,37.5,6,0,0,4),(2999,28,5,5,6,0,0,0),(3000,28,5,10,6,0,0,0),(3001,28,5,15,6,0,0,0),(3002,28,5,20,6,0,0,0),(3003,28,5,25,6,0,0,0),(3004,28,5,33.5,6,0,0,1),(3005,28,5,33.5,6,0,0,2),(3006,28,5,33.5,6,0,0,3),(3007,28,5,33.5,6,0,0,4),(3008,29,3,20,6,0,0,0),(3009,29,3,30,6,0,0,0),(3010,29,3,40,6,0,0,0),(3011,29,3,50,6,0,0,0),(3012,29,3,60,6,0,0,0),(3013,29,3,70,6,0,0,0),(3014,29,3,80,6,0,0,0),(3015,29,3,90,6,0,0,0),(3016,29,3,100,6,0,0,0),(3017,29,3,120,6,0,0,1),(3018,29,3,120,6,0,0,2),(3019,29,3,120,6,0,0,3),(3020,29,3,120,6,0,0,4),(3021,29,6,20,6,0,0,0),(3022,29,6,25,6,0,0,0),(3023,29,6,30,6,0,0,0),(3024,29,6,35,6,0,0,0),(3025,29,6,40,6,0,0,0),(3026,29,6,45,6,0,0,0),(3027,29,6,50,6,0,0,0),(3028,29,6,60,6,0,0,1),(3029,29,6,60,6,0,0,2),(3030,29,6,60,6,0,0,3),(3031,29,6,60,6,0,0,4),(3032,29,8,20,6,0,0,0),(3033,29,8,25,6,0,0,0),(3034,29,8,30,6,0,0,1),(3035,29,8,30,6,0,0,2),(3036,29,8,30,6,0,0,3),(3037,29,8,30,6,0,0,4),(3038,29,5,5,6,0,0,0),(3039,29,5,15,6,0,0,0),(3040,29,5,26,6,0,0,1),(3041,29,5,26,6,0,0,2),(3042,29,5,26,6,0,0,3),(3043,29,5,26,6,0,0,4),(3044,30,3,20,6,0,0,0),(3045,30,3,30,6,0,0,0),(3046,30,3,40,6,0,0,0),(3047,30,3,50,6,0,0,0),(3048,30,3,60,6,0,0,0),(3049,30,3,70,6,0,0,0),(3050,30,3,80,6,0,0,0),(3051,30,3,90,6,0,0,0),(3052,30,3,100,6,0,0,0),(3053,30,3,110,6,0,0,0),(3054,30,3,122.5,6,0,0,1),(3055,30,3,122.5,6,0,0,2),(3056,30,3,122.5,6,0,0,3),(3057,30,3,122.5,6,0,0,4),(3058,30,6,20,6,0,0,0),(3059,30,6,25,6,0,0,0),(3060,30,6,30,6,0,0,0),(3061,30,6,35,6,0,0,0),(3062,30,6,40,6,0,0,0),(3063,30,6,45,6,0,0,0),(3064,30,6,50,6,0,0,0),(3065,30,6,55,6,0,0,0),(3066,30,6,62.5,6,0,0,1),(3067,30,6,62.5,6,0,0,2),(3068,30,6,62.5,6,0,0,3),(3069,30,6,62.5,6,0,0,4),(3070,30,8,20,6,0,0,0),(3071,30,8,25,6,0,0,0),(3072,30,8,32.5,6,0,0,1),(3073,30,8,32.5,6,0,0,2),(3074,30,8,32.5,6,0,0,3),(3075,30,8,32.5,6,0,0,4),(3076,30,5,5,6,0,0,0),(3077,30,5,15,6,0,0,0),(3078,30,5,28.5,6,0,0,1),(3079,30,5,28.5,6,0,0,2),(3080,30,5,28.5,6,0,0,3),(3081,30,5,28.5,6,0,0,4),(3082,31,3,20,6,0,0,0),(3083,31,3,30,6,0,0,0),(3084,31,3,40,6,0,0,0),(3085,31,3,50,6,0,0,0),(3086,31,3,60,6,0,0,0),(3087,31,3,70,6,0,0,0),(3088,31,3,80,6,0,0,0),(3089,31,3,90,6,0,0,0),(3090,31,3,100,6,0,0,0),(3091,31,3,110,6,0,0,0),(3092,31,3,125,6,0,0,1),(3093,31,3,125,6,0,0,2),(3094,31,3,125,6,0,0,3),(3095,31,3,125,6,0,0,4),(3096,31,6,20,6,0,0,0),(3097,31,6,25,6,0,0,0),(3098,31,6,30,6,0,0,0),(3099,31,6,35,6,0,0,0),(3100,31,6,40,6,0,0,0),(3101,31,6,45,6,0,0,0),(3102,31,6,50,6,0,0,0),(3103,31,6,55,6,0,0,0),(3104,31,6,65,6,0,0,1),(3105,31,6,65,6,0,0,2),(3106,31,6,65,6,0,0,3),(3107,31,6,65,6,0,0,4),(3108,31,8,20,6,0,0,0),(3109,31,8,25,6,0,0,0),(3110,31,8,30,6,0,0,0),(3111,31,8,35,6,0,0,1),(3112,31,8,35,6,0,0,2),(3113,31,8,35,6,0,0,3),(3114,31,8,35,6,0,0,4),(3115,31,5,5,6,0,0,0),(3116,31,5,17.5,6,0,0,0),(3117,31,5,31,6,0,0,1),(3118,31,5,31,6,0,0,2),(3119,31,5,31,6,0,0,3),(3120,31,5,31,6,0,0,4),(3121,32,3,20,6,0,0,0),(3122,32,3,30,6,0,0,0),(3123,32,3,40,6,0,0,0),(3124,32,3,50,6,0,0,0),(3125,32,3,60,6,0,0,0),(3126,32,3,70,6,0,0,0),(3127,32,3,80,6,0,0,0),(3128,32,3,90,6,0,0,0),(3129,32,3,100,6,0,0,0),(3130,32,3,110,6,0,0,0),(3131,32,3,127.5,6,0,0,1),(3132,32,3,127.5,6,0,0,2),(3133,32,3,127.5,6,0,0,3),(3134,32,3,127.5,6,0,0,4),(3135,32,6,20,6,0,0,0),(3136,32,6,30,6,0,0,0),(3137,32,6,40,6,0,0,0),(3138,32,6,50,6,0,0,0),(3139,32,6,60,6,0,0,0),(3140,32,6,67.5,6,0,0,1),(3141,32,6,67.5,6,0,0,2),(3142,32,6,67.5,6,0,0,3),(3143,32,6,67.5,6,0,0,4),(3144,32,8,20,6,0,0,0),(3145,32,8,25,6,0,0,0),(3146,32,8,30,6,0,0,0),(3147,32,8,37.5,6,0,0,1),(3148,32,8,37.5,6,0,0,2),(3149,32,8,37.5,6,0,0,3),(3150,32,8,37.5,6,0,0,4),(3151,32,5,5,6,0,0,0),(3152,32,5,10,6,0,0,0),(3153,32,5,15,6,0,0,0),(3154,32,5,20,6,0,0,0),(3155,32,5,25,6,0,0,0),(3156,32,5,33.5,6,0,0,1),(3157,32,5,33.5,6,0,0,2),(3158,32,5,33.5,6,0,0,3),(3159,32,5,33.5,6,0,0,4),(3160,41,2,20,6,0,0,0),(3161,41,2,25,6,0,0,0),(3162,41,2,30,6,0,0,0),(3163,41,2,35,6,0,0,0),(3164,41,2,40,6,0,0,1),(3165,41,2,40,6,0,0,2),(3166,41,2,40,6,0,0,3),(3167,41,2,40,6,0,0,4),(3168,42,1,20,6,0,0,0),(3169,42,1,25,6,0,0,0),(3170,42,1,30,6,0,0,0),(3171,42,1,35,6,0,0,0),(3172,42,1,42.5,6,0,0,1),(3173,42,1,42.5,6,0,0,2),(3174,42,1,42.5,6,0,0,3),(3175,42,1,42.5,6,0,0,4),(3176,43,1,20,6,0,0,0),(3177,43,1,25,6,0,0,0),(3178,43,1,30,6,0,0,0),(3179,43,1,35,6,0,0,0),(3180,43,1,40,6,0,0,0),(3181,43,1,45,6,0,0,1),(3182,43,1,45,6,0,0,2),(3183,43,1,45,6,0,0,3),(3184,43,1,45,6,0,0,4),(3185,44,1,20,6,0,0,0),(3186,44,1,25,6,0,0,0),(3187,44,1,30,6,0,0,0),(3188,44,1,35,6,0,0,0),(3189,44,1,40,6,0,0,0),(3190,44,1,47.5,6,0,0,1),(3191,44,1,47.5,6,0,0,2),(3192,44,1,47.5,6,0,0,3),(3193,44,1,47.5,6,0,0,4);
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
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workout`
--

LOCK TABLES `workout` WRITE;
/*!40000 ALTER TABLE `workout` DISABLE KEYS */;
INSERT INTO `workout` VALUES (1,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-22',1),(2,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-11-29',0),(3,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-06',0),(4,'LowerBody Heavy','Leg','Strength Training Focus On Lower Body. Mainly using squart and deadlift','2016-12-13',0),(5,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-11-23',1),(6,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-11-30',0),(7,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-12-07',0),(8,'UpperBody Heavy','Upper Body','Strength Training Focused on Upperbody. Chess press.','2016-12-14',0),(9,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-07',0),(10,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-14',0),(11,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-21',0),(12,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-28',0),(13,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-07',0),(14,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-14',0),(15,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-21',0),(16,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-28',0),(17,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-07',0),(18,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-14',0),(19,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-21',0),(20,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-28',0),(21,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-07',0),(22,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-14',0),(23,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-21',0),(24,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-28',0),(25,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-07',0),(26,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-14',0),(27,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-21',0),(28,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-28',0),(29,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-10',0),(30,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-17',0),(31,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-24',0),(32,'WarmSetTesting','Chest','This is Warm Set Generate Testing','2016-12-31',0),(33,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-12',0),(34,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-19',0),(35,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-26',0),(36,'BugFixing','Bug','Fixing the Bug of Weekly problem','2017-01-02',0),(37,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-12',0),(38,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-19',0),(39,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-26',0),(40,'BugFixing','Bug','Fixing the Bug of Weekly problem','2017-01-02',0),(41,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-12',0),(42,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-19',0),(43,'BugFixing','Bug','Fixing the Bug of Weekly problem','2016-12-26',0),(44,'BugFixing','Bug','Fixing the Bug of Weekly problem','2017-01-02',0);
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

-- Dump completed on 2016-12-08 13:50:32
