import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { jwtDecode } from 'jwt-decode';
import { ConfigService } from '../config/config.service';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private currentUserSubject: BehaviorSubject<any>;
  public currentUser: Observable<any>;
  private isAuthenticatedSubject: BehaviorSubject<boolean>;
  public isAuthenticated: Observable<boolean>;

  private readonly url: string;
  private readonly token: string;

  constructor(private http: HttpClient, private configService: ConfigService) {
    const storedUser = localStorage.getItem('currentUser');
    this.currentUserSubject = new BehaviorSubject<any>(
      storedUser ? JSON.parse(storedUser) : null
    );
    this.currentUser = this.currentUserSubject.asObservable();

    this.isAuthenticatedSubject = new BehaviorSubject<boolean>(false);
    this.isAuthenticated = this.isAuthenticatedSubject.asObservable();

    this.checkTokenValidity();

    this.url = this.configService.apiBaseUrl;
    this.token = this.configService.apiToken;
  }

  public get currentUserValue(): any {
    return this.currentUserSubject.value;
  }

  public get isAuthenticatedValue(): boolean {
    return this.isAuthenticatedSubject.value;
  }

  login(email: string, password: string): Observable<any> {
    return this.http.post<any>(`${this.url}/auth/login`, { email, password }).pipe(
      map((response) => {
        this.setToken(response);
        this.isAuthenticatedSubject.next(true);
        return response;
      })
    );
  }

  register(body: object): Observable<any> {
    return this.http.post<any>(`${this.url}/auth/register`, { ...body }).pipe(
      map((response) => {
        return response;
      })
    );
  }

  setToken(response: any): void {
    const decodedToken = jwtDecode<{
      sub: string;
      exp: number;
      role: string;
    }>(response.token);

    if (decodedToken.exp * 1000 < Date.now()) {
      this.logout();
      throw new Error('Token has expired');
    }

    localStorage.setItem('currentUser', JSON.stringify(response.token));
    this.currentUserSubject.next(response);
    this.isAuthenticatedSubject.next(true);
  }

  logout(): void {
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
    this.isAuthenticatedSubject.next(false);
  }

  getToken(): string | null {
    const currentUser = JSON.parse(localStorage.getItem('currentUser') || '{}');
    return currentUser || null;
  }

  private checkTokenValidity(): void {
    const token = this.getToken();

    if (!token) {
      this.isAuthenticatedSubject.next(false);
      return;
    }

    try {
      const decodedToken = jwtDecode<{ exp: number }>(token);
      if (decodedToken.exp * 1000 < Date.now()) {
        this.logout();
      } else {
        this.isAuthenticatedSubject.next(true);
      }
    } catch (error) {
      this.logout();
    }
  }

  public getDecodedToken(): any {
    const token = this.getToken();

    if (!token || Object.keys(token).length == 0) {
      this.isAuthenticatedSubject.next(false);
      return null;
    }

    const decodedToken = jwtDecode<{
      id: string;
      username: string;
      exp: number;
      role: string;
      email: string;
    }>(token);

    return decodedToken
  }

}