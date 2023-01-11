import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Flagment } from "react";
import DefaultLayout from "./components/layout";

import "./main.css";
import { privateRoute, publicRoute } from "./routes";

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          {publicRoute.map((route, index) => {
            let Layout = DefaultLayout;

            if (route.layout) {
              Layout = route.layout;
            }
            if (route.layout === null) {
              Layout = Flagment;
            }

            const Page = <Layout>{route.element}</Layout>;

            return <Route key={index} path={route.path} element={<Page />} />;
          })}
          {privateRoute.map((route, index) => {
            let Layout = DefaultLayout;

            if (route.layout) {
              Layout = route.layout;
            }
            if (route.layout === null) {
              Layout = Flagment;
            }

            const Page = <Layout>{route.element}</Layout>;

            return <Route key={index} path={route.path} element={<Page />} />;
          })}
        </Routes>
      </Router>
    </div>
  );
}

export default App;
