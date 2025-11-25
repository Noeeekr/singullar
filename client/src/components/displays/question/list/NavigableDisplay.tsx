import QuestionList from "./DefaultDisplay"

import { Fragment, type JSX } from "react"
import type { QuestionListDisplayProps  } from "@components/displays/question/list"
import { Link } from "react-router-dom"

export interface NavigableQuestionListDisplayProps extends QuestionListDisplayProps {
    navigable?: boolean 
}

export default ({ navigable, ...props }: NavigableQuestionListDisplayProps): JSX.Element => {
    return(
        <Fragment>
        {
            navigable 
            ? <Link style={{textDecoration: "none"}} to={String(props.list.id)}>
                <QuestionList navigable {...props} />
            </Link> 
            : <QuestionList {...props} />
        }
        </Fragment>
    )
}