import { 
    createContext, 
    useEffect, 
    useState,
} from 'react';
import { useLocation } from 'react-router-dom';
import { IUser } from '../types/user';

/** 
* The context used for providing user information utilized in protected routes authorization, login and logout hooks. 
*/
export const AuthContext = createContext<[
    // this type is just a mock because of dispatch null behavior before its first instace is created
    userState: null | IUser,
    setUserState: Function,
]>([
    null,
    () => {},
]);

/** 
* The authorization provider that wraps information for protected routes authorization login and logout
* 
* @remarks  Import UseAuth Hook to use the context provided.
*/
export const AuthContextProvider = ({ children }: { children: JSX.Element}) => {
    const location = useLocation();
    const user = localStorage.getItem('user');

    const [userState, setUserState] = useState(user ? JSON.parse(user) : null);

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        
        setUserState(storedUser ? JSON.parse(storedUser) : null)        
    }, [location.pathname]); 

    useEffect(() => {
        localStorage.setItem('user', JSON.stringify(userState));
    }, [userState]);

    return (
        <AuthContext.Provider value={[ userState, setUserState ]}>
            {children}
        </AuthContext.Provider>
    );
};