import { useState } from 'react'

import { updateUser } from '../slices/userSlice'
import { useAppDispatch } from '../slices/store'

import { SERVER_ADDR } from '../configs'

const useSignOut = () => {
    const dispatch = useAppDispatch()

    const [isSigningOut, setIsSigningOut] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const signout = async () => {
        setIsSigningOut(true)

        try {
            const response = await fetch(`${SERVER_ADDR}/api/auth/`, {
                credentials: "include"
            })

            if (response.ok) {
                dispatch(updateUser(null))
            } else {
                throw new Error("Falha ao deslogar usuário")
            }
        } catch {
            setError("Falha ao requisitar deslogamento do usuário")
        }
    }
    
    return { signout, isSigningOut, error } 
}
export default useSignOut