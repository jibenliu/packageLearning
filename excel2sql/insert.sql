
DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('安徽管益生新材料科技有限公司', '招商加盟', '水管招商', '百度信息流', 12, '管益生-管益生-无', '龚忠顺', 51, '武汉', 18, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '管益生-管益生-无', '百度信息流', '中视电广', 30, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('安徽管益生新材料科技有限公司', '招商加盟', '不锈钢水管', '百度搜索', 8, '管益生-管益生-无', '龚忠顺', 51, '武汉', 10, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '管益生-管益生-无', '百度搜索', '中视电广', 48, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('北京车驰天下汽车科技有限公司', '招商加盟', '汽配', '今日头条', 0, '车美驰-车美驰-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '车美驰-车美驰-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('北京车驰天下汽车科技有限公司', '招商加盟', '汽配', '快手', 15, '车美驰-车美驰-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '车美驰-车美驰-无', '快手', '北京零距离', 35, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('北京车驰天下汽车科技有限公司', '招商加盟', '汽配', '百度信息流', 30, '车美驰-车美驰-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '车美驰-车美驰-无', '百度信息流', '百川互动', 48, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('北京万里达建筑装饰工程有限公司', '装修设计', '装修', '今日头条', 2, '屋美-全屋通-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '屋美-全屋通-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('北京万里达建筑装饰工程有限公司', '招商加盟', '装修', '今日头条', 2, '屋美-全屋通-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '屋美-全屋通-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('采悠科技（武汉）有限公司', '招商加盟', '民宿招商', '今日头条', 0, '采悠-悠家-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '采悠-悠家-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('成都晟世中荣实业有限公司', '招商加盟', '轻钢别墅', '今日头条', 0, '中荣-圣百年-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中荣-圣百年-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('成都晟世中荣实业有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '中荣实业-中荣-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中荣实业-中荣-无', '广点通', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('成都晟世中荣实业有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '中荣实业-中荣-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中荣实业-中荣-无', '广点通', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('成都双流亿瑞鑫贸易有限公司', '招商加盟', '汽配加盟', '今日头条', 0, '双流-亿瑞鑫-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '双流-亿瑞鑫-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('鄂州中嘉装饰材料有限公司', '招商加盟', '墙板建材', '今日头条', 0, '欧橡-简美家-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '欧橡-简美家-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('鄂州中嘉装饰材料有限公司', '招商加盟', '墙板建材', '今日头条', 0, '欧橡-北欧空间-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '欧橡-北欧空间-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('鄂州中嘉装饰材料有限公司', '招商加盟', '墙板建材', '今日头条', 0, '中创-摩登空间-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中创-摩登空间-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('佛山尚美陶瓷有限公司', '招商加盟', '瓷砖加盟', '今日头条', 0, '尚美-德克波罗-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '尚美-德克波罗-无', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东缔造者新型材料有限公司', '招商加盟', '集成墙板', '今日头条', 2, '缔造者-皇家首府-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '缔造者-皇家首府-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '直播', '今日头条', 3, '长洲-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '长洲-无-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '游戏陪玩', '今日头条', 3, '一如-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '一如-无-无', '今日头条', '广州好商汇', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '发热涂料', '今日头条', 3, '几夫庄-居尔-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '几夫庄-居尔-无', '今日头条', '北京零距离', 9, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '发热涂料', '今日头条', 3, '中凯-浩邦特-无', '龚忠顺', 51, '武汉', 0, '是', 0, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中凯-浩邦特-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '设备', '今日头条', 20, '巨蚁-巨蚁-无', '龚忠顺', 51, '武汉', 0, '是', 0, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '巨蚁-巨蚁-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '集成墙板', '今日头条', 3, '民邦-碧豪-无', '龚忠顺', 51, '武汉', 0, '是', 0, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '民邦-碧豪-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '集成墙板', '今日头条', 3, '民邦-宅居易-无', '龚忠顺', 51, '武汉', 0, '是', 0, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '民邦-宅居易-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东络盟网络科技有限公司', '招商加盟', '短视频', '今日头条 ', 3, '络盟-无-无', '龚忠顺', 51, '武汉', 0, '是', 0, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '络盟-无-无', '今日头条 ', '广州好商汇', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东省华晟鼎创建材有限公司', '招商加盟', '集成墙板', '广点通', 0, '华晟-领袖饰界-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '华晟-领袖饰界-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东省华晟鼎创建材有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '华晟鼎创-领袖饰界-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '华晟鼎创-领袖饰界-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广东省华晟鼎创建材有限公司', '招商加盟', '集成墙板', '今日头条', 0, '华晟-领袖饰界-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '华晟-领袖饰界-无', '今日头条', '东信时代', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广州合适家装配式建筑有限公司', '招商加盟', '轻钢别墅', '今日头条', 3, '合适家-合适家-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '合适家-合适家-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('广州赛博驰汽车用品有限公司', '招商加盟', '汽配', '今日头条', 0, '赛博驰-无-无', '龚忠顺', 51, '武汉', 1, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '赛博驰-无-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('贵州御农良品农业科技有限公司', '招商加盟', '羊肚菌', '百度搜索', 15, '御农-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '御农-无-无', '百度搜索', '百川互动', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('合肥荣事达电子电器集团有限公司', '招商加盟', '家电清洗', '今日头条', 3, '荣事达-荣事达-无 ', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '荣事达-荣事达-无 ', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北诚创华环环保装饰材料有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '诚创华环-皇室春天-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '诚创华环-皇室春天-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北弘旭机械设备有限公司', '招商加盟', '粉墙机', '今日头条', 0, '弘旭-弘宇-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '弘旭-弘宇-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北弘旭机械设备有限公司', '招商加盟', '粉墙机', '今日头条', 0, '弘宇-泵-二', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '弘宇-泵-二', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北弘旭机械设备有限公司', '招商加盟', '粉墙机', '广点通', 0, '弘旭-无-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '弘旭-无-无', '广点通', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北梦扬文化传媒有限公司', '招商加盟', '外卖运营', '今日头条', 0, '梦杨-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '梦杨-无-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北喜蛙蛙生态农业科技发展有限公司', '招商加盟', '养殖', '快手', 15, '喜丰-喜丰-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '喜丰-喜丰-无', '快手', '江苏品致', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北鲜乐园农业科技有限公司', '招商加盟', '对虾养殖2', '快手', 23, '鲜乐园-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鲜乐园-无-无', '快手', '江苏品致', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('湖北中工南科智能机械有限公司', '招商加盟', '制砂机', '今日头条', 0, '中工-中工南科-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中工-中工南科-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('华踏实业 华踏(湖北)实业有限公司', '招商加盟', '家电清洗', '今日头条', 0, '华踏-妙博士-无', '龚忠顺', 51, '武汉', 3, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '华踏-妙博士-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('江苏世宜森新型建材有限公司', '招商加盟', '集成墙板', '今日头条', 0, '兴家邦-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '兴家邦-无-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('江苏世宜森新型建材有限公司', '招商加盟', '墙板', '今日头条', 0, '世宜森-兴家邦', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '世宜森-兴家邦', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('来客旅游(武汉)有限公司', '商务服务', '旅游', '百度搜索', 10, '来客-无-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '来客-无-无', '百度搜索', '神雕侠侣', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('南京中博士新材料有限公司', '招商加盟', '轻钢别墅', '今日头条', 0, '中博-中博士-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中博-中博士-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('南京中博士新材料有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '中博-中博士-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中博-中博士-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('南京中博士新材料有限公司', '招商加盟', '墙板招商', '今日头条', 0, '中博士配齐-中博士-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中博士配齐-中博士-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('山东鲁工润屋装配式工程有限公司', '招商加盟', '轻钢别墅', '快手', 15, '鲁工-美伦盛装-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鲁工-美伦盛装-无', '快手', '丹山网络', 30, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('山东鲁工润屋装配式工程有限公司', '招商加盟', '轻钢别墅', '广点通', 0, '鲁工-美伦盛装-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鲁工-美伦盛装-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('山东鲁工润屋装配式工程有限公司', '招商加盟', '轻钢', '百度搜索', 15, '鲁工-美伦盛家-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鲁工-美伦盛家-无', '百度搜索', '百川互动', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('陕西维佳乐装修装饰工程有限公司', '招商加盟', '集成墙板', '今日头条', 3, '维佳乐-维佳乐-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '维佳乐-维佳乐-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('上海诚灸实业有限公司', '招商加盟', '餐饮加盟', '今日头条', 0, '诚灸-火正新-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '诚灸-火正新-无', '今日头条', '广州好商汇', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('上海品奇投资管理有限公司', '招商加盟', '餐饮加盟', '快手', 20, '品奇-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '品奇-无-无', '快手', '丹山网络', 30, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('上海品奇投资管理有限公司', '招商加盟', '汉堡加盟', '广点通', 3, '品奇-麦华仕-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '品奇-麦华仕-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('上海品奇投资管理有限公司', '招商加盟', '餐饮', '今日头条', 3, '品奇-辣么疯狂-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '品奇-辣么疯狂-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('四川万博鑫建材有限公司', '招商加盟', '轻钢别墅', '今日头条', 0, '万博鑫-法维家-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '万博鑫-法维家-无', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('唐山市唯森新材料有限公司', '招商加盟', '集成墙板', '今日头条', 3, '唯森-唯森-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '唯森-唯森-无', '今日头条', '百川互动', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('天娇生物科技（上海）有限公司', '招商加盟', '化妆品', '今日头条', 3, '天娇-百黛天娇-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '天娇-百黛天娇-无', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉保诺驰汽车科技有限公司', '招商加盟', '汽车美容', '今日头条', 0, '保诺驰-保诺驰-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '保诺驰-保诺驰-无', '今日头条', '东信时代', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉倍思凯尔信息技术有限公司', '招商加盟', '洗衣', '广点通', 0, '倍思-优洗-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '倍思-优洗-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉海味缘饮食文化传播有限公司', '招商加盟', '鸭脖加盟', '百度搜索', 15, '海味缘-绝味-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '海味缘-绝味-无', '百度搜索', '北京零距离', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉见歌食品开发有限公司', '招商加盟', '餐饮', '百度搜索', 15, '见歌-吾禾-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '见歌-吾禾-无', '百度搜索', '百川互动', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械设备', '今日头条', 0, '杰仕利-杰仕利-关', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-杰仕利-关', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械设备', '今日头条', 0, '杰仕利-杰仕利-五', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-杰仕利-五', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械', '今日头条', 0, '杰仕利-泵-叶', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-泵-叶', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械', '今日头条', 0, '杰仕利-粉墙机-叶', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-粉墙机-叶', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械', '今日头条', 0, '杰仕利-杰仕利-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-杰仕利-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '二次构造', '今日头条', 0, '杰仕利-泵-关', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-泵-关', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '机械-泵', '今日头条', 0, '杰仕利-泵-关', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-泵-关', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '粉墙机', '今日头条', 0, '杰仕利-杰仕利-刘', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-杰仕利-刘', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '二次构造', '今日头条', 0, '杰仕利-泵-四', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-泵-四', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '粉墙机', '今日头条', 0, '杰仕利-粉墙机-三', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-粉墙机-三', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '石膏机', '今日头条', 0, '杰仕利-杰仕利-关', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-杰仕利-关', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰仕利机械设备有限公司', '招商加盟', '二次构造', '今日头条', 0, '杰仕利-泵-一', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰仕利-泵-一', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉杰斯特重工装备有限责任公司', '招商加盟', '机械设备', '今日头条', 0, '杰斯特-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '杰斯特-无-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉聚佰品机电设备有限公司', '招商加盟', '汽配', '今日头条', 0, '聚佰品-保士捷-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '聚佰品-保士捷-无', '今日头条', '深圳九星', 9, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉聚佰品机电设备有限公司', '招商加盟', '汽配', '广点通', 0, '聚佰品-保士捷-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '聚佰品-保士捷-无', '广点通', '百川互动', 7, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉聚尚客网络科技有限公司', '招商加盟', '直播', '今日头条', 0, '聚尚客-无-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '聚尚客-无-无', '今日头条', '广州好商汇', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉渠渠信息科技有限公司', '招商加盟', '奶茶招商', '广点通', 3, '渠渠-百货-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '渠渠-百货-无', '广点通', '无', 0, 41, '黄颖', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉荣诚方昇文化传播有限公司', '电商', '图书', '快手金牛', 27, '荣诚-跟谁学-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '荣诚-跟谁学-无', '快手金牛', '无', 0, 41, '黄颖', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉荣诚方昇文化传播有限公司', '电商', ' 图书', '快手', 27, ' 荣诚-跟谁学-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, ' 荣诚-跟谁学-无', '快手', '无', 0, 41, '黄颖', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉盛晨机电设备有限公司', '招商加盟', '五金', '百度信息流', 30, '盛晨-妙博士-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '盛晨-妙博士-无', '百度信息流', '北京中视', 50, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉盛晨机电设备有限公司', '招商加盟', '五金机电', '百度', 18, '盛晨-妙博士-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '盛晨-妙博士-无', '百度', '北京中视', 30, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉鑫顺通机械设备有限公司', '招商加盟', '机械设备', '今日头条', 0, '鑫顺-鑫顺通-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鑫顺-鑫顺通-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉鑫顺通机械设备有限公司', '招商加盟', '粉墙机', '今日头条', 0, '鑫顺-鑫顺通-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '鑫顺-鑫顺通-无', '今日头条', '河南优众', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('云南佳泽投资有限公司', '房产', '房地产', '今日头条', 0, '佳泽-凯旋怡景-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '佳泽-凯旋怡景-无', '今日头条', '武汉艾普', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('云南嘉农年华农业科技有限公司', '招商加盟', '种植', '百度搜索', 15, '嘉农-滇圣草-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '嘉农-滇圣草-无', '百度搜索', '晟创意', 28, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('浙江奥威狮集成家居科技有限公司', '招商加盟', '集成墙板', '今日头条', 3, '奥威狮-奥翼-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '奥威狮-奥翼-无', '今日头条', '百川互动', 9, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('浙江东晟康达新材料有限公司', '招商加盟', '轻钢别墅', '今日头条', 0, '康达-匠百年', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '康达-匠百年', '今日头条', '北京零距离', 9, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('中山华建美家新型材料有限公司', '招商加盟', '集成墙板', '今日头条', 3, '华建-宅尚美-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '华建-宅尚美-无', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('中亚构家(湖北)实业发展有限责任公司', '招商加盟', '全铝家居', '今日头条', 0, '中创-宽宅-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中创-宽宅-无', '今日头条', '北京中视', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('中亚构家(湖北)实业发展有限责任公司', '招商加盟', '全铝', '今日头条', 0, '中亚-中亚-无', '龚忠顺', 51, '武汉', 2, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '中亚-中亚-无', '今日头条', '武汉艾普', 0, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('重庆好食客食品机械有限公司', '招商加盟', '豆腐机', '今日头条', 20, '好食客-豆宴坊-无', '龚忠顺', 51, '武汉', 0, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '好食客-豆宴坊-无', '今日头条', '武汉有米来', 5, 41, '洪亮', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;

DELIMITER $$
DROP PROCEDURE IF EXISTS insertData $$
CREATE PROCEDURE insertData()
BEGIN
    DECLARE customerId INT(10) DEFAULT 0;
	START TRANSACTION;
	INSERT INTO customers  (`customer_name`, `industry`, `second_industry`, `platform`, `account_strategy`, `customer_nickname`, `seller`, `seller_id`, `mechanism`, `make_costs`, `operate_type`, `remark_rate`, `created_at`, `updated_at`) VALUES  ('武汉网柚未来科技有限公司', '电商', '培训', '百度', 10, '网柚-拼多多-无', '龚忠顺', 51, '武汉', 5, '是', 1, '2021-01-21 19:12:25', '2021-01-21 19:12:25');
	SET customerId = (SELECT LAST_INSERT_ID());
	SET @insertSql = CONCAT(REPLACE("INSERT INTO customer_infos  (`customer_id`, `customer_nickname`, `platform`, `upper_name`, `platform_strategy`, `medium_id`, `medium`, `seller`, `seller_id`, `created_at`, `updated_at`) VALUES  (customerId, '网柚-拼多多-无', '百度', '无', 0, 41, '黄颖', '龚忠顺', 51, '2021-01-21 19:12:25', '2021-01-21 19:12:25')",'customerId', customerId));
	PREPARE itemSql FROM @insertSql;
	EXECUTE itemSql;
	COMMIT;
END $$
DELIMITER ;

CALL insertData();
DROP PROCEDURE insertData;
