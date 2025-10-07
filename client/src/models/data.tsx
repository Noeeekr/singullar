import { User } from './server';

export interface UserState {
    user: null | User,
    mostVisitedUrls: { [key: string]: number },
}