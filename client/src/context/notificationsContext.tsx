import {
    createContext,
} from 'react'

export const NotificationContext = createContext<{
    notifications: {
        notifications: object[],
        help: object[],
    },
}>({
    notifications: {
        notifications: [{}],
        help: [{}],
    },
})