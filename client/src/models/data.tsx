import type { User } from './server/user';

export interface UserState {
    user: null | User,
    mostVisitedUrls: { [key: string]: number },
}