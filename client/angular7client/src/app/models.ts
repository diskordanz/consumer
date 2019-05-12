import { Time } from '@angular/common';
import { Timestamp } from 'rxjs/internal/operators/timestamp';

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
  consumer_id: number;
  franchise_id: number;
  time: Date;
  total: number;
  status: string;
}

export class OrderItem{
  id: number;
  order_id: number;
  product_id: number;
  count: number;
}

export class OrderItemWithProduct{
  id: number;
  order_id: number;
  product: Product;
  count: number;
}

export class OrderWithItems{
  order: Order;
  items: Array<OrderItem>;
}

export class OrderWithItemsAndProducts{
  order: Order;
  items: Array<OrderItemWithProduct>;
}

export class Consumer{
  id: number;
  first_name: string;
  last_name: string;
  phone_number: string;
  city: string;
  address: string;
  login: string; 
  mail: string;
  password: string;
}

export class LoginUser {
  id: number;
  login: string;
  password: string;
  token: string;
}

export class CartItem {
  id: number;
  consumer_id: number;
  product: Product;
  count: number;
}

export class CartItemByFranchise {
  franchise_id: number;
  franchise_name: string;
  cart_items: Array<CartItem>;
  total: number;
}