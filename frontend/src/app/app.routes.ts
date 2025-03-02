import { Routes } from '@angular/router';
import { DefaultLayoutComponent } from './layout';

export const routes: Routes = [
  // children: [
  //   {
  //     path: 'theme',
  //     loadChildren: () => import('../shared/base/theme/routes').then((m) => m.routes)
  //   },
  //   {
  //     path: 'base',
  //     loadChildren: () => import('../shared/base/routes').then((m) => m.routes)
  //   },
  //   {
  //     path: 'pages',
  //     loadChildren: () => import('./routes').then((m) => m.routes)
  //   }
  // ]

  {
    path: 'home',
    component: DefaultLayoutComponent,
    data: {
      title: 'Home'
    }
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