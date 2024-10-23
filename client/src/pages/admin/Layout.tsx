import { Outlet } from 'react-router-dom'

const Layout = () => {

    return(
        <div>
            Layout page
            <Outlet/>
        </div>
    )
}

export default Layout;