import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { AuthService } from '../../src/shared/services/auth/auth.service';
import { ConfigService } from '../../src/shared/services/config/config.service';

describe('AuthService', () => {
    let service: AuthService;
    let httpMock: HttpTestingController;
    let configService: ConfigService;
    let localStorageSpy: jest.SpyInstance;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
            providers: [
                AuthService,
                {
                    provide: ConfigService,
                    useValue: {
                        apiBaseUrl: 'http://api.example.com'
                    }
                }
            ]
        });

        service = TestBed.inject(AuthService);
        httpMock = TestBed.inject(HttpTestingController);
        configService = TestBed.inject(ConfigService);

        // Mock localStorage
        localStorageSpy = jest.spyOn(Storage.prototype, 'setItem');
        jest.spyOn(Storage.prototype, 'getItem');
        jest.spyOn(Storage.prototype, 'removeItem');
    });

    afterEach(() => {
        httpMock.verify();
        jest.clearAllMocks();
    });

    it('should login and set token', (done) => {
        const mockResponse = { token: 'mock-token' };

        service.login('test@test.com', 'password').subscribe({
            next: response => {
                expect(response).toEqual(mockResponse);
                expect(service.isAuthenticatedValue).toBe(true);
                expect(localStorageSpy).toHaveBeenCalledWith('currentUser', 'mock-token');
                done();
            },
            error: err => {
                fail(err);
                done();
            }
        });

        const req = httpMock.expectOne(`${configService.apiBaseUrl}/login`);
        expect(req.request.method).toBe('POST');
        expect(req.request.body).toEqual({
            email: 'test@test.com',
            password: 'password'
        });
        req.flush(mockResponse);
    });

    it('should handle login failure', (done) => {
        const errorResponse = { status: 401, statusText: 'Unauthorized' };

        service.login('wrong@test.com', 'wrong').subscribe({
            next: () => {
                fail('Expected login to fail');
                done();
            },
            error: error => {
                expect(error.status).toBe(401);
                expect(service.isAuthenticatedValue).toBe(false);
                expect(localStorageSpy).not.toHaveBeenCalled();
                done();
            }
        });

        const req = httpMock.expectOne(`${configService.apiBaseUrl}/login`);
        req.flush('Invalid credentials', errorResponse);
    });
});