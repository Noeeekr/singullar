import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

import PgAuthentication from "./pages/auth/Auth";
import {
  PgInstitutionSelection as PgInstitutionSelection,
  LoaderInstitutionSelection,
  PgMyClasses,
  PgClassesCreate,
  PgStudentsSearch,
  PgStudentsCreate,
} from "./pages/admin";

import Layout from "./components/Layout";
import Root from "./components/DefaultRoot";
import ProtectedRoutes from "@components/ProtectedRoutes";

// This and route provider component might become App.tsx file
// so I can put redux store and things like that there.
/*const router = createBrowserRouter(
  createRoutesFromElements(
    <Routes>
      <Route path="/" element={<ProtectedRoutes />}>
        <Route path="auth" element={<PgAuthentication />} />
        <Route path="home" element={<Layout />}>
          <Route index element={<div>Home root page</div>} />
          <Route path="*" element={<div>Not found specific home</div>} />
        </Route>
        <Route path="teacher" element={<Root />}>
          <Route index element={<div>Index page teach</div>} />
        </Route>
        <Route path="supervisor" element={<div>Layout</div>}>
          <Route index element={<Root />} />
        </Route>
        <Route path="admin" element={<Layout />}>
          <Route index element={<Root />} />
          <Route path="search" element={<PgInstitutionSelection />} loader={InstitutionSelectionLoader} />
          <Route path="classes/create" element={<PgCreate />} />
          <Route path="classes/:id" element={<PgClasses />} />
          <Route path="*" element={<div>Not index page adm</div>} />
        </Route>
        <Route path="*" element={<div>Not found general page</div>} />
      </Route>
    </Routes>
  )
)
  */

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
                element: <PgStudentsSearch />,
              },
              {
                path: "students/create",
                element: <PgStudentsCreate />,
              },
              {
                path: "classes/search",
                element: <PgInstitutionSelection nextUrl="/admin/classes"/>,
                loader: LoaderInstitutionSelection,
              },
              {
                path: "classes/create",
                element: <PgClassesCreate />,
              },
              {
                path: "classes",
                element: <PgMyClasses />,
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
