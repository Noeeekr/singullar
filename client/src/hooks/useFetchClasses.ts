// Features
import useContextAwareFetch, { defaultRequestInit, useContextAwareFetchReturn } from '@hooks/useContextAwareFetch';

// Models
import type { Class, UserSegments } from '@models/server';
import { SERVER_ADDR } from '../configs';

export interface SearchClassFilters {
    id?: number
    class_name?: string
    teacher_name?: string
    student_name?: string
    segment?: UserSegments
    series?: string
    creation_year?: Date
}

export interface useFetchClassesRequest {
    filters: SearchClassFilters[]
    loadStudents: boolean
    loadTeacher: boolean
}

const useFetchClasses = (): useContextAwareFetchReturn<Class[], SearchClassFilters[]> => {
    return useContextAwareFetch<Class[], SearchClassFilters[]>(
        `${SERVER_ADDR}/api/classes`,
        { ...defaultRequestInit, method: "POST" },
    )
}

export default useFetchClasses