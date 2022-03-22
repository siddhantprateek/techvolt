
import React, { FC } from "react";
import { LOGO } from "../../static";
import './header.styles.css';
import { Link } from "react-router-dom";
const Header: FC = () => {
    return(
        <header className="header">
            <div className="brand-logo">
                <Link to="/">
                    <img src={LOGO} alt="" />
                </Link>
            </div>
            <div className="nav-option">
                <ul className="nav-bar">
                    <li>
                        <input type="text" placeholder="Search" />
                    </li>
                    <li>
                        <Link to="/">
                            Explore
                        </Link>
                    </li>
                    <li>
                        <Link to="/collections">
                            Collections
                        </Link>
                    </li>
                    <li>
                        <Link to="/stats">
                            Stats
                        </Link>
                    </li>
                </ul>
            </div>
        </header>
    )
}

export default Header;