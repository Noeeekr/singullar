import { IUser } from './user';

export interface UserState {
    user: null | IUser,
    mostVisitedUrls: { [key: string]: number },
}