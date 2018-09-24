import React from 'react'
import axios from 'axios'

const URL = "http://localhost:8080/notes"

export default class NoteDisplay extends React.Component {
    state = {
        notes: []
    }

    componentDidMount() {
        axios.get(URL)
        .then(resp => this.setState({...this.state, notes: resp.data}))
    }

    removeNote(e) {
        const id = e.target.parentNode.getAttribute("data-id")
        axios.delete(`${URL}/${id}`)
        .then(resp => window.location.reload(true))
    }

    updateNote(e) {
        const id = e.target.parentNode.getAttribute("data-id")
        this.props.history.push(`/notes/${id}`)
    }

    render() {
        if (this.state.notes){
            return (
                <div className="container">
                    <table className="table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Headline</th>
                                <th>CreatedAt</th>
                                <th>Reminder</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                this.state.notes.map(el => {
                                    return (
                                        <tr key={el.id}>
                                            <td>{el.id}</td>
                                            <td>{el.headline}</td>
                                            <td>{el.created}</td>
                                            <td>{el.reminder}</td>
                                            <td data-id={el.id}>
                                                <button onClick={e => this.removeNote(e)}
                                                className="btn btn-danger mr-2">
                                                    <i className="fa fa-trash"></i>
                                                </button>
                                                <button  className="btn btn-primary"onClick={this.updateNote.bind(this)}>
                                                    <i className="fa fa-edit"></i>
                                                </button>
                                            </td>
                                        </tr>
                                    )   
                                })
                            }
                        </tbody>
                    </table>
                </div>
            )

        } else {
            return (
                <div className="container">
                    <h1>Não existem dados!</h1>
                </div>
            )
        }
    }
}
