import { SERVER_ADDR } from "../configs";

import useContextAwareFetch from "./useContextAwareFetch";

import type { ResponseUtilities } from "./useContextAwareFetch";
import type { User, UserRoles } from "@models/server";

interface UserRequest {
	"accepted_roles"?:   UserRoles[]
    "filters"?: UserFilter[]
    "offset"?: number
}
interface UserFilter {
	"id"?: number
}

// Role argument must be memoized or declared outside components to avoid unnecessary re-renders.
function useFetchUsers(): ResponseUtilities<User[], UserRequest> {
    return useContextAwareFetch<User[], UserRequest>(
        `${SERVER_ADDR}/api/users`,
        {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            credentials: "include",
        }
    );
}

export default useFetchUsers;
