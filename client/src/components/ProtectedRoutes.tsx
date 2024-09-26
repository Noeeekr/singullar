import { useLocation, Outlet, Navigate } from 'react-router-dom'
import useAuth from '../hooks/useAuthContext'

import routes from '../routes'

const ProtectedRoutes = () => {
    const url = useLocation().pathname;

    // Might become hash maps if there are too many routes so speed go to O(1)
    const isAuthRoute = routes.auth.some((route) => (url.includes(route)))
    const isPrivateRoute = routes.private.some((route) => (url.includes(route))) 

    const isSigned = (useAuth().user != null)

    console.log("Auth state (user) and (user != null):")
    console.log(useAuth().user)
    console.log(useAuth().user != null)

    // Root page
    if (url == "/") return <Navigate to="/home"/>;
    // Handle private route redirect
    if (!isSigned && isPrivateRoute) return <Navigate to="/auth"></Navigate>;
    // Handle auth route redirect (for signed users to not signin again)
    if (isSigned && isAuthRoute) return <Navigate to="/home"></Navigate>
    // Serves the routes
    return <Outlet/>
}

export default ProtectedRoutes