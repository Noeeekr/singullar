import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

import { lazy, Suspense } from "react";

const Layout = lazy(() => import("./components/Layout")) 
const Root = lazy(() => import("./components/DefaultRoot")) 
const ProtectedRoutes = lazy(() => import("./components/ProtectedRoutes"))
const PgStudentsSearch = lazy(() => import("./pages/admin/students/Search")) 
const PgStudentsCreate = lazy(() => import("./pages/admin/students/pages/Create")) 
const PgClassesSearch = lazy(() => import("./pages/admin/classes/Page")) 
const PgClassesCreate = lazy(() => import("./pages/admin/classes/create/Page")) 
const PgTeachersSearch = lazy(() => import("./pages/admin/teachers/Search"))
const PgAuthentication = lazy(() => import("./pages/auth/Auth"))

const r2 = createBrowserRouter([
  {
    path: "/",
    element: <Suspense fallback={<div>Loading Route Guard</div>}><ProtectedRoutes /></Suspense>,
    children: [
      { 
        path: "/auth",
        element: <Suspense fallback={<div>Loading Authentication Page</div>}><PgAuthentication /></Suspense>,
        errorElement: <div>Error element auth</div>,
      },
      {
        path: "/",
        element: <Suspense fallback={<div>Loading Layout Page</div>}><Layout /></Suspense>,
        children: [
          {
            index: true,
            element: <div>Home root page</div>,
          },
          {
            path: "teacher",
            children: [
              {
                index: true,
                element: <div>Teacher root page</div>,
              },
              {
                path: "*",
                element: <div>Teacher path home page</div>,
              },
            ],
          },
          {
            path: "supervisor",
            children: [
              {
                index: true,
                element: <div>Supervisor root page</div>,
              },
              {
                path: "*",
                element: <div>Supervisor path home page</div>,
              },
            ],
          },
          {
            path: "admin",
            children: [
              {
                index: true,
                element: <Root/>,
              },
              {
                path: "students",
                element: <Suspense fallback={<div>Loading Students</div>}><PgStudentsSearch /></Suspense>,
              },
              {
                path: "students/create",
                element: <Suspense fallback={<div>Loading Student Create</div>}><PgStudentsCreate /></Suspense>,
              },
              {
                path: "classes",
                element: <Suspense fallback={<div>Loading Classes</div>}><PgClassesSearch /></Suspense>,
              },
              {
                path: "classes/create",
                element: <Suspense fallback={<div>Loading Classes Create</div>}><PgClassesCreate /></Suspense>,
              },
              {
                path: "teachers",
                element: <Suspense fallback={<div>Loading Teachers</div>}><PgTeachersSearch /></Suspense>,
              },
              {
                path: "*",
                element: <div>Any path admin page</div>,
              },
            ],
          },
          {
            path: "*",
            element: <div>Any route default page</div>,
          },
        ],
      },
    ],
  },
  {
    path: "*",
    element: <div>Any path home page</div>,
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <ThemeProvider theme={lightTheme}>
          <CssBaseline />
          <GlobalStyles styles={{ ul: { padding: 0, margin: 0 } }} />
            <RouterProvider router={r2} />
        </ThemeProvider>
      </PersistGate>
    </Provider>
  </StrictMode>
);
