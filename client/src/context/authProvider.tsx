import { createContext, useReducer, useEffect } from 'react';

export const AuthContext = createContext<any>({ user: null, dispatch: null });

const authReducer = (state: any, action: any) => {
    switch (action.type) {
        case "login":
            return { user: action.payload };
        case "logout":
            return { user: null };
        default:
            return state;
    }
};

export const AuthContextProvider = ({ children }: { children: JSX.Element}) => {
    const [state, dispatch] = useReducer(authReducer, { 
        user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')) : null,
    });

    useEffect(() => {
        localStorage.setItem('user', JSON.stringify(state.user));
    }, [state.user]);

    return (
        <AuthContext.Provider value={{ ...state, dispatch }}>
            {children}
        </AuthContext.Provider>
    );
};