import React from 'react'
import "./note.css"
import Form from './noteForm'
import axios from 'axios'

const URL = `http://localhost:8080/notes` 
const inputs = {
    id: "",
    headline: "",
    content: "",
    reminder: "",
}


export default class Note extends React.Component {
    state = {
        inputs
    }

    onChange(e) {
        const {state} = this
        const input = e.target.name 
        state.inputs[input] =  e.target.value 
        this.setState(state)
    }
    clear() {
        const {state} = this
        state.inputs = {
            id: "",
            headline: "",
            content: "",
            reminder: "",
        }
        this.setState(state)
    }

    onClick(e) {
        e.preventDefault()
        const {state} = this
        try{
            axios.post(URL, state.inputs)
            .then(resp =>{ 
                this.clear()
                this.props.history.push("/notes")
            })
            .catch(resp => console.log(resp))

        }catch(e){
            alert(e)
        }
        
    }
    
    render() {
        return (
                <div className="container">
                    <div className="row align-items-center">
                        <Form onChange={this.onChange.bind(this)} 
                        {...this.state.inputs}
                        onClick={this.onClick.bind(this)} />
                    </div>
                </div>
                )
    }
} 