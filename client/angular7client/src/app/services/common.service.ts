import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Consumer, CartItem, OrderItem, Order } from '../models';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CommonService {

  uri = 'http://localhost:8083/api';

  constructor(private http: HttpClient) { }
  
  listFranchises(count: number, offset: number) { 
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/franchises?count=${count}&offset=${offset}`);
  }

  getFranchise(id: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/franchises/${id}`);
  }

  listLocations(id: number, count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/franchises/${id}/locations?count=${count}&offset=${offset}`);
  }

  listCategories() {
    return this.http.get(`${this.uri}/categories`);
  }

  listProducts(name: string, count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/products?count=${count}&offset=${offset}&name=${name}`);
  }

  listProductsByCategory(id: number, name: string,count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/categories/${id}/products?count=${count}&offset=${offset}&name=${name}`);
  }
 
  listProductsOfFranchise(id: number,count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/franchises/${id}/products?count=${count}&offset=${offset}`);
  }

  listProductsOfFranchiseAndCategory(id: number, id_category: number,count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/franchises/${id}/categories/${id_category}/products?count=${count}&offset=${offset}`);
  }

  listOrders(id: number, count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/consumers/${id}/orders?count=${count}&offset=${offset}`);
  }

  getOrder(id: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/orders/${id}`);
  }

  getProfile(id : number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/consumers/${id}`);
  }

  getCart(id : number, count: number, offset: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/consumers/${id}/cart?count=${count}&offset=${offset}`);
  }

  getProduct(id: number) {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get(`${this.uri}/products/${id}`);
  }

  updateConsumer(user: Consumer): Observable<Consumer> {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.put<Consumer>(`${this.uri}/consumers/${user.id}`, user);
  }

  createConsumer(user : Consumer): Observable<Consumer>{    
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.post<Consumer>(`${this.uri}/consumers`, user);
  }

  createCartItem(item : CartItem): Observable<CartItem>{   
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.post<CartItem>(`${this.uri}/consumers/${item.consumer_id}/cart`, item);
  }

  updateCartItem(item : CartItem): Observable<CartItem>{    
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.put<CartItem>(`${this.uri}/consumers/${item.consumer_id}/cart/${item.id}`, item);
  }

  getCartItem(item : CartItem): Observable<CartItem> {
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.get<CartItem>(`${this.uri}/consumers/${item.consumer_id}/cart?product_id=${item.product.id}`);
  }
  deleteCartItem(id){
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.delete(`${this.uri}/cart/${id}`);
  }

  createOrderItem(item : OrderItem): Observable<OrderItem>{   
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.post<OrderItem>(`${this.uri}/orders/${item.order_id}`, item);
  }

  createOrder(item : Order): Observable<Order>{   
    let tokenHeader = {headers: {'Token':localStorage.getItem('token')}}
    return this.http.post<Order>(`${this.uri}/orders`, item);
  }


}
