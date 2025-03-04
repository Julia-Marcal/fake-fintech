import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, catchError,throwError } from 'rxjs';
import { map } from 'rxjs/operators';
import { jwtDecode } from 'jwt-decode';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private currentUserSubject = new BehaviorSubject<any>(null);

  constructor(private http: HttpClient) {
    const storedUser = localStorage.getItem('currentUser');
    if (storedUser) {
      this.currentUserSubject.next(JSON.parse(storedUser));
    }
  }

  login(email: string, password: string) {
    return this.http.post<any>('http://localhost:8080/api/login', { email, password }).pipe(
      map((response) => {
        const decodedToken = jwtDecode<{ id: string; username: string; role: string; email: string }>(
          response.access_token
        );
        response.user = decodedToken;
        localStorage.setItem('currentUser', JSON.stringify(response));
        this.currentUserSubject.next(response);
        return response;
      })
    );
  }

  register(body: object) {
    return this.http.post<any>('http://localhost:8080/api/users', body).pipe(
      map((response) => {
        const decodedToken = jwtDecode<{ id: string; username: string; role: string; email: string }>(
          response.access_token
        );
  
        response.user = decodedToken;
        localStorage.setItem('currentUser', JSON.stringify(response));
        this.currentUserSubject.next(response);
  
        return response;
      }),
      catchError((error) => {
        console.error('Registration error', error);
        return throwError(error);
      })
    );
  }

  logout() {
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
  }

  get currentUserValue() {
    return this.currentUserSubject.value;
  }

  getToken(): string | null {
    const currentUser = JSON.parse(localStorage.getItem('currentUser') || '{}');
    return currentUser && currentUser.access_token ? currentUser.access_token : null;
  }
}