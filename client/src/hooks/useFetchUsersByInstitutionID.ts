import { useMemo, useState } from "react"

import { SERVER_ADDR } from "../configs"

import type { User } from "../types/server"

import useContextAwareFetch from "./useContextAwareFetch"

// Fetch users using their institutionID
function useFetchUsersByInstitutionID(id: number): [User[], boolean, string | null, () => void] {
    const [error, setError] = useState<string>("")

    const options: RequestInit = useMemo(() => (
            {
                "headers": {
                    "Content-Type": "application/json"
                },
                "method": "POST",
                "credentials": "include",
                "body": JSON.stringify({id: id}),
            }
    ),[id])
    const [response, isLoading, fetch] = useContextAwareFetch<User[]>(
        `${SERVER_ADDR}/api/institution/users`,
        options
    )

    if (response?.error) {
        setError(response.error)
    }

    return [response?.data == null ? [] : response.data, isLoading, error, fetch];
}

export default useFetchUsersByInstitutionID