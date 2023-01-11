import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Flagment } from "react";

import "./main.css";
import 'bootstrap/dist/css/bootstrap.min.css';
import DefaultLayout from "./components/layout";
import { privateRoute, publicRoute } from "./routes";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {publicRoute.map((route, index) => {
          let Layout = DefaultLayout;

          if (route.layout) {
            Layout = route.layout;
          }
          if (route.layout === null) {
            Layout = Flagment;
          }

          const Page = route.element;

          return (
            <Route
              key={index}
              path={route.path}
              element={
                <Layout>
                  <Page />
                </Layout>
              }
            />
          );
        })}
        {privateRoute.map((route, index) => {
          let Layout = DefaultLayout;

          if (route.layout) {
            Layout = route.layout;
          }
          if (route.layout === null) {
            Layout = Flagment;
          }

          const Page = route.element;

          return (
            <Route
              key={index}
              path={route.path}
              element={
                <Layout>
                  <Page />
                </Layout>
              }
            />
          );
        })}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
