// Utilities
import { SERVER_ADDR } from "../configs"
import useContextAwareFetch, { defaultRequestInit, useContextAwareFetchReturn } from "./useContextAwareFetch"

// Models
import type { QuestionList } from "@models/server/question"

export interface FilterOptions {
    filters: {
        fields?: {
            Title: undefined | string
            SubjectId: undefined | number
            DifficultyLevel: undefined | number
        }[]
        ids?: number[]
    }
    offset: number
}

const defaultBody: FilterOptions = {
    filters: {},
    offset: 0,
}

export default (): useContextAwareFetchReturn<QuestionList[], FilterOptions> => {
    let values = useContextAwareFetch<QuestionList[], FilterOptions>(
        `${SERVER_ADDR}/api/question/list`,
        { ...defaultRequestInit, method: "POST" }
    )

    const defaultSend = values.send

    values.send = (body?: FilterOptions) => {
        if (body == undefined) body = defaultBody
        defaultSend(body)
    }
    return values
}