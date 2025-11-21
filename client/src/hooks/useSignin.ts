import { useState } from 'react'

import { DefaultResponse } from '../models/server'

import { updateUser } from '../slices/userSlice'
import { AppDispatch } from '../slices/store'
import { useDispatch } from 'react-redux'

import { SERVER_ADDR } from '../configs'
import type { User } from '../models/server'

const useSignIn = (): {
    signin: typeof signin, 
    isLoading: boolean, 
    signinError: string | null
} => {
    const [signinError, setSigninError] = useState<null | string>(null)
    const [isLoading, setIsLoading] = useState(false);

    const dispatch = useDispatch<AppDispatch>()

    const signin = async (email: string, password: string) => {
        setIsLoading(true)
        setSigninError(null)

        try {
            const response = await fetch(`${SERVER_ADDR}/api/auth`,{
                method: "POST",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    password: password,
                    email: email,
                })
            })

            const res: DefaultResponse<User> = await response.json()

            if (response.ok) {
                dispatch(updateUser(res.data))
                return true
            } else  {
                setSigninError(res.error)
                return false
            }
        } catch {
            setSigninError("Falha ao logar o usuário.")
            return false
        } finally {
            setIsLoading(false)
        }
    }

    return { signin, isLoading, signinError } 
}

export default useSignIn