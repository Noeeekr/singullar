import { useState } from 'react'

import { IUserRequest } from '../types/user'

import { updateUser } from '../slices/authSlice'
import { AppDispatch } from '../slices/store'
import { useDispatch } from 'react-redux'

const useSignin = (): [typeof signin, boolean, string |null] => {
    const [ error, setError] = useState<null | string>(null)
    const [ isLoading, setIsLoading] = useState(false);

    const dispatch = useDispatch<AppDispatch>()

    const signin = async (email: string, password: string) => {
        setIsLoading(true)
        setError(null)
        
        try {
            const response = await fetch("http://localhost:8000/api/auth/signin",{
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
                setError(res.error)
            } else  {
                dispatch(updateUser(res.data))
            }
        } catch (err) {
            setError("Falha ao logar o usuário.")
        } finally {
            setIsLoading(false)
        }
    }

    return [signin, isLoading, error]
}

export default useSignin