import { useState } from 'react'

import { IUserRequest } from '../types/user'

import { updateUser } from '../slices/authSlice'
import { AppDispatch } from '../slices/store'
import { useDispatch } from 'react-redux'

const useSignin = (): {
    signin: typeof signin, 
    isLoading: boolean, 
    signinError: string |null
} => {
    const [signinError, setSigninError] = useState<null | string>(null)
    const [isLoading, setIsLoading] = useState(false);

    const dispatch = useDispatch<AppDispatch>()

    const signin = async (email: string, password: string) => {
        setIsLoading(true)
        setSigninError(null)
        
        try {
            const response = await fetch("http://localhost:8000/api/user/signin",{
                method: "POST",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    password: password,
                    email: email,
                })
            })

            const res: IUserRequest = await response.json()

            if (!response.ok) {
                setSigninError(res.error)
                return false
            } else  {
                dispatch(updateUser(res.data))
                return true
            }
        } catch (err) {
            setSigninError("Falha ao logar o usuário.")
            return false
        } finally {
            setIsLoading(false)
        }
    }

    return { signin, isLoading, signinError } 
}

export default useSignin