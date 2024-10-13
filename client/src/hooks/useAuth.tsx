import {
    useRef,
    useState,
    useEffect,
} from 'react';
import {
    useLocation
} from 'react-router-dom'

import type { AppDispatch } from '../slices/store';
import { useDispatch } from 'react-redux'
import { updateUser } from '../slices/authSlice'

const useAuth = (isAuthRoute: boolean, isPrivateRoute: boolean): {
    isSigned: boolean,
    isLoading: boolean,
} => {
    const url = useLocation().pathname;
    
    const [isLoading, setIsLoading] = useState(true);
    const [isSigned, setIsSigned] = useState(false);

    const dispatch = useDispatch<AppDispatch>()
    // Problem: isLoading state starts as false after a redirect making the
    // actual component appear instead for a brief second before setIsloading(true)

    // Solution: 
    // 0 : There's no redirect that needs isLoading to be true in meanwhile 
    // 1 : A redirect got identified and it is waiting for it to trigger this component again
    // 2 : The redirect triggered the component again and now isLoading will be set to true
    const isLoadingShouldBeTrue = useRef(0)

    if (isLoadingShouldBeTrue.current === 2) {
        setIsLoading(true)
        isLoadingShouldBeTrue.current = 0
    }
    if (isLoadingShouldBeTrue.current === 1) {
        isLoadingShouldBeTrue.current = 2
    }

    useEffect(() => {
        const refreshToken = async () => {
            setIsLoading(true)
            try {
                // Request to refresh end-point. Returns user data if signed
                const response = await fetch("http://localhost:8000/api/user/auth", {
                    credentials: "include"
                })

                const data = await response.json()

                // If there's data refresh user state. Otherwise signout bc
                // there's no user or there was a fail validating user.
                if (response.ok && data.data) {
                    dispatch(updateUser(data.data))
                    setIsSigned(true)
                } else {
                    // THE PROBLEM LIES HERE, WHERE DISPATCH IS RE RENDERING COMPONENT PROBABLY
                    dispatch(updateUser(null))
                    setIsSigned(false)
                }
            } catch (err) {
                // If the request above gives an unauthorized error, tries
                // to signout user..
                const response = await fetch("http://localhost:8000/api/auth/signout", {
                    credentials: "include"
                })

                if (response.ok) {
                    dispatch(updateUser(null))
                    setIsSigned(false)
                }
            } finally {
                if (!isLoading && ((url == "/" && isSigned) || (url == "/" && !isSigned) || (!isSigned && isPrivateRoute) || (isSigned && isAuthRoute))) {
                    isLoadingShouldBeTrue.current = 1;
                }
                console.log("called")
                setIsLoading(false)
            }
        }

        refreshToken()
    }, [url])

    return { isSigned, isLoading }
}

export default useAuth;