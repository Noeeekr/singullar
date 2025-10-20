// Non-specific models 
export interface IdentificationField {
    "id": number
}
export interface DefaultFields extends IdentificationField {
    created_at: string,
    deleted_at: string | null,
    updated_at: string,
}

// User related models
export type { 
    User, 
    UserRoles, 
    UserSegments,
} from "./user"
export {
    ROLE_ADMIN,
    ROLE_STUDENT,
    ROLE_SUPERVISOR,
    ROLE_TEACHER, 
    EF1, 
    EF2, 
    EM 
} from "./user"
// Class related models
export type { Class } from "./class"

// Institution related models
export type { Institution } from "./institution"

// Response related models
export type { DefaultResponse } from "./response"
