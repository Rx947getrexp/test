ALTER TABLE t_user_op_log
    ADD COLUMN interfaceUrl varchar(255) NULL COMMENT '接口地址' AFTER content,
    ADD COLUMN serverCode   varchar(64)  NULL COMMENT '后端状态码' AFTER interfaceUrl,
    ADD COLUMN httpCode     varchar(64)  NULL COMMENT 'HTTP状态码' AFTER serverCode,
    ADD COLUMN traceId      varchar(255) NULL COMMENT '请求标识' AFTER httpCode;
