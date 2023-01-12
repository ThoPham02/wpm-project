import React from "react";
import { Link } from "react-router-dom";
import { AiOutlineHome, AiOutlineUser, AiOutlineLogin } from "react-icons/ai";
import { GiQueenCrown } from "react-icons/gi";

function Header() {
  return (
    <header className="header">
      <div className="wapper">
        <div className="header-logo">
          <img src="./logo.png" alt="logo" />
        </div>
        <div className="header-content">
          <h1 className="header-name">Typing Flash</h1>
          <nav className="header-nav">
            <Link className="header-nav-item" to="/home">
              <AiOutlineHome />
              <span>Home</span>
            </Link>
            <Link className="header-nav-item" to="/rank">
              <GiQueenCrown />
              <span>Rank</span>
            </Link>
            <Link className="header-nav-item" to="/user">
              <AiOutlineUser />
              <span>User</span>
            </Link>
            <Link className="header-nav-item" to="/login">
              <AiOutlineLogin />
              <span>Login</span>
            </Link>
          </nav>
        </div>
      </div>
    </header>
  );
}

export default Header;
