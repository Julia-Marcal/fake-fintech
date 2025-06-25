import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ConfigService } from '../../../shared/services/config/config.service';

export interface User {
  Id: string;
  name: string;
  last_name: string;
  email: string;
  age: number;
}

interface ApiResponse {
  data: User;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private readonly url: string;
  private readonly token: string;

  constructor(private http: HttpClient, private configService: ConfigService) {
    this.url = this.configService.apiBaseUrl;
    this.token = this.configService.apiToken;
  }

  /**
   * Fetches the current user's data.
   * Replace with actual implementation, e.g., fetching based on logged-in user ID.
   */
  getCurrentUser(id: string): Observable<User> {
    return this.http.get<ApiResponse>(`${this.url}/users/${id}`, {
      headers: {
        Authorization: `Bearer ${this.token}`
      }
    }).pipe(
      map((response: ApiResponse) => {
        return response.data;
      })
    );
  }

}