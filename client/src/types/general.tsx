import { IUser } from './user';

export interface UserState {
    user: null | IUser,
    mostVisitedUrls: { [key: string]: number },
}

export interface InstitutionsState {
    id: number,
    created_at: Date,
    deleted_at: Date | null,
    updated_at: Date,
    name: string,
}