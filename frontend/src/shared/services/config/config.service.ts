import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  constructor() {}

  get apiBaseUrl(): string {
    return 'http://localhost:8080/api';
  }

  get apiToken(): string {
    return 'token';
  }
}
