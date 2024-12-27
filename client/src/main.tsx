import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import { Provider } from 'react-redux'
import store from './slices/store'

import {
  ThemeProvider,
  CssBaseline,
  GlobalStyles,
} from '@mui/material'
import {
  lightTheme
} from './themes'

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import PgAuthentication from './pages/auth/Auth';
import {
  PgInstitutionSelection as PgInstitutionSelection, LoaderInstitutionSelection,
  PgMyClasses,
  PgClassesCreate,

  PgStudentsSearch,
  PgStudentsCreate,
} from './pages/admin';

import ProtectedRoutes from './components/ProtectedRoutes'
import Layout from './components/Layout'
import Root from './components/DefaultRoot';

const GlobalStyle = () => (
  <GlobalStyles
    styles={{
      'ul': {
        padding: 0,
        margin: 0,
      },
    }}
  />
);

// This and route provider component might become App.tsx file
// so I can put redux store and things like that there.
/*const router = createBrowserRouter(
  createRoutesFromElements(
    <Routes>
      <Route path="/" element={<ProtectedRoutes />}>
        <Route path="auth" element={<PgAuthentication />} />
        <Route path="home" element={<Layout />}>
          <Route index element={<div>Home root page</div>} />
          <Route path="*" element={<div>Not found specific home</div>} />
        </Route>
        <Route path="teacher" element={<Root />}>
          <Route index element={<div>Index page teach</div>} />
        </Route>
        <Route path="supervisor" element={<div>Layout</div>}>
          <Route index element={<Root />} />
        </Route>
        <Route path="admin" element={<Layout />}>
          <Route index element={<Root />} />
          <Route path="search" element={<PgInstitutionSelection />} loader={InstitutionSelectionLoader} />
          <Route path="classes/create" element={<PgCreate />} />
          <Route path="classes/:id" element={<PgClasses />} />
          <Route path="*" element={<div>Not index page adm</div>} />
        </Route>
        <Route path="*" element={<div>Not found general page</div>} />
      </Route>
    </Routes>
  )
)
  */
 
const r2 = createBrowserRouter([
  {
    path: "/",
    element: <ProtectedRoutes />,
    errorElement: <div>Error element1</div>,
    children: [
      {
        path: "auth",
        element: <PgAuthentication/>
      },
      {
        path: "home",
        element: <Layout />,
        children: [
          {
            index: true, 
            element: <div>Home root page</div>,
          },
          {
            path: "*",
            element: <div>Any path home page</div>,
          }
        ]
      },
      {
        path: "teacher",
        element: <Layout />,
        children: [
          {
            index: true, 
            element: <div>Teacher root page</div>,
          },
          {
            path: "*",
            element: <div>Teacher path home page</div>,
          }
        ]
      },
      {
        path: "supervisor",
        element: <Layout />,
        children: [
          {
            index: true, 
            element: <div>Supervisor root page</div>,
          },
          {
            path: "*",
            element: <div>Supervisor path home page</div>,
          }
        ]
      },
      {
        path: "admin",
        element: <Layout />,
        children: [
          {
            index: true,
            element: <Root/>,
          },
          {
            path: "students/search",
            element: <PgInstitutionSelection nextUrl="/admin/students" />,
            loader: LoaderInstitutionSelection,
          },
          {
            path: "students",
            element: <PgStudentsSearch />
          },
          {
            path: "students/create",
            element: <PgStudentsCreate />
          },
          {
            path: "classes/search",
            element: <PgInstitutionSelection nextUrl="/admin/classes" />,
            loader: LoaderInstitutionSelection,
          },
          {
            path: "classes/create",
            element: <PgClassesCreate />,
          },
          {
            path: "classes",
            element: <PgMyClasses />
          },
          {
            path: "*",
            element: <div>Any path admin page</div>,
          },
        ]
      },
      {
        path: "*",
        element: <div>Any route default page</div>,
      }
    ]
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={lightTheme}>
      <CssBaseline />
      <GlobalStyle />
      <Provider store={store}>
        <RouterProvider router={r2} />
      </Provider>
    </ThemeProvider>
  </StrictMode>
)
