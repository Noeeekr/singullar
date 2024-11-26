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
    const { isSigned, isLoading, user } = useAuth()
    
    // Gets the default page (home) url for every role
    const DefaultUserRouteByRole = (() => {
        switch(user?.role) {
            case "student":
                return routes.student[0]
            case "admin":
                return routes.admin[0]
            case "supervisor":
                return routes.supervisor[0]
            case "teacher":
                return routes.teacher[0]
            default:
                return routes.auth[0]
        }
    })();

    const isUserRoleRoute = (() => {
        switch(user?.role) {
            case "student":
                return routes.student.some(route => (url.includes(route)))
            case "teacher":
                return routes.teacher.some(route => (url.includes(route)))
            case "supervisor":
                return routes.supervisor.some(route => (url.includes(route)))
            case "admin":
                return routes.admin.some(route => (url.includes(route)))
            default:
                return false
        }
    })();

    // Serves the routes
    if (isLoading) {        
        return <div>Redirecting...</div>
    }
    
    // redirect cases
    if (!isSigned) {
        if (url === "/" || isPrivateRoute) {
          return <Navigate to="/auth" />;
        }
      }
    
      if (isSigned) {
        if (url === "/") {
          return <Navigate to={DefaultUserRouteByRole} />;
        }
        if (isAuthRoute || !isUserRoleRoute) {
          return <Navigate to={DefaultUserRouteByRole} />;
        }
      }
          
    return (
        <Outlet/>
    )
}

export default ProtectedRoutes;