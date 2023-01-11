import React from "react";
import Footer from "./Footer";
import Header from "./Header";

function DefaultLayout({ chilren }) {
  return (
    <div>
      <Header />
      <div className="content">{chilren}</div>
      <Footer />
    </div>
  );
}

export default DefaultLayout;
