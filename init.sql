/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `registers` (
  `teacher_id` bigint unsigned NOT NULL,
  `teacher_email` varchar(191) NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  `student_email` varchar(191) NOT NULL,
  PRIMARY KEY (`teacher_id`,`teacher_email`,`student_id`,`student_email`),
  KEY `fk_registers_student` (`student_id`,`student_email`),
  CONSTRAINT `fk_registers_student` FOREIGN KEY (`student_id`, `student_email`) REFERENCES `students` (`id`, `email`),
  CONSTRAINT `fk_registers_teacher` FOREIGN KEY (`teacher_id`, `teacher_email`) REFERENCES `teachers` (`id`, `email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(191) NOT NULL,
  `name` longtext,
  `suspended` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`,`email`),
  KEY `idx_students_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `teachers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(191) NOT NULL,
  `name` longtext,
  PRIMARY KEY (`id`,`email`),
  KEY `idx_teachers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `registers` (`teacher_id`, `teacher_email`, `student_id`, `student_email`) VALUES
(1, 'teacher1@gmail.com', 1, 'student1@gmail.com'),
(2, 'teacher2@gmail.com', 1, 'student1@gmail.com');

INSERT INTO `students` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `name`, `suspended`) VALUES
(1, '2023-03-25 18:25:57.785', '2023-03-27 13:28:48.437', NULL, 'student1@gmail.com', 'Student 1', 0),
(2, '2023-03-25 18:25:57.790', '2023-03-27 14:09:19.149', NULL, 'student2@gmail.com', 'Student 2', 0),
(3, '2023-03-25 18:25:57.798', '2023-03-27 14:09:06.853', NULL, 'student3@gmail.com', 'Student 3', 0);

INSERT INTO `teachers` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `name`) VALUES
(1, '2023-03-25 18:25:57.752', '2023-03-27 12:48:28.292', NULL, 'teacher1@gmail.com', 'Teacher 1'),
(2, '2023-03-25 18:25:57.771', '2023-03-27 14:17:08.261', NULL, 'teacher2@gmail.com', 'Teacher 2'),
(3, '2023-03-25 18:25:57.777', '2023-03-25 18:25:57.777', NULL, 'teacher3@gmail.com', 'Teacher 3');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;