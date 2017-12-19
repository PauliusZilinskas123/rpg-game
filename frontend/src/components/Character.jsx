import React from 'react';
import * as utils from '../functions.js'
import { confirmAlert } from 'react-confirm-alert'; // Import
import 'react-confirm-alert/src/react-confirm-alert.css'
import {
    BrowserRouter as Router,
    Route,
    Link
  } from 'react-router-dom';
  import AllCharacters from './AllCharacters'
  import NewCharacter from './NewCharacter'

class Character extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
          characters: []
      };
      this.getCharacter = this.getCharacter.bind(this);
      this.getCharacters = this.getCharacters.bind(this);
      this.updateCharacter = this.updateCharacter.bind(this);
      this.deleteCharacter = this.deleteCharacter.bind(this);
      this.charactersToList = this.charactersToList.bind(this);
    }
    deleteCharacter(char) {
        fetch("http://127.0.0.1:8080/characters/"+encodeURIComponent(char)+"?state="+encodeURIComponent(utils.getState()), {method: 'DELETE'})
        .then(
            (result) => {
                fetch("http://127.0.0.1:8080/characters?state="+encodeURIComponent(utils.getState()))
                .then(res => res.json())
                .then(
                    (result) => {
                        if(result.body != null) {
                            this.setState({
                                characters: result.body
                              });
                        } else {
                            this.setState({
                                characters: []
                            })
                        }
                    },
                    (error) => {
                      console.log("api error");
                      this.setState({
                          characters: []
                      })
                    }
                  )
            },
            (error) => {
              console.log("api error on delete operation "+error);
              this.setState({
                  characters: []
              })
            }
          )
    }
    getCharacter() {
        return "";
    }
    getCharacters() {
        fetch("http://127.0.0.1:8080/characters?state="+encodeURIComponent(utils.getState()))
        .then(res => res.json())
        .then(
            (result) => {
                if(result.body != null) {
                    this.setState({
                        characters: result.body
                      });
                } else {
                    this.setState({
                        characters: []
                    })
                }
            },
            (error) => {
              console.log("api error");
              this.setState({
                  characters: []
              })
            }
          )
    }
    updateCharacter(character) {
        return "";
    }
    charactersToList(characters) {
        return characters.array.forEach(element => {
            <li>{element.name} {element.title}</li>
        });
    }
    render() {
      return (
        <Router>
            <div className="character">
                <ul>
                    <li><Link to="/character">All Characters</Link></li>
                    <li><Link to="/character/newcharacter">New Character</Link></li>
                </ul>
                <div className="container">
                    <Route exact path="/character" component={AllCharacters} />
                    <Route path="/character/newcharacter" component={NewCharacter} />
                </div>
            </div>
        </Router>
        )
    }
  }

export default Character