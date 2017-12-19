import React from 'react';
import Login from './components/Login';
import {Link} from 'react-router-dom';
import 'normalize.css';
import 'styles/index.scss';

class App extends React.Component {
  render() {
    return (
      <div className='App'>
        <div className="page-header">
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
        </div>
        <div className="container">
          {this.props.children}
        </div>
      </div>
    );
  }
}

export default App;
