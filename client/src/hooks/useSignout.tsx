import { useState } from 'react'
import useChangeUserState from './useChangeUserState';

const useSignOut = () => {
    const { setUser } = useChangeUserState();

    const [isSigningOut, setIsSigningOut] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const signout = async () => {
        setIsSigningOut(true)

        try {
            const response = await fetch("http://localhost:8000/api/auth/signout", {
                credentials: "include"
            })
            
            if (response.ok) {
                console.log("dispatched logout")
                setUser(null)
            } else {
                let res = await response.json()
                setError(res.Error)
            }

        } catch (err: unknown) {
            setError("Falha ao requisitar deslogamento do usuário")
        } finally {
            setIsSigningOut(false)
        }
    
    }

    return { signout, isSigningOut, error } 
}
export default useSignOut