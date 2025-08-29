import { IUser } from './server';

export interface UserState {
    user: null | IUser,
    mostVisitedUrls: { [key: string]: number },
}