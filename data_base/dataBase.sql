CREATE DATABASE IF NOT EXISTS `championship` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `championship`;
-- MySQL dump 10.13  Distrib 5.6.30, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: championship
-- ------------------------------------------------------
-- Server version	5.6.30-1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE = @@TIME_ZONE */;
/*!40103 SET TIME_ZONE = '+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

--
-- Table structure for table `championship`
--

DROP TABLE IF EXISTS `championship`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `championship` (
  `Id`   VARCHAR(6)  NOT NULL,
  `Name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `championship`
--

LOCK TABLES `championship` WRITE;
/*!40000 ALTER TABLE `championship`
  DISABLE KEYS */;
INSERT INTO `championship` VALUES ('DD1F7E', 'Universidad De Cuenca');
/*!40000 ALTER TABLE `championship`
  ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `getEliminatoria`
--

DROP TABLE IF EXISTS `getEliminatoria`;
/*!50001 DROP VIEW IF EXISTS `getEliminatoria`*/;
SET @saved_cs_client = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `getEliminatoria` AS
  SELECT
    1 AS `Stage_id`,
    1 AS `Team_winner`,
    1 AS `Gol_equipo1`,
    1 AS `Team_losser`,
    1 AS `Gol_equipo2`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `get_eliminatoria`
--

DROP TABLE IF EXISTS `get_eliminatoria`;
/*!50001 DROP VIEW IF EXISTS `get_eliminatoria`*/;
SET @saved_cs_client = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `get_eliminatoria` AS
  SELECT
    1 AS `Stage_id`,
    1 AS `Team_winner`,
    1 AS `Gol_equipo1`,
    1 AS `Team_losser`,
    1 AS `Gol_equipo2`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `get_winner`
--

DROP TABLE IF EXISTS `get_winner`;
/*!50001 DROP VIEW IF EXISTS `get_winner`*/;
SET @saved_cs_client = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `get_winner` AS
  SELECT
    1 AS `Id_team`,
    1 AS `Championship_id`,
    1 AS `Round`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `goal`
--

