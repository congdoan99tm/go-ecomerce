-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_crm_user_c` (
`usr_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
`usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
`usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
`usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'User name',
`usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
`usr_created_at` int NOT NULL DEFAULT '0' COMMENT 'PCreation time',
`usr_update_at` int NOT NULL DEFAULT '0' COMMENT 'Update time',
`usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
`usr_last_login_at` int NOT NULL DEFAULT '0' COMMENT 'last login time',
`usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
`usr_login_times` int NOT NULL DEFAULT '0' COMMENT 'Login times',
`usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable, -1:deleted',
PRIMARY KEY (`usr_id`),
KEY `idx_email` (`usr_email`),
KEY `idx_phone` (`usr_phone`),
KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Account';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table IF EXISTS `pre_go_crm_user_c`
-- +goose StatementEnd
