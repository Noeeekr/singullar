import React, { useContext } from 'react';
import { AuthContext } from '../context/authProvider';
import { IUser } from '../types/user'

interface IuseAuthResponse {
    dispatch: React.Dispatch<{
        payload: IUser | null,
        type: string
    }>,
    user: IUser | null,
}
/**
*   Returns user profile and a dispatcher for creation of other hooks
*   that manipulate authentication.
*/
const useAuth = (): IuseAuthResponse => {
    const context = useContext(AuthContext);

    if (!context) {
        throw new Error("useAuth (hook) must be called inside a AuthProvider context")
    }
    if (context.dispatch === null) {
        throw new Error("useAuth (hook) got null instead of a dispatcher")
    }

    return {
        dispatch: context.dispatch as React.Dispatch<{
            payload: IUser | null,
            type: string
        }>,
        user: context.user,
    }
};

export default useAuth;