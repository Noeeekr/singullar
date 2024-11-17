import { IUser } from './server';
import { IInstitutions } from './server';

export interface UserState {
    user: null | IUser,
    mostVisitedUrls: { [key: string]: number },
}

export interface InstitutionsState extends IInstitutions {}