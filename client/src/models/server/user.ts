import type { DefaultFields } from ".";

export type UserRoles = typeof ROLE_ADMIN | typeof ROLE_STUDENT | typeof ROLE_TEACHER | typeof ROLE_SUPERVISOR;
export const ROLE_ADMIN = "admin"
export const ROLE_STUDENT = "student"
export const ROLE_TEACHER = "teacher"
export const ROLE_SUPERVISOR = "supervisor"

export type UserSegments = typeof EF1 | typeof EF2 | typeof EM | "";
export const EF1 = "ensino fundamental 1";
export const EF2 = "ensino fundamental 2";
export const EM = "ensino medio";

export interface User extends DefaultFields {
    "email": string,
    "role": UserRoles,
    "name": string,
    "profile_picture": string,
    "institution_id"?: number,
    "segment": UserSegments | null,
}