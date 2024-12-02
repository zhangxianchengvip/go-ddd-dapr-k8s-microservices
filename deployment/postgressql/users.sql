CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT ,                  -- 使用 UUID 作为主键
    loginname VARCHAR(255) UNIQUE NOT NULL,        -- 登录名，唯一且不能为空
    password VARCHAR(255) NOT NULL,                -- 密码，不能为空
    realname VARCHAR(255),                         -- 真实姓名
    email VARCHAR(255),                            -- 邮箱
    phone VARCHAR(255),                            -- 手机号
    role_id UUID                                   -- 角色ID
);