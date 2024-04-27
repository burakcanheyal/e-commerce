import { uniqueId } from 'lodash';
import DashboardCustomizeIcon from '@mui/icons-material/DashboardCustomize';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import LoginIcon from '@mui/icons-material/Login';
import HowToRegIcon from '@mui/icons-material/HowToReg';
import Person2Icon from '@mui/icons-material/Person2';
import WalletIcon from '@mui/icons-material/Wallet';
const MenuItems = [
  {
    navlabel: true,
    subheader: 'Home',
  },

  {
    id: uniqueId(),
    title: 'Dashboard',
    icon: DashboardCustomizeIcon,
    href: '/dashboard',

  },
  {
    id: uniqueId(),
    title: 'Offered Trips',
    icon: LocalOfferIcon,
    href: '/offeredtrips',
  },
  {
    id: uniqueId(),
    title: 'Routes Page',
    href: '/routespage',
  },
  {
    id: uniqueId(),
    title: 'Survey',
    href: '/survey',
  },
  {
    id: uniqueId(),
    title: 'Wallet',
    href: '/wallet',
    icon: WalletIcon,
  },
  {
    navlabel: true,
    subheader: 'Auth',
  },
  {
    id: uniqueId(),
    title: 'Login',
    icon: LoginIcon,
    href: '/auth/login',
  },
  {
    id: uniqueId(),
    title: 'Register',
    href: '/auth/register',
    icon: HowToRegIcon,
  },
  {
    navlabel: true,
    subheader: 'Extra',
  },

  {
    id: uniqueId(),
    title: 'User Profile',
    icon: Person2Icon,
    href: '/sample-page',
  },
];

export default MenuItems;
