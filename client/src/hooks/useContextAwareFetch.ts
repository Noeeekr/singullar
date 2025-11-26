import {
    useCallback,
    useState
} from "react"
import {
    DefaultResponse
} from "../models/server"

import {
    useNavigate
} from "react-router-dom"

import {
    updateUser as actionUpdateUser,
} from "@slices/userSlice"
import {
    useAppDispatch
} from "@slices/store"

export interface ResponseObject<ResponseData> {
    response: ResponseData | null,
    error: string,
    status: number,
}
export interface OnResponseCallbackProps<ResponseData, RequestBody> extends ResponseObject<ResponseData> {
    body?: RequestBody
}

export type OnResponseCallback<ResponseData, RequestBody> = (props: OnResponseCallbackProps<ResponseData, RequestBody>) => void

export interface useContextAwareFetchReturn<ResponseData, RequestBody = void> extends ResponseObject<ResponseData> {
    isLoading: boolean,

    send: (body?: RequestBody) => void,
    registerOnResponseCallback: (cb: OnResponseCallback<ResponseData, RequestBody>) => void
}

export const defaultRequestInit: RequestInit = {
    method: "GET",
    headers: {
        "Content-Type": "application/json",
    },
    credentials: 'include',
    cache: 'no-cache',
}

function appendRequestBodyIfNecessary<RequestBody>(init: RequestInit | undefined, body: RequestBody) {
    if (init && init.method != "GET") {
        init.body = JSON.stringify(body)
    }
}
// useContextAwareFetch is a wrapper around fetch that checks the responses from server for specific events in each call. It returns a JSON
// useContextAwareFetch will cause unecessary rerenders if its arguments are non-memoized objects
function useContextAwareFetch<ResponseData, RequestBody = null>(
    input: string | URL | globalThis.Request,
    init?: RequestInit,
): useContextAwareFetchReturn<ResponseData, RequestBody> {
    const [response, setResponse] = useState<ResponseData | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string>("");
    const [status, setStatus] = useState(0);
    const [onResponseCallbacks, setOnResponseCallbacks] = useState<OnResponseCallback<ResponseData, RequestBody>[]>([])

    const navigate = useNavigate()
    const dispatch = useAppDispatch()

    const send = useCallback(async (body?: RequestBody) => {
        setResponse(null)
        setIsLoading(true)
        setError("")
        try {
            appendRequestBodyIfNecessary(init, body)
            const res = await fetch(input, init)

            setStatus(res.status)
            if (res.status == 401) {
                dispatch(actionUpdateUser(null))
                navigate("/auth")
                return
            }

            const responseBody: DefaultResponse<ResponseData> = await res.json()
            if (responseBody.data != null) {
                setResponse(responseBody.data)
            } else {
                setError(responseBody.error)
            }

            onResponseCallbacks.forEach((cb) => cb({
                body: body,
                error: responseBody.error,
                status,
                response: responseBody.data
            }))
        } catch (e) {
            console.log(e)
            setError("Falha ao processar a requisição")
        } finally {
            setIsLoading(false)
        }
    }, [init, input, navigate, dispatch]);

    const registerOnResponseCallback = (cb: OnResponseCallback<ResponseData, RequestBody>) => {
        setOnResponseCallbacks((callbacks) => {
            callbacks.push(cb)
            return callbacks
        })
    }

    return {
        response,
        isLoading,
        error,
        send,
        status,
        registerOnResponseCallback,
    }
}

export default useContextAwareFetch;