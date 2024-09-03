import { checkAndRedirect } from './redirect';

export const loader = () => {
    return checkAndRedirect("any")
}

export default function RootPage() {

    return(
        <> 
            Redirect Failed
        </>
    )
}