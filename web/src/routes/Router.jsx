import React, { lazy } from 'react';
import { Navigate } from 'react-router-dom';
import Loadable from '../layouts/full/shared/loadable/Loadable';

/* ***Layouts**** */
const FullLayout = Loadable(lazy(() => import('../layouts/full/FullLayout')));
const BlankLayout = Loadable(lazy(() => import('../layouts/blank/BlankLayout')));

/* ****Pages***** */
const Dashboard = Loadable(lazy(() => import('../views/dashboard/Dashboard.jsx')))
const SamplePage = Loadable(lazy(() => import('../views/account/UserProfile.jsx')))
const ErrorPage = Loadable(lazy(() => import('../views/authentication/Error.jsx')));
const Register = Loadable(lazy(() => import('../views/authentication/Register.jsx')));
const Login = Loadable(lazy(() => import('../views/authentication/Login.jsx')));
const OfferedTrips = Loadable(lazy(() => import('../views/OfferedTrips/offeredtrips.jsx')));
const Survey = Loadable(lazy(() => import('../views/survey/surveypage.jsx')));
const Wallet = Loadable(lazy(() => import('../views/wallet/walletdetails.jsx')));
const Admin = Loadable(lazy(() => import('../views/authentication/Admin.jsx')));
const RoutesPage = Loadable(lazy(() => import('../views/dashboard/components/routespage.jsx')));
const Recommendation = Loadable(lazy(() => import('../views/recommendation/airecom.jsx')));
const Router = [
  {
    path: '/',
    element: <FullLayout />,
    children: [
      { path: '/', element: <Navigate to="/auth/login" /> },
      { path: '/dashboard', exact: true, element: <Dashboard /> },
      { path: '/sample-page', exact: true, element: <SamplePage /> },
      { path: '/offeredtrips', exact:true, element: <OfferedTrips/>},
      { path: '/airecom', exact:true, element: <Recommendation/>},
      { path: '/survey', exact:true, element: <Survey/>},
      { path: '/wallet', exact:true, element: <Wallet/>},
      { path: '/routespage', exact:true, element: <RoutesPage/>},
      { path: '*', element: <Navigate to="/auth/404" /> },
    ],
  },
  {
    path: '/auth',
    element: <BlankLayout />,
    children: [
      { path: '404', element: <ErrorPage /> },
      { path: '/auth/register', element: <Register /> },
      { path: '/auth/login', element: <Login /> },
      { path: '/auth/admin', element: <Admin /> },
      { path: '*', element: <Navigate to="/auth/404" /> },
    ],
  },
];

export default Router;
