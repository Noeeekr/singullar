import { UserRoles } from './models/server/server';

const studentRoutes = [
    "/home"
]
const teacherRoutes = [
    "/teacher"
]
const adminRoutes = [
    "/admin",
    "/learning"
]
const supervisorRoutes = [
    "/supervisor"
]

/**
*   Holds the urls permitted in important situations
*   
*   Important: the first url of every object should be the default redirect local
*/

type Routes = {
    role: UserRoles | null
    routes: string[]
}

const routes: Routes[] = [
    {
        role: null,
        routes: [
            "/auth"
        ]
    },
    {
        role: "student",
        routes: studentRoutes,
    },
    {
        role: "teacher",
        routes: teacherRoutes,
    },
    {
        role: "supervisor",
        routes: supervisorRoutes,
    },
    {
        role: "admin",
        routes: adminRoutes,
    },
]

export default routes