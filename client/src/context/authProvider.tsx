import { createContext, useReducer, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { IUser } from '../types/user';

const authReducer = (state: any, action: { payload?: IUser, type: string }) => {
    switch (action.type) {
        case "login":
            return { user: action.payload };
        case "logout":
            return { user: null };
        default:
            return state;
    }
};
/** 
* The context used for providing user information utilized in protected routes authorization, login and logout hooks. 
*/
export const AuthContext = createContext<{
    // this type is just a mock because of dispatch null behavior before its first instace is created
    user: null | IUser
    dispatch: null | React.Dispatch<{
        payload?: IUser,
        type: string
    }> 
}>({
    user: null,
    dispatch: null,
});

/** 
* The authorization provider that wraps information for protected routes authorization login and logout
* 
* @remarks  Import UseAuth Hook to use the context provided.
*/
export const AuthContextProvider = ({ children }: { children: JSX.Element}) => {
    const location = useLocation();
    const user = localStorage.getItem('user');

    const [state, dispatch] = useReducer(authReducer, { 
        user: user ? JSON.parse(user) : null,
    });

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        
        dispatch({
            type: 'login',
            payload: storedUser ? JSON.parse(storedUser) : null
        });
        
    }, [location.pathname]); 

    useEffect(() => {
        localStorage.setItem('user', JSON.stringify(state.user));
    }, [state.user]);

    return (
        <AuthContext.Provider value={{ ...state, dispatch }}>
            {children}
        </AuthContext.Provider>
    );
};