import { useLocation, Outlet, useNavigate } from "react-router-dom";
import { useAppSelector } from "@slices/store";
import useSignOut from "@hooks/useSignout";

import RouteGuard, { routes } from "../../routes";
import { useEffect, useState } from "react";
import { UserRoles } from "../../models/server";

/**
 * Protected routes checks if user is logged and manages the authorization
 * redirects.
 *
 * Renders children routes.
 *
 * @remarks  It must be implemented with a redux toolkit store provider
 */
const ProtectedRoutes = () => {
    let role: UserRoles | null;
    {
        const user = useAppSelector((store) => store.user.user);
        role = user ? user.role : null
    }
    const routeGuard = new RouteGuard(routes);
    
    const navigate = useNavigate();
    
    const { pathname } = useLocation();
    const { signout } = useSignOut();
    
    const [isValidating, setIsValidating] = useState<boolean>(false);

    useEffect(() => {
        setIsValidating(true)
        const ok = routeGuard.validateRoute(pathname, role)
        if (!ok) {
            if (role == null) {
                signout() 
                navigate("/auth") 
                return
            }
            navigate("/home")
            return
        }
        setIsValidating(false)
    }, [pathname])

    if (isValidating) return <div>Loading Page...</div>
    return <Outlet/>
};

export default ProtectedRoutes;