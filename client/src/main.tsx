import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

import PgAuthentication from "./pages/auth/Auth";

import Layout from "./components/Layout";
import Root from "./components/DefaultRoot";
import ProtectedRoutes from "@components/ProtectedRoutes";

import { lazy, Suspense } from "react";

const PgStudentsSearch = lazy(() => import("./pages/admin/students/Search")) 
const PgStudentsCreate = lazy(() => import("./pages/admin/students/pages/Create")) 
const PgMyClasses = lazy(() => import("./pages/admin/classes/pages/Create")) 
const PgClassesCreate = lazy(() => import("./pages/admin/classes/pages/MyClasses")) 

const r2 = createBrowserRouter([
  {
    path: "/",
    element: <ProtectedRoutes />,
    children: [
      {
        path: "/auth",
        element: <PgAuthentication />,
        errorElement: <div>Error element auth</div>,
      },
      {
        path: "/",
        element: <Layout />,
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
                element: <Suspense fallback={<div>Loading...</div>}><PgStudentsSearch /></Suspense>,
              },
              {
                path: "students/create",
                element: <Suspense fallback={<div>Loading...</div>}><PgStudentsCreate /></Suspense>,
              },
              {
                path: "classes/create",
                element: <Suspense fallback={<div>Loading...</div>}><PgClassesCreate /></Suspense>,
              },
              {
                path: "classes",
                element: <Suspense fallback={<div>Loading...</div>}><PgMyClasses /></Suspense>,
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
