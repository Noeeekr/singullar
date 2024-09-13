import { 
    redirect,
} from 'react-router-dom'

import routes from '../routes'

function checkAndRedirect( // this could be two functions in a wrapper
    r: Request,
): Response | null {

    // Public route logic

    const location = r.url
    const routeFormat = /(:?(?:https|http):\/\/\w+(:?(:?(:?\.\w+)+)|(?::\d+)))/g
    
    const currentRoute = location.replace(routeFormat,"")
        
    const isPublicRoute = routes.public.some(
        (publicRoute) => currentRoute == publicRoute
    )
        
    if (isPublicRoute) return null;

    // Authentication route logic

    const isSigned = false
    const isAuthRoute = false
    
    if (isSigned) {
        if (isAuthRoute) {
            return redirect("home")
        }
        return null
    } else {
        return redirect(routes.auth.main)
    }
}

// Might create a response type for checkAndRedirect

export { checkAndRedirect }