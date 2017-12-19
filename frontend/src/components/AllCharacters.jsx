import React from 'react';
import * as utils from '../functions.js'
import { confirmAlert } from 'react-confirm-alert'; // Import
import 'react-confirm-alert/src/react-confirm-alert.css'

class AllCharacters extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
          characters: []
      };
      this.getCharacters = this.getCharacters.bind(this);
      this.deleteCharacter = this.deleteCharacter.bind(this);
      this.getCharacters()
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
    submit = (charName) => {
        confirmAlert({
          title: 'Confirm to submit',                        // Title dialog
          message: 'Are you sure to do delete '+charName+'?',               // Message dialog
          childrenElement: () => {},       // Custom UI or Component
          confirmLabel: 'Confirm',                           // Text button confirm
          cancelLabel: 'Cancel',                             // Text button cancel
          onConfirm: () => this.deleteCharacter(charName),    // Action after Confirm
          onCancel: () => {},      // Action after Cancel
        })
      };
    render() {
        return (
            <ul>
            {
                this.state.characters.map(function(char, index){
                    return <li key={index}>{char.name} {char.title} {char.level} {char.experience} <button id={index} onClick={() => this.submit(char.name)}>delete</button></li>;
                }, this)
            }
            </ul>
        )
    }
}

export default AllCharacters