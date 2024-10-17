import { useLocation, Outlet, Navigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'

import routes from '../routes'

/**
* Protected routes checks if user is logged and manages the authorization
* redirects. 
* 
* Renders children routes.
* 
* @remarks  It must be implemented with a redux toolkit store provider 
*/
const ProtectedRoutes = () => {
    const url = useLocation().pathname;
    
    // Changing to maps is a option to improve speed if there are too many routes
    const isAuthRoute = routes.auth.some((route) => (url.includes(route)))
    const isPrivateRoute = routes.private.some((route) => (url.includes(route))) 
    
    // Authenticate user in every protected route
    const { isSigned, isLoading } = useAuth()
    
    // Serves the routes
    if (isLoading) {        
        return <div>Redirecting...</div>
    }
    
    // redirect cases
    if (url == "/" && isSigned) return <Navigate to="/home"/>;
    if (url == "/" && !isSigned) return <Navigate to="/auth"/>;

    if (!isSigned && isPrivateRoute) return <Navigate to="/auth"/>;

    if (isSigned && isAuthRoute) return <Navigate to="/home"/>;
    
    return (
        <Outlet/>
    )
}

export default ProtectedRoutes