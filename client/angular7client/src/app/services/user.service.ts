import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Consumer } from '../models';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ConsumerService {
  private URL = 'http://localhost:3000';
  constructor(private http: HttpClient) {
  }

  getConsumers(): Observable<Consumer[]> {
    return this.http.get<Consumer[]>(`${this.URL}/consumers`);
  }

  getConsumerById(id: number): Observable<Consumer> {
    return this.http.get<Consumer>(`${this.URL}/consumers/${id}`);
  }

  getAuthorizedConsumer(): Observable<Consumer> {
    return this.http.get<Consumer>(`${this.URL}/consumersprofile`);
  }

  addConsumer(Consumer: Consumer): Observable<Consumer> {
    return this.http.post<Consumer>(`${this.URL}/consumers/add`, Consumer);
  }

  deleteConsumer(Consumer: Consumer): Observable<Consumer> {
    return this.http.delete<Consumer>(`${this.URL}/consumers/${Consumer.id}`);
  }

  updateConsumer(Consumer: Consumer): Observable<Consumer> {
    return this.http.put<Consumer>(`${this.URL}/consumers/${Consumer.id}`, Consumer);
  }

  searchConsumers(term: string): Observable<Consumer[]> {
    if (!term.trim()) {
      // if not search term, return all Consumers
      return this.http.get<Consumer[]>(`${this.URL}/consumers`);
    }
    return this.http.get<Consumer[]>(`${this.URL}/search/${term}`);
  }

}