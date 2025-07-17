import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class ThemeService {
    private static readonly THEME_KEY = 'coreui-free-angular-admin-template-theme-default';
    private themeSubject: BehaviorSubject<string | null>;

    public isDarkTheme$: Observable<boolean>;

    constructor() {
        const savedTheme = localStorage.getItem(ThemeService.THEME_KEY);
        this.themeSubject = new BehaviorSubject<string | null>(JSON.parse(savedTheme || 'null'));
        this.isDarkTheme$ = this.themeSubject.asObservable().pipe(map(theme => theme === 'dark'));
        this.setupStorageListener();
    }

    private setupStorageListener(): void {
        window.addEventListener('storage', this.handleStorageChange.bind(this));

        const originalSetItem = localStorage.setItem;
        localStorage.setItem = function (key: string, value: string) {
            originalSetItem.apply(this, [key, value]);
            if (key === ThemeService.THEME_KEY) {
                window.dispatchEvent(new StorageEvent('storage', {
                    key: key,
                    newValue: value,
                    storageArea: localStorage
                }));
            }
        };
    }

    private handleStorageChange(event: StorageEvent): void {
        if (event.key === ThemeService.THEME_KEY) {
            this.themeSubject.next(JSON.parse(event.newValue || 'null'));
        }
    }
}