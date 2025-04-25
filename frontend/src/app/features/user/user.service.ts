import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ConfigService } from '../../../shared/services/config/config.service';

export interface User {
  id?: number;
  username: string;
  email: string;
  bio?: string;
  urls?: {
    website?: string;
    twitter?: string;
  };
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
    return this.http.get<User>(`${this.url}/v1/users/${id}`, {
      headers: {
      Authorization: `Bearer ${this.token}`
      }
    });
  }

}