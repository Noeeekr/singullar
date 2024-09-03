import { redirect } from 'react-router-dom'

function checkAndRedirect(when: string): Response | null {

    switch(when) {
        case "any": 
        return redirect("login")
        default:
        return null
    }
}

export { checkAndRedirect }