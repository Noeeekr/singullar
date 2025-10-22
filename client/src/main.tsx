import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from "react-router-dom";

import { lazy, Suspense } from "react";

const ProtectedRoutes = lazy(() => import("./components/layout/ProtectedRoutes"))
const Layout = lazy(() => import("./components/layout/Layout"))
const Home = lazy(() => import("./components/Home"))

const PgStudentsSearch = lazy(() => import("./pages/admin/students/Search"))
const PgStudentsCreate = lazy(() => import("./pages/admin/students/pages/Create"))
const PgClassesSearch = lazy(() => import("./pages/admin/classes/Page"))
const PgClass = lazy(() => import("./pages/admin/classes/class/Page"))
const PgClassesCreate = lazy(() => import("./pages/admin/classes/create/Page"))
const PgTeachersSearch = lazy(() => import("./pages/admin/teachers/Search"))
const PgAuthentication = lazy(() => import("./pages/auth/Auth"))
const PgDashboard = lazy(() => import("./pages/admin/dashboard/Page"))
const PgQuestionList = lazy(() => import("./pages/platform/questions/Page"))
const PgQuestionListCreate = lazy(() => import("./pages/platform/questions/create/Page"))

const routerRoutes = createRoutesFromElements(
  <>
    <Route path="/"
      element={<Suspense fallback={<div>Loading Route Guard</div>}><ProtectedRoutes /></Suspense>}
    >
      <Route path="auth"
        element={<Suspense fallback={<div>Loading Authentication Page</div>}><PgAuthentication /></Suspense>}
        errorElement={<div>Error element auth</div>}
      />
      <Route path="/"
        element={<Suspense fallback={<div>Loading Layout Page</div>}><Layout /></Suspense>}
      >
        <Route index element={<div>Home root page</div>} />
        <Route path="platform">
          <Route path="question">
            <Route path="list">
              <Route index
                element={<Suspense fallback={<div>Loading Questions Page</div>}><PgQuestionList /></Suspense>}
              />
              <Route path="create"
                element={<Suspense fallback={<div>Loading Questions Page</div>}><PgQuestionListCreate /></Suspense>}
              />
              <Route path=":listId"
                element={<Suspense fallback={<div>Loading Questions Page</div>}><PgQuestionListCreate /></Suspense>}
              />
            </Route>
          </Route>
        </Route>
        <Route path="teacher">
          <Route index element={<div>Teacher root page</div>} />
          <Route path="*" element={<div>Teacher path home page</div>} />
        </Route>
        <Route path="home">
          <Route index element={<Home />} />
        </Route>
        <Route path="supervisor">
          <Route index element={<div>Supervisor root page</div>} />
          <Route path="*" element={<div>Supervisor path home page</div>} />
        </Route>
        <Route path="admin">
          <Route path="dashboard"
            element={<Suspense fallback={<div>Loading Dashboard Page</div>}><PgDashboard /></Suspense>}
          />
          <Route path="students">
            <Route index
              element={<Suspense fallback={<div>Loading Students</div>}><PgStudentsSearch /></Suspense>}
            />
            <Route path="create"
              element={<Suspense fallback={<div>Loading Student Create</div>}><PgStudentsCreate /></Suspense>}
            />
          </Route>
          <Route path="classes">
            <Route index element={<Suspense fallback={<div>Loading Classes</div>}><PgClassesSearch /></Suspense>} />
            <Route path=":id" element={<Suspense fallback={<div>Loading Class Info</div>}><PgClass /></Suspense>} />
            <Route path="create"
              element={<Suspense fallback={<div>Loading Classes Create</div>}><PgClassesCreate /></Suspense>}
            />
          </Route>
          <Route path="teachers">
            <Route index
              element={<Suspense fallback={<div>Loading Teachers</div>}><PgTeachersSearch /></Suspense>}
            />
          </Route>
          <Route path="*" element={<div>Any path admin page</div>} />
        </Route>
        <Route path="*" element={<div>Any route default page</div>} />
      </Route>
    </Route>
    <Route path="*" element={<div>Any path home page</div>} />
  </>
);
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <ThemeProvider theme={lightTheme}>
          <CssBaseline />
          <GlobalStyles styles={{ ul: { padding: 0, margin: 0 } }} />
          <RouterProvider router={createBrowserRouter(routerRoutes)} />
        </ThemeProvider>
      </PersistGate>
    </Provider>
  </StrictMode>
);
