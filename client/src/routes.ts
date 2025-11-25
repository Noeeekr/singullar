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
 * Made assuming a recursive type structure
 */
interface RouteConfig {
    PermitedRoles?: RouterRoles[]
    // Any other key in RouteConfig is treated as a subroute
    [key: string]: RouteConfig | RouterRoles[] | undefined
}

/**
 * The structure accepted by the route guard validation object
 */
export type Routes = Record<string, RouteConfig>;

/**
* Empty permited roles defaults to ANY_AUTHENTICATED_ROLE
*/
export const routes: Routes = {
    "auth": { PermitedRoles: [UNAUTHENTICATED_ROLE] },
    "admin": {
        PermitedRoles: [ROLE_ADMIN],
        "*": { PermitedRoles: [ROLE_ADMIN] }
    },
    "home": { PermitedRoles: [ANY_AUTHENTICATED_ROLE] },
    "platform": {
        "subjects": {
            PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR],
            "*": { PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR] },
        },
        "question": {
            "list": {
                PermitedRoles: undefined,
                "create": { PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR] },
                "*": { PermitedRoles: [ANY_AUTHENTICATED_ROLE] }
            },
            "create": { PermitedRoles: [ROLE_ADMIN, ROLE_SUPERVISOR] },
            "*": { PermitedRoles: [ANY_AUTHENTICATED_ROLE] }
        }
    }
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
    public validateRoute(pathname: string, targetRole: UserRoles | null): boolean {
        // Clean the pathname and split into segments
        const paths = pathname.split("/").filter(p => p.length > 0);

        if (paths.length === 0) {
            // The solicited path is not registered. 
            // Meaning the route could not be found to validate.
            return false;
        }

        let currentConfig: RouteConfig = this.routes;

        // Traverse all path components
        for (let i = 0; i < paths.length; i++) {
            const segment = paths[i];

            // Check if the current level has defined a role rule.
            if (i > 0) {
                const parentRouteNode = currentConfig;
                // If the current level has a defined role rule then check the user role agaisnt it
                if (parentRouteNode.PermitedRoles !== undefined) {
                    // If the validation fails, disable user to proceed to this route
                    if (!this.validateRole(parentRouteNode.PermitedRoles, targetRole)) {
                        return false;
                    }
                }
            }

            // If rule validati on doesn't fail or doesn't exist, move to the next segment directly
            // There is a **Noticeable** vulnerability here that the pathname could be the name of a reserved map key like
            // "PermitedRoles" which would lead to unexpected behavior.
            let nextConfig: RouteConfig | undefined = currentConfig[segment] as RouteConfig;

            // If next segment is found, move deeper and continue the loop.
            if (nextConfig !== undefined) {
                currentConfig = nextConfig;
                continue;
            }

            // If no segment match found. Check for a wildcard.

            const wildcardConfig: RouteConfig | undefined = currentConfig["*"] as RouteConfig;

            // If a wildcard exists, use its permissions and we are DONE for this path.
            if (wildcardConfig !== undefined) {
                // If a wildcard is found, it applies to this segment AND any segments that follow.
                // We validate the wildcard's role and return immediately.
                return this.validateRole(wildcardConfig.PermitedRoles, targetRole);
            }

            // No direct match and no wildcard match. The route is NOT defined.
            return false;
        }

        // 6. Loop finished: We reached the end of the path.
        // The final segment's own PermitedRoles must be validated.
        // Example: If path is "/home", loop finds "home". Now we validate "home"'s PermitedRoles.

        // `currentConfig` holds the configuration for the final segment (e.g., "platform.question.list")
        return this.validateRole(currentConfig.PermitedRoles, targetRole);
    }
}

export default RouteGuard