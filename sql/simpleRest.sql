/*!40101 SET NAMES utf8mb4 */;
create database `simpleRest` default character set utf8mb4 collate utf8mb4_general_ci;
use simpleRest;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '唯一識別',
  `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父評論Id',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '評論內容',
  `author` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `favorite` int(1) NULL DEFAULT NULL COMMENT '是否是熱門評論',
  `update_time` datetime NOT NULL COMMENT '更新時間',
  PRIMARY KEY (`uuid`) USING BTREE,
  INDEX `parent_id`(`parent_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('3b8c1748-bcb1-4833-9175-8f1e20131b5b', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:54');
INSERT INTO `comment` VALUES ('926332b4-1a77-4252-a99e-7f954086f128', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:55');
INSERT INTO `comment` VALUES ('30fe3e88-dc2c-4ed5-817f-f61d76de7201', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:56');
INSERT INTO `comment` VALUES ('77cc994b-18b2-402c-85d6-83b6cce5062e', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:57');
INSERT INTO `comment` VALUES ('42b9aa3b-46db-4f9e-9813-4a6409f57fd2', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:57');
INSERT INTO `comment` VALUES ('95b995e0-9350-4490-aa8b-e378698f4890', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:57');
INSERT INTO `comment` VALUES ('8191fb0b-88b7-42ba-b872-36cafa2eed76', '', '根據中央氣象局地震測報中心地震報告，這起規模...', '氣象局網站', 1, '2022-07-22 01:31:57');

