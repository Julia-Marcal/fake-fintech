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

  private readonly url: string;
  private readonly token: string;

  constructor(private http: HttpClient, private configService: ConfigService) {
    const storedUser = localStorage.getItem('currentUser');
    this.currentUserSubject = new BehaviorSubject<any>(
      storedUser ? JSON.parse(storedUser) : null
    );
    this.currentUser = this.currentUserSubject.asObservable();

    this.url = this.configService.apiBaseUrl;
    this.token = this.configService.apiToken;
  }

  public get currentUserValue(): any {
    return this.currentUserSubject.value;
  }

  login(email: string, password: string): Observable<any> {
    return this.http
      .post<any>(`${this.url}/login`, { email, password })
      .pipe(
        map((response) => {
          const decodedToken = jwtDecode<{
            id: string;
            username: string;
            role: string;
            email: string;
          }>(response.token);

          response.user = decodedToken;

          localStorage.setItem('currentUser', JSON.stringify(response));

          this.currentUserSubject.next(response);

          return response;
        })
      );
  }


  logout(): void {
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
  }

  getToken(): string | null {
    const currentUser = JSON.parse(localStorage.getItem('currentUser') || '{}');
    return currentUser?.access_token || null;
  }
}
