import { useCallback } from "react";
import { SERVER_ADDR } from "../configs";

import useContextAwareFetch from "./useContextAwareFetch";

import type { Response } from "./useContextAwareFetch";
import type { User, UserRoles } from "../models/server";
// Role argument must be memoized or declared outside components to avoid unnecessary re-renders.
function useFetchUsers(): Response<User[], UserRoles[]> {
    const [response, isLoading, error, send] = useContextAwareFetch<User[], UserRoles[]>(
        `${SERVER_ADDR}/api/users`,
        {
            headers: {
                "Content-Type": "application/json",
            },
            method: "POST",
            credentials: "include",
        }
    );

    const sendRequest = useCallback((body: UserRoles[]) => {
        send(body)
    }, [send])
    return [response, isLoading, error, sendRequest];
}

export default useFetchUsers;
