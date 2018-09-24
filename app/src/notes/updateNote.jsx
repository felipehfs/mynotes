import React from 'react'
import axios from 'axios'
import NoteForm from './noteForm'


const  note = {
    id: "",
    headline: "",
    content: "",
    reminder: ""
}
const URL = "http://localhost:8080/notes"

export default class UpdateNote extends React.Component {

    state = {
       note
    }

    componentDidMount(){
        const id = this.props.match.params.id
        if (id) {
            axios.get(`${URL}/${id}`)
            .then(resp => this.setState({...this.state, note: resp.data }))
            .catch(resp => alert(resp.data))
        }
    }

    onChange(e) {
        const {state} = this
        const input = e.target.name 
        state.note[input] =  e.target.value  
        this.setState(state)
    }
    onClick(e) {
        const id = this.props.match.params.id

        axios.put(`${URL}/${id}`, this.state.note)
        .then(resp => window.location.hash = "#/notes")
    }
    render() {
        return (
            <div className="container">
                <div className="row align-items-center">
                    <NoteForm onChange={e => this.onChange(e)}
                     {...this.state.note}
                     onClick={e => this.onClick(e)}
                    />
                </div>
            </div>
        )
    }
}