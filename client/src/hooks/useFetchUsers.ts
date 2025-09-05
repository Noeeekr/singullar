import { useMemo } from "react";
import { SERVER_ADDR } from "../configs";

import type { UserRoles, User } from "../types/server";

import useContextAwareFetch from "./useContextAwareFetch";

// Role argument must be memoized or declared outside components to avoid unnecessary re-renders.
function useFetchUsers(
    roles: UserRoles[]
): [User[], boolean, string, () => void] {
    const options: RequestInit = useMemo(() => ({
        headers: {
            "Content-Type": "application/json",
        },
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
            target_roles: roles,
        }),
    }), [roles]);

    const [response, isLoading, error, fetch] = useContextAwareFetch<User[]>(
        `${SERVER_ADDR}/api/institution/users`,
        options
    );

    return [response?.data == null ? [] : response.data, isLoading, error, fetch];
}

export default useFetchUsers;
