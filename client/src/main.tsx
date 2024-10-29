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
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom"

import ProtectedRoutes from './components/ProtectedRoutes'
import AuthPage from './pages/auth/Auth'
import Layout from './components/Layout'

import RootAdmin from './pages/admin/Root';

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
          <Route index element={<div>Home root page</div>} />
          <Route path="*" element={<div>Not found specific home</div>} />
        </Route>  
        <Route path="auth" element={<AuthPage />} />
        <Route path="teacher" element={<div>Layout</div>}>
          <Route index element={<div>Index page teach</div>} />
        </Route>
        <Route path="supervisor" element={<div>Layout</div>}>
          <Route index element={<div>Index page superv</div>} />
        </Route>
        <Route path="admin" element={<Layout />}>
          <Route index element={<RootAdmin/>} />
          <Route path="*" element={<div>Not index page adm</div>} />
        </Route>
        <Route path="*" element={<div>Not found general page</div>} />
      </Route>
    </Routes>
  )
}


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={lightTheme}>
      <CssBaseline />
      <GlobalStyle />
      <Provider store={store}>
        <Router>
          <App />
        </Router>
      </Provider>
    </ThemeProvider>
  </StrictMode>
)
