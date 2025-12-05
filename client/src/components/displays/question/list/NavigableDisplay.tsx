// Components
import Display, { QuestionListDisplayVariants } from "./Display"
import { Link } from "react-router-dom"
import { Fragment } from "react"

// Models
import type { JSX } from "react"
import type { QuestionListDisplayProps } from "@components/displays/question/list"

export interface NavigableQuestionListDisplayProps extends QuestionListDisplayProps, QuestionListDisplayVariants {
    navigable?: boolean
}

const LinkWrapper = ({ to, children }: { to: string, children: JSX.Element }): JSX.Element => (
    <Link style={{ textDecoration: "none" }} to={to}>
        {children}
    </Link>
)

export default ({ navigable, children, ...props }: NavigableQuestionListDisplayProps): JSX.Element => {
    return (
        <Fragment>
            {
                navigable
                    ? <LinkWrapper to={String(props.list.id)}>
                        <Display navigable {...props} />
                    </LinkWrapper>
                    : <Display {...props} />
            }
        </Fragment>
    )
}