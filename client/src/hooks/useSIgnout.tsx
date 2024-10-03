import { useState } from 'react'
import useAuth from './useAuth';

export const useSignOut = () => {
    const { dispatch } = useAuth();

    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const signout = async () => {
        setIsLoading(true)

        try {
            const response = await fetch("http://localhost:8000/auth/signout", {
                credentials: "include"
            })
            
            if (response.ok) {
                dispatch({
                    type: "logout",
                })
            } else {
                let res = await response.json()
                setError(res.Error)
            }

        } catch (err: unknown) {
            setError("Falha ao requisitar deslogamento do usuário")
        } finally {
            setIsLoading(false)
        }
    
    }

    return [signout, isLoading, error]
}