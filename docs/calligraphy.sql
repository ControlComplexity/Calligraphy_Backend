/*
Navicat MySQL Data Transfer

Source Server         : 101.43.39.188
Source Server Version : 80031
Source Host           : 101.43.39.188:3306
Source Database       : calligraphy

Target Server Type    : MYSQL
Target Server Version : 80031
File Encoding         : 65001

Date: 2023-02-12 09:48:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for carousel
-- ----------------------------
DROP TABLE IF EXISTS `carousel`;
CREATE TABLE `carousel` (
  `id` int NOT NULL,
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT '轮播图编号',
  `title` longtext,
  `url` longtext,
  `time` datetime(3) NOT NULL COMMENT '文章内容',
  `type` varchar(128) NOT NULL DEFAULT '' COMMENT '图片类型',
  `link` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '超链接',
  `hits` bigint NOT NULL DEFAULT '0' COMMENT '访问量',
  `like` bigint NOT NULL DEFAULT '0' COMMENT '喜欢次数',
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of carousel
-- ----------------------------
INSERT INTO `carousel` VALUES ('1', '', 'carousel1', 'http://101.43.39.188/image/calligraphy/carousel/carousel1.jpg', '2023-02-10 07:07:49.000', '', 'www.baidu.com', '0', '0', '2023-02-11 19:35:35', '2023-02-11 19:35:35');
INSERT INTO `carousel` VALUES ('2', '', 'carousel2', 'http://101.43.39.188/image/calligraphy/carousel/carousel2.jpg', '2023-02-10 07:07:49.000', '', 'https://baijiahao.baidu.com/s?id=1716100790698670975&wfr=spider&for=pc', '0', '0', '2023-02-11 19:35:46', '2023-02-11 19:35:46');
INSERT INTO `carousel` VALUES ('3', '', 'carousel3', 'http://101.43.39.188/image/calligraphy/carousel/carousel3.jpg', '2023-02-10 07:07:49.000', '', 'https://baijiahao.baidu.com/s?id=1716100790698670975&wfr=spider&for=pc', '0', '0', '2023-02-11 19:35:47', '2023-02-11 19:35:47');
INSERT INTO `carousel` VALUES ('4', '', 'carousel4', 'http://101.43.39.188/image/calligraphy/carousel/carousel4.jpg', '2023-02-10 07:07:49.000', '', 'https://baijiahao.baidu.com/s?id=1716100790698670975&wfr=spider&for=pc', '0', '0', '2023-02-11 19:35:48', '2023-02-11 19:35:48');
INSERT INTO `carousel` VALUES ('5', '', 'carousel5', 'http://101.43.39.188/image/calligraphy/carousel/carousel5.jpg', '2023-02-10 07:07:49.000', '', 'https://baijiahao.baidu.com/s?id=1716100790698670975&wfr=spider&for=pc', '0', '0', '2023-02-11 19:35:48', '2023-02-11 19:35:48');
INSERT INTO `carousel` VALUES ('6', '', 'carousel6', 'http://101.43.39.188/image/calligraphy/carousel/carousel6.jpg', '2023-02-10 07:07:49.000', '', 'https://baijiahao.baidu.com/s?id=1716100790698670975&wfr=spider&for=pc', '0', '0', '2023-02-11 19:35:49', '2023-02-11 19:35:49');

-- ----------------------------
-- Table structure for essay
-- ----------------------------
DROP TABLE IF EXISTS `essay`;
CREATE TABLE `essay` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT '文章编号',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '文章标题',
  `abstract` varchar(255) DEFAULT NULL,
  `content` varchar(10000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文章内容',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT 'URL',
  `time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '文章内容',
  `type` varchar(128) NOT NULL DEFAULT '' COMMENT '文章类型',
  `image` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `home_image` varchar(255) DEFAULT NULL,
  `writerdynasty` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `contentdynasty` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `calligrapher` varchar(255) DEFAULT NULL,
  `hits` int(8) unsigned zerofill DEFAULT NULL COMMENT '访问量',
  `like` int(8) unsigned zerofill DEFAULT NULL COMMENT '喜欢次数',
  `display_home_page` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0' COMMENT '是否在首页展示。1为是，0为否',
  `created_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '删除日期',
  PRIMARY KEY (`id`,`display_home_page`),
  KEY `idx_essay_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of essay
-- ----------------------------
INSERT INTO `essay` VALUES ('1', 'sdfdsf', '董其昌字画欣赏《香光书画真迹》', '明代董其昌山水十二开，高清书画作品图片14张。 香光书画真迹，韩荣光珍藏，子孙永保。 清代狄曼农书法题跋...', '明代董其昌山水十二开，高清书画作品图片14张。\r\n\r\n董其昌字画欣赏《香光书画真迹》\r\n\r\n香光书画真迹，韩荣光珍藏，子孙永保。\r\n\r\n\r\n董其昌字画欣赏《香光书画真迹》\r\n该图片原图过大，需单击该处才可查看到原图。\r\n董其昌字画欣赏《香光书画真迹》', '', '2023-02-12 09:06:17.864', 'shufa', 'http://101.43.39.188/image/calligraphy/carousel/carousel6.jpg', 'http://101.43.39.188/image/calligraphy/carousel/carousel6.jpg', '当代', '明朝', '董其昌', null, null, '0', '2023-02-12 09:06:17.864', '2023-02-12 09:06:17.864', '2023-02-12 09:06:17.864');
INSERT INTO `essay` VALUES ('2', 'dsgvfdsg', '彦辕谈书法：诗书创作与五律诗谱（一）', '书法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。', '法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。\r\n\r\n法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。\r\n法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。\r\n\r\n法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。\r\n\r\n法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。\r\n\r\n书法和诗词都是中华民族的优秀传统文化，都需要我们来继承。尤其是书法爱好者们。在书法学习到了一定程度之后，都有创作作品的欲望。那么，在内容上除选择古诗词外，也都很希望书写自己创作的诗词。然后写成书法作品，供自己，家人，朋友，大众来欣赏。然而，也看到有些朋友写的诗，内容上挺好的，但从格律上，不符合规则，也不符合古体诗的规则，也是一首“顺口溜”。那为什么不把它按照格律诗的规则来写呢？究其原因，是因为有许多书法爱好者，对格律诗没有接触过的缘故。为了使大家对格律诗有些了解，在这里介绍一下格律诗的诗谱和古诗诗例并附李砚园小楷书古诗图例；又附李砚园先生依谱所做的诗和小楷书法图例，供大家参考，或许能帮助上大家，在创作上有所启迪。\r\n\r\n\r\n五言律诗诗谱：\r\n五言律诗每句五个字，一共八句话，共40字。\r\n五言律诗分平起（一二声的字）仄起（三四声的字）两大类。每类又分首句不押韵和首句押韵，所以就形成了四种诗谱谱式，分述于下。\r\n\r\n\r\n（一）.仄起--（首句第二个字为仄声）首句不押韵古诗例（首句最后一个字不押韵）\r\n对仗：在五律中第三四句；五六句中，上下句要对账（平仄.词性等对仗）\r\n\r\n\r\n1.例子-----（杜甫《春望》-古韵）\r\n\r\n国破山河在，城春草木深。感时花溅泪，恨别鸟惊心。烽火连三月，家书抵万金。白头搔更短，浑欲不胜簪。', '', '2023-02-12 09:48:14.670', 'shufa ', 'http://101.43.39.188/image/calligraphy/carousel/carousel5.jpg', 'http://101.43.39.188/image/calligraphy/carousel/carousel5.jpg', null, null, null, null, null, '1', '2023-02-12 09:48:14.670', '2023-02-12 09:48:14.670', '2023-02-12 09:48:14.670');

-- ----------------------------
-- Table structure for theory
-- ----------------------------
DROP TABLE IF EXISTS `theory`;
CREATE TABLE `theory` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT '文章编号',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '文章标题',
  `abstract` varchar(255) DEFAULT NULL,
  `content` varchar(8000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文章内容',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT 'URL',
  `time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '文章内容',
  `type` varchar(128) NOT NULL DEFAULT '' COMMENT '文章类型',
  `image` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `home_image` varchar(255) DEFAULT NULL,
  `hits` int(8) unsigned zerofill DEFAULT NULL COMMENT '访问量',
  `like` int(8) unsigned zerofill DEFAULT NULL COMMENT '喜欢次数',
  `display_home_page` tinyint(1) unsigned zerofill NOT NULL COMMENT '是否在首页展示。1为是，0为否',
  `created_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '删除日期',
  PRIMARY KEY (`id`,`display_home_page`),
  KEY `idx_essay_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of theory
-- ----------------------------
INSERT INTO `theory` VALUES ('1', 'sdfdsfds', '书谱', '夫自古之善书者，汉魏有钟张之绝，晋末称二王之妙。王羲之云：“顷寻诸名书，钟张信为绝伦，其馀不足观。”可谓钟张云没，而羲献继之。又云：“吾书比之钟张，钟当抗行，或谓过之。张草犹当雁行。然张精熟，池水尽墨，假令寡人耽之若此，未必谢之。”此乃推张迈钟之意也。考其专擅，虽未果于前规；摭以兼通，故无惭于即事。', '夫自古之善书者，汉魏有钟张之绝，晋末称二王之妙。王羲之云：“顷寻诸名书，钟张信为绝伦，其馀不足观。”可谓钟张云没，而羲献继之。又云：“吾书比之钟张，钟当抗行，或谓过之。张草犹当雁行。然张精熟，池水尽墨，假令寡人耽之若此，未必谢之。”此乃推张迈钟之意也。考其专擅，虽未果于前规；摭以兼通，故无惭于即事。\r\n孙过庭《书谱》\r\n孙过庭《书谱》\r\n评者云：“彼之四贤，古今特绝；而今不逮古，古质而今妍。”夫质以代兴，妍因俗易。虽书契之作，适以记言；而淳醨一迁，质文三变，驰骛沿革，物理常然。贵能古不乖时，今不同弊，所谓“文质彬彬。然后君子。”何必易雕宫于穴处，反玉辂于椎轮者乎！\r\n又云：“子敬之不及逸少，犹逸少之不及钟张。”意者以为评得其纲纪，而未详其始卒也。且元常专工于隶书，伯英尤精于草体，彼之二美，而逸少兼之。拟草则馀真，比真则长草，虽专工小劣，而博涉多优；总其终始，匪无乖互。\r\n谢安素善尺牍，而轻子敬之书。子敬尝作佳书与之，谓必存录，安辄题后答之，甚以为恨。安尝问敬：“卿书何如右军？”答云：“故当胜。”安云：“物论殊不尔。”子敬又答：“时人那得知！”敬虽权以此辞折安所鉴，自称胜父，不亦过乎！且立身扬名，事资尊显，胜母之里，曾参不入。以子敬之豪翰，绍右军之笔札，虽复粗传楷则，实恐未克箕裘。况乃假托神仙，耻崇家范，以斯成学，孰愈面墙！后羲之往都，临行题壁。子敬密拭除之，辄书易其处，私为不恶。羲之还见，乃叹曰：“吾去时真大醉也！”敬乃内惭。是知逸少之比钟张，则专博斯别；子敬之不及逸少，无或疑焉。\r\n余志学之年，留心翰墨，味钟张之余烈，挹羲献之前规，极虑专精，时逾二纪。有乖入木之术，无间临池之志。观夫悬针垂露之异，奔雷坠石之奇，鸿飞兽骇之资，鸾舞蛇惊之态，绝岸颓峰之势，临危据槁之形。或重若崩云，或轻如蝉翼。导之则泉注，顿之则山安。\r\n纤纤乎似初月之出天涯，落落乎犹众星之列河汉。同自然之妙有，非力运之能成。信可谓“智巧兼优，心手双畅；翰不虚动，下必有由。”\r\n一画之间，变起伏于锋杪；一点之内，殊衄挫于毫芒。况云积其点画，乃成其字。曾不傍窥尺牍，俯习寸阴。引班超以为辞，援项籍而自满。任笔为体，聚墨成形。心昏拟效之方，手迷挥运之理，求其妍妙，不亦谬哉！\r\n然君子立身，务修其本。杨雄谓：“诗赋小道，壮夫不为。”况复溺思毫厘，沦精翰墨者也！\r\n夫潜神对弈，犹标坐隐之名；乐志垂纶，尚体行藏之趣。讵若功宣（一说“定）礼乐，妙拟神仙，犹挻埴之罔穷，与工炉而并运。好异尚奇之士，玩体势之多方；穷微测妙之夫，得推移之奥赜。著述者假其糟粕，藻鉴者挹其菁华，固义理之会归，信贤达之兼善者矣。存精寓赏，岂徒然欤？\r\n而东晋士人，互相陶染。至于王谢之族，郗庾之伦，纵不尽其神奇，咸亦挹其风味。去之滋永，斯道逾微。方复闻疑称疑，得末行末，古今阻绝，无所质问；设有所会，缄秘已深；遂令学者茫然，莫知领要，徒见成功之美，不悟所致之由。或乃就分布于累年，向规矩而犹远，图真不悟，习草将迷。假令薄解草书，粗传隶法，则好溺偏固，自阂通规。讵知心手会归，若同源而异派；转用之术，犹共树而分条者乎？\r\n加以趋变适时，行书为要；题勒方畐（幅），真乃居先。草不兼真，殆于专谨；真不通草，殊非翰札。真以点画为形质，使转为情性；草以点画为情性，使转为形质。\r\n草乖使转，不能成字；真亏点画，犹可记文。回互虽殊，大体相涉。故亦傍通二篆，俯贯八分，包括篇章，涵泳飞白。若毫厘不察，则胡越殊风者焉。\r\n至如钟繇隶奇，张芝草圣，此乃专精一体，以致绝伦。伯英不真，而点画狼藉；元常不草，使转纵横。自兹己降，不能兼善者，有所不逮，非专精也。虽篆隶草章，工用多变，济成厥美，各有攸宜：篆尚婉而通，隶欲精而密，草贵流而畅，章务检而便。然后凛之以风神，温之以妍润，鼓之以枯劲，和之以闲雅。故可达其情性，形其哀乐，验燥湿之殊节，千古依然；体老壮之异时，百龄俄顷。嗟乎，不入其门，讵窥其奥者也！\r\n又一时而书，有乖有合，合则流媚，乖则雕疏，略言其由，各有其五：\r\n神怡务闲，一合也；感惠徇知，二合也；时和气润，三合也；纸墨相发，四合也；偶然欲书，五合也。心遽体留，一乖也；意违势屈，二乖也；风燥日炎，三乖也；纸墨不称，四乖也；情怠手阑，五乖也。乖合之际，优劣互差。得时不如得器，得器不如得志。若五乖同萃，思遏手蒙；五合交臻，神融笔畅。畅无不适，蒙无所从。当仁者得意忘言，罕陈其要；企学者希风叙妙，虽述犹疏。徒立其工，未敷厥旨。不揆庸昧，辄效所明；庶欲弘既往之风规，导将来之器识，除繁去滥，睹迹明心者焉。\r\n代有《笔阵图》七行，中画执笔三手，图貌乖舛，点画湮讹。顷见南北流传，疑是右军所制。虽则未详真伪，尚可发启童蒙。既常俗所存，不藉编录。至于诸家势评，多涉浮华，莫不外状其形，内迷其理，今之所撰，亦无取焉。\r\n若乃师宜官之高名，徒彰史牒；邯郸淳之令范，空著缣缃。暨乎崔、杜以来，萧、羊已往，代祀绵远，名氏滋繁。或藉甚不渝，人亡业显；或凭附增价，身谢道衰。加以糜蠹不传，搜秘将尽，偶逢缄赏，时亦罕窥，优劣纷纭，殆难覼缕。其有显闻当代，遗迹见存，无俟抑扬，自标先后。\r\n且六文之作，肇自轩辕；八体之兴，始于嬴政。其来尚矣，厥用斯弘。但今古不同，妍质悬隔。既非所习，又亦略诸。复有龙蛇云露之流，龟鹤花英之类，乍图真于率尔，或写瑞于当年。巧涉丹青，工亏翰墨，异夫楷式，非所详焉。\r\n代传羲之与子敬笔势论十章，文鄙理疏，意乖言拙，详其旨趣，殊非右军。且右军位重才高，调清词雅，声尘未泯，翰牍仍存。观夫致一书，陈一事，造次之际，稽古斯在；岂有贻谋令嗣，道叶义方，章则顿亏，一至于此！又云与张伯英同学，斯乃更彰虚诞。若指汉末伯英，时代全不相接；必有晋人同号，史传何其寂寥！非训非经，宜从弃择。\r\n夫心之所达，不易尽于名言；言之所通，尚难形于纸墨。粗可仿佛其状，纲纪其辞。冀酌希夷，取会佳境。阙而末逮，请俟将来。今撰执、使、转、用之由，以祛未悟。\r\n执，谓深浅长短之类是也；使，谓纵横牵掣之类是也；转，谓钩环盘纡之类是也；用，谓点画向背之类是也。方复会其数法，归于一途；编列众工，错综群妙。举前人之未及，启后学于成规；窥其根源，析其枝派。贵使文约理赡，迹显心通；披卷可明，下笔无滞。诡辞异说，非所详焉。\r\n然今之所陈，务裨学者。但右军之书，代多称习，良可据为宗匠，取立指归。岂惟会古通今，亦乃情深调合。致使摹拓日广，研习岁滋，先后著名，多从散落；历代孤绍，非其效欤？试言其由，略陈数意：\r\n止如《乐毅论》、《黄庭经》、《东方朔画赞》、《太师箴》、《兰亭集序》、《告誓文》，斯并代俗所传，真行绝致者也。写《乐毅》则情多怫郁；书《画赞》则意涉瑰奇；《黄庭经》则怡怿虚无；《太师箴》又纵横争折；暨乎《兰亭》兴集，思逸神超；私门诫誓，情拘志惨。所谓涉乐方笑，言哀已叹。岂惟驻想流波，将贻啴喛之奏；驰神睢涣，方思藻绘之文。虽其目击道存，尚或心迷议舛。莫不强名为体，共习分区。岂知情动形言，取会风骚之意；阳舒阴惨，本乎天地之心。既失其情，理乖其实，原夫所致，安有体哉！\r\n夫运用之方，虽由己出。规模所设，信属目前。差之一豪，失之千里。苟知其术，适可兼通。心不厌精，手不忘熟。若运用尽于精熟，规矩谙于胸襟，自然容与徘徊，意先笔后，潇洒流落，翰逸神飞，亦犹弘羊之心，预乎无际；庖丁之目，不见全牛。\r\n尝有好事，就吾求习，吾乃粗举纲要，随而授之，无不心悟手从，言忘意得，纵未穷于众术，断可极于所诣矣。\r\n若思通楷则，少不如老；学成规矩，老不如少。思则老而愈妙，学乃少而可勉。勉之不已，抑有三时；时然一变，极其分矣。至如初学分布，但求平正；既知平正，务追险绝，既能险绝，复归平正。初谓未及，中则过之，后乃通会。通会之际，人书俱老。仲尼云：“五十知命”、“七十从心”。故以达夷险之情，体权变之道，亦犹谋而后动，动不失宜；时然后言，言必中理矣。\r\n是以右军之书，末年多妙，当缘思虑通审，志气和平，不激不厉，而风规自远。子敬已下，莫不鼓努为力，标置成体，岂独工用不侔，亦乃神情悬隔者也。或有鄙其所作，或乃矜其所运。自矜者将穷性域，绝于诱进之途；自鄙者尚屈情涯，必有可通之理。嗟乎，盖有学而不能，未有不学而能者也。考之即事，断可明焉。\r\n然消息多方，性情不一，乍刚柔以合体，忽劳逸而分躯。或恬憺雍容，内涵筋骨；或折挫槎枿，外曜锋芒。察之者尚精，拟之者贵似。况拟不能似，察不能精，分布犹疏，形骸未检；跃泉之态，未睹其妍，窥井之谈，已闻其丑。纵欲唐突羲献，诬罔钟张，安能掩当年之目，杜将来之口！慕习之辈，尤宜慎诸。\r\n至有未悟淹留，偏追劲疾；不能迅速，翻效迟重。夫劲速者，超逸之机；迟留者，赏会之致。将反其速，行臻会美之方；专溺于迟，终爽绝伦之妙。能速不速，所谓淹留；因迟就迟，讵名赏会！非其心闲手敏，难以兼通者焉。\r\n假令众妙攸归，务存骨气；骨既存矣，而遒润加之。亦犹枝干扶疏，凌霜雪而弥劲；花叶鲜茂，与云日而相晖。如其骨力偏多，遒丽盖少，则若枯槎架险，巨石当路，虽妍媚云阙，而体质存焉。若遒丽居优，骨气将劣，譬夫芳林落蕊，空照灼而无依；兰沼漂萍，徒青翠而奚托。是知偏工易就，尽善难求。\r\n虽学宗一家，而变成多体，莫不随其性欲，便以为姿。质直者则径侹不遒，刚佷者又倔强无润；矜敛者弊于拘束，脱易者失于规矩；温柔者伤于软缓，躁勇者过于剽迫；狐疑者溺于滞涩，迟重者终于蹇钝；轻琐者淬于俗吏。斯皆独行之士，偏玩所乖。\r\n《易》曰：“观乎天文，以察时变；观乎人文，以化成天下。”况书之为妙，近取诸身。假令运用未周，尚亏工于秘奥；而波澜之际，已浚发于灵台。必能傍通点画之情，博究始终之理，镕铸虫篆，陶均草隶。体五材之并用，仪形不极；象八音之迭起，感会无方。至若数画并施，其形各异；众点齐列，为体互乖。一点成一字之规，一字乃终篇之准。违而不犯，和而不同；留不常迟，遣不恒疾；带燥方润，将浓遂枯；泯规矩于方圆，遁钩绳之曲直；乍显乍晦，若行若藏；穷变态于毫端，合情调于纸上；无间心手，忘怀楷则；自可背羲献而无失，违钟张而尚工。\r\n譬夫绛树青琴，殊姿共艳；随珠和璧，异质同妍。何必刻鹤图龙，竟惭真体；得鱼获兔，犹吝筌蹄。\r\n闻夫家有南威之容，乃可论于淑媛；有龙泉之利，然后议于断割。语过其分，实累枢机。\r\n吾尝尽思作书，谓为甚合，时称识者，辄以引示。其中巧丽，曾不留目；或有误失，翻被嗟赏。既昧所见，尤喻所闻；或以年职自高，轻致陵诮。余乃假之以缃缥，题之以古目，则贤者改观，愚夫继声；竞赏毫末之奇，罕议锋端之失；犹惠侯之好伪，似叶公之惧真。是知伯子之息流波，盖有由矣。\r\n夫蔡邕不谬赏，孙阳不妄顾者，以其玄鉴精通，故不滞于耳目也。向使奇音在爨，庸听惊其妙响；逸足伏枥，凡识知其绝群，则伯喈不足称，伯乐未可尚也。至若老姥遇题扇，初怨而后请；门生获书几，父削而子懊。知与不知也。夫士屈于不知己，而申于知己；彼不知也，曷足怪乎！故庄子曰：“朝菌不知晦朔，蟪蛄不知春秋。”老子云：“下士闻道，大笑之；不笑之则不足以为道也。岂可执冰而咎夏虫哉！”\r\n自汉魏已来，论书者多矣，妍蚩杂糅，条目纠纷：或重述旧章，了不殊于既往；或苟兴新说，竟无益于将来；徒使繁者弥繁，阙者仍阙。今撰为六篇，分成两卷，第其工用，名曰书谱，庶使一家后进，奉以规模；四海知音，或存观省。缄秘之旨，余无取焉。\r\n垂拱三年写记', 'www.baidu.com', '2023-02-11 18:52:40.736', 'scenery', 'http://101.43.39.188/image/calligraphy/carousel/carousel2.jpg', 'https://qcloud.dpfile.com/pc/CvQn0s_joCtJrChcojbdNr5897bL0hOqAqe3031_Y903PHbpIrPl0eqmuwzaVMJCNZvK3RMEnSZusP-oHmeixOpEuctC7NDQ04iV_zcg7BUYPeq-iX-oD-RLc_uA0GHhp6cjSpPcu-http://101.43.39.188/image/calligraphy/carousel/carousel2.jpg', '00000083', '00000032', '1', '2023-02-11 18:52:40.736', '2023-02-11 18:52:40.736', '2023-02-11 18:52:40.736');
INSERT INTO `theory` VALUES ('2', 'sdfdsfds', '书法教育教学的理论与推广研究', '继承与弘扬中华民族的优秀传统文化，为全面实施素质教育提供了深刻的精神养料和文化素材。将书法教育引入中小学教学体系，不仅能够培养学生基本的书写技能，而且能够帮助学生建立更加完整的汉语言认知体系，更能够在艺术的层面上生发对抽象美的体察与感悟。', '中国传统文化是中华民族悠远发展历程中优秀思想的呈示和智慧心识的展现，蕴含着丰富的人文精神，对于当下中国社会的发展和聚心凝力的社会意识的形成起到了十分重要的作用。然而伴随着社会经济的高速发展和科学技术的日新月异，人们对于物质的关注和欲望变得越来越强烈，而对于传统文化精神基因的漠视倾向则逐渐显化，这无疑对传统文化的传承带来了极大的冲击。传统文化与时代的群体意识似乎陷入了愈加疏离的状态，人们对于传统文化的精神内核与外延表征的认识都存在着极大的不足。尤其是在青少年当中，对于潮流文化的追捧使得他们逐渐失去了来自民族意识基因当中的精神气质和文化涵养，甚至对中国传统文化表现出了冷漠和排斥的态度。而这种思想倾向，也越来越受到各界人士的密切关注。由于专家、学者多次在两会上的广泛提议，民族优秀文化的传承与弘扬这一问题再次成为社会话题和民众热议的焦点，传统文化智库的汲取和文化基因的传承得到了进一步的关注。\r\n\r\n书法艺术是我国传统文化艺术的一个重要支脉，既是一门传统艺术，又是一门综合艺术。书法之美是根植于汉民族传统文化的审美认知，它是在由中国哲学和价值观所组成的近乎闭合的逻辑体系中得以生发的审美系统，其表现出物质与精神双重层次上的特殊性，它最集中地体现出了中国人的美学观和艺术观。通过对书法艺术的深入学习和理解能够更好的把握中国传统哲学与文化的深刻内涵与特质。\r\n\r\n书法是关乎汉字的艺术，汉字构形本身为书法的艺术创作和审美宣发提供了物料，汉字形态中具有显化的对于平衡、阴阳、节奏等东方意象思维的考量，它为书法美的表现提供了丰富的物化基础。书法是中国艺术的独绝，是“一之舞蹈，气之流行，道之飞动”。因此，书法又被誉为“无言的诗，无行的舞；无图的画，无声的乐”，书法通过其独特的美，带给了人们艺术的享受。“学书明理”书法的学习能够增强学生的民族认同感，能够让学生在书法学习过程中发现中国优秀传统文化的魅力，进而做到中国优秀传统文化的传承。书法艺术是人的主观精神的一种物质显现，书品的高低取决于书者的精神风貌及其全方位修养，书法的学习能够帮助学生修身立品，能够完善学生的人格。除此之外，书法的学习能够陶冶学生的情操，变化学生的气质，能够增强学生的自信心。在学习书法的过程中，能够锻炼学生左右脑协同能力，可以开发大脑的整体思维。在临摹碑帖的过程中，学生还可以通过碑帖中的文字内容了解更多的历史，这样既能丰富学生的知识储备又能开阔学生的眼界。\r\n\r\n继承与弘扬中华民族的优秀传统文化，为全面实施素质教育提供了深刻的精神养料和文化素材。将书法教育引入中小学教学体系，不仅能够培养学生基本的书写技能，而且能够帮助学生建立更加完整的汉语言认知体系，更能够在艺术的层面上生发对抽象美的体察与感悟。可以说，书法教育是对学生从技能训练到知识认知再到审美情趣的全面育人过程，中华民族传统文化的思想精髓贯穿始终，形成了一个寓情怀于知识的沉浸式体验过程。在这个过程中，学生会自然而然的感受到中国传统文化和审美意识的独特魅力，继而产生出发自内心的文化自信和由衷的爱国情怀。因而，无论是从提高学生自身素质的层面还是从弘扬中华民族优秀传统文化的角度，在中小学全面开展书法教育都有着十分重要的价值与意义。李再湘主编的《书法教育的实施与推广》（湖南师范大学出版社）一书主要介绍了书法教育的具体实施步骤以及推广方法。本书首先向读者介绍了书法的发展历程，在汉字的产生阶段，中国的书法艺术就已经开始显现，虽说当时的文字刻画形态较为简单，但是已经具备了一定的审美情趣。在漫漫历史长河之中，书法字体也在不断的变革。起初是甲骨文和金文，之后慢慢演变为小篆和隶书，到了东汉魏晋时期，又出现了草书、楷书以及行书等书法字体，书法风格种类日渐繁多，散发着其独特的艺术魅力。除了介绍书法的历史之外，该书为了让读者们更加深入的了解书法，在书中向读者们详细介绍了书法的内涵。介绍了书法所使用的工具，也就是我们常说的文房四宝（笔墨纸砚），筆墨纸砚是书法艺术的基础，而作为笔墨纸砚诸如此类中国特有的书写工具的使用，使得书法中的点画线条呈现出旖旎而多变的视觉效果，它为书法美的表现提供了更多的路径与可能性。除此之外，在书中还向读者介绍了书法艺术的本体，介绍了笔法、字法、章法、墨法等内容，让读者对于书法艺术的理解更加的饱满和立体化。与此同时，书法之美的生长可谓具有丰饶的物性土壤，而如何结出瑰丽的思想之花，则在更大程度上取决于中国独特传统文化审美的影响。作者介绍了书法艺术的背景，也就是中国传统文化。中国优秀的传统文化是书法赖以生存以及发展的基础，书法根植在中国优秀传统文化的土壤之上，才能够更好的得到茁壮的发展。\r\n\r\n除了有对书法的发展历程以及书法内涵等的详细介绍之外，此书更多的还是介绍了针对书法教育方面的内容和现实意义。书法教育的开展，能够让学生在学习书法的过程中，将中国优秀传统文化进一步发扬光大，能够让学生更好地感受到中华民族的伟大，惊叹于中国人民的智慧，这会加深学生对于祖国的热爱，增强学生们的文化自信。书法教育还可以提高学生的艺术素养，学生在临摹优秀书法作品时，不仅能够提高自己的书写水平，更能够获得视觉上的审美以及心灵上的享受，提高审美能力以及艺术鉴赏能力。除此之外，书法教育对于学生良好习惯的养成以及认知方法的丰富具有重要的意义，学生们在临摹优秀字帖、进行文字书写过程中，能够养成仔细观察、脚踏实地、勇敢创新等优秀习惯和品质，这对于学生个人的发展有着十分重要的意义。\r\n\r\n另外，此书重点介绍了当前书法教育存在的不足，比如说当前社会对于书法教育的重视度不够。如今信息化以及网络化迅速发展，人们与电脑、手机的接触逐渐增多，也越来越少去拿笔写字，这就造成了人们的书写水平不断下降，更不要说能够掌握传统书法文化。当前中小学书法教育的政策以及制度还不够健全，中小学的书法教育就如同美术、音乐课程一样，经常被主课占用，并没有受到足够的重视。书法教育的教材也不够规范，学生的书法学习缺少合适的书法教育教材。学校对于学生的书法教育的重视度也是不够，学校将更多的精力都放在了应试教育之上，忽视了学生的素质教育。虽说学校每天都在强调素质教育的重要性，但是素质教育一直停留在口头上，并没有实质性的进展以及提升，这在很大程度上阻碍了书法教育的进行。除此之外，当前学校内部的书法教育体制还不够健全，学校还缺少专业化的书法教师，教师专业知识掌握不够熟练，并不能很好的进行学生的书法教育教学。\r\n\r\n针对这些不足以及其他问题，此书提出了一系列促进书法教育工作顺利进行的措施。首先，我们应当做好书法教育的推广，向社会、学校、教师、学生家长以及学生普及书法常识，让大家明白书法的意义所在，明白书法教育的重要性所在。比如说，可以以展示书法优秀作品的方式来推广书法，来向大家宣传书法。也可以请书法专业的名家到学校来进行书法讲座，让学校的教师以及学生认识到书法学习的重要性，认识到书法教育的必要性。同时，教育管理部门应当科学编写书法教育教学教材，严格把控学校书法教育时长。避免一些学校只是把书法教育放在口头上，却并没有付诸行动这一现象的出现。学校应当严格按照教育部门规定的课时标准执行书法教育教学。另外，学校在进行书法教师选拔时，应当选拔高素质且高水平的书法专业教师，要注重书法教师的创作能力、理论分析能力以及教学技巧等方面的考核。注重引导学生理论学习与技法学习相结合，注重培养学生的审美情感，注重提高学生的综合素质，注重学生创作能力的提高。并在之后的教学过程中，时刻监督书法教师的教学情况，提高书法教师的书法能力以及教学水平。除了这些措施之外，学校可以定期举办书法大赛，表现优异的同学可以获得一定的奖励，这能够进一步激发学生们学习书法的热情。\r\n\r\n总之，《书法教育的实施与推广》一书让读者认识到了书法的魅力，了解到了书法教育的意义以及价值，同时为当前的书法教育教学提供了一定的指导，对于书法教学工作的开展，具有十分重要的价值。\r\n\r\n（山东艺术学院 贾长庆）', 'www.baidu.com', '2023-02-11 18:56:40.560', 'scenery', 'http://101.43.39.188/image/calligraphy/carousel/carousel3.jpg', 'https://qcloud.dpfile.com/pc/CvQn0s_joCtJrChcojbdNr5897bL0hOqAqe3031_Y903PHbpIrPl0eqmuwzaVMJCNZvK3RMEnSZusP-oHmeixOpEuctC7NDQ04iV_zcg7BUYPeq-iX-oD-RLc_uA0GHhp6cjSpPcu-http://101.43.39.188/image/calligraphy/carousel/carousel3.jpg', '00000083', '00000032', '1', '2023-02-11 18:56:40.560', '2023-02-11 18:56:40.560', '2023-02-11 18:56:40.560');
INSERT INTO `theory` VALUES ('3', 'sdfdsfds', '欧阳中石先生对书法教育教学体系的构建', '在欧阳中石先生生前，虽也有“少无大志，见异思迁，不务正业，无家可归”的戏谑自评，但他最钟情的工作还是教书。每逢有人真诚地称他“大师”，他说：“我不是‘大师’，我比‘大师’可大多了，我是‘老师’。”先生自中学毕业之后即开始了他七十年的教育生涯。可以说，教育既是他的职业，也是他的事业，更是他的坚守和情怀。尤其晚年以来，他专力从事书法教育，作出了伟大的贡献，取得了丰硕的成果。', '在欧阳中石先生生前，虽也有“少无大志，见异思迁，不务正业，无家可归”的戏谑自评，但他最钟情的工作还是教书。每逢有人真诚地称他“大师”，他说：“我不是‘大师’，我比‘大师’可大多了，我是‘老师’。”先生自中学毕业之后即开始了他七十年的教育生涯。可以说，教育既是他的职业，也是他的事业，更是他的坚守和情怀。尤其晚年以来，他专力从事书法教育，作出了伟大的贡献，取得了丰硕的成果。\r\n\r\n\r\n欧阳先生一生都在从事教育工作，他经常说：“我不是大师，我是老师！”于是就有了“中石自吟”诗句\r\n\r\n笔者自2009年进入首都师范大学博士后流动站，跟随欧阳先生从事汉字和书法的研究，之后又留校任教。在与先生朝夕相处和随时请教的日子，我真实地感受到他对于书法教育体系深入的思考和全方位的构建。\r\n\r\n一 他构建了完整的书法教育人才培养体系\r\n\r\n人类文化的任何成果都必须通过承传才能推动社会的发展和文明的进步，而承传则需要每一个身处其中的人艰深的探索和切实的努力。上个世纪30年代和60年代，沈尹默、陆维钊等先生都曾为高校开设书法专业作出过艰深的努力。这是先生经常提及的前辈们的奋斗历程，至于后来的成绩，他则更多说到上级主管部门的鼎力支持、同道们的共同协作和与学生相处时的各种趣事。\r\n\r\n欧阳先生于1981年调入首都师范大学(当时叫北京师范学院)，1985年在全国率先开办书法成人大专班，以后又发展到本科、硕士。1993年，经国务院学位委员会批准，设立我国第一个以书法为研究方向的博士点。1998年，经国家人事部批准，又设立我国第一个以书法为研究方向的博士后流动站。经过三十余年坚持不懈的努力，欧阳先生首次在我国高校构建了从专科、本科直至硕士、博士、博士后完整的高等书法教育人才培养体系，涵盖了全日制、同等学历、访问学者、第二学历辅修、研究生主要课程进修班、成人继续教育、书法自考等各种教育模式。\r\n\r\n1994年，经首都师范大学批准，创建中国书法艺术研究所(后来更名为中国书法文化研究所)。2005年，扩建为中国书法文化研究院，同时成立中国大学中第一个“书法文化博物馆”。以欧阳先生为旗帜和带头人的首都师大书法学科持续发展壮大，成为教育部第一个书法类艺术师资人才培养培训基地、中国书法学科中唯一国家重点培育学科、第一个省部级重点学科等。\r\n\r\n先生连续担任五届全国政协委员。20多年来，他作为个人或联名提交了有关教育、文化、书法等方面的提案97件，对于推动书法文化在当代的繁荣与发展发挥了重要作用。在从事高等书法教育之余，他时刻关注并亲力亲为地推动中小学书法教育。2013年，欧阳先生已经85岁高龄了。他作为首席专家为教育部主持制定了《中小学书法教育指导纲要》。这些显著成绩，极大促进了当代书法事业的发展，必将成为中国书法史上光辉的一页。\r\n\r\n二 他提出了系统的书法教育理念\r\n\r\n1985年，欧阳先生在首都师范大学创办书法学科时概括出“书面文心”的理念，用以强调书法艺术与文辞内容之间的关系。在长期的教学、研究过程中，他提出了“书法文化”的概念，后来，他将之发展为“作字行文，文以载道，以书焕采，切时如需”，认为“作字”是为了“行文”，“行文”则基于“载道”，至于“书”，主要起到“焕采”的作用，当然，所有这一切都需要“切时如需”。较之“书面文心”的说法，欧阳先生在“文”之上增添了“道”，又强调了“切时如需”。这十六字构建出当代书法事业与我国伟大文艺传统以及时代使命之间的关联，不仅表现出我国传统文人的家国情怀和时代担当，也描绘了历代伟大书法家们共同的品格。\r\n\r\n\r\n在长期的书法教育教学探索中，欧阳先生总结出“作字行文，文以载道，以书焕采，切时如需”的书学理论，成为首都师范大学中国书法文化研究院的院训\r\n\r\n及至晚年，欧阳先生对于书学的研究重点希望可以从汉字认知研究入手，跳出文字学和书法研究的惯性思维，探讨书法与汉字、艺术与文化之间的关系。他补充说，人类生活在天地之间，最主要、最直接的思想交流是通过语言来完成的。但是，语言以声音为媒介，往往受到时、空的限制。文字通过对语言的记录有效弥补了语言的这一缺憾，从而为人类文明的承传和文化的发展作出了巨大的贡献。为此，他专门申报了国家社科基金重大项目“汉字教育与书法表现”，力求将书法文化的研究根植于中国字、中国话和中华优秀传统文化的丰厚土壤之中。\r\n\r\n欧阳先生的这些理念不但切中了书法的内核，更推进了今天的文化自信建设。汉字通过对汉语言的记录，为中华文化的承传作出了不可磨灭的贡献。千百年来，汉字的这一社会功能主要是通过书写来实现的。在这期间，人们又将生活的体验、审美的诉求和文化的提升融入其中，逐渐衍生出中国独有的书法艺术形式。可以看到，从“阐典坟之大猷、成国家之盛业”的历史使命，到“能者加之以玄妙”的“翰墨之道”，书法本身由于具有多个层面、多个侧面的属性，便成为窥视中华文化的窗口，也成为中华文化的集中体现。\r\n\r\n基于以上这些思考，他在主持制定《中小学书法教育指导纲要》时，开篇即谈到，“汉字和以汉字为载体的中国书法是中华民族的文化瑰宝，是人类文明的宝贵财富，”那么，在中小学开展书法教育，“对培养学生的书写能力、审美能力和文化品质具有重要作用。”\r\n\r\n三 他规划了全面的教学内容和教材\r\n\r\n欧阳先生一直主张将书法置于文化的大背景下构建书法文化学科。在长期的教育教学实践中，他提出了以“书法文化”为龙头，“书法史论”为两翼，以汉字、碑帖、文献、鉴藏、材料研究等为外围的“书学”概念，长期致力于“以期成一充实广阔而渊深之学科”。\r\n\r\n早在上个世纪八十年代，欧阳先生参与了中国书画函授大学的创办，承担了主要领导工作。他主编了中国书画函授大学书法教材，包括《书学导论》《书法基础》《中国书法史略》《书论会要》《古文体知识与诗词创作》《行书浅鉴》《草书浅鉴》《楷书浅鉴》《篆书浅鉴》《隶书浅鉴》以及各种字体的书写实践，大致勾勒出高等书法教育中的教学内容体系。\r\n\r\n在首都师范大学创办书法专业后，他主编或撰写《书法与中国文化》《中国书法史鉴》《书法教程》《中国的书法》《学书津梁丛书》《大学本科书法专业系列教材》《书法专业系列教材》《(新编)书法教程》等数十部教材，在有关学科领域都有很重要的影响。其中《书法与中国文化》2002年获北京市哲学社会科学优秀成果二等奖，《(新编)书法教程》获2008年国家级精品教材奖。他在1998年至2008年期间主持国家外文局和美国耶鲁大学合作的国际项目《中国书法艺术》的研究，2008年中英文版在中国和美国同时出版，在国内外学术界引起很大反响。体现欧阳先生主要文化理念的《文化漫谭》，由人民日报出版社出版以来广受好评。\r\n\r\n\r\n欧阳中石先生在听作者汇报科研项目的进展情况\r\n\r\n至于中小学，在他主持的《纲要》中谈到，“书法教育以语文课程中识字和写字教学为基本内容，以提高汉字书写能力为基本目标，以书写实践为基本途径，适度融入书法审美和书法文化教育。”他还主持编写书法教材，举行各种师资培养讲座，身体力行地推动书法教育在全国范围内的中小学校开展。\r\n\r\n四 他设计了科学而又系统的学习方法\r\n\r\n书法教育教学离不开对中华文化传统的把握，当然也离不开在书写实践中对历代经典法书的学习。欧阳先生长期关注这一问题，他曾在早期的大专班上对学生讲：“临帖无‘我’，创作靡‘人’，”意思是说，在通过临帖学习古人时，不要考虑“我”作为个人的存在，但在创作作品时，因为所学已经化为自身的素养，便不要再拘泥于古人的条条框框。这句话强调了在临帖和创作过程中“我”和“人”的关系，辩证而又合理，但至晚年，他又进一步修正说：“其实自己写字时也是没有‘我’的。”为了加以强调，他甚至不肯使用“创作”这个术语。\r\n\r\n\r\n欧阳中石先生在教作者如何临帖\r\n\r\n关于“学”与“练”，欧阳先生认为，“练固然不可少，但是掌握正确的方法更为重要。‘学’则是把人家已经公认为成功的东西拿到自己手中来，这是多么‘合算’的事！为此，首先要能辨认‘好’与‘不好’。在学的过程中要遵守公允的社会历史共识。‘取法乎上’，是‘便宜’的路程。认定目标之后，就要一点不差、惟妙惟肖地把对象学到自己手上来。这样，第一‘练’好了你的眼睛；第二‘练’好了你的手，看得既准，就能写得出，你的书写能力就了不起了。”所以“学”必须注重“实效”。他认为，书法实践不是一件很难的事，想达到一般的水平并不难。\r\n\r\n在此基础上，他提出“打圆心”的说法，也即抓住事物的本质和做事的主要目标。即以临帖而言，“像”或者“似”只是手段，为了把某个字写好，必须仔细观察和体味，还需要考虑如何通过“手”操控“笔”，这样，自然而然就学到了古人的创作技巧和本领。实际上，欧阳先生一生无所不临，博采众家之长，由楷而至行书、篆隶，由唐楷而至魏晋、北碑，铸就了雄浑朴厚、大气开张的艺术风格。\r\n\r\n当代以来，由于各种思潮的影响，人们对书法审美的要求产生了各种认识的差异，因而提出“标准多元化”的概念，又反过来进一步加剧了技术操作的要求。欧阳先生抓住传统书法论著中对“势”的论述，充分发挥他对戏曲艺术研究之所得，有意识地思考书法创作的体势安排和总体布局。每逢创作之前，他总是预想“为书之体，须入其形”，使之“纵横有可象”，然后作书，往往一次成功、一气呵成，鲜见失败之作。他经常说，写字是没有穿上行头的戏剧。为了让学生听得明白，他谈到京剧中的“子午相”，要求演员在亮相的时候，面部、眼神、胸部、双臂、腿脚的摆放应处于合适的位置，趋于合适的方向，而又能互相映衬，从而达到整个身体的动态平衡。直至晚年，他去青岛出差时还亮了一个身段，右足独立，左足高抬，左臂似向后，右臂前扬，果真是修短适宜、收放有度、开合有矩！\r\n\r\n五 余论\r\n\r\n欧阳先生去世了！他依依不舍地离开了长期为之奋斗的讲坛！记得2007年，他获得全国模范教师的荣誉称号时，曾说：“模范教师这是我们这个集体的荣誉，他应当属于我们这个集体！”如今，在全国范围内，已经将近两百所高校开设书法专业，中小学书法教育也渐次有序的展开！我们在怀念先生的同时，将他和前辈们开创的事业继续开拓和发展！\r\n\r\n但是，先生一生并非单纯从事书法教育教学。先生去世后，讣告中称他为“著名学者、教育家、书法家、京剧艺术家”。即以教育而论，他早年主持了中国逻辑与语言函授大学的逻辑教学，撰写了逻辑学的专著和教材，培养了一大批逻辑学教学与研究的人才，同时，作为奚派传人，他为京剧表演事业作出了极大的贡献，撰写了《中国戏曲表现体系研究》等著述，及至晚年，被聘为中国艺术研究院戏曲表演专业的博士生导师。所有这些，我不敢论及。\r\n\r\n单以书法教育而论，我作为最为年轻的学生之一，也不能真正全面而深入地体察他当年奋斗的艰辛、思想的深刻和影响的广远。加之当时的写作，心绪凌乱，啰唆不能成文，恐怕未及先生万一。而蒙《中国教育报》不弃，作全文发表。今先生去世一周年了！在我从事工作和研究的同时，不断整理先生的资料，越加体会到他的魅力、贡献和影响！于是，写下“师恩难忘，教泽长新”的话！我想，我一定会牢记您的教诲，将全身心的学习、工作和研究投入的“老师”的事业中去！\r\n\r\n\r\n孙学峰书：“师恩难忘，教泽常新”\r\n\r\n(作者简介：孙学峰，男，1978年2月生于山东滨州。首都师范大学中国书法文化研究院教授，硕士研究生导师。北京市教育学会书法教育专业委员会常务理事长。师从欧阳中石、秦永龙、国光红诸位先生。为中国书法家协会会员、中国文艺评论家协会会员、中国文字学会会员、中国训诂学研究会会员。\r\n\r\n近年来，以传统小学、美学和书法史论的研究为学术支撑，着眼于书法创作理论本身，力求书法教育教学的科学化、体系化。在《人民日报》《光明日报》《文艺研究》《中国书法》等报纸、期刊发表学术文章40余篇。参与国家社科重大项目“汉字教育与书法表现”，主持北京市社科项目“北京市中小学书法教育的现状调查与对策研究”。出版《汉字形态论》等著述。“中国书法用笔五步训练法”“中国书法‘占格’五步训练法”“中国书法笔毫调控训练法”“中国书法结字四维训练法”“中国书法‘重心’塑造训练法”获得国家版权认证。“‘高参小’对中小学书法教育体系的构建”荣获2017年北京市高等教育教学成果二等奖。)', 'www.baidu.com', '2023-02-11 18:59:38.249', 'scenery', 'http://101.43.39.188/image/calligraphy/carousel/carousel4.jpg', 'https://qcloud.dpfile.com/pc/CvQn0s_joCtJrChcojbdNr5897bL0hOqAqe3031_Y903PHbpIrPl0eqmuwzaVMJCNZvK3RMEnSZusP-oHmeixOpEuctC7NDQ04iV_zcg7BUYPeq-iX-oD-RLc_uA0GHhp6cjSpPcu-http://101.43.39.188/image/calligraphy/carousel/carousel4.jpg', '00000083', '00000032', '1', '2023-02-11 18:59:38.249', '2023-02-11 18:59:38.249', '2023-02-11 18:59:38.249');
