import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Consumer, CartItem } from '../models';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CommonService {

  uri = 'http://localhost:8080/api';

  constructor(private http: HttpClient) { }
  
  listFranchises() {
    return this
           .http
           .get(`${this.uri}/franchises?count=10&offset=0`);
  }

  getFranchise(id: number) {
    return this
           .http
           .get(`${this.uri}/franchises/${id}`);
  }

  listLocations(id: number) {
    return this
           .http
           .get(`${this.uri}/franchises/${id}/locations?count=10&offset=0`);
  }

  listCategories() {
    return this
           .http
           .get(`${this.uri}/categories`);
  }

  listProducts(name: string) {
    return this
           .http
           .get(`${this.uri}/products?count=10&offset=0&name=${name}`);
  }

  listProductsByCategory(id: number, name: string) {
    return this
           .http
           .get(`${this.uri}/categories/${id}/products?count=10&offset=0&name=${name}`);
  }

 
  listProductsOfFranchise(id: number) {
    return this
           .http
           .get(`${this.uri}/franchises/${id}/products?count=10&offset=0`);
  }

  listProductsOfFranchiseAndCategory(id: number, id_category: number) {
    return this
           .http
           .get(`${this.uri}/franchises/${id}/categories/${id_category}/products?count=10&offset=0`);
  }

  listOrders() {
    return this
           .http
           .get(`${this.uri}/orders?count=10&offset=0`);
  }

  getOrder(id: number) {
    return this
           .http
           .get(`${this.uri}/orders/${id}`);
  }

  getProfile(id : number) {
    return this
           .http
           .get(`${this.uri}/consumers/${id}`);
  }

  getCart(id : number) {
    return this
           .http
           .get(`${this.uri}/consumers/${id}/cart?count=10&offset=0`);
  }

  getProduct(id: number) {
    return this
           .http
           .get(`${this.uri}/products/${id}`);
  }

  updateConsumer(user: Consumer): Observable<Consumer> {
    return this.http.put<Consumer>(`${this.uri}/consumers/${user.id}`, Consumer);
  }

  createConsumer(user : Consumer): Observable<Consumer>{    
    return this.http.post<Consumer>(`${this.uri}/consumers`, user);
  }

  createCartItem(item : CartItem): Observable<CartItem>{    
    return this.http.post<CartItem>(`${this.uri}/consumers/${item.id_consumer}/cart`, item);
  }

  updateCartItem(item : CartItem): Observable<CartItem>{    
    return this.http.put<CartItem>(`${this.uri}/consumers/${item.id_consumer}/cart/${item.id}`, item);
  }

  getCartItem(item : CartItem) {
    return this
           .http
           .get(`${this.uri}/consumers/${item.id_consumer}/cart?product_id=${item.product.id}`);
  }
}
