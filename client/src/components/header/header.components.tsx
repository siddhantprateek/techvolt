
import React, { FC } from "react";
import './header.styles.css';

const Header: FC = () => {
    return(
        <header>
            <ul className="nav-bar">
                <li>Explore</li>
                <li>Collections</li>
                <li>Stats</li>
            </ul>
        </header>
    )
}

export default Header;