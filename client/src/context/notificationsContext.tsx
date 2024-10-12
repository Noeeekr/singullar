import {
    createContext,
} from 'react'

export const NotificationContext = createContext<{
    hasNotifications: boolean,
    setHasNotifications: Function,
}>({
    hasNotifications: false,
    setHasNotifications: () => { return },
})