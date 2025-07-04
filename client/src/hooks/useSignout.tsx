import { useState } from 'react'

import { useDispatch } from 'react-redux'
import { updateUser } from '../slices/userSlice'
import { AppDispatch } from '../slices/store'

const useSignOut = () => {
    const dispatch = useDispatch<AppDispatch>()

    const [isSigningOut, setIsSigningOut] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const signout = async () => {
        setIsSigningOut(true)

        try {
            const response = await fetch("http://localhost:8000/api/auth/", {
                credentials: "include"
            })
            
            if (response.ok) {
                dispatch(updateUser(null))
            } else {
                const res = await response.json()
                setError(res.Error)
            }

        } catch {
            setError("Falha ao requisitar deslogamento do usuário")
        } finally {
            setIsSigningOut(false)
        }
    
    }

    return { signout, isSigningOut, error } 
}
export default useSignOut