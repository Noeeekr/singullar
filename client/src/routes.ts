const studentRoutes = [
    "/home"
]
const teacherRoutes = [
    "/teacher"
]
const adminRoutes = [
    "/admin"
]
const supervisorRoutes = [
    "/supervisor"
]

/**
*   Holds the urls permitted in important situations
*   
*   Important: the first url of every object should be the default redirect local
*/
const routes = {
    auth: [
        "/auth",
    ],
    private: [
        ...studentRoutes,
        ...teacherRoutes,
        ...supervisorRoutes,
        ...adminRoutes,
    ],
    student: studentRoutes,
    teacher: teacherRoutes,
    supervisor: supervisorRoutes,
    admin: adminRoutes,
}

export default routes