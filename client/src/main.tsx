import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import { 
  ThemeProvider,
  CssBaseline
} from '@mui/material'
import { lightTheme } from './theme'

import { 
  createBrowserRouter,
  RouterProvider,
  Outlet
} from "react-router-dom"

import routes from './routes'

import RootPage, { loader as rootLoader } from './pages/root/RootPage'
import AuthPage from './pages/auth/AuthPage'

// This and route provider component might become App.tsx file
// so I can put redux store and things like that there.

const router = createBrowserRouter([
  { // Root's a redirect middleware for logged/non-logged
    path: "/",
    element: <Outlet></Outlet>,
    children: [
      {
        path: "/",
        element: <RootPage></RootPage>,
        loader: rootLoader
      },
      {
        path: "/*",
        element: <RootPage></RootPage>,
        loader: rootLoader
      },
      {
        path: routes.auth.main,
        element: <AuthPage></AuthPage>
      },
      {
        path: routes.auth.recover,
        element: <div>abc</div>
      }
    ]
  }
])


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={lightTheme}>
      <CssBaseline></CssBaseline>
      <RouterProvider router={router}></RouterProvider>
    </ThemeProvider>
  </StrictMode>,
)
