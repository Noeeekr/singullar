import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles, Box } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from "react-router-dom";

import { lazy } from "react";

import Suspense from "./components/layout/suspense"

const Home = lazy(() => import("./components/Home"))
const Layout = lazy(() => import("./components/layout/Layout"))

const RouteGuard = lazy(() => import("./components/layout/RouteGuard"))
const PgAuthentication = lazy(() => import("./pages/auth/Auth"))

const PgDashboard = lazy(() => import("./pages/admin/dashboard/Page"))

const PgStudentsSearch = lazy(() => import("./pages/admin/students/Search"))
const PgStudentsCreate = lazy(() => import("./pages/admin/students/pages/Create"))

const PgClass = lazy(() => import("./pages/admin/classes/class/Page"))
const PgClassesSearch = lazy(() => import("./pages/admin/classes/Page"))
const PgClassesCreate = lazy(() => import("./pages/admin/classes/create/Page"))

const PgSubjectCreate = lazy(() => import("./pages/platform/subjects/create/Page"))

const PgTeachersSearch = lazy(() => import("./pages/admin/teachers/Page"))

const PgQuestion = lazy(() => import("./pages/platform/question/Page"))
const PgQuestionCreate = lazy(() => import("./pages/platform/question/create/Page"))
const PgQuestionList = lazy(() => import("./pages/platform/question/list/Page"))
const PgQuestionListSearch = lazy(() => import("./pages/platform/question/list/search/Page"))
const PgQuestionListCreate = lazy(() => import("./pages/platform/question/list/create/Page"))

const routerRoutes = createRoutesFromElements(
  <>
    <Route path="/"
      element={<Box sx={{ width: "100vw", height: "100vh" }}><Suspense to={<RouteGuard />} /></Box>}
    >
      <Route path="auth"
        element={<Suspense to={<PgAuthentication />} />}
        errorElement={<div>Error element auth</div>}
      />
      <Route path="/"
        element={<Suspense to={<Layout />} />}
      >
        <Route index element={<div>Home root page</div>} />
        <Route path="platform">
          <Route path="subjects">
            <Route path="create"
              element={<Suspense children={<PgSubjectCreate />} />} />
          </Route>
          <Route path="question">
            <Route path="create"
              element={<Suspense to={<PgQuestionCreate />} />}
            />
            <Route path=":question_id"
              element={<Suspense to={<PgQuestion />} />}
            />
            <Route path="list">
              <Route index
                element={<Suspense to={<PgQuestionListSearch />} />}
              />
              <Route path=":list_id"
                element={<Suspense to={<PgQuestionList />} />}
              />
              <Route path="create"
                element={<Suspense to={<PgQuestionListCreate />} />}
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
            element={<Suspense to={<PgDashboard />} />}
          />
          <Route path="students">
            <Route index
              element={<Suspense to={<PgStudentsSearch />} />}
            />
            <Route path="create"
              element={<Suspense to={<PgStudentsCreate />} />}
            />
          </Route>
          <Route path="classes">
            <Route index element={<Suspense to={<PgClassesSearch />} />} />
            <Route path=":id" element={<Suspense to={<PgClass />} />} />
            <Route path="create"
              element={<Suspense to={<PgClassesCreate />} />}
            />
          </Route>
          <Route path="teachers">
            <Route index
              element={<Suspense to={<PgTeachersSearch />} />}
            />
          </Route>
          <Route path="*" element={<div>Any path admin page</div>} />
        </Route>
        <Route path="*" element={<div>Em construção</div>} />
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
