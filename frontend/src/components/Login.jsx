import React from 'react';
import * as utils from '../functions.js'

class Login extends React.Component {
    constructor(props) {
      super(props); 
      this.state = {
        error: null,
        isLoaded: false,
        loginURL: "",
        user: ""
      };
      this.logout = this.logout.bind(this);
    }
    logout() {
        utils.deleteUser();
        utils.deleteState();
        fetch("http://127.0.0.1:8080/getlogin")
        .then(res => res.json())
        .then(
          (result) => {
            utils.setState(result.state);
            this.setState({
              isLoaded: true,
              loginURL: result.url,
              user: null
            });
          },
          (error) => {
            console.log("error api");
            this.setState({
              error: error,
              isLoaded: true,
              loginURL: null,
              user: null
            });
          }
        )
    }
    componentDidMount() {
      if(utils.getState() != null) {
        fetch("http://127.0.0.1:8080/getlogin?state="+encodeURIComponent(utils.getState()))
        .then(res => res.json())
        .then(
          (result) => {
            if(result.user != "") {
                utils.setUser(result.user);
            }
            this.setState({
              isLoaded: true,
              loginURL: result.url,
              user: result.user
            });
          },
          (error) => {
            console.log("error api");
            this.setState({
              error: error,
              isLoaded: true,
              loginURL: null,
              user: null
            });
          }
        )
      } else {
        fetch("http://127.0.0.1:8080/getlogin")
        .then(res => res.json())
        .then(
          (result) => {
            utils.setState(result.state);
            this.setState({
              isLoaded: true,
              loginURL: result.url,
              user: null
            });
          },
          (error) => {
            console.log("error api");
            this.setState({
              error: error,
              isLoaded: true,
              loginURL: null,
              user: null
            });
          }
        )
      }
    }
    render() {
        const { error, isLoaded, loginURL, user } = this.state;
        if(utils.getUser() == null) {
            return (
                <ul className="nav navbar-nav navbar-right">
                    <li><a href={loginURL}><span className="glyphicon glyphicon-log-in"></span> Login</a></li>
                </ul>
            )
        } else {
            return (
                <ul className="nav navbar-nav navbar-right">
                    <li><span className="navbar-text">Logged in as {user}</span></li>
                    <li><a href="#" onClick={this.logout}><span className="glyphicon glyphicon-log-in"></span> Logout</a></li>
                </ul>
            )
        }
    }
}

export default Login;