// Utilities
import { SERVER_ADDR } from "../configs"
import useContextAwareFetch, { defaultRequestInit } from "./useContextAwareFetch"

// Models
import type { Question } from "@models/server"
import type { useContextAwareFetchReturn } from "./useContextAwareFetch"
import { QuestionFilters } from "../pages/platform/question/list/components/Search/Filters"

export interface FilterOptions {
    question_list_ids?: number[]
    fields?: QuestionFilters[]
    offset?: number
}

const DEFAULT_BODY: FilterOptions = {
    question_list_ids: [],
    fields: [],
    offset: 0,
}

export default (): useContextAwareFetchReturn<Question[], FilterOptions> => {
    let values = useContextAwareFetch<Question[], FilterOptions>(
        `${SERVER_ADDR}/api/question`,
        { ...defaultRequestInit, method: "POST" }
    )

    const defaultSend = values.send

    values.send = (body?: FilterOptions) => {
        if (body == undefined) body = DEFAULT_BODY
        defaultSend(body)
    }
    
    return values
}