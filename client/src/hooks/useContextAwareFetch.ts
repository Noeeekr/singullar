import { 
    useCallback, 
    useState 
} from "react"
import { 
    DefaultResponse 
} from "../types/server"

import { 
    useNavigate 
} from "react-router-dom"

import {
    updateUser as actionUpdateUser,
} from "@slices/userSlice"
import {
    useAppDispatch
} from "@slices/store"

// useContextAwareFetch is a wrapper around fetch that checks the responses from server for specific events in each call. It returns a JSON
// useContextAwareFetch will cause unecessary rerenders if its arguments are non-memoized objects
function useContextAwareFetch<ServerResponseDataType>(
    input: string | URL | globalThis.Request,
    init?: RequestInit,
): [DefaultResponse<ServerResponseDataType> | null, boolean, string, () => void] {
    const [response, setResponse] = useState<DefaultResponse<ServerResponseDataType> | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string>("");

    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const fetchAndValidate = useCallback(async () => {
        setIsLoading(true)
        try {
            const response = await fetch(input, init)
            if (response.status == 401) {
                dispatch(actionUpdateUser(null))
                navigate("/auth")
                return
            }
            const responseBody: DefaultResponse<ServerResponseDataType> = await response.json()
            setResponse(responseBody)
        } catch(e: unknown) {
            if (e instanceof Error) {
                setError(e.message)
            }
        } finally {
            setIsLoading(false)
        }
    }, [init, input, navigate, dispatch]); 

    return [response, isLoading, error, fetchAndValidate]
}

export default useContextAwareFetch;