// Components
import Display from "./Display"
import { Link } from "react-router-dom"
import { Fragment } from "react"

// Models
import type { JSX } from "react"
import type { QuestionListDisplayProps  } from "@components/displays/question/list"

export interface NavigableQuestionListDisplayProps extends QuestionListDisplayProps {
    navigable?: boolean 
}

export default ({ navigable, children, ...props }: NavigableQuestionListDisplayProps): JSX.Element => {
    return(
        <Fragment>
        {
            navigable 
            ? <Link style={{textDecoration: "none"}} to={String(props.list.id)}>
                <Display navigable {...props} />
            </Link> 
            : <Display {...props} />
        }
        </Fragment>
    )
}