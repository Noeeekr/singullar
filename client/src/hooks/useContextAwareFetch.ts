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

type Response<ResponseData> = [
    DefaultResponse: DefaultResponse<ResponseData> | null,
    isLoading: boolean,
    error: string,
    doFetch: () => void,
]

// useContextAwareFetch is a wrapper around fetch that checks the responses from server for specific events in each call. It returns a JSON
// useContextAwareFetch will cause unecessary rerenders if its arguments are non-memoized objects
function useContextAwareFetch<ResponseData = unknown>(
    input: string | URL | globalThis.Request,
    init?: RequestInit,
): Response<ResponseData> {
    const [response, setResponse] = useState<DefaultResponse<ResponseData> | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string>("");

    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const doFetch = useCallback(async () => {
        setIsLoading(true)
        setError("")
        try {
            const response = await fetch(input, init)
            if (response.status == 401) {
                dispatch(actionUpdateUser(null))
                navigate("/auth")
                return
            }
            const responseBody: DefaultResponse<ResponseData> = await response.json()
            setResponse(responseBody)
        } catch {
            setError("Falha ao processar a requisição")
        } finally {
            setIsLoading(false)
        }
    }, [init, input, navigate, dispatch]); 

    return [response, isLoading, error, doFetch]
}

export default useContextAwareFetch;