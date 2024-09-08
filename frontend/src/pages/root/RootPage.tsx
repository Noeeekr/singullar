import { checkAndRedirect } from '../../components/redirect';

const loader = ({ request }: { request: Request }): Response | null => {    
    return checkAndRedirect(request)
}

const RootPage = (): JSX.Element => {

    return(
        <>
            404 (Not Found)
            This page might not exist or the access to it is not guaranted 
        </>
    )
}

export default RootPage
export { loader }