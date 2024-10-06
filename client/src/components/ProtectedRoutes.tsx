import { useLocation, Outlet, Navigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'

import routes from '../routes'

/**
* Protected routes checks if user is logged and manages the authorization
* redirects. 
* 
* Renders children routes.
* 
* @remarks  It must have a AuthProvider as its parent
*/
const ProtectedRoutes = () => {
    const url = useLocation().pathname;

    // Might become hash maps if there are too many routes so speed go to O(1)
    const isAuthRoute = routes.auth.some((route) => (url.includes(route)))
    const isPrivateRoute = routes.private.some((route) => (url.includes(route))) 

    const { user } = useAuth();

    const isSigned = Boolean(user);

    // Root page
    if (url == "/") return <Navigate to="/home"/>;
    // Handle private route redirect
    if (!isSigned && isPrivateRoute) return <Navigate to="/auth"/>;
    // Handle auth route redirect (for signed users to not signin again)
    if (isSigned && isAuthRoute) return <Navigate to="/home"/>
    // Serves the routes

    return (
        <Outlet/>
    )
}

export default ProtectedRoutes