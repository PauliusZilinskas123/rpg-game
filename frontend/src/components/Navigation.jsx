import React from 'react';
import {Link} from 'react-router-dom';
import Login from './Login.jsx'

class Navigation extends React.Component {
  render() {
    return (
    <nav className="navbar navbar-inverse">
      <div className="container-fluid">
        <div className="navbar-header">
          <a className="navbar-brand" href="#">RPG-game</a>
        </div>
        <ul className="nav navbar-nav">
          <li><Link to="/">Home</Link></li>
          <li><Link to="/characters">Characters</Link></li>
          <li><Link to="/inventory">Inventory</Link></li>
        </ul>
        <Login/>
      </div>
    </nav>
    )
  }
}

export default Navigation;
