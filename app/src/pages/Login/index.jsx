import React from "react";
import { Link } from "react-router-dom"
import FormLogin from "./components/FormLogin";

function Login() {
  return (
    <div className="login">
      <h2 className="login-title">Login</h2>
      <FormLogin />
      <Link to="/register">Create an account!</Link>
    </div>
  );
}

export default Login;
