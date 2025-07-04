import { useLocation, Outlet, useNavigate } from "react-router-dom";
import routes from "../routes";

import { useAppSelector } from "@slices/store";
import useSignOut from "@hooks/useSignout";

import { useEffect, useState } from "react";
/**
 * Protected routes checks if user is logged and manages the authorization
 * redirects.
 *
 * Renders children routes.
 *
 * @remarks  It must be implemented with a redux toolkit store provider
 */
const ProtectedRoutes = () => {
    const user = useAppSelector((store) => store.user.user);
    const { signout } = useSignOut();
    const [isAllowed, setIsAllowed] = useState<boolean>(false);

    const url = useLocation().pathname;
    const roleRoutes = routes.find((route) => route.role == user?.role);

    const navigate = useNavigate();

    useEffect(() => {   
        // if user == null || role not found || role == null
        if (!roleRoutes) {
            signout();
            navigate("/auth");
            return;
        }

        // Check if the route is inside one of the permited routes
        const route = roleRoutes.routes.some((route) => url.toString().startsWith(route));

        if (!route) {
            // if there's another route to redirect
            if (roleRoutes.routes.length) {
                navigate(roleRoutes.routes[0]);
                return
            }
            // if there's no routes on that role you shouldn't even be using it
            signout();
            navigate("/auth/");
            return
        }
        if (!isAllowed) {
            setIsAllowed(true)
        }
    }, [url, roleRoutes, navigate, signout, isAllowed]);
    if (isAllowed) {
        return <Outlet />;
    } else {
        return <div>Loading...</div>;
    }
};

export default ProtectedRoutes;
