create database platform owner postgres;
\connect platform

CREATE TABLE franchises(
  id SERIAL PRIMARY KEY, 
  name TEXT NOT NULL, 
  description TEXT NOT NULL,
  image TEXT,
  country_id INT
);

CREATE TABLE locations(
  id SERIAL PRIMARY KEY, 
  name TEXT NOT NULL, 
  description TEXT NOT NULL, 
  address TEXT NOT NULL, 
  latitude TEXT NOT NULL, 
  longitude TEXT NOT NULL, 
  franchise_id INT, 
  FOREIGN KEY(franchise_id) REFERENCES franchises(id)
);

CREATE TABLE users(
  id SERIAL PRIMARY KEY, 
  username TEXT NOT NULL, 
  password TEXT NOT NULL, 
  role VARCHAR(15) NOT NULL, 
  franchise_id INT NOT NULL, 
  FOREIGN KEY(franchise_id) REFERENCES franchises(id)
);

CREATE TABLE categories(
  id SERIAL PRIMARY KEY, 
  name TEXT NOT NULL
);

CREATE TABLE products(
  id SERIAL PRIMARY KEY, 
  name TEXT NOT NULL, 
  description TEXT NOT NULL,
  image TEXT NOT NULL,
  franchise_id INT NOT NULL,
  category_id INT NOT NULL,
  count INT NOT NULL,
  price FLOAT NOT NULL,
  FOREIGN KEY(franchise_id) REFERENCES franchises(id),
  FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE consumers(
  id SERIAL PRIMARY KEY, 
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  phone_number TEXT,
  city TEXT,
  address TEXT,
  mail TEXT,
  login TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE orders(
  id SERIAL PRIMARY KEY, 
  consumer_id INT NOT NULL,
  franchise_id INT NOT NULL,
  time TIMESTAMP,
  total FLOAT,
  status VARCHAR(15),
  FOREIGN KEY(consumer_id) REFERENCES consumers(id),
  FOREIGN KEY(franchise_id) REFERENCES franchises(id)
);

CREATE TABLE order_items(
  id SERIAL PRIMARY KEY, 
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  count INT, 
  FOREIGN KEY(order_id) REFERENCES orders(id),
  FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE cart_items(
  id SERIAL PRIMARY KEY, 
  consumer_id INT NOT NULL,
  product_id INT NOT NULL,
  count INT, 
  FOREIGN KEY(consumer_id) REFERENCES consumers(id),
  FOREIGN KEY(product_id) REFERENCES products(id)
);

INSERT INTO categories (name)
VALUES 
      ('Все категории'),
      ('Продукты'),
      ('Товары для детей'),
      ('Косметика, бытовая химия'),
      ('Техника, электроника'),
      ('Товары для дома'),
      ('Одежда, обувь, акскссуары'),
      ('Зоотовары'),
      ('Товары для авто'),
      ('Спорт и отдых'),
      ('Товары для дачи'),
      ('Подарки, сувениры');

INSERT INTO consumers (first_name, last_name, phone_number, city, address, login, mail, password)
VALUES
      ('Nadya', 'Barsukova', '9999999', 'Gomel', 'a/g', 'login', 'n@gmail.com','password');

INSERT INTO franchises (id, country_id, name, description, image)
VALUES 
      (1, 1, 'Evroopt', 'Eto evrik', 'logo_evroopt.jpg'),
      (2, 1, 'Gippo', 'Eto kopeechka', 'logo_gippo.jpg'),
      (3, 2, 'Dobryk', 'Eto dobryk', 'logo_dobronom.jpg'),
      (4, 1, '5 element', 'Eto Monetochka', 'logo_5element.png'),
      (5, 2, 'Electrosila', 'Eto Bobryk', 'logo_electrosila.jpg'),
      (6, 1, 'Karusel', 'Eto Kosmetychka', 'logo_karusel.png');


INSERT INTO products (id, name, description, image, franchise_id, category_id, count, price)
VALUES 
      (1, 'Evroopt2', 'Eto evrik', 'r', 1, 2, 0, 1.0),
      (2, 'Kopeechka2', 'Eto kopeechka', 'r', 1, 2, 0, 1.0),
      (3, 'Dobryk3', 'Eto dobryk', 'r', 2, 3, 0, 1.0),
      (4, 'Monetoc3hka', 'Eto Monetochka', 'r', 2, 3, 0, 1.0),
      (5, 'Bobryk3', 'Eto Bobryk', 'r', 3, 4, 0, 1.0),
      (6, 'Kosmety4chka', 'Eto Kosmetychka', 'r', 3, 4, 0, 1.0),
      (7, 'Game4sh', 'Eto Gamesh', 'r', 4, 5, 0, 1.0),
      (8, 'Ugam4e', 'Eto Ugame', 'r', 4, 5, 0, 1.0),
      (9, 'Many4Opt', 'Eto ManyOpt', 'r', 5, 6, 0, 1.0),
      (10, 'A5pelsyn', 'Eto Apelsyn', 'r', 5, 6, 0, 1.0),
      (11, 'Se5crect', 'Eto Secrect', 'r', 6, 7, 0, 1.0),
      (12, 'Evi6lMag', 'Eto EvilMag', 'r', 6, 7, 0, 1.0),
      (13, 'Mand5arin', 'Eto Mandarin', 'r', 6, 8, 0, 1.0);

INSERT INTO cart_items (consumer_id, product_id, count)
VALUES
      (1, 1, 1),
      (1, 2, 2),
      (1, 3, 3),
      (1, 4, 4);

INSERT INTO orders (consumer_id, franchise_id, time, total, status)
VALUES
      (1, 1, '2004-10-19 10:23:54', 10.0, 'TEST');

INSERT INTO order_items (order_id, product_id, count)
VALUES
      (1, 1, 1),
      (1, 2, 2);