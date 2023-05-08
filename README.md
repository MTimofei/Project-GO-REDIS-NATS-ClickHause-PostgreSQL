<h1 align='center'>API Service</h1>

![go](https://img.shields.io/badge/Go-F5F5F5?style=for-the-badge&logo=Go&logoColor=#61DAFB)
![postgres](https://img.shields.io/badge/postgres-F5F5F5?style=for-the-badge&logo=postgres&logoColor=#61DAFB)
![SQL](https://img.shields.io/badge/sql-F5F5F5?style=for-the-badge&logo=sql&logoColor=#61DAFB)
![docker](https://img.shields.io/badge/docker-F5F5F5?style=for-the-badge&logo=docker&logoColor=#61DAFB)
![clickhouse](https://img.shields.io/badge/clickhouse-F5F5F5?style=for-the-badge&logo=clickhouse&logoColor=#61DAFB)
![redis](https://img.shields.io/badge/redis-F5F5F5?style=for-the-badge&logo=redis&logoColor=#61DAFB)
![nats](https://img.shields.io/badge/nats-F5F5F5?style=for-the-badge&logo=nats&logoColor=#61DAFB)


<h3>Настройка : ClickHouse </h3>

- Добавление кода для настройки для табличного движка;
  

```html
<nats>
    <user>click</user>
    <password>house</password>
    <token>clickhouse</token>
</nats>
```

- Создание таблиц для получения,отоброжения данных NATS;

```sql
CREATE TABLE LogStor (
    id UInt64,
    log String
  ) ENGINE = NATS 
   SETTINGS nats_url = 'service_nats:4222',
             nats_subjects = 'log',
             nats_format = 'JSONEachRow',
             date_time_input_format = 'best_effort';

  CREATE TABLE daily (id UInt64, log String)
    ENGINE = MergeTree() ORDER BY id;

  CREATE MATERIALIZED VIEW consumer TO daily
    AS SELECT id, log FROM LogStor;
```

- Создание таблиц для миграции данных из PostgreSQL;

```sql
CREATE TABLE migrations (
  Id Int32,
  Campaignid Int32,
  Name String,
  Description String,
  Priority Int32,
  Removed UInt8,
  EventTime Datetime
)ENGINE = MergeTree()
ORDER BY (Id);
```

<h3>Настройка : PostgreSQL </h3>

- Создание таблици для отслеживания миграции;

```sql
postgres
CREATE TABLE CAMPAIGNS (
  id SERIAL PRIMARY KEY,
  name CHAR(12) NOT NULL
);
```

- Создание таблици для хронения данных;

```sql
CREATE TABLE ITEMS (
  id SERIAL PRIMARY KEY,
  campaign_id INT NOT NULL REFERENCES CAMPAIGNS(id),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  priority INT NOT NULL  ,
  removed BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```
-Создание создание тригера и функции для автоматического определиния приориета;

```sql
CREATE OR REPLACE FUNCTION set_priority_on_insert()
RETURNS TRIGGER AS $$
BEGIN
  NEW.priority := COALESCE((SELECT MAX(priority) FROM items), 0) + 1;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_priority_trigger
BEFORE INSERT ON items
FOR EACH ROW
EXECUTE FUNCTION set_priority_on_insert();
```

<h3>Настройка : Docker </h3>
- Создание и настройка контейнеров описана в docker-compose.yaml;

```
NATS,REDIS используют настройки по умолчанию
```
