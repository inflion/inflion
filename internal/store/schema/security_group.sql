CREATE TABLE security_group(
  id BIGSERIAL PRIMARY KEY,
  security_group_id VARCHAR(255) NOT NULL,
  security_group_name VARCHAR(255) NOT NULL
);