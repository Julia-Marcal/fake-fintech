import { Routes } from '@angular/router';
import { DefaultLayoutComponent } from './layout';

export const routes: Routes = [
  {
    path: '',
    component: DefaultLayoutComponent,
    data: {
      title: 'Home'
    },
    children: [
      {
        path: 'user',
        loadComponent: () => import('./features/user/user.component').then(m => m.UserComponent),
        data: {
          title: 'User Profile'
        }
      },
      {
        path: 'wallet',
        loadComponent: () => import('./features/wallet/wallet.component').then(m => m.WalletComponent),
        data: {
          title: 'Wallets'
        }
      },
      {
        path: 'wallet/:id/stocks',
        loadComponent: () => import('./features/stock/stock.component').then(m => m.StockComponent),
        data: {
          title: 'Wallet Stocks'
        }
      },
    ]
  },
  {
    path: '404',
    loadComponent: () => import('./features/page404/page404.component').then(m => m.Page404Component),
    data: {
      title: 'Page 404'
    }
  },
  {
    path: 'login',
    loadComponent: () => import('./features/login/login.component').then(m => m.LoginComponent),
    data: {
      title: 'Login Page'
    }
  },
  {
    path: 'register',
    loadComponent: () => import('./features/register/register.component').then(m => m.RegisterComponent),
    data: {
      title: 'Register Page'
    }
  },
  { path: '**', redirectTo: 'home' }
];