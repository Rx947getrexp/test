CREATE TABLE IF NOT EXISTS speed_report.t_daily_ad_statistics (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    ad_id INT NOT NULL COMMENT '广告ID',
    ad_name VARCHAR(255) NOT NULL COMMENT '广告名称',
    date INT NOT NULL COMMENT '统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日',
    exposure INT DEFAULT 0 COMMENT '广告的曝光量，默认值为0，表示当天的曝光次数',
    clicks INT DEFAULT 0 COMMENT '广告的点击量，默认值为0，表示当天的点击次数',
    rewards INT DEFAULT 0 COMMENT '广告完播后获赠时长的用户数，默认值为0，表示当天广告完播后获赠时长的用户数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认值为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间，默认值为当前时间，并在每次更新时自动更新',
    UNIQUE KEY unique_ad_date (ad_id, DATE) COMMENT '唯一约束，确保每个广告每天的数据唯一'
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每日广告统计数据表，存储按天统计的广告曝光量和点击量';