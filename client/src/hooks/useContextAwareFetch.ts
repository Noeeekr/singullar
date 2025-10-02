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

type Response<ResponseData, RequestBody> = [
    response: ResponseData | null,
    isLoading: boolean,
    error: string,
    doFetch: (body: RequestBody) => void,
]

// useContextAwareFetch is a wrapper around fetch that checks the responses from server for specific events in each call. It returns a JSON
// useContextAwareFetch will cause unecessary rerenders if its arguments are non-memoized objects
function useContextAwareFetch<ResponseData = unknown, RequestBody = unknown>(
    input: string | URL | globalThis.Request,
    init?: RequestInit,
): Response<ResponseData, RequestBody> {
    const [response, setResponse] = useState<ResponseData | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string>("");

    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const send = useCallback(async (body: RequestBody) => {
        setIsLoading(true)
        setError("")
        try {
            if (init != null) {
                init.body = JSON.stringify(body);
            }
            const response = await fetch(input, init)
            if (response.status == 401) {
                dispatch(actionUpdateUser(null))
                navigate("/auth")
                return
            }
            const responseBody: DefaultResponse<ResponseData> = await response.json()
            if (responseBody.data != null) {
                setResponse(responseBody.data)
            } else {
                setError(responseBody.error)
            }
        } catch {
            setError("Falha ao processar a requisição")
        } finally {
            setIsLoading(false)
        }
    }, [init, input, navigate, dispatch]); 

    return [response, isLoading, error, send]
}

export default useContextAwareFetch;