DROP TABLE IF EXISTS `goal`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goal` (
  `Player_id`    VARCHAR(10) NOT NULL,
  `Match_id`     VARCHAR(6)  NOT NULL,
  `Number_goals` INT(1)      NOT NULL,
  PRIMARY KEY (`Player_id`, `Match_id`),
  KEY `fk_goal_1_idx` (`Player_id`),
  KEY `fk_goal_2_idx` (`Match_id`),
  CONSTRAINT `fk_goal_1` FOREIGN KEY (`Player_id`) REFERENCES `player` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_goal_2` FOREIGN KEY (`Match_id`) REFERENCES `match` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `match`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `match` (
  `Id`          VARCHAR(6) NOT NULL,
  `Team_winner` VARCHAR(6) DEFAULT NULL,
  `Team_losser` VARCHAR(6) DEFAULT NULL,
  `Is_draw`     INT(1)     DEFAULT NULL,
  `Stage_id`    VARCHAR(6) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `fk_match_2_idx` (`Stage_id`),
  CONSTRAINT `fk_match_2` FOREIGN KEY (`Stage_id`) REFERENCES `stage` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `Id`        VARCHAR(10) NOT NULL,
  `Name`      VARCHAR(45) NOT NULL,
  `Lastname`  VARCHAR(45) NOT NULL,
  `Direction` VARCHAR(60) NOT NULL,
  `Number`    INT(2)      NOT NULL,
  `Position`  VARCHAR(20) NOT NULL,
  PRIMARY KEY (`Id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `player_team`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_team` (
  `Player_id`       VARCHAR(10) NOT NULL,
  `Team_id`         VARCHAR(6)  NOT NULL DEFAULT '',
  `Championship_id` VARCHAR(6)  NOT NULL DEFAULT '',
  PRIMARY KEY (`Player_id`, `Team_id`, `Championship_id`),
  KEY `fk_player_idx` (`Player_id`),
  KEY `fk_player_team_1_idx` (`Championship_id`),
  KEY `fk_team` (`Team_id`),
  CONSTRAINT `fk_player` FOREIGN KEY (`Player_id`) REFERENCES `player` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_team_1` FOREIGN KEY (`Championship_id`) REFERENCES `championship` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_team` FOREIGN KEY (`Team_id`) REFERENCES `team` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `posicion`;
/*!50001 DROP VIEW IF EXISTS `posicion`*/;
SET @saved_cs_client = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `posicion` AS
  SELECT
    1 AS `Stage_id`,
    1 AS `Name`,
    1 AS `Id_team`,
    1 AS `Pj`,
    1 AS `Points`,
    1 AS `Pg`,
    1 AS `Pp`,
    1 AS `Pe`,
    1 AS `Gf`,
    1 AS `Gc`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `stage`
--

DROP TABLE IF EXISTS `stage`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stage` (
  `Id`              VARCHAR(6) NOT NULL,
  `Round`           INT(1)     NOT NULL,
  `Is_stage`        INT(1)     NOT NULL,
  `Championship_id` VARCHAR(6) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `fk_stage_1_idx` (`Championship_id`),
  CONSTRAINT `fk_stage_1` FOREIGN KEY (`Championship_id`) REFERENCES `championship` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `team`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `team` (
  `Id`   VARCHAR(6)  NOT NULL,
  `Name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name_UNIQUE` (`Name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `team_championship`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `team_championship` (
  `Team_id`         VARCHAR(6) NOT NULL,
  `Championship_id` VARCHAR(6) NOT NULL,
  PRIMARY KEY (`Team_id`, `Championship_id`),
  KEY `fk_team_championship_2_idx` (`Championship_id`),
  CONSTRAINT `fk_team_championship_1` FOREIGN KEY (`Team_id`) REFERENCES `team` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_team_championship_2` FOREIGN KEY (`Championship_id`) REFERENCES `championship` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


DROP TABLE IF EXISTS `team_stage`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `team_stage` (
  `Team_id`  VARCHAR(6) NOT NULL,
  `Stage_id` VARCHAR(6) NOT NULL,
  PRIMARY KEY (`Team_id`, `Stage_id`),
  KEY `fk_team_stage_3_idx` (`Stage_id`),
  CONSTRAINT `fk_team_stage_2` FOREIGN KEY (`Team_id`) REFERENCES `team` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_team_stage_3` FOREIGN KEY (`Stage_id`) REFERENCES `stage` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'championship'
--
/*!50003 DROP FUNCTION IF EXISTS `getGoalLosser` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `getGoalLosser`(
  `Inteam`     VARCHAR(6),
  `InStage_id` VARCHAR(6)
)
  RETURNS INT(11)
  BEGIN
    DECLARE Vmatch_id VARCHAR(6);
    DECLARE Vnumber_goal INT(11) DEFAULT 0;
    DECLARE Vteam_winner VARCHAR(6);
    DECLARE Vteam_losser VARCHAR(6);
    DECLARE flag_championship VARCHAR(6);
    DECLARE fin INTEGER DEFAULT 0;

    DECLARE point_cursor CURSOR FOR
      SELECT
        Id,
        Team_winner,
        Team_losser
      FROM championship.match
      WHERE championship.match.Stage_id = `InStage_id` AND
            (championship.match.Team_winner = `Inteam` OR championship.match.Team_losser = `Inteam`);
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;

    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vmatch_id, Vteam_winner, Vteam_losser;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSE
        SELECT Championship_id
        INTO flag_championship
        FROM stage
        WHERE stage.Id = `InStage_id`;
        IF `Inteam` = Vteam_winner
        THEN
          SET Vnumber_goal = Vnumber_goal + getGoalMatch(Vteam_losser, flag_championship, Vmatch_id);
        ELSE
          SET Vnumber_goal = Vnumber_goal + getGoalMatch(Vteam_winner, flag_championship, Vmatch_id);
        END IF;
      END IF;
    END LOOP;
    CLOSE point_cursor;

    RETURN Vnumber_goal;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `getGoalMatch` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `getGoalMatch`(
  `Inteam`            VARCHAR(6),
  `Inchampionship_id` VARCHAR(6),
  `Inmatch_id`        VARCHAR(6)
)
  RETURNS INT(11)
  BEGIN
    DECLARE goles INT DEFAULT 0;
    DECLARE vPlayer VARCHAR(10);
    DECLARE vNumber_goals INT(1);
    DECLARE vMatch VARCHAR(6);

    DECLARE flag_player INT DEFAULT 0;

    DECLARE fin INTEGER DEFAULT 0;

    DECLARE point_cursor CURSOR FOR
      SELECT
        Player_id,
        Match_id,
        Number_goals
      FROM `goal`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;

    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO vPlayer, vMatch, vNumber_goals;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF vMatch = `Inmatch_id`
        THEN
          SELECT count(*)
          INTO flag_player
          FROM player_team
          WHERE Player_id = vPlayer AND Championship_id = `Inchampionship_id` AND Team_id = `Inteam`;
          IF flag_player = 1
          THEN
            SET goles = goles + 1;
          END IF;

      END IF;
    END LOOP;
    CLOSE point_cursor;
    RETURN goles;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `getGoalWinner` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `GETGOALWINNER`(
  `Inteam`     VARCHAR(6),
  `InStage_id` VARCHAR(6)
)
  RETURNS INT(11)
  BEGIN
    DECLARE Vmatch_id VARCHAR(6);
    DECLARE Vnumber_goal INT(11) DEFAULT 0;
    DECLARE flag_championship VARCHAR(6);
    DECLARE fin INTEGER DEFAULT 0;

    DECLARE point_cursor CURSOR FOR
      SELECT Id
      FROM championship.match
      WHERE championship.match.Stage_id = `InStage_id`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;

    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vmatch_id;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSE
        SELECT Championship_id
        INTO flag_championship
        FROM stage
        WHERE stage.Id = `InStage_id`;
        SET Vnumber_goal = Vnumber_goal + getGoalMatch(`Inteam`, flag_championship, Vmatch_id);
      END IF;
    END LOOP;
    CLOSE point_cursor;

    RETURN Vnumber_goal;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `getMatch` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `getMatch`(
  `InStage_id` VARCHAR(6)
)
  RETURNS INT(11)
  BEGIN
    DECLARE Vround INT(1) DEFAULT 0;

    SELECT Round
    INTO Vround
    FROM stage
    WHERE stage.Id = `InStage_id`;

    RETURN Vround;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `GETMATCHDRAW` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `GETMATCHDRAW`(
  team VARCHAR(6), id_stage VARCHAR(6))
  RETURNS INT(11)
  BEGIN
    DECLARE numDraw INT;
    DECLARE Vstage_id VARCHAR(6);
    DECLARE Vwinner_id VARCHAR(6);
    DECLARE Vlosser_id VARCHAR(6);
    DECLARE VDraw INT;
    DECLARE fin INTEGER DEFAULT 0;
    DECLARE point_cursor CURSOR FOR
      SELECT
        Team_winner,
        Team_losser,
        Stage_id,
        Is_draw
      FROM `match`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;
    SET numDraw = 0;
    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vwinner_id, Vlosser_id, Vstage_id, VDraw;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF (Vwinner_id = team OR Vlosser_id = team) AND Vstage_id = id_stage
        THEN
          IF VDraw = 1
          THEN
            SET numDraw = numDraw + 1;
          END IF;
      END IF;
    END LOOP;
    CLOSE point_cursor;

    RETURN numDraw;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `GETMATCHLOSSER` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `GETMATCHLOSSER`(
  team VARCHAR(6), id_stage VARCHAR(6))
  RETURNS INT(11)
  BEGIN
    DECLARE numLosser INT;
    DECLARE Vteam_losser VARCHAR(6);
    DECLARE Vstage_id VARCHAR(6);
    DECLARE VDraw INT;
    DECLARE fin INTEGER DEFAULT 0;
    DECLARE point_cursor CURSOR FOR
      SELECT
        Team_losser,
        Stage_id,
        Is_draw
      FROM `match`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;
    SET numLosser = 0;
    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vteam_losser, Vstage_id, VDraw;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF Vteam_losser = team AND Vstage_id = id_stage
        THEN
          IF VDraw = 0
          THEN
            SET numLosser = numLosser + 1;
          END IF;
      END IF;
    END LOOP;
    CLOSE point_cursor;

    RETURN numLosser;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `GETMATCHWINNER` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `GETMATCHWINNER`(
  team VARCHAR(6), id_stage VARCHAR(6))
  RETURNS INT(11)
  BEGIN
    DECLARE numWinner INT;
    DECLARE Vteam_winner VARCHAR(6);
    DECLARE Vstage_id VARCHAR(6);
    DECLARE VDraw INT;
    DECLARE fin INTEGER DEFAULT 0;
    DECLARE point_cursor CURSOR FOR
      SELECT
        Team_winner,
        Stage_id,
        Is_draw
      FROM `match`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;
    SET numWinner = 0;
    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vteam_winner, Vstage_id, VDraw;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF Vteam_winner = team AND Vstage_id = id_stage
        THEN
          IF VDraw = 0
          THEN
            SET numWinner = numWinner + 1;
          END IF;
      END IF;
    END LOOP;
    CLOSE point_cursor;

    RETURN numWinner;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `getNameTeam` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `getNameTeam`(
  `InTeam_id` VARCHAR(6)
)
  RETURNS VARCHAR(45)
  CHARSET latin1
  BEGIN
    DECLARE VNameTeam VARCHAR(45);

    SELECT team.Name
    INTO VnameTeam
    FROM team
    WHERE team.Id = `InTeam_id`;
    RETURN VNameTeam;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP FUNCTION IF EXISTS `GETPUNTOS` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` FUNCTION `GETPUNTOS`(team VARCHAR(6), id_stage VARCHAR(6))
  RETURNS INT(11)
  BEGIN
    DECLARE points INT;
    DECLARE Vteam_winner VARCHAR(6);
    DECLARE Vteam_losser VARCHAR(6);
    DECLARE Vstage_id VARCHAR(6);
    DECLARE VDraw INT;
    DECLARE fin INTEGER DEFAULT 0;
    DECLARE point_cursor CURSOR FOR
      SELECT
        Team_winner,
        Team_losser,
        Stage_id,
        Is_draw
      FROM `match`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;
    SET points = 0;
    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vteam_winner, Vteam_losser, Vstage_id, VDraw;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF (Vteam_winner = team OR Vteam_losser = team) AND Vstage_id = id_stage
        THEN
          IF Vteam_winner = team
          THEN
            IF VDraw = 0
            THEN
              SET points = points + 3;
            ELSE
              SET points = points + 1;
            END IF;
          END IF;

          IF Vteam_losser = team AND Vdraw = 1
          THEN
            SET points = points + 1;
          END IF;
      END IF;
    END LOOP;
    CLOSE point_cursor;


    RETURN points;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP PROCEDURE IF EXISTS `getPoints` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` PROCEDURE `getPoints`(`team` VARCHAR(6), `id_stage` VARCHAR(6))
  BEGIN
    DECLARE points INT;
    DECLARE Vteam_winner VARCHAR(6);
    DECLARE Vstage_id VARCHAR(6);
    DECLARE fin INTEGER DEFAULT 0;
    DECLARE point_cursor CURSOR FOR
      SELECT
        Team_winner,
        Stage_id
      FROM `match`;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin = 1;
    SET points = 0;
    OPEN point_cursor;
    get_points: LOOP
      FETCH point_cursor
      INTO Vteam_winner, Vstage_id;
      IF fin = 1
      THEN
        LEAVE get_points;
      ELSEIF Vteam_winner = team AND Vstage_id = id_stage
        THEN
          SET points = points + 1;
      END IF;
    END LOOP;
    CLOSE point_cursor;

    SELECT points * 3;

  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP PROCEDURE IF EXISTS `getRaundHigher` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` PROCEDURE `getRaundHigher`(
  championship_id VARCHAR(6)
)
  BEGIN

    SELECT max(stage.Round)
    FROM stage
    WHERE stage.Championship_id = championship_id;

  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;
/*!50003 DROP PROCEDURE IF EXISTS `table_position` */;
/*!50003 SET @saved_cs_client = @@character_set_client */;
/*!50003 SET @saved_cs_results = @@character_set_results */;
/*!50003 SET @saved_col_connection = @@collation_connection */;
/*!50003 SET character_set_client = utf8 */;
/*!50003 SET character_set_results = utf8 */;
/*!50003 SET collation_connection = utf8_general_ci */;
/*!50003 SET @saved_sql_mode = @@sql_mode */;
/*!50003 SET sql_mode = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */;
DELIMITER ;;
CREATE DEFINER =`root`@`localhost` PROCEDURE `table_position`(`championship` VARCHAR(6))
  BEGIN
    DECLARE stage_id VARCHAR(6) DEFAULT '';
    DECLARE fin_init INTEGER DEFAULT 0;
    DECLARE stage_cursor CURSOR FOR
      SELECT Id
      FROM `stage`
      WHERE Championship_id = championship;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin_init = 1;

    OPEN stage_cursor;
    get_stage: LOOP
      FETCH stage_cursor
      INTO stage_id;

      IF fin_init = 1
      THEN
        LEAVE get_stage;

      END IF;

        BLOCK2: BEGIN
        DECLARE points INT DEFAULT 0;
        DECLARE team_winner VARCHAR(6);
        DECLARE team_winner_alt VARCHAR(6) DEFAULT '';
        DECLARE fin_match INTEGER DEFAULT 0;
        DECLARE match_cursor CURSOR FOR
          SELECT team_winner
          FROM `match`
          WHERE Stage_id = stage_id;
        DECLARE CONTINUE HANDLER FOR NOT FOUND SET fin_match = 1;

        OPEN match_cursor;
        get_match: LOOP
          FETCH match_cursor
          INTO team_winner;
          IF fin_match = 1
          THEN
            LEAVE get_match;
          ELSE
            IF team_winner != team_winner_alt
            THEN
              SET team_winner_alt = team_winner;
              IF points != 0
              THEN
                SELECT
                  team_winner,
                  points;
              END IF;
            ELSE
              SET points = points + 1;
            END IF;

          END IF;
        END LOOP;

      END;

    END LOOP get_stage;
    CLOSE stage_cursor;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode = @saved_sql_mode */;
/*!50003 SET character_set_client = @saved_cs_client */;
/*!50003 SET character_set_results = @saved_cs_results */;
/*!50003 SET collation_connection = @saved_col_connection */;

--
-- Final view structure for view `getEliminatoria`
--

/*!50001 DROP VIEW IF EXISTS `getEliminatoria`*/;
/*!50001 SET @saved_cs_client = @@character_set_client */;
/*!50001 SET @saved_cs_results = @@character_set_results */;
/*!50001 SET @saved_col_connection = @@collation_connection */;
/*!50001 SET character_set_client = utf8 */;
/*!50001 SET character_set_results = utf8 */;
/*!50001 SET collation_connection = utf8_general_ci */;
/*!50001 CREATE ALGORITHM = UNDEFINED */
  /*!50013 DEFINER =`root`@`localhost`
  SQL SECURITY DEFINER */
  /*!50001 VIEW `getEliminatoria` AS
  SELECT
    `match`.`Stage_id`                                         AS `Stage_id`,
    `getNameTeam`(`match`.`Team_winner`)                       AS `Team_winner`,
    `GETGOALWINNER`(`match`.`Team_winner`, `match`.`Stage_id`) AS `Gol_equipo1`,
    `getNameTeam`(`match`.`Team_losser`)                       AS `Team_losser`,
    `GETGOALWINNER`(`match`.`Team_losser`, `match`.`Stage_id`) AS `Gol_equipo2`
  FROM `match` */;
/*!50001 SET character_set_client = @saved_cs_client */;
/*!50001 SET character_set_results = @saved_cs_results */;
/*!50001 SET collation_connection = @saved_col_connection */;

--
-- Final view structure for view `get_eliminatoria`
--

/*!50001 DROP VIEW IF EXISTS `get_eliminatoria`*/;
/*!50001 SET @saved_cs_client = @@character_set_client */;
/*!50001 SET @saved_cs_results = @@character_set_results */;
/*!50001 SET @saved_col_connection = @@collation_connection */;
/*!50001 SET character_set_client = utf8 */;
/*!50001 SET character_set_results = utf8 */;
/*!50001 SET collation_connection = utf8_general_ci */;
/*!50001 CREATE ALGORITHM = UNDEFINED */
  /*!50013 DEFINER =`root`@`localhost`
  SQL SECURITY DEFINER */
  /*!50001 VIEW `get_eliminatoria` AS
  SELECT
    `match`.`Stage_id`                                         AS `Stage_id`,
    `GETNAMETEAM`(`match`.`Team_winner`)                       AS `Team_winner`,
    `GETGOALWINNER`(`match`.`Team_winner`, `match`.`Stage_id`) AS `Gol_equipo1`,
    `GETNAMETEAM`(`match`.`Team_losser`)                       AS `Team_losser`,
    `GETGOALWINNER`(`match`.`Team_losser`, `match`.`Stage_id`) AS `Gol_equipo2`
  FROM `match` */;
/*!50001 SET character_set_client = @saved_cs_client */;
/*!50001 SET character_set_results = @saved_cs_results */;
/*!50001 SET collation_connection = @saved_col_connection */;

--
-- Final view structure for view `get_winner`
--

/*!50001 DROP VIEW IF EXISTS `get_winner`*/;
/*!50001 SET @saved_cs_client = @@character_set_client */;
/*!50001 SET @saved_cs_results = @@character_set_results */;
/*!50001 SET @saved_col_connection = @@collation_connection */;
/*!50001 SET character_set_client = utf8 */;
/*!50001 SET character_set_results = utf8 */;
/*!50001 SET collation_connection = utf8_general_ci */;
/*!50001 CREATE ALGORITHM = UNDEFINED */
  /*!50013 DEFINER =`root`@`localhost`
  SQL SECURITY DEFINER */
  /*!50001 VIEW `get_winner` AS
  SELECT
    `posicion`.`Id_team`      AS `Id_team`,
    `stage`.`Championship_id` AS `Championship_id`,
    `stage`.`Round`           AS `Round`
  FROM (`posicion`
    JOIN `stage` ON ((`stage`.`Id` = `posicion`.`Stage_id`))) */;
/*!50001 SET character_set_client = @saved_cs_client */;
/*!50001 SET character_set_results = @saved_cs_results */;
/*!50001 SET collation_connection = @saved_col_connection */;

--
-- Final view structure for view `posicion`
--

/*!50001 DROP VIEW IF EXISTS `posicion`*/;
/*!50001 SET @saved_cs_client = @@character_set_client */;
/*!50001 SET @saved_cs_results = @@character_set_results */;
/*!50001 SET @saved_col_connection = @@collation_connection */;
/*!50001 SET character_set_client = utf8 */;
/*!50001 SET character_set_results = utf8 */;
/*!50001 SET collation_connection = utf8_general_ci */;
/*!50001 CREATE ALGORITHM = UNDEFINED */
  /*!50013 DEFINER =`root`@`localhost`
  SQL SECURITY DEFINER */
  /*!50001 VIEW `posicion` AS
  SELECT
    `team_stage`.`Stage_id`                                              AS `Stage_id`,
    `team`.`Name`                                                        AS `Name`,
    `team`.`Id`                                                          AS `Id_team`,
    `championship`.`getMatch`(`team_stage`.`Stage_id`)                   AS `Pj`,
    `GETPUNTOS`(`team`.`Id`, `team_stage`.`Stage_id`)                    AS `Points`,
    `GETMATCHWINNER`(`team`.`Id`, `team_stage`.`Stage_id`)               AS `Pg`,
    `GETMATCHLOSSER`(`team`.`Id`, `team_stage`.`Stage_id`)               AS `Pp`,
    `GETMATCHDRAW`(`team`.`Id`, `team_stage`.`Stage_id`)                 AS `Pe`,
    `championship`.`getGoalWinner`(`team`.`Id`, `team_stage`.`Stage_id`) AS `Gf`,
    `championship`.`getGoalLosser`(`team`.`Id`, `team_stage`.`Stage_id`) AS `Gc`
  FROM ((`team`
    LEFT JOIN `team_stage` ON ((`team`.`Id` = `team_stage`.`Team_id`))) LEFT JOIN `match`
      ON (((`team`.`Id` = `match`.`Team_winner`) AND (`team`.`Id` = `match`.`Team_losser`))))
  ORDER BY `GETPUNTOS`(`team`.`Id`, `team_stage`.`Stage_id`) DESC,
    `championship`.`getGoalWinner`(`team`.`Id`, `team_stage`.`Stage_id`) DESC */;
/*!50001 SET character_set_client = @saved_cs_client */;
/*!50001 SET character_set_results = @saved_cs_results */;
/*!50001 SET collation_connection = @saved_col_connection */;
/*!40103 SET TIME_ZONE = @OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE = @OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES = @OLD_SQL_NOTES */;

-- Dump completed on 2017-01-31  6:34:59
