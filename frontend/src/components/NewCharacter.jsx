import React from 'react';
import * as utils from '../functions.js'
import { confirmAlert } from 'react-confirm-alert'; // Import
import 'react-confirm-alert/src/react-confirm-alert.css'

class NewCharacter extends React.Component {
    constructor(props) {
      super(props);
      this.createCharacter = this.createCharacter.bind(this);
    }

    createCharacter(event) {
        event.preventDefault();
        
        var formdata = {
            name : document.getElementById("name").value,
            title : document.getElementById("title").value,
            attributes : [document.getElementById("attributes").value],
            level : document.getElementById("level").value,
            experience : document.getElementById("experience").value,
        }
        console.log(formdata);
        var payload = {
            a: 1,
            b: 2
        };
        
        console.log(JSON.stringify( formdata ))
        fetch("http://127.0.0.1:8080/characters?state="+utils.getState(), {
            method: 'post',
            body: JSON.stringify( formdata ),
        })
        .then((
            (result) => {
                if(result != null) {
                    console.log(result)
                } else {
                    console.log("null result")
                }
            },
            (error) => {
              console.log("api error");
            }
          ));
    }

    /*type Character struct {
	UserId		string 		`json:"user-id"`
	Name    	string 		`json:"name"`
	Title 		string 		`json:"title"`
	Attributes 	[]string 	`json:"attributes"`
	Level   	int 		`json:"level"`
	Experience	int 		`json:"experience"`
}*/

    render() {
        return (
            <form onSubmit={this.createCharacter} id="mainForm" method="post">
                <div className="form-group">
                    <label htmlFor="name">Name</label>
                    <input type="text" className="form-control" id="name" name="name"/>
                </div>
                <div className="form-group">
                    <label htmlFor="title">Title</label>
                    <input type="text" className="form-control" id="title" name="title"/>
                </div>
                <div className="form-group">
                    <label htmlFor="attributes">Attributes</label>
                    <input type="text" className="form-control" id="attributes" name="attributes"/>
                </div>
                <div className="form-group">
                    <label htmlFor="level">Level</label>
                    <input type="text" className="form-control" id="level" name="level"/>
                </div>
                <div className="form-group">
                    <label htmlFor="experience">Experience</label>
                    <input type="text" className="form-control" id="experience" name="experience"/>
                </div>
                <input type="submit" value="Submit" />
            </form>
        )
    }
}

export default NewCharacter