import { INavData } from '@coreui/angular';

export const navItems: INavData[] = [
  // {
  //   name: 'Dashboard',
  //   url: '/dashboard',
  //   iconComponent: { name: 'cil-speedometer' },
  //   badge: {
  //     color: 'info',
  //     text: 'NEW'
  //   }
  // },
  // {
  //   title: true,
  //   name: 'Extras'
  // },
  {
    name: 'Finance',
    url: '/finance',
    iconComponent: { name: 'cil-dollar' },
    children: [
      {
        name: 'Wallet',
        url: '/wallet',
        icon: 'nav-icon-bullet'
      },


    ]
  },

];
