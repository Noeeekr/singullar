import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import { 
  ThemeProvider,
  CssBaseline
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
import AuthPage from './pages/auth/AuthPage'

// This and route provider component might become App.tsx file
// so I can put redux store and things like that there.

const App = (): JSX.Element => {
  return(
    <Routes>
      <Route path="/" element={<ProtectedRoutes/>}>
        <Route path="/home" element={<div>Home page</div>}/>
        <Route path="/auth" element={<AuthPage/>} />
      </Route>
    </Routes>
  )
}


createRoot(document.getElementById('root')!).render(
  <AuthContextProvider>
  <ThemeProvider theme={lightTheme}>
    <CssBaseline />
    <Router>
      <App />
    </Router>
  </ThemeProvider>
</AuthContextProvider>
)
