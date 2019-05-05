export class Location {
  id: number;
  franchise_id: number;
  city: string;
  inhabited_locality: string;
  opening_hours: string;
  adress: string;
}

export class Franchise {
  id: number;
  name: string;
  country_id: number;
  image: string;
}

export class Category {
  id: number;
  name: string;
}

export class Product{
  id: number;
  name: string;
  description: string;
  image: string;
  franchise_id: number;
  //location_id: Int32Array; or count on each location
  count: number;
  price: number;
}

export class Order{
  id: number;
  user_id: number;
  products: Map<number, number>;
  date: Date;
  total: number;
  status: string;
}

export class Consumer{
  id: number;
  first_name: string;
  last_name: string;
  phone_number: string;
  city: string;
  adress: string;
  cart: Map<number, number>;
  login: string;
  mail: string;
  password: string;
}