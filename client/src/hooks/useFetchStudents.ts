import { User, UserSegments } from "@models/server"
import useContextAwareFetch, { useContextAwareFetchReturn } from "./useContextAwareFetch"
import { SERVER_ADDR } from "../configs"

export interface StudentFilters {
    name?: string,
    id?: number,
    class_id?: number,
    email?: string,
    segment?: UserSegments,
}

const useFetchStudents = (): useContextAwareFetchReturn<User[], StudentFilters[]> => {
    return useContextAwareFetch<User[], StudentFilters[]>(
        `${SERVER_ADDR}/api/students`,
        {
            method: 'POST',
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            cache: 'no-cache',
        },
    )
}

export default useFetchStudents;