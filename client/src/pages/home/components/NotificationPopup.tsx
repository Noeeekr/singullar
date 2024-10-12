import {
    useContext
} from 'react'
import {
    NotificationContext
} from '../../../context/notificationsContext'

const Notification = (): JSX.Element => {

    return(
        <div>
            Divit Euseter Loreau Chirden
        </div>
    )
}
const NotificationPopup = (): JSX.Element => {
    const { setHasNotifications } = useContext(NotificationContext)
    
    const fetchNotifications = Boolean(10)
    setHasNotifications(fetchNotifications)
    
    return(
        <div>
            Divit Lor
        </div>
    )
}
export default NotificationPopup