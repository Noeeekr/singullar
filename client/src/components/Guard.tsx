// Utilities
import { useAppSelector } from "@slices/store"

// Models
import type { BoxProps } from "@mui/material/Box"
import type { UserRoles } from "@models/server"
import type { JSX } from "react"

export interface GuardProps extends BoxProps {
    roles: UserRoles[] | UserRoles
    children: JSX.Element
}

// Guard shows the component only when the user role matches the specified role
export default ({ roles, children }: GuardProps) => {
    const targetRole = useAppSelector(store => store.user.user?.role)
    if (typeof roles == "string") return roles == targetRole ? children : <></>  
    if (roles.some(role => role == targetRole)) return children
    return <></>
}