import { useContext } from 'react';
import { AuthContext } from '../context/authProvider';
import { IUser } from '../types/user'

interface IuseAuthResponse {
    user: IUser | null,
    setUser: Function,
} 

/**
*   Returns user profile and a dispatcher for creation of other hooks
*   that manipulate authentication.
*/
const useAuth = (): IuseAuthResponse => {
    const [user, setUser] = useContext(AuthContext);

    return {
        user,
        setUser
    }
};

export default useAuth;