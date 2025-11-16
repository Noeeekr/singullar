import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { PersistGate } from "redux-persist/integration/react";
import { Provider } from "react-redux";
import { store, persistor } from "./slices/store";

import { ThemeProvider, CssBaseline, GlobalStyles, Box } from "@mui/material";
import { lightTheme } from "./themes";

import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from "react-router-dom";

import { lazy } from "react";

import DefaultSuspense from "./components/layout/DefaultSuspense"

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

const PgTeachersSearch = lazy(() => import("./pages/admin/teachers/Search"))

const PgQuestionCreate = lazy(() => import("./pages/platform/question/create/Page"))
const PgQuestionList = lazy(() => import("./pages/platform/question/list/Page"))
const PgQuestionListCreate = lazy(() => import("./pages/platform/question/list/create/Page"))

const routerRoutes = createRoutesFromElements(
  <>
    <Route path="/"
      element={<Box sx={{ width: "100vw", height: "100vh" }}><DefaultSuspense><RouteGuard/></DefaultSuspense></Box>}
    >
      <Route path="auth"
        element={<DefaultSuspense><PgAuthentication /></DefaultSuspense>}
        errorElement={<div>Error element auth</div>}
      />
      <Route path="/"
        element={<DefaultSuspense><Layout /></DefaultSuspense>}
      >
        <Route index element={<div>Home root page</div>} />
        <Route path="platform">
          <Route path="question">
            <Route path="create"
              element={<DefaultSuspense><PgQuestionCreate /></DefaultSuspense>}
            />
            <Route path="list">
              <Route index
                element={<DefaultSuspense><PgQuestionList /></DefaultSuspense>}
              />
              <Route path="create"
                element={<DefaultSuspense><PgQuestionListCreate /></DefaultSuspense>}
              />
              <Route path=":listId"
                element={<DefaultSuspense><PgQuestionListCreate /></DefaultSuspense>}
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
            element={<DefaultSuspense><PgDashboard /></DefaultSuspense>}
          />
          <Route path="students">
            <Route index
              element={<DefaultSuspense><PgStudentsSearch /></DefaultSuspense>}
            />
            <Route path="create"
              element={<DefaultSuspense><PgStudentsCreate /></DefaultSuspense>}
            />
          </Route>
          <Route path="classes">
            <Route index element={<DefaultSuspense><PgClassesSearch /></DefaultSuspense>} />
            <Route path=":id" element={<DefaultSuspense><PgClass /></DefaultSuspense>} />
            <Route path="create"
              element={<DefaultSuspense><PgClassesCreate /></DefaultSuspense>}
            />
          </Route>
          <Route path="teachers">
            <Route index
              element={<DefaultSuspense><PgTeachersSearch /></DefaultSuspense>}
            />
          </Route>
          <Route path="*" element={<div>Any path admin page</div>} />
        </Route>
        <Route path="*" element={<div>Rot</div>} />
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
