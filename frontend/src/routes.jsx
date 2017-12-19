import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom';
import Welcome from './components/Welcome';
import 'styles/index.scss';
import Character from './components/Character'
import Inventory from './components/Inventory'
import Login from './components/Login'
import ImageLoader from 'react-load-image'; 

function Preloader(props) {
  return <img src="spinner.gif" />;
}

//<span className="navbar-text" href="#">RPG-game</span>
const Routes = () => (
  <Router>
    <div>
      <div className='App'>
        <div className="page-header">
          <nav className="navbar navbar-inverse navbar-static-top" role="navigation">
            <div className="container-fluid">
              <div className="navbar-header">
                <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                  <span className="sr-only">Toggle navigation</span>
                  <span className="icon-bar"></span>
                  <span className="icon-bar"></span>
                </button>
                <ImageLoader
                  src="/rpg-game.svg" className="logo-img"
                >
                  <img />
                  <div>Error!</div>
                  <Preloader />
                </ImageLoader>
                
              </div>
              <div className="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                  <ul className="nav navbar-nav">
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/character">Character</Link></li>
                    <li><Link to="/inventory">Inventory</Link></li>
                  </ul>
              </div>
              <Login/>
            </div>
          </nav>
        </div>
      </div>
      <div className="container">
        <Route exact path="/" component={Welcome} />
        <Route path="/character" component={Character} />
        <Route path="/inventory" component={Inventory} />
      </div>
    </div>
  </Router>
);

export default Routes;
