export interface IUser {
    id: number,
    name: string,
    email: string,
    surname: string,
    createdAt: Date,
    profileImg: string
}

export interface IUserRequest {
    data: null | IUser,
    error: null | string,
}