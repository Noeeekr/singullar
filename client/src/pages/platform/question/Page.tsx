// Components
import { Fragment } from "react"
import Student from "./Student"
import Admin from "./Admin"

// Utilities
import { useAppSelector } from "@slices/store"
import { UserState } from "@models/data"
import { ROLE_ADMIN, ROLE_STUDENT, ROLE_SUPERVISOR, ROLE_TEACHER } from "@models/server"

// Models
import { type JSX } from "react"

export default (): JSX.Element => {
    const userdata: UserState = useAppSelector((store) => store.user)
    const Page: JSX.Element = (() => {
        switch (userdata.user?.role) {
            case ROLE_ADMIN || ROLE_SUPERVISOR || ROLE_TEACHER:
                return <Admin />
            case ROLE_STUDENT:
                return <Student />
            default:
                return <></>
        }
    })()
    return (
        <Fragment>
            {Page}
        </Fragment>
    )
}