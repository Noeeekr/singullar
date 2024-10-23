import {
    useRef,
    useState,
    useEffect,
} from 'react';
import {
    useLocation
} from 'react-router-dom'

import type { IUser } from '../types/user'
import type { AppDispatch } from '../slices/store';
import { useDispatch } from 'react-redux'
import { updateUser } from '../slices/authSlice'

const useAuth = (): {
    isSigned: boolean,
    isLoading: boolean,
    user: IUser | null,
} => {
    const url = useLocation().pathname;
    const lastUrl = useRef(url);

    const [isLoading, setIsLoading] = useState(true);
    const [isSigned, setIsSigned] = useState(false);
    const [user, setUser] = useState<null | IUser>(null);

    // Prevents isLoading from "flicking"
    //
    // Explanation: State starts true and then becomes false in the end of first use
    // so content can be loaded, but in the next use it'll start as false and then become
    // true in the next second because of useEffect, making it "flick"
    if (lastUrl.current !== url) {
        lastUrl.current = url;
        setIsLoading(true)
    }
    
    const dispatch = useDispatch<AppDispatch>()

    useEffect(() => {
        const refreshToken = async () => {
            setIsLoading(true)
            try {
                // Request to refresh end-point. Returns user data if signed
                const response = await fetch("http://localhost:8000/api/user/auth", {
                    credentials: "include"
                })

                const data = await response.json()

                // If there's data refresh user state otherwise signs out user 
                // because we could verify its authenticity
                if (response.ok && data.data) {
                    dispatch(updateUser(data.data))
                    setIsSigned(true)
                    setUser(data.data)
                } else {
                    dispatch(updateUser(null))
                    setIsSigned(false)
                    setUser(null)
                }
            } catch (err) {
                // Tries to signout user..
                const response = await fetch("http://localhost:8000/api/auth/signout", {
                    credentials: "include"
                })

                if (response.ok) {
                    dispatch(updateUser(null))
                    setIsSigned(false)
                    setUser(null)
                }
            } finally {
                setIsLoading(false)
            }
        }

        refreshToken()
    }, [url])

    return { isSigned, isLoading, user }
}

export default useAuth;