import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import {
  ThemeProvider,
  CssBaseline,
  GlobalStyles,
} from '@mui/material'
import {
  lightTheme
} from './themes'

import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom"

import { AuthContextProvider } from './context/authProvider'
import ProtectedRoutes from './components/ProtectedRoutes'
import AuthPage from './pages/auth/Auth'
import Layout from './pages/home/Layout'

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
const App = (): JSX.Element => {
  return (
    <Routes>
      <Route path="/" element={<ProtectedRoutes />}>
        <Route path="home" element={<Layout />}>
          <Route path="" element={<div>Home Root page</div>} />
        </Route>
        <Route path="auth" element={<AuthPage />} />
        <Route path="*" element={<div>not found</div>} />
      </Route>
    </Routes>
  )
}


createRoot(document.getElementById('root')!).render(
      <ThemeProvider theme={lightTheme}>
          <CssBaseline />
          <GlobalStyle />
          <Router>
            <AuthContextProvider>
              <App />
            </AuthContextProvider>
          </Router>
      </ThemeProvider>
)
