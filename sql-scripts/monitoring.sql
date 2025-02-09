-- Docker контейнеры
CREATE TABLE containers (
    -- IPv4 адрес контейнера
    ipv4 VARCHAR(15) PRIMARY KEY,
    -- дата последней успешной попытки
    success_ping_time VARCHAR(50) NOT NULL
);

-- Результаты запусков ping
CREATE TABLE ping_results (
    id SERIAL PRIMARY KEY,
    -- время пинга
    ping_time VARCHAR(15) NOT NULL,
    container_ipv4 VARCHAR(15) NOT NULL,
    FOREIGN KEY (container_ipv4) REFERENCES containers(ipv4)
);

