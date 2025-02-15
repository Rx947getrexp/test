#每日按渠道统计充值金额总数
CREATE TABLE IF NOT EXISTS speed_report.t_daily_payment_total_by_channel (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    date INT NOT NULL COMMENT '统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日',
    channel VARCHAR(255) NOT NULL COMMENT '支付渠道名称',
    amount DECIMAL(10, 2) NOT NULL COMMENT '支付金额统计',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认值为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间，默认值为当前时间，并在每次更新时自动更新',
    UNIQUE KEY unique_date_channel (date, channel) COMMENT '确保每个日期和支付渠道的组合是唯一的',
    INDEX idx_date (date) COMMENT '加速按日期查询',
    INDEX idx_channel (channel) COMMENT '加速按支付渠道查询'
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每日按渠道统计充值金额总数';

#每日按渠道统计日活月活充值金额
CREATE TABLE IF NOT EXISTS speed_report.t_channel_registration_pay_daily (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    date INT NOT NULL COMMENT '统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日',
    channel VARCHAR(128) NOT NULL COMMENT '渠道id',
    new_users INT NOT NULL COMMENT '新增用户数量',
    daily_active_users INT NOT NULL COMMENT '日活用户数量',
    monthly_active_users INT NOT NULL COMMENT '月活用户数量',
    total_recharge_users INT NOT NULL COMMENT '充值用户数量',
    total_recharge_amount VARCHAR(128) NOT NULL COMMENT '付费金额数量',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认值为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间，默认值为当前时间，并在每次更新时自动更新',
    UNIQUE KEY unique_date_channel (date, channel) COMMENT '确保每个日期和渠道的组合是唯一的',
    INDEX idx_date (date) COMMENT '加速按日期查询',
    INDEX idx_channel (channel) COMMENT '加速按渠道查询'
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每日按渠道统计日活月活充值金额';

#每日按渠道统计留存
CREATE TABLE IF NOT EXISTS speed_report.t_channel_retaind_daily (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    date INT NOT NULL COMMENT '统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日',
    channel VARCHAR(128) NOT NULL COMMENT '渠道id',
    new_users INT NOT NULL COMMENT '新增用户数量',
    day_2_retained INT NOT NULL COMMENT '新增用户次日留存数量',
    day_7_retained INT NOT NULL COMMENT '新增用户7日留存数量',
    day_15_retained INT NOT NULL COMMENT '新增用户15日留存数量',
    day_30_retained INT NOT NULL COMMENT '新增用户30日留存数量',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认值为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间，默认值为当前时间，并在每次更新时自动更新',
    UNIQUE KEY unique_date_channel (date, channel) COMMENT '确保每个日期和渠道的组合是唯一的',
    INDEX idx_date (date) COMMENT '加速按日期查询',
    INDEX idx_channel (channel) COMMENT '加速按渠道查询'
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每日按渠道统计留存';

#每月按渠道统计充值续费人数金额
CREATE TABLE IF NOT EXISTS speed_report.t_channel_recharge_renewals_monthly (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    month INT NOT NULL COMMENT '统计数据月份，整数类型，格式为 YYYYMM，例如202501',
    channel VARCHAR(128) NOT NULL COMMENT '渠道id',
    recharge_users INT NOT NULL COMMENT '付费用户数量',
    recharge_amount VARCHAR(128) NOT NULL COMMENT '付费用户充值总金额',
    retained INT NOT NULL COMMENT '充值用户次月留存数量',
    renewals_users INT NOT NULL COMMENT '次月续费人数',
    renewals_amount VARCHAR(128) NOT NULL COMMENT '次月续费充值总金额',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间，默认值为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间，默认值为当前时间，并在每次更新时自动更新',
    UNIQUE KEY unique_month_channel (month, channel) COMMENT '确保每个日期和渠道的组合是唯一的',
    INDEX idx_month (month) COMMENT '加速按日期查询',
    INDEX idx_channel (channel) COMMENT '加速按渠道查询'
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每月按渠道统计充值续费人数金额';
