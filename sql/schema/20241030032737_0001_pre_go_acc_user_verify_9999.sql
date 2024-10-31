-- +goose Up
-- +goose StatementBegin
-- DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
CREATE TABLE `pre_go_acc_user_verify_9999`(
  `verify_id` int NOT NULL AUTO_INCREMENT,
  `verify_otp` varchar(6) not null ,
  `verify_key` varchar(255) not null ,
  `verify_key_hash` varchar(255) not null ,
  `verify_type` int DEFAULT '1',                    -- 1: Email, 2: Phone
  `is_verified` int DEFAULT '0',                    -- 0: no, 1:yes
  `is_deleted` int default  '0',                    -- 0: no, 1:yes
  `verify_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `verify_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`verify_id`),
  UNIQUE KEY `unique_verify_key` (`verify_key`),
  KEY `idx_verify_otp` (`verify_otp`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='account_user_verify'
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
-- +goose StatementEnd
