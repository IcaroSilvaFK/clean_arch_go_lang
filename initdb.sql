CREATE TABLE IF NOT EXISTS orders (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  price FLOAT NOT NULL,
  tax FLOAT NOT NULL,
  final_price FLOAT NOT NULL,
) 