import { ROLE_ADMIN, ROLE_SUPERVISOR, UserRoles } from "./models/server";

/**
 * Utilized by route guard to allow route to everyone
 */
const ANY_ROLE = "any_role"
/**
 * Utilized by route guard to allow route to everyone but unauthenticated users.
 */
const ANY_AUTHENTICATED_ROLE = "any_authenticated_role"
/**
 * Utilized by route guard to only allow route to unauthenticated users.
 */
const UNAUTHENTICATED_ROLE = null

/**
 * All routes that can be providen to route guard rules
 */
export type RouterRoles = UserRoles
    | typeof ANY_ROLE
    | typeof ANY_AUTHENTICATED_ROLE
    | typeof UNAUTHENTICATED_ROLE

/**
 * The structure accepted by the route guard validation object
 */
export interface Routes {
    [index: string]: {
        PermitedRoles?: RouterRoles[]
        Routes?: Routes,
    }
}

/**
* Empty permited roles defaults to ANY_AUTHENTICATED_ROLE
*/
export const routes: Routes = {
    "auth": { PermitedRoles: [UNAUTHENTICATED_ROLE] },
    "admin": { 
        Routes: {
            "*": { PermitedRoles: [ROLE_ADMIN] }
        },
        PermitedRoles: [ROLE_ADMIN] 
    },
    "home": { PermitedRoles: [ANY_AUTHENTICATED_ROLE] },
    "platform": {
        Routes: {
            "subjects": {
                PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR],
                Routes: {
                    "*": { PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR]},
                },
            },
            "question": {
                Routes: {
                    "list": {
                        Routes: {
                            "create": { PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR] },
                        },
                    },
                    "create": {
                        PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR],
                    },
                    "*": { PermitedRoles: [ANY_AUTHENTICATED_ROLE] }
                }
            }
        }
    },
}

class RouteGuard {
    private routes: Routes

    constructor(routes: Routes) {
        this.routes = routes
    }

    public validateRole(permitedRoles: RouterRoles[] | undefined, targetRole: UserRoles | null): boolean {
        if (permitedRoles == undefined) return targetRole != null;
        if (targetRole == null) return permitedRoles.some((roles) => roles == null || roles == ANY_ROLE)
        if (permitedRoles.some((role) => role == ANY_AUTHENTICATED_ROLE || role == ANY_ROLE || role == targetRole)) {
            return true
        }
        return false
    }
    /**
    * 
    * @param pathname The current location url pathname part.
    *  
    * @param targetRole The roles the user contains to be checked agaisnt the allowed roles. 
    * @returns A boolean that is [true] if the user is permited on the route and [false] if not.
    */
    public validateRoute(pathname: string, targetRole: UserRoles | null): boolean {
        const paths = pathname.split("/");
        let route: Routes = this.routes

        for (let i = 1; i < paths.length; i++) {
            let currentRoute = route[paths[i]]

            // No further routes registered, no wildcard registered.
            if (currentRoute == undefined) return false;
            // Found the deepest route needed. Trigger role validation and liberate if allowed
            if (paths.length == i + 1) return this.validateRole(currentRoute.PermitedRoles, targetRole);
            // There are no further routes registered. Since there will be a further check that will fail, fail now instead.
            if (currentRoute.Routes == undefined) return false;
            // There are further routes, but the wanted route is not registered for validation.
            if (currentRoute.Routes[paths[i + 1]] == undefined) {
                // Tries to find a wildcard, if found, validates through wildcard's route guard
                if (currentRoute.Routes["*"] != undefined) {
                    return this.validateRole(currentRoute.Routes["*"].PermitedRoles, targetRole)
                }
                return false
            }

            // Validates if the user is allowed through this route guard
            if (!this.validateRole(currentRoute.PermitedRoles, targetRole)) {
                return false
            }

            // Updates the route
            route = currentRoute.Routes
        }
        return true
    }
}

export default RouteGuard