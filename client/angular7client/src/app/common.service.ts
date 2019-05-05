import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

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

  listProducts() {
    return this
           .http
           .get(`${this.uri}/products?count=10&offset=0`);
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

  listProductsByCategory(id: number) {
    return this
           .http
           .get(`${this.uri}/categories/${id}/products?count=10&offset=0`);
  }

  listProductsByCategoryAndName(id: number, name: string) {
    return this
           .http
           .get(`${this.uri}/categories/${id}/products?name=${name}`);
  }

  listProductsByName(name: string) {
    return this
           .http
           .get(`${this.uri}/products?name=${name}`);
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

  getProfile() {
    return this
           .http
           .get(`${this.uri}/profile`);
  }

  getProduct(id: number) {
    return this
           .http
           .get(`${this.uri}/products/${id}`);
  }
}
