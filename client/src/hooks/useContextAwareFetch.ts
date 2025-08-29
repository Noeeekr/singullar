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
// useEventAwareFetch is a wrapper around fetch that checks the responses from server for specific events in each call. It returns a JSON
function useContextAwareFetch<ServerResponseDataType>(
    input: string | URL | globalThis.Request,
    init?: RequestInit,
): [DefaultResponse<ServerResponseDataType> | null, boolean, () => void] {
    const [response, setResponse] = useState<DefaultResponse<ServerResponseDataType> | null>(null)
    const [isLoading, setIsLoading] = useState(false)

    const navigate = useNavigate()
    const dispatch = useAppDispatch()
    console.log("res null, loading", response == null, isLoading)
    console.log("Parent", init, input)
    const fetchAndValidate = useCallback(async () => {
        console.log("Child")
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
        } finally {
            setIsLoading(false)
        }
    }, [init, input, navigate, dispatch]); 

    return [response, isLoading, fetchAndValidate]
}

export default useContextAwareFetch;