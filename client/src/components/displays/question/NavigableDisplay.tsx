import { Fragment } from "react"

// Models
import type { DisplayerProps, DisplayerVariants } from "."
import type { JSX } from "react"
import Display from "./Display"
import { Link } from "react-router-dom"

export type NavigableDisplayProps<IsSelected extends boolean> = DisplayerProps<IsSelected> & DisplayerVariants & {
    navigable?: boolean
}

export default <IsSelected extends boolean>({ navigable, ...props }: NavigableDisplayProps<IsSelected>): JSX.Element => {
    return (
        <Fragment>
            {
                navigable 
                ? <Link style={{ textDecoration: "none" }} to={"/platform/question/" + String(props.question.id)}>
                    <Display navigable {...props} />
                </Link>
                : <Display {...props} />
            }
        </Fragment>
    )
